package domain_services

import (
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
func (s *RateGroupService) Create(model *models.RateGroup) (*models.RateGroup, error) {

	return s.Repository.Create(model)
}

// Update updates RateGroup.
func (s *RateGroupService) Update(model *models.RateGroup) (*models.RateGroup, error) {

	return s.Repository.Update(model)
}

// Find returns RateGroup and if it does not find the RateGroup, it returns nil.
func (s *RateGroupService) Find(id uint64) (*models.RateGroup, error) {

	return s.Repository.Find(id)
}

// FindAll returns paginates list of RateGroups.
func (s *RateGroupService) FindAll(input *dto.PaginationFilter) (*commons.PaginatedResult, error) {

	return s.Repository.FindAll(input)
}

// Delete removes RateGroup  by given id.
func (s *RateGroupService) Delete(id uint64) error {

	return s.Repository.Delete(id)
}
