package repositories

import "gorm.io/gorm"

type ReservationRepository struct {
	DB *gorm.DB
}

func NewReservationRepository(db *gorm.DB) *ReservationRepository {
	return &ReservationRepository{DB: db}
}
