package domain_services

import (
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
	"reservation-api/internal/repositories"
)

type RateCodeService struct {
	Repository *repositories.RateCodeRepository
}

// NewRateCodeService returns new RateCodeService
func NewRateCodeService(r *repositories.RateCodeRepository) *RateCodeService {
	return &RateCodeService{Repository: r}
}

// Create creates new RateCode.
func (s *RateCodeService) Create(model *models.RateCode) (*models.RateCode, error) {

	return s.Repository.Create(model)
}

// Update updates RateCode.
func (s *RateCodeService) Update(model *models.RateCode) (*models.RateCode, error) {

	return s.Repository.Update(model)
}

// Find returns RateCode and if it does not find the RateCode, it returns nil.
func (s *RateCodeService) Find(id uint64) (*models.RateCode, error) {

	return s.Repository.Find(id)
}

// FindAll returns paginates list of RateCodes.
func (s *RateCodeService) FindAll(input *dto.PaginationFilter) (*commons.PaginatedResult, error) {

	return s.Repository.FindAll(input)
}

// Delete removes RateCode  by given id.
func (s *RateCodeService) Delete(id uint64) error {

	return s.Repository.Delete(id)
}
