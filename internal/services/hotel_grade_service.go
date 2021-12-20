package services

import (
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
	"reservation-api/internal/repositories"
)

type HotelGradeService struct {
	Repository *repositories.HotelGradeRepository
}

// NewhotelGradeService returns new HotelGradeService
func NewhotelGradeService() *HotelGradeService {
	return &HotelGradeService{}
}

// Create creates new HotelGrade.
func (s *HotelGradeService) Create(hotelGrade *models.HotelGrade) (*models.HotelGrade, error) {

	return s.Repository.Create(hotelGrade)
}

// Update updates HotelGrade.
func (s *HotelGradeService) Update(hotelGrade *models.HotelGrade) (*models.HotelGrade, error) {

	return s.Repository.Update(hotelGrade)
}

// Find returns HotelGrade and if it does not find the HotelGrade, it returns nil.
func (s *HotelGradeService) Find(id uint64) (*models.HotelGrade, error) {

	return s.Repository.Find(id)
}

// FindAll returns paginates list of hotel grades.
func (s *HotelGradeService) FindAll(input *dto.PaginationInput) (*commons.PaginatedList, error) {

	return s.Repository.FindAll(input)
}

// Delete removes hotel type by given id.
func (s *HotelGradeService) Delete(id uint64) error {

	return s.Repository.Delete(id)
}
