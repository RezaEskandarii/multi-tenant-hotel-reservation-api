package services

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
func NewHotelTypeService() *HotelTypeService {
	return &HotelTypeService{}
}

// Create creates new HotelType.
func (s *HotelTypeService) Create(hotelType *models.HotelType) (*models.HotelType, error) {

	return s.Repository.Create(hotelType)
}

// Update updates HotelType.
func (s *HotelTypeService) Update(hotelType *models.HotelType) (*models.HotelType, error) {

	return s.Repository.Update(hotelType)
}

// Find returns HotelType and if it does not find the HotelType, it returns nil.
func (s *HotelTypeService) Find(id uint64) (*models.HotelType, error) {

	return s.Repository.Find(id)
}

// FindAll returns paginates list of hotel types.
func (s *HotelTypeService) FindAll(input *dto.PaginationInput) (*commons.PaginatedList, error) {

	return s.Repository.FindAll(input)
}

// Delete removes hotel type by given id.
func (s *HotelTypeService) Delete(id uint64) error {

	return s.Repository.Delete(id)
}
