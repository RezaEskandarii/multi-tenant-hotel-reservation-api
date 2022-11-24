package domain_services

import (
	"context"
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
	"reservation-api/internal/repositories"
)

type RateGroupService struct {
	Repository *repositories.RateGroupRepository
}

// NewRateGroupService returns new RateGroupService
func NewRateGroupService(r *repositories.RateGroupRepository) *RateGroupService {
	return &RateGroupService{Repository: r}
}

// Create creates new RateGroup.
func (s *RateGroupService) Create(ctx context.Context, model *models.RateGroup) (*models.RateGroup, error) {

	return s.Repository.Create(ctx, model)
}

// Update updates RateGroup.
func (s *RateGroupService) Update(ctx context.Context, model *models.RateGroup) (*models.RateGroup, error) {

	return s.Repository.Update(ctx, model)
}

// Find returns RateGroup and if it does not find the RateGroup, it returns nil.
func (s *RateGroupService) Find(ctx context.Context, id uint64) (*models.RateGroup, error) {

	return s.Repository.Find(ctx, id)
}

// FindAll returns paginates list of RateGroups.
func (s *RateGroupService) FindAll(ctx context.Context, filter *dto.PaginationFilter) (*commons.PaginatedResult, error) {

	return s.Repository.FindAll(ctx, filter)
}

// Delete removes RateGroup  by given id.
func (s *RateGroupService) Delete(ctx context.Context, id uint64) error {

	return s.Repository.Delete(ctx, id)
}
