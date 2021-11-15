package services

import (
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
	"reservation-api/internal/repositories"
	"reservation-api/pkg/cache"
)

type CityService struct {
	Repository   *repositories.CityRepository
	CacheManager cache.Manager
}

// NewCityService returns new CityService
func NewCityService() *CityService {
	return &CityService{}
}

// Create creates new city.
func (s *CityService) Create(city *models.City) (*models.City, error) {
	s.CacheManager.Set("", city, nil)
	return s.Repository.Create(city)
}

// Update updates city.
func (s *CityService) Update(city *models.City) (*models.City, error) {

	return s.Repository.Update(city)
}

// Find returns city and if it does not find the city, it returns nil.
func (s *CityService) Find(id uint64) (*models.City, error) {

	return s.Repository.Find(id)
}

// FindAll returns paginates list of cities.
func (s *CityService) FindAll(input *dto.PaginationInput) (*commons.PaginatedList, error) {

	return s.Repository.FindAll(input)
}
