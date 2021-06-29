package services

import (
	"hotel-reservation/internal/commons"
	"hotel-reservation/internal/dto"
	"hotel-reservation/internal/models"
	"hotel-reservation/internal/repositories"
)

type CountryService struct {
	Repository repositories.CountryRepository
}

// NewCountryService returns new countryService
func NewCountryService() *CountryService {
	return &CountryService{}
}

// Create creates new country.
func (s *CountryService) Create(country *models.Country) (*models.Country, error) {

	return s.Repository.Create(country)
}

// Update updates country.
func (s *CountryService) Update(country *models.Country) (*models.Country, error) {

	return s.Repository.Update(country)
}

// Find returns country and if it does not find the country, it returns nil.
func (s *CountryService) Find(id uint64) (*models.Country, error) {

	return s.Repository.Find(id)
}

// FindAll returns paginates list of countries.
func (s *CountryService) FindAll(input *dto.PaginationInput) (*commons.PaginatedList, error) {

	return s.Repository.FindAll(input)
}

// GetProvinces returns provinces by given countryId.
func (s *CountryService) GetProvinces(countryId uint64) ([]*models.Province, error) {

	return s.Repository.GetProvinces(countryId)
}
