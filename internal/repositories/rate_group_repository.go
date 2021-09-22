package repositories

import "gorm.io/gorm"

type RateGroupRepository struct {
	DB *gorm.DB
}

// NewRateGroupRepository returns new RateGroupRepository.
func NewRateGroupRepository(db *gorm.DB) *RateGroupRepository {

	return &RateGroupRepository{DB: db}
}
