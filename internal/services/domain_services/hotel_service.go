package domain_services

import (
	"context"
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
func (s *HotelService) Create(ctx context.Context, hotel *models.Hotel) (*models.Hotel, error) {

	return s.Repository.Create(ctx, hotel)
}

// Update updates Hotel.
func (s *HotelService) Update(ctx context.Context, hotel *models.Hotel) (*models.Hotel, error) {

	return s.Repository.Update(ctx, hotel)
}

// Find returns Hotel and if it does not find the Hotel, it returns nil.
func (s *HotelService) Find(ctx context.Context, id uint64) (*models.Hotel, error) {

	return s.Repository.Find(ctx, id)
}

// FindAll returns paginates list of hotels
func (s *HotelService) FindAll(ctx context.Context, filter *dto.PaginationFilter) (*commons.PaginatedResult, error) {

	return s.Repository.FindAll(ctx, filter)
}

// Delete removes hotel type by given id.
func (s *HotelService) Delete(ctx context.Context, id uint64) error {

	return s.Repository.Delete(ctx, id)
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
