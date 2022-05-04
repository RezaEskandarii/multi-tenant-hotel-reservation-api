package repositories

import (
	"bytes"
	"crypto/rand"
	"database/sql"
	"fmt"
	"gorm.io/gorm"
	"math"
	"math/big"
	"reservation-api/internal/commons"
	"reservation-api/internal/config"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
	"reservation-api/internal/utils"
	"strings"
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

func (r *ReservationRepository) CreateReservationRequest(requestDto *dto.RoomRequestDto) (*models.ReservationRequest, error) {

	// read from default config
	expireTime := config.RoomDefaultLockDuration
	buffer := bytes.Buffer{}

	// default hour
	checkInDate := time.Date(requestDto.CheckInDate.Year(), requestDto.CheckInDate.Month(), requestDto.CheckInDate.Day(), 12, 0, 0, 0, requestDto.CheckInDate.Location())
	checkOutDate := time.Date(requestDto.CheckOutDate.Year(), requestDto.CheckOutDate.Month(), requestDto.CheckOutDate.Day(), 12, 0, 0, 0, requestDto.CheckOutDate.Location())

	// fill checkin date to converted checkinDate with default hour.
	requestDto.CheckInDate = &checkInDate
	// fill checkin date to converted checkinDate with default hour.
	requestDto.CheckOutDate = &checkOutDate

	// get random number.
	rnd, err := rand.Int(rand.Reader, big.NewInt(5))

	// generate reservation request key.
	requestKey := utils.GenerateSHA256(fmt.Sprintf("%s%s%s%s", expireTime, buffer.String(), requestDto.CheckInDate.String(), requestDto.CheckOutDate.String()))

	if err == nil {
		buffer.WriteString(rnd.String())
	}

	requestModel := models.ReservationRequest{
		RoomId:       requestDto.RoomId,
		ExpireTime:   expireTime,
		RequestKey:   requestKey,
		CheckOutDate: requestDto.CheckOutDate,
		CheckInDate:  requestDto.CheckInDate,
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
	reservation.Id = id
	tx := r.DB.Begin()
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

	result, err := r.Find(id)

	return result, err
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

	if status == models.Checkout && reservation.CheckoutDate.After(time.Now()) {
		checkoutDate := time.Now()
		reservation.CheckoutDate = &checkoutDate
	}

	if err := r.DB.Model(&models.Reservation{}).Where("id=?", id).Update("check_status", status).Error; err != nil {
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

func (r *ReservationRepository) HasConflict(request *dto.RoomRequestDto, reservation *models.Reservation) (bool, error) {

	var reservationRequestCount int64 = 0

	if err := r.DB.Model(&models.ReservationRequest{}).
		Where("room_id=? AND check_in_date >=? AND check_out_date<=? ",
			request.RoomId, request.CheckInDate, request.CheckOutDate).Count(&reservationRequestCount).Error; err != nil {
		return false, err
	}

	if reservationRequestCount > 0 {
		return true, nil
	}

	if request.RequestType == dto.CreateReservation {
		hasReservationConflict, err := r.HasReservationConflict(request.CheckInDate, request.CheckOutDate, request.RoomId)
		if err != nil {
			return false, err
		}
		return hasReservationConflict, nil
	}

	if request.RequestType == dto.UpdateReservation && reservation != nil {
		// check if reservation checkin or checkout date changes in update operation
		// and prevent to conflict with other reservations in update operations.
		if request.CheckInDate.Before(*reservation.CheckinDate) || request.CheckInDate.After(*request.CheckInDate) ||
			request.CheckOutDate.Before(*reservation.CheckoutDate) || request.CheckInDate.After(*reservation.CheckoutDate) {

			hasReservationConflict, err := r.HasReservationConflict(request.CheckInDate, request.CheckOutDate, request.RoomId)
			if err != nil {
				return false, err
			}
			return hasReservationConflict, nil

		}

	}

	return false, nil
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

func (r *ReservationRepository) FindReservationRequest(requestKey string) (*models.ReservationRequest, error) {
	reservationRequest := models.ReservationRequest{}
	if err := r.DB.Where("request_key=?", requestKey).Find(&reservationRequest).Error; err != nil {
		return nil, err
	}
	if reservationRequest.Id == 0 {
		return nil, nil
	}
	return &reservationRequest, nil
}

func (r *ReservationRepository) RemoveExpiredReservationRequests() error {
	err := r.DB.Where("expire_time < ?", time.Now()).Delete(&models.ReservationRequest{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *ReservationRepository) FindAll(filter *dto.ReservationFilter) (error, *commons.PaginatedResult) {

	reservations := make([]*models.Reservation, 0)

	query := r.DB.Model(&models.Reservation{})
	query = r.getReservationFilteredQuery(query, filter)

	if err := query.Scan(&reservations).Error; err != nil {
		return err, nil
	}

	return nil, paginateWithFilter(query, reservations, filter, filter.Page, filter.PerPage)

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

func (r *ReservationRepository) getReservationFilteredQuery(query *gorm.DB, filter *dto.ReservationFilter) *gorm.DB {

	if filter.CreatedFrom != nil {
		query = query.Where("created_at >= ?", filter.CreatedFrom)
	}

	if filter.CreatedTo != nil {
		query = query.Where("created_at <= ?", filter.CreatedTo)
	}

	if filter.CheckInFrom != nil {
		query = query.Where("checkin_date >= ?", filter.CheckInFrom)
	}

	if filter.CheckInTo != nil {
		query = query.Where("checkin_date >= ?", filter.CheckInTo)
	}

	if strings.TrimSpace(filter.GuestName) != "" {
		query = query.Where("supervisor.first_name LIKE '%?%' OR "+
			"supervisor.middle_name LIKE '%?%' OR "+
			"supervisor.last_name LIKE '%?%'",
			filter.GuestName, filter.GuestName, filter.GuestName)
	}

	if filter.RateCodeId != 0 {
		query = query.Where("rate_code_id=?", filter.RateCodeId)
	}

	if filter.RoomId != 0 {
		query = query.Where("room_id=?", filter.RateCodeId)
	}

	if filter.CheckStatus != nil {
		query = query.Where("check_status=?", filter.CheckStatus)
	}

	return query
}
