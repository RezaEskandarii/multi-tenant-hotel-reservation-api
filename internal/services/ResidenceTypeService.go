package services

import (
	"hotel-reservation/internal/commons"
	"hotel-reservation/internal/dto"
	"hotel-reservation/internal/models"
	"hotel-reservation/internal/repositories"
)

type ResidenceTypeService struct {
	Repository *repositories.ResidenceTypeRepository
}

// NewResidenceTypeService returns new ResidenceTypeService
func NewResidenceTypeService() *ResidenceTypeService {
	return &ResidenceTypeService{}
}

// Create creates new ResidenceType.
func (s *ResidenceTypeService) Create(residenceType *models.ResidenceType) (*models.ResidenceType, error) {

	return s.Repository.Create(residenceType)
}

// Update updates ResidenceType.
func (s *ResidenceTypeService) Update(residenceType *models.ResidenceType) (*models.ResidenceType, error) {

	return s.Repository.Update(residenceType)
}

// Find returns ResidenceType and if it does not find the ResidenceType, it returns nil.
func (s *ResidenceTypeService) Find(id uint64) (*models.ResidenceType, error) {

	return s.Repository.Find(id)
}

// FindAll returns paginates list of currencies
func (s *ResidenceTypeService) FindAll(input *dto.PaginationInput) (*commons.PaginatedList, error) {

	return s.Repository.FindAll(input)
}
