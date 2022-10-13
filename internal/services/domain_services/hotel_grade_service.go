package domain_services

import (
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
	"reservation-api/internal/repositories"
)

type HotelGradeService struct {
	Repository *repositories.HotelGradeRepository
}

// NewHotelGradeService returns new HotelGradeService
func NewHotelGradeService(r *repositories.HotelGradeRepository) *HotelGradeService {
	return &HotelGradeService{Repository: r}
}

// Create creates new HotelGrade.
func (s *HotelGradeService) Create(hotelGrade *models.HotelGrade, tenantID uint64) (*models.HotelGrade, error) {

	return s.Repository.Create(hotelGrade, tenantID)
}

// Update updates HotelGrade.
func (s *HotelGradeService) Update(hotelGrade *models.HotelGrade, tenantID uint64) (*models.HotelGrade, error) {

	return s.Repository.Update(hotelGrade, tenantID)
}

// Find returns HotelGrade and if it does not find the HotelGrade, it returns nil.
func (s *HotelGradeService) Find(id uint64, tenantID uint64) (*models.HotelGrade, error) {

	return s.Repository.Find(id, tenantID)
}

// FindAll returns paginates list of hotel grades.
func (s *HotelGradeService) FindAll(input *dto.PaginationFilter) (*commons.PaginatedResult, error) {

	return s.Repository.FindAll(input)
}

// Delete removes hotel type by given id.
func (s *HotelGradeService) Delete(id uint64, tenantID uint64) error {

	return s.Repository.Delete(id, tenantID)
}
