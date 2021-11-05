package services

import (
	"reservation-api/internal/repositories"
)

type SeederService struct {
	Repository *repositories.SeederRepository
}

func NewSeederService() *SeederService {
	return &SeederService{}
}

// Seed gives the model and address of the json file from the input and
// fills given model's fields with given json and bulk inserts given model.
func (s *SeederService) Seed(jsonFilePath string, model []interface{}) error {

	return s.Repository.Seed(jsonFilePath, model)
}
