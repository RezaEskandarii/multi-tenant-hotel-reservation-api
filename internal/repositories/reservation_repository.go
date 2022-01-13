package repositories

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"gorm.io/gorm"
	"math/big"
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

//func (r *ReservationRepository) CalculatePrice(reservation *models.Reservation) (float64, error) {
//	db := r.DB
//	var defaultPrice float64 = -1
//	rateCodeDetails := make([]models.RateCodeDetail, 0)
//	prices := make(map[time.Time]float64, 0)
//
//	if err := db.Where("room_id=?", reservation.RoomId).Preload("RateCodeDetailPrice").Find(&rateCodeDetails).Error; err != nil {
//		return defaultPrice, err
//	}
//
//	if rateCodeDetails == nil || (rateCodeDetails != nil && len(rateCodeDetails) == 0) {
//		return defaultPrice, errors.New("rate code not found matching with reservation information")
//	}
//
//	for _, details := range rateCodeDetails {
//		if details.RatePrices != nil {
//			if reservation.CheckinDate.After(*details.DateStart) && reservation.CheckoutDate.Before(*details.DateEnd) {
//				for _, priceDetail := range details.RatePrices {
//				}
//			}
//		}
//	}
//}

/**

select
details.created_at,details.room_id,details.id,details.rate_code_id,details.date_start,details.date_end,
prices.id,prices.price,prices.guest_count
from rate_code_details details
join rate_code_detail_prices prices
on prices.rate_code_detail_id=details.id where details.room_id=1
and prices.guest_count = 1 and details.min_nights >=1 and details.max_nights <=10
and details.date_start >='2019-02-02' and details.date_end <='2025-01-01'

*/
