package domain_services

import (
	"context"
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
func (s *RateCodeService) Create(ctx context.Context, model *models.RateCode) (*models.RateCode, error) {

	return s.Repository.Create(ctx, model)
}

// Update updates RateCode.
func (s *RateCodeService) Update(ctx context.Context, model *models.RateCode) (*models.RateCode, error) {

	return s.Repository.Update(ctx, model)
}

// Find returns RateCode and if it does not find the RateCode, it returns nil.
func (s *RateCodeService) Find(ctx context.Context, id uint64) (*models.RateCode, error) {

	return s.Repository.Find(ctx, id)
}

// FindAll returns paginates list of RateCodes.
func (s *RateCodeService) FindAll(ctx context.Context, filter *dto.PaginationFilter) (*commons.PaginatedResult, error) {

	return s.Repository.FindAll(ctx, filter)
}

// Delete removes RateCode  by given id.
func (s *RateCodeService) Delete(ctx context.Context, id uint64) error {

	return s.Repository.Delete(ctx, id)
}
