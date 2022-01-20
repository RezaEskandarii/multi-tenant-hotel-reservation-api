package domain_services

import (
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
	"reservation-api/internal/repositories"
)

type RateCodeDetailService struct {
	Repository *repositories.RateCodeDetailRepository
}

// NewRateCodeDetailService returns new RateCodeDetailService
func NewRateCodeDetailService() *RateCodeDetailService {
	return &RateCodeDetailService{}
}

// Create creates new RateCodeDetail.
func (s *RateCodeDetailService) Create(model *models.RateCodeDetail) (*models.RateCodeDetail, error) {

	return s.Repository.Create(model)
}

// Update updates RateCodeDetail.
func (s *RateCodeDetailService) Update(model *models.RateCodeDetail) (*models.RateCodeDetail, error) {

	return s.Repository.Update(model)
}

// Find returns RateCodeDetail and if it does not find the RateCodeDetail, it returns nil.
func (s *RateCodeDetailService) Find(id uint64) (*models.RateCodeDetail, error) {

	return s.Repository.Find(id)
}

// FindAll returns paginates list of RateCodeDetails.
func (s *RateCodeDetailService) FindAll(input *dto.PaginationInput) (*commons.PaginatedList, error) {

	return s.Repository.FindAll(input)
}

// Delete removes RateCodeDetail  by given id.
func (s *RateCodeDetailService) Delete(id uint64) error {

	return s.Repository.Delete(id)
}

// FindPrice returns RateCodePrice by given id.
func (s *RateCodeDetailService) FindPrice(id uint64) (*models.RateCodeDetailPrice, error) {

	return s.Repository.FindPrice(id)
}
