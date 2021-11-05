package services

import (
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
	"reservation-api/internal/repositories"
)

type ResidenceService struct {
	Repository *repositories.ResidenceRepository
}

// NewResidenceService returns new ResidenceService
func NewResidenceService() *ResidenceService {

	return &ResidenceService{}
}

// Create creates new Residence.
func (s *ResidenceService) Create(residence *models.Residence) (*models.Residence, error) {

	return s.Repository.Create(residence)
}

// Update updates Residence.
func (s *ResidenceService) Update(residence *models.Residence) (*models.Residence, error) {

	return s.Repository.Update(residence)
}

// Find returns Residence and if it does not find the Residence, it returns nil.
func (s *ResidenceService) Find(id uint64) (*models.Residence, error) {

	return s.Repository.Find(id)
}

// FindAll returns paginates list of residences
func (s *ResidenceService) FindAll(input *dto.PaginationInput) (*commons.PaginatedList, error) {

	return s.Repository.FindAll(input)
}

// Delete removes residence type by given id.
func (s *ResidenceService) Delete(id uint64) error {

	return s.Repository.Delete(id)
}

func (s ResidenceService) Map(givenModel *models.Residence, returnModel *models.Residence) *models.Residence {

	returnModel.Name = givenModel.Name
	returnModel.ResidenceTypeId = givenModel.ResidenceTypeId
	returnModel.Address = givenModel.Address
	returnModel.ResidenceGradeId = givenModel.ResidenceGradeId
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
