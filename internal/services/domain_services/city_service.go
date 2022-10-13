package domain_services

import (
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
	"reservation-api/internal/repositories"
	"reservation-api/internal/services/common_services"
	"reservation-api/internal/utils"
)

type CityService struct {
	Repository   *repositories.CityRepository
	CacheManager common_services.CacheManager
}

// NewCityService returns new CityService
func NewCityService(repository *repositories.CityRepository, cm common_services.CacheManager) *CityService {
	return &CityService{
		Repository:   repository,
		CacheManager: cm,
	}
}

// Create creates new city.
func (s *CityService) Create(city *models.City, tenantID uint64) (*models.City, error) {

	city, err := s.Repository.Create(city, tenantID)

	if err != nil && city != nil {
		s.CacheManager.Set(getCityCacheKey(city, tenantID), city, nil)
	}

	return city, err
}

// Update updates city.
func (s *CityService) Update(city *models.City, tenantID uint64) (*models.City, error) {
	s.CacheManager.Update(getCityCacheKey(city, tenantID), city)
	return s.Repository.Update(city, tenantID)
}

// Find returns city and if it does not find the city, it returns nil.
func (s *CityService) Find(id uint64, tenantID uint64) (*models.City, error) {

	return s.Repository.Find(id, tenantID)
}

// Delete delete city by given id.
func (s *CityService) Delete(id uint64, tenantID uint64) error {

	city, err := s.Find(id, tenantID)

	if err == nil {
		s.CacheManager.Del(getCityCacheKey(city, tenantID))
		return s.Repository.Delete(id, tenantID)
	}
	return err
}

// FindAll returns paginates list of cities.
func (s *CityService) FindAll(input *dto.PaginationFilter) (*commons.PaginatedResult, error) {

	return s.Repository.FindAll(input)

}

func getCityCacheKey(city *models.City, tenantID uint64) string {
	return utils.GenerateCacheKey(city.Id, city.Name, city.ProvinceId, tenantID)
}
