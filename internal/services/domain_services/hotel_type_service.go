package domain_services

import (
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
func (s *HotelTypeService) Create(hotelType *models.HotelType, tenantID uint64) (*models.HotelType, error) {

	return s.Repository.Create(hotelType, tenantID)
}

// Update updates HotelType.
func (s *HotelTypeService) Update(hotelType *models.HotelType, tenantID uint64) (*models.HotelType, error) {

	return s.Repository.Update(hotelType, tenantID)
}

// Find returns HotelType and if it does not find the HotelType, it returns nil.
func (s *HotelTypeService) Find(id uint64, tenantID uint64) (*models.HotelType, error) {

	return s.Repository.Find(id, tenantID)
}

// FindAll returns paginates list of hotel types.
func (s *HotelTypeService) FindAll(input *dto.PaginationFilter) (*commons.PaginatedResult, error) {

	return s.Repository.FindAll(input)
}

// Delete removes hotel type by given id.
func (s *HotelTypeService) Delete(id uint64, tenantID uint64) error {

	return s.Repository.Delete(id, tenantID)
}
