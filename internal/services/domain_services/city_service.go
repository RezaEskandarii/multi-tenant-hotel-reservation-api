package domain_services

import (
	"context"
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
func (s *CityService) Create(ctx context.Context, city *models.City) (*models.City, error) {

	city, err := s.Repository.Create(ctx, city)

	if err != nil && city != nil {
		s.CacheManager.Set(getCityCacheKey(city), city, nil)
	}

	return city, err
}

// Update updates city.
func (s *CityService) Update(ctx context.Context, city *models.City) (*models.City, error) {
	s.CacheManager.Update(getCityCacheKey(city), city)
	return s.Repository.Update(ctx, city)
}

// Find returns city and if it does not find the city, it returns nil.
func (s *CityService) Find(ctx context.Context, id uint64) (*models.City, error) {

	return s.Repository.Find(ctx, id)
}

// Delete delete city by given id.
func (s *CityService) Delete(ctx context.Context, id uint64) error {

	city, err := s.Find(ctx, id)

	if err == nil {
		s.CacheManager.Del(getCityCacheKey(city))
		return s.Repository.Delete(ctx, id)
	}
	return err
}

// FindAll returns paginates list of cities.
func (s *CityService) FindAll(ctx context.Context, input *dto.PaginationFilter) (*commons.PaginatedResult, error) {

	return s.Repository.FindAll(ctx, input)

}

func getCityCacheKey(city *models.City) string {
	return utils.GenerateCacheKey(city.Id, city.Name, city.ProvinceId)
}
