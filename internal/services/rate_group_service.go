package services

import (
	"hotel-reservation/internal/commons"
	"hotel-reservation/internal/dto"
	"hotel-reservation/internal/models"
	"hotel-reservation/internal/repositories"
)

type RateGroupService struct {
	Repository repositories.RateGroupRepository
}

// NewRateGroupService returns new RateGroupService
func NewRateGroupService() *RateGroupService {
	return &RateGroupService{}
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
func (s *RateGroupService) FindAll(input *dto.PaginationInput) (*commons.PaginatedList, error) {

	return s.Repository.FindAll(input)
}

// Delete removes RateGroup  by given id.
func (s *RateGroupService) Delete(id uint64) error {

	return s.Repository.Delete(id)
}
