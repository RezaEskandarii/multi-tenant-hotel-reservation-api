package services

import (
	"hotel-reservation/internal/commons"
	"hotel-reservation/internal/dto"
	"hotel-reservation/internal/models"
	"hotel-reservation/internal/repositories"
)

type CityService struct {
	Repository repositories.CityRepository
}

func NewCityService() *CityService {
	return &CityService{}
}

func (s *CityService) Create(city *models.City) (*models.City, error) {

	return s.Repository.Create(city)
}

func (s *CityService) Update(city *models.City) (*models.City, error) {

	return s.Repository.Update(city)
}

func (s *CityService) Find(id uint64) (*models.City, error) {

	return s.Repository.Find(id)
}

func (s *CityService) FindAll(input *dto.PaginationInput) (*commons.PaginatedList, error) {

	return s.Repository.FindAll(input)
}
