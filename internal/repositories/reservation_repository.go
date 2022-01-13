package repositories

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"gorm.io/gorm"
	"math/big"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
	"reservation-api/internal/utils"
	"time"
)

type ReservationRepository struct {
	DB *gorm.DB
}

// NewReservationRepository returns new ReservationRepository
func NewReservationRepository(db *gorm.DB) *ReservationRepository {
	return &ReservationRepository{DB: db}
}

func (r *ReservationRepository) CreateReservationRequest(roomId uint64) (*models.ReservationRequest, error) {

	expireTime := time.Now().Add(time.Minute * 20)
	buffer := bytes.Buffer{}
	rnd, err := rand.Int(rand.Reader, big.NewInt(5))
	if err == nil {
		buffer.WriteString(rnd.String())
	}
	requestModel := models.ReservationRequest{
		RoomId:     roomId,
		ExpireTime: expireTime,
		RequestKey: utils.GenerateSHA256(fmt.Sprintf("%s%s", expireTime, buffer.String())),
	}

	if err := r.DB.Create(&requestModel).Error; err != nil {
		return nil, err
	}

	return &requestModel, nil
}

func (r *ReservationRepository) Create(reservation *models.Reservation) (*models.Reservation, error) {
	db := r.DB

	reservationRequest := models.ReservationRequest{}
	if err := db.Where("request_key=? AND room_id=?", reservation.RequestKey, reservation.RoomId).Find(&reservationRequest).Error; err != nil {

	}
	// check if exists.
	if reservationRequest.Id != 0 {
		if time.Now().Before(reservationRequest.ExpireTime) {

		}
	}

	return nil, nil
}

func (r *ReservationRepository) Update() (*models.Reservation, error) {
	panic("not implemented")
}

func (r ReservationRepository) CheckIn(model *models.Reservation) error {
	panic("not implemented")
}

func (r ReservationRepository) CheckOut(model *models.Reservation) error {
	panic("not implemented")
}

func (r *ReservationRepository) CalculatePrice(priceDto *dto.GetRatePriceDto) ([]*dto.RateCodePricesDto, error) {

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
