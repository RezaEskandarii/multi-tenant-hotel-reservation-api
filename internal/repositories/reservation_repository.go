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
		LockKey:    utils.GenerateSHA256(fmt.Sprintf("%s%s", expireTime, buffer.String())),
	}

	if err := r.DB.Create(&requestModel).Error; err != nil {
		return nil, err
	}

	return &requestModel, nil
}

func (r *ReservationRepository) Create() (*models.Reservation, error) {
	panic("not implemented")
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
