package services

import (
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
	"reservation-api/internal/repositories"
)

type ResidenceGradeService struct {
	Repository *repositories.ResidenceGradeRepository
}

// NewResidenceGradeService returns new ResidenceGradeService
func NewResidenceGradeService() *ResidenceGradeService {
	return &ResidenceGradeService{}
}

// Create creates new ResidenceGrade.
func (s *ResidenceGradeService) Create(residenceGrade *models.ResidenceGrade) (*models.ResidenceGrade, error) {

	return s.Repository.Create(residenceGrade)
}

// Update updates ResidenceGrade.
func (s *ResidenceGradeService) Update(residenceGrade *models.ResidenceGrade) (*models.ResidenceGrade, error) {

	return s.Repository.Update(residenceGrade)
}

// Find returns ResidenceGrade and if it does not find the ResidenceGrade, it returns nil.
func (s *ResidenceGradeService) Find(id uint64) (*models.ResidenceGrade, error) {

	return s.Repository.Find(id)
}

// FindAll returns paginates list of residence grades.
func (s *ResidenceGradeService) FindAll(input *dto.PaginationInput) (*commons.PaginatedList, error) {

	return s.Repository.FindAll(input)
}

// Delete removes residence type by given id.
func (s *ResidenceGradeService) Delete(id uint64) error {

	return s.Repository.Delete(id)
}
