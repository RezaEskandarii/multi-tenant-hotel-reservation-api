package domain_services

import (
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
	"reservation-api/internal/repositories"
)

type CountryService struct {
	Repository *repositories.CountryRepository
}

// NewCountryService returns new CountryService
func NewCountryService(r *repositories.CountryRepository) *CountryService {
	return &CountryService{Repository: r}
}

// Create creates new country.
func (s *CountryService) Create(country *models.Country, tenantID uint64) (*models.Country, error) {

	return s.Repository.Create(country, tenantID)
}

// Update updates country.
func (s *CountryService) Update(country *models.Country, tenantID uint64) (*models.Country, error) {

	return s.Repository.Update(country, tenantID)
}

// Find returns country and if it does not find the country, it returns nil.
func (s *CountryService) Find(id uint64, tenantID uint64) (*models.Country, error) {

	return s.Repository.Find(id, tenantID)
}

// FindAll returns paginates list of countries.
func (s *CountryService) FindAll(input *dto.PaginationFilter) (*commons.PaginatedResult, error) {

	return s.Repository.FindAll(input)
}

// GetProvinces returns provinces by given countryId.
func (s *CountryService) GetProvinces(countryId uint64, tenantID uint64) ([]*models.Province, error) {

	return s.Repository.GetProvinces(countryId, tenantID)
}
