package domain_services

import (
	"context"
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
func (s *CountryService) Create(ctx context.Context, country *models.CountryCreateUpdate) (*models.Country, error) {
	model := &models.Country{
		Name:  country.Name,
		Alias: country.Alias,
	}
	return s.Repository.Create(ctx, model)
}

// Update updates country.
func (s *CountryService) Update(ctx context.Context, country *models.Country) (*models.Country, error) {

	return s.Repository.Update(ctx, country)
}

// Find returns country and if it does not find the country, it returns nil.
func (s *CountryService) Find(ctx context.Context, id uint64) (*models.Country, error) {

	return s.Repository.Find(ctx, id)
}

// FindAll returns paginates list of countries.
func (s *CountryService) FindAll(ctx context.Context, input *dto.PaginationFilter) (*commons.PaginatedResult, error) {

	return s.Repository.FindAll(ctx, input)
}

// GetProvinces returns provinces by given countryId.
func (s *CountryService) GetProvinces(ctx context.Context, countryId uint64) ([]*models.Province, error) {

	return s.Repository.GetProvinces(ctx, countryId)
}
