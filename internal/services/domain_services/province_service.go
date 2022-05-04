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
func (s *ProvinceService) Create(Province *models.Province) (*models.Province, error) {

	return s.Repository.Create(Province)
}

// Update updates province.
func (s *ProvinceService) Update(Province *models.Province) (*models.Province, error) {

	return s.Repository.Update(Province)
}

// Find returns province and if it does not find the province, it returns nil.
func (s *ProvinceService) Find(id uint64) (*models.Province, error) {

	return s.Repository.Find(id)
}

// FindAll returns paginates list of provinces.
func (s *ProvinceService) FindAll(input *dto.PaginationFilter) (*commons.PaginatedResult, error) {

	return s.Repository.FindAll(input)
}

// GetCities returns the list of cities that belong to the given province ID.
func (s *ProvinceService) GetCities(ProvinceId uint64) ([]*models.City, error) {

	return s.Repository.GetCities(ProvinceId)
}
