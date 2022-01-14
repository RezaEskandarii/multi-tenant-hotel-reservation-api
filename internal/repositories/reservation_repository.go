package repositories

import (
	"bytes"
	"crypto/rand"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"math/big"
	"reservation-api/internal/config"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
	"reservation-api/internal/utils"
	"strings"
	"time"
)

var (
	InvalidReservationKeyErr = errors.New("invalid reservation key")
)

type ReservationRepository struct {
	DB *gorm.DB
}

// NewReservationRepository returns new ReservationRepository
func NewReservationRepository(db *gorm.DB) *ReservationRepository {
	return &ReservationRepository{DB: db}
}

func (r *ReservationRepository) CreateReservationRequest(dto *dto.RoomRequestDto) (*models.ReservationRequest, error) {

	if dto.CheckInDate == nil {
		return nil, errors.New("checkInDate is empty")
	}
	if dto.CheckOutDate == nil {
		return nil, errors.New("checkOutDate is empty")
	}
	expireTime := config.RoomDefaultLockDuration
	buffer := bytes.Buffer{}
	rnd, err := rand.Int(rand.Reader, big.NewInt(5))
	requestKey := utils.GenerateSHA256(fmt.Sprintf("%s%s%s%s", expireTime, buffer.String(), dto.CheckInDate.String(), dto.CheckOutDate.String()))
	if err == nil {
		buffer.WriteString(rnd.String())
	}
	requestModel := models.ReservationRequest{
		RoomId:       dto.RoomId,
		ExpireTime:   expireTime,
		RequestKey:   requestKey,
		CheckoutDate: dto.CheckOutDate,
		CheckInDate:  dto.CheckInDate,
	}

	if err := r.DB.Create(&requestModel).Error; err != nil {
		return nil, err
	}

	return &requestModel, nil
}

func (r *ReservationRepository) Create(reservation *models.Reservation) (*models.Reservation, error) {
	db := r.DB
	if strings.TrimSpace(reservation.RequestKey) == "" {
		return nil, InvalidReservationKeyErr
	}
	reservationRequest := models.ReservationRequest{}
	if err := db.Where("request_key=? AND room_id=?", reservation.RequestKey, reservation.RoomId).Find(&reservationRequest).Error; err != nil {
		return nil, err
	}
	// check if exists.
	if reservationRequest.Id == 0 {
		return nil, InvalidReservationKeyErr
	}
	if time.Now().After(reservationRequest.ExpireTime) {
		return nil, InvalidReservationKeyErr
	}
	if reserveErr := db.Create(&reservation).Error; reserveErr != nil {
		return nil, reserveErr
	}
	return nil, nil
}

func (r *ReservationRepository) Update(model *models.Reservation) (*models.Reservation, error) {
	panic("not implemented")
}

func (r ReservationRepository) CheckIn(model *models.Reservation) error {
	panic("not implemented")
}

func (r ReservationRepository) CheckOut(model *models.Reservation) error {
	panic("not implemented")
}

func (r *ReservationRepository) GetRelatedRateCodes(priceDto *dto.GetRatePriceDto) ([]*dto.RateCodePricesDto, error) {

	db := r.DB
	ratePrices := make([]*dto.RateCodePricesDto, 0)

	db.Table("rate_code_details details").Select(`
	   parent.name as rate_code_name,
       details.rate_code_id,
       details.created_at,
       details.room_id,
       details.id,
       details.date_start,
       details.date_end,
       prices.price,
       prices.guest_count
	`).Joins(`
         join rate_code_detail_prices prices
              on prices.rate_code_detail_id = details.id
         join rate_codes parent 
              on details.rate_code_id = parent.id
	`).Where(`
		  details.room_id = ?
		  and prices.guest_count = ?
		  and details.min_nights >= ?
		  and details.date_start >= ?
		  and details.date_end <= ?
	`, priceDto.RoomId, priceDto.GuestCount, priceDto.NightCount, priceDto.DateStart, priceDto.DateEnd).Scan(&ratePrices)

	return ratePrices, nil
	//if len(ratePrices) == 0 {
	//	return nil, errors.New("rate code details not found")
	//}
	//
	//if len(ratePrices) == 1 {
	//	return ratePrices[0], nil
	//}
	//
	//// get latest inserted rate code details price by created at time.
	//result := ratePrices[0]
	//for _, v := range ratePrices {
	//	if v.CreatedAt.After(*result.CreatedAt) {
	//		result = v
	//	}
	//}
	//
	//return result, nil
}
