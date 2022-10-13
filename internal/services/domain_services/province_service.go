package domain_services

import (
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
func (s *ProvinceService) Create(province *models.Province, tenantID uint64) (*models.Province, error) {

	return s.Repository.Create(province, tenantID)
}

// Update updates province.
func (s *ProvinceService) Update(province *models.Province, tenantID uint64) (*models.Province, error) {

	return s.Repository.Update(province, tenantID)
}

// Find returns province and if it does not find the province, it returns nil.
func (s *ProvinceService) Find(id uint64, tenantID uint64) (*models.Province, error) {

	return s.Repository.Find(id, tenantID)
}

// FindAll returns paginates list of provinces.
func (s *ProvinceService) FindAll(input *dto.PaginationFilter) (*commons.PaginatedResult, error) {

	return s.Repository.FindAll(input)
}

// GetCities returns the list of cities that belong to the given province ID.
func (s *ProvinceService) GetCities(provinceId uint64, tenantID uint64) ([]*models.City, error) {

	return s.Repository.GetCities(provinceId, tenantID)
}
