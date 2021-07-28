package services

import (
	"hotel-reservation/internal/commons"
	"hotel-reservation/internal/dto"
	"hotel-reservation/internal/models"
	"hotel-reservation/internal/repositories"
)

type ResidenceService struct {
	Repository *repositories.ResidenceRepository
}

// NewResidenceService returns new ResidenceService
func NewResidenceService() *ResidenceService {

	return &ResidenceService{}
}

// Create creates new Residence.
func (s *ResidenceService) Create(residence *models.Residence) (*models.Residence, error) {

	return s.Repository.Create(residence)
}

// Update updates Residence.
func (s *ResidenceService) Update(residence *models.Residence) (*models.Residence, error) {

	return s.Repository.Update(residence)
}

// Find returns Residence and if it does not find the Residence, it returns nil.
func (s *ResidenceService) Find(id uint64) (*models.Residence, error) {

	return s.Repository.Find(id)
}

// FindAll returns paginates list of currencies
func (s *ResidenceService) FindAll(input *dto.PaginationInput) (*commons.PaginatedList, error) {

	return s.Repository.FindAll(input)
}

// Delete removes residence type by given id.
func (s *ResidenceService) Delete(id uint64) error {

	return s.Repository.Delete(id)
}
