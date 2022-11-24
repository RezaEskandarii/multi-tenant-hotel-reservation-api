package domain_services

import (
	"context"
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
	"reservation-api/internal/repositories"
)

type RateCodeDetailService struct {
	Repository *repositories.RateCodeDetailRepository
}

// NewRateCodeDetailService returns new RateCodeDetailService
func NewRateCodeDetailService(repository *repositories.RateCodeDetailRepository) *RateCodeDetailService {
	return &RateCodeDetailService{Repository: repository}
}

// Create creates new RateCodeDetail.
func (s *RateCodeDetailService) Create(ctx context.Context, model *models.RateCodeDetail) (*models.RateCodeDetail, error) {

	return s.Repository.Create(ctx, model)
}

// Update updates RateCodeDetail.
func (s *RateCodeDetailService) Update(ctx context.Context, model *models.RateCodeDetail) (*models.RateCodeDetail, error) {

	return s.Repository.Update(ctx, model)
}

// Find returns RateCodeDetail and if it does not find the RateCodeDetail, it returns nil.
func (s *RateCodeDetailService) Find(ctx context.Context, id uint64) (*models.RateCodeDetail, error) {

	return s.Repository.Find(ctx, id)
}

// FindAll returns paginates list of RateCodeDetails.
func (s *RateCodeDetailService) FindAll(ctx context.Context, filter *dto.PaginationFilter) (*commons.PaginatedResult, error) {

	return s.Repository.FindAll(ctx, filter)
}

// Delete removes RateCodeDetail  by given id.
func (s *RateCodeDetailService) Delete(ctx context.Context, id uint64) error {

	return s.Repository.Delete(ctx, id)
}

// FindPrice returns RateCodePrice by given id.
func (s *RateCodeDetailService) FindPrice(ctx context.Context, id uint64) (*models.RateCodeDetailPrice, error) {

	return s.Repository.FindPrice(ctx, id)
}
