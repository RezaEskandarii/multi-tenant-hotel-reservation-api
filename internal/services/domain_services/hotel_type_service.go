package domain_services

import (
	"context"
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
	"reservation-api/internal/repositories"
)

type HotelTypeService struct {
	Repository *repositories.HotelTypeRepository
}

// NewHotelTypeService returns new HotelTypeService
func NewHotelTypeService(r *repositories.HotelTypeRepository) *HotelTypeService {
	return &HotelTypeService{Repository: r}
}

// Create creates new HotelType.
func (s *HotelTypeService) Create(ctx context.Context, hotelType *models.HotelType) (*models.HotelType, error) {

	return s.Repository.Create(ctx, hotelType)
}

// Update updates HotelType.
func (s *HotelTypeService) Update(ctx context.Context, hotelType *models.HotelType) (*models.HotelType, error) {

	return s.Repository.Update(ctx, hotelType)
}

// Find returns HotelType and if it does not find the HotelType, it returns nil.
func (s *HotelTypeService) Find(ctx context.Context, id uint64) (*models.HotelType, error) {

	return s.Repository.Find(ctx, id)
}

// FindAll returns paginates list of hotel types.
func (s *HotelTypeService) FindAll(ctx context.Context, filter *dto.PaginationFilter) (*commons.PaginatedResult, error) {

	return s.Repository.FindAll(ctx, filter)
}

// Delete removes hotel type by given id.
func (s *HotelTypeService) Delete(ctx context.Context, id uint64) error {

	return s.Repository.Delete(ctx, id)
}
