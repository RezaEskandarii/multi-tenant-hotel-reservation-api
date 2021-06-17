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

func NewCountryService() *CountryService {
	return &CountryService{}
}

func (s *CountryService) Create(country *models.Country) (*models.Country, error) {

	return s.Repository.Create(country)
}

func (s *CountryService) Update(country *models.Country) (*models.Country, error) {

	return s.Repository.Update(country)
}

func (s *CountryService) Find(id uint64) (*models.Country, error) {

	return s.Repository.Find(id)
}

func (s *CountryService) FindAll(input *dto.PaginationInput) (*commons.PaginatedList, error) {

	return s.Repository.FindAll(input)
}

func (s *CountryService) GetProvinces(countryId uint64) ([]*models.City, error) {

	return s.Repository.GetProvinces(countryId)
}
