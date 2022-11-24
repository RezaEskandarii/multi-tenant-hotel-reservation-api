package domain_services

import (
	"context"
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
func (s *HotelGradeService) Create(ctx context.Context, hotelGrade *models.HotelGrade) (*models.HotelGrade, error) {

	return s.Repository.Create(ctx, hotelGrade)
}

// Update updates HotelGrade.
func (s *HotelGradeService) Update(ctx context.Context, hotelGrade *models.HotelGrade) (*models.HotelGrade, error) {

	return s.Repository.Update(ctx, hotelGrade)
}

// Find returns HotelGrade and if it does not find the HotelGrade, it returns nil.
func (s *HotelGradeService) Find(ctx context.Context, id uint64) (*models.HotelGrade, error) {

	return s.Repository.Find(ctx, id)
}

// FindAll returns paginates list of hotel grades.
func (s *HotelGradeService) FindAll(ctx context.Context, filter *dto.PaginationFilter) (*commons.PaginatedResult, error) {

	return s.Repository.FindAll(ctx, filter)
}

// Delete removes hotel type by given id.
func (s *HotelGradeService) Delete(ctx context.Context, id uint64) error {

	return s.Repository.Delete(ctx, id)
}
