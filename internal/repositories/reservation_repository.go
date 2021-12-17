package repositories

import (
	"fmt"
	"gorm.io/gorm"
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

func (r *ReservationRepository) CreateReservationRequest(model models.Reservation) (*models.ReservationRequest, error) {

	requestModel := models.ReservationRequest{
		RoomId:                  model.RoomId,
		GuestId:                 model.SupervisorId,
		RateCodeId:              model.RateCodeId,
		ExpireTime:              time.Now().Add(time.Minute * 20),
		LockKey:                 utils.GenerateSHA256(fmt.Sprintf("%d_%d_%d_%s_%s", model.RoomId, model.SupervisorId, model.RateCodeId, model.CheckinDate, model.CheckoutDate)),
		ReservationCheckinDate:  model.CheckinDate,
		ReservationCheckoutDate: model.CheckoutDate,
	}

	if err := r.DB.Create(&requestModel).Error; err != nil {
		return nil, err
	}

	return &requestModel, nil
}
