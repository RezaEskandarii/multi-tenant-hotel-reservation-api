package services

import (
	"hotel-reservation/internal/models"
	"hotel-reservation/internal/repositories"
)

type CountryService struct {
	Repository repositories.CountryRepository
}

func NewCountryService() *CountryService {
	return &CountryService{}
}

func (s *CountryService) Create(country *models.Country) (*models.Country, error) {
	return s.Repository.Create(country)
}
