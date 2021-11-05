package services

import (
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
	"reservation-api/internal/repositories"
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

// FindAll returns paginates list of residence types.
func (s *ResidenceTypeService) FindAll(input *dto.PaginationInput) (*commons.PaginatedList, error) {

	return s.Repository.FindAll(input)
}

// Delete removes residence type by given id.
func (s *ResidenceTypeService) Delete(id uint64) error {

	return s.Repository.Delete(id)
}
