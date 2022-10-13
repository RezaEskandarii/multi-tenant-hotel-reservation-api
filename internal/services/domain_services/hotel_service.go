package domain_services

import (
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
	"reservation-api/internal/repositories"
)

type HotelService struct {
	Repository *repositories.HotelRepository
}

// NewHotelService returns new HotelService
func NewHotelService(r *repositories.HotelRepository) *HotelService {

	return &HotelService{Repository: r}
}

// Create creates new Hotel.
func (s *HotelService) Create(hotel *models.Hotel, tenantID uint64) (*models.Hotel, error) {

	return s.Repository.Create(hotel, tenantID)
}

// Update updates Hotel.
func (s *HotelService) Update(hotel *models.Hotel, tenantID uint64) (*models.Hotel, error) {

	return s.Repository.Update(hotel, tenantID)
}

// Find returns Hotel and if it does not find the Hotel, it returns nil.
func (s *HotelService) Find(id uint64, tenantID uint64) (*models.Hotel, error) {

	return s.Repository.Find(id, tenantID)
}

// FindAll returns paginates list of hotels
func (s *HotelService) FindAll(input *dto.PaginationFilter) (*commons.PaginatedResult, error) {

	return s.Repository.FindAll(input)
}

// Delete removes hotel type by given id.
func (s *HotelService) Delete(id uint64, tenantID uint64) error {

	return s.Repository.Delete(id, tenantID)
}

func (s HotelService) Map(givenModel *models.Hotel, returnModel *models.Hotel) *models.Hotel {

	returnModel.Name = givenModel.Name
	returnModel.HotelTypeId = givenModel.HotelTypeId
	returnModel.Address = givenModel.Address
	returnModel.HotelGradeId = givenModel.HotelGradeId
	returnModel.Description = givenModel.Description
	returnModel.ProvinceId = givenModel.ProvinceId
	returnModel.CityId = givenModel.CityId
	returnModel.EmailAddress = givenModel.EmailAddress
	returnModel.FaxNumber = givenModel.FaxNumber
	returnModel.Latitude = givenModel.Latitude
	returnModel.Longitude = givenModel.Longitude
	returnModel.OwnerId = givenModel.OwnerId
	returnModel.PhoneNumber1 = givenModel.PhoneNumber1
	returnModel.PhoneNumber2 = givenModel.PhoneNumber2

	return returnModel
}
