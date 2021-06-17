package services

import (
	"hotel-reservation/internal/commons"
	"hotel-reservation/internal/dto"
	"hotel-reservation/internal/models"
	"hotel-reservation/internal/repositories"
)

type ProvinceService struct {
	Repository repositories.ProvinceRepository
}

func NewProvinceService() *ProvinceService {
	return &ProvinceService{}
}

func (s *ProvinceService) Create(Province *models.Province) (*models.Province, error) {

	return s.Repository.Create(Province)
}

func (s *ProvinceService) Update(Province *models.Province) (*models.Province, error) {

	return s.Repository.Update(Province)
}

func (s *ProvinceService) Find(id uint64) (*models.Province, error) {

	return s.Repository.Find(id)
}

func (s *ProvinceService) FindAll(input *dto.PaginationInput) (*commons.PaginatedList, error) {

	return s.Repository.FindAll(input)
}

func (s *ProvinceService) GetCities(ProvinceId uint64) ([]*models.City, error) {

	return s.Repository.GetCities(ProvinceId)
}
