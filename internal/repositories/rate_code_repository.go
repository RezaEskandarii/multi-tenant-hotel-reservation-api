package repositories

import "gorm.io/gorm"

type RateCodeRepository struct {
	DB *gorm.DB
}

// NewRateCodeRepository returns new RateCodeRepository.
func NewRateCodeRepository(db *gorm.DB) *RateCodeRepository {

	return &RateCodeRepository{DB: db}
}
