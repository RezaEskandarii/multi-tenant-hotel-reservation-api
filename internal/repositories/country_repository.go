package repositories

import (
	"gorm.io/gorm"
	"hotel-reservation/internal/models"
	"hotel-reservation/pkg/application_loger"
)

type CountryRepository struct {
	DB *gorm.DB
}

func NewCountryRepository(db *gorm.DB) *CountryRepository {
	return &CountryRepository{
		DB: db,
	}
}

func (r *CountryRepository) Create(country *models.Country) (*models.Country, error) {

	if err := r.DB.Create(&country); err != nil {
		application_loger.LogError(err.Error)
		return nil, err.Error
	}

	return country, nil
}
