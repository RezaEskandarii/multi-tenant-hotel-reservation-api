package domain_services

import (
	"context"
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
	"reservation-api/internal/repositories"
)

type ProvinceService struct {
	Repository *repositories.ProvinceRepository
}

// NewProvinceService returns new ProvinceService
func NewProvinceService(r *repositories.ProvinceRepository) *ProvinceService {
	return &ProvinceService{Repository: r}
}

// Create creates new province.
func (s *ProvinceService) Create(ctx context.Context, province *models.Province) (*models.Province, error) {

	return s.Repository.Create(ctx, province)
}

// Update updates province.
func (s *ProvinceService) Update(ctx context.Context, province *models.Province) (*models.Province, error) {

	return s.Repository.Update(ctx, province)
}

// Find returns province and if it does not find the province, it returns nil.
func (s *ProvinceService) Find(ctx context.Context, id uint64) (*models.Province, error) {

	return s.Repository.Find(ctx, id)
}

// FindAll returns paginates list of provinces.
func (s *ProvinceService) FindAll(ctx context.Context, filter *dto.PaginationFilter) (*commons.PaginatedResult, error) {

	return s.Repository.FindAll(ctx, filter)
}

// GetCities returns the list of cities that belong to the given province ID.
func (s *ProvinceService) GetCities(ctx context.Context, provinceId uint64) ([]*models.City, error) {

	return s.Repository.GetCities(ctx, provinceId)
}
