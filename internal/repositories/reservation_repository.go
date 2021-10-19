package repositories

import "gorm.io/gorm"

type ReservationRepository struct {
	DB *gorm.DB
}

// NewReservationRepository returns new ReservationRepository
func NewReservationRepository(db *gorm.DB) *ReservationRepository {
	return &ReservationRepository{DB: db}
}
