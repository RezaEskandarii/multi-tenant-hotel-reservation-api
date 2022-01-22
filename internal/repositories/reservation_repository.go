package repositories

import (
	"bytes"
	"crypto/rand"
	"database/sql"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"math"
	"math/big"
	"reservation-api/internal/config"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
	"reservation-api/internal/utils"
	"time"
)

type ReservationRepository struct {
	DB                 *gorm.DB
	RateCodeRepository *RateCodeDetailRepository
}

// NewReservationRepository returns new ReservationRepository
func NewReservationRepository(db *gorm.DB, rateCodeRepository *RateCodeDetailRepository) *ReservationRepository {
	return &ReservationRepository{
		DB:                 db,
		RateCodeRepository: rateCodeRepository,
	}
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
		CheckOutDate: dto.CheckOutDate,
		CheckInDate:  dto.CheckInDate,
	}

	if err := r.DB.Create(&requestModel).Error; err != nil {
		return nil, err
	}

	return &requestModel, nil
}

func (r *ReservationRepository) Create(reservation *models.Reservation) (*models.Reservation, error) {

	r.setReservationCalcFields(reservation)

	option := sql.TxOptions{
		Isolation: sql.LevelDefault,
		ReadOnly:  false,
	}

	tx := r.DB.Begin(&option)

	if err := tx.Create(&reservation).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	// remove reservation request after create reservation.
	if err := tx.Where("request_key=?", reservation.RequestKey).Delete(models.ReservationRequest{}).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return reservation, nil
}

func (r *ReservationRepository) Update(id uint64, reservation *models.Reservation) (*models.Reservation, error) {

	r.setReservationCalcFields(reservation)

	option := sql.TxOptions{
		Isolation: sql.LevelDefault,
		ReadOnly:  false,
	}

	tx := r.DB.Begin(&option)
	// remove old sharers and replace with new sharers.
	if err := tx.Where("reservation_id=?", id).Delete(&models.Sharer{}).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Where("id=?", id).Updates(&reservation).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	// remove reservation request after create reservation.
	if err := tx.Where("request_key=?", reservation.RequestKey).Delete(models.ReservationRequest{}).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return reservation, nil
}

// ChangeStatus changes the reservation check status.
func (r ReservationRepository) ChangeStatus(id uint64, status models.ReservationCheckStatus) (*models.Reservation, error) {
	reservation := models.Reservation{}
	if err := r.DB.Find(&reservation, id).Error; err != nil {
		return nil, err
	}
	if reservation.Id == 0 {
		return nil, nil
	}

	if err := r.DB.Model(&models.Reservation{}).Where("id=?", id).Update("CheckStatus", status).Error; err != nil {
		return nil, err
	}
	return &reservation, nil
}

func (r *ReservationRepository) GetRecommendedRateCodes(priceDto *dto.GetRatePriceDto) ([]*dto.RateCodePricesDto, error) {

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
       prices.guest_count,
       prices.id as rate_price_id
	`).Joins(`
         join rate_code_detail_prices prices
              on prices.rate_code_detail_id = details.id
         join rate_codes parent 
              on details.rate_code_id = parent.id
	`).Where(`
		  details.room_id = ?
		  and prices.guest_count = ?
		  and details.min_nights <= ?
		  and details.date_start <= ?
		  and details.date_end >= ?
          and details.rate_code_id=?
	`, priceDto.RoomId, priceDto.GuestCount, priceDto.NightCount, priceDto.DateStart, priceDto.DateEnd, priceDto.RateCodeId).Scan(&ratePrices)

	return ratePrices, nil
}

func (r *ReservationRepository) HasConflict(request *dto.RoomRequestDto) (bool, error) {

	var reservationRequestCount int64 = 0

	if err := r.DB.Model(&models.ReservationRequest{}).
		Where("room_id=? AND check_in_date >=? AND check_out_date<=? AND expire_time >=?",
			request.RoomId, request.CheckInDate, request.CheckOutDate, time.Now()).Count(&reservationRequestCount).Error; err != nil {
		return false, err
	}

	if reservationRequestCount > 0 {
		return true, nil
	}

	hasReservationConflict, err := r.HasReservationConflict(request.CheckInDate, request.CheckOutDate, request.RoomId)
	if err != nil {
		return false, err
	}

	return hasReservationConflict, nil
}

func (r *ReservationRepository) HasReservationConflict(checkInDate *time.Time, checkOutDate *time.Time, roomId uint64) (bool, error) {
	var count int64 = 0
	if err := r.DB.Model(&models.Reservation{}).
		Where("room_id=? AND checkin_date >=? AND checkout_date<=?",
			roomId, checkInDate, checkOutDate).Count(&count).Error; err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}

func (r *ReservationRepository) DeleteReservationRequest(requestKey string) error {
	var count int64 = 0
	if err := r.DB.Model(models.ReservationRequest{}).Where("requestKey=?", requestKey).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		if err := r.DB.Where("requestKey=?", requestKey).Delete(&models.ReservationRequest{}).Error; err != nil {
			return err
		}
	}
	return nil
}

func (r *ReservationRepository) Find(id uint64) (*models.Reservation, error) {
	reservation := models.Reservation{}
	db := r.DB.Model(models.Reservation{})
	db = r.preloadReservationRelations(db)
	if err := db.Where("id=?", id).Find(&reservation).Error; err != nil {
		return nil, err
	}
	if reservation.Id == 0 {
		return nil, nil
	}
	return &reservation, nil
}

func (r *ReservationRepository) FindReservationRequest(requestKey string, roomId uint64) (*models.ReservationRequest, error) {
	reservationRequest := models.ReservationRequest{}
	if err := r.DB.Where("request_key=? AND room_id=?", requestKey, roomId).Find(&reservationRequest).Error; err != nil {
		return nil, err
	}
	if reservationRequest.Id == 0 {
		return nil, nil
	}
	return &reservationRequest, nil
}

/*================= private functions ===========================================================*/

func (r *ReservationRepository) preloadReservationRelations(query *gorm.DB) *gorm.DB {
	return query.Preload("Room").Preload("Supervisor").Preload("RateCode").
		Preload("Sharers").Preload("Sharers.Guest")
}

func (r *ReservationRepository) calculatePrice(reservation *models.Reservation) float64 {

	priceDto := &dto.GetRatePriceDto{
		RoomId:     reservation.RoomId,
		NightCount: reservation.Nights,
		GuestCount: reservation.GuestCount,
		DateStart:  reservation.CheckinDate,
		DateEnd:    reservation.CheckoutDate,
		RateCodeId: reservation.RateCodeId,
	}

	prices, err := r.GetRecommendedRateCodes(priceDto)
	if err != nil {
		return 0
	}
	if len(prices) == 0 {
		return 0
	}
	if len(prices) == 1 {
		return prices[0].Price * reservation.Nights
	}
	defaultPrice := prices[0]
	for _, price := range prices {
		// get latest inserted.
		if price.CreatedAt.After(*defaultPrice.CreatedAt) {
			defaultPrice = price
		}
	}
	return defaultPrice.Price * reservation.Nights
}

// fill calculation fields
func (r *ReservationRepository) setReservationCalcFields(reservation *models.Reservation) {
	reservation.Nights = math.Round(reservation.CheckoutDate.Sub(*reservation.CheckinDate).Hours() / 24)
	reservation.GuestCount = uint64(len(reservation.Sharers))
	reservation.Price = r.calculatePrice(reservation)
}
