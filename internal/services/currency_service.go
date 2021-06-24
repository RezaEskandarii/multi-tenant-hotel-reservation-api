package services

import (
	"hotel-reservation/internal/commons"
	"hotel-reservation/internal/dto"
	"hotel-reservation/internal/models"
	"hotel-reservation/internal/repositories"
)

type CurrencyService struct {
	Repository *repositories.CurrencyRepository
}

func NewCurrencyService() *CurrencyService {
	return &CurrencyService{}
}

func (s *CurrencyService) Create(currency *models.Currency) (*models.Currency, error) {

	return s.Repository.Create(currency)
}

func (s *CurrencyService) Update(currency *models.Currency) (*models.Currency, error) {

	return s.Repository.Update(currency)
}

func (s *CurrencyService) Find(id uint64) (*models.Currency, error) {

	return s.Repository.Find(id)
}

func (s *CurrencyService) FindAll(input *dto.PaginationInput) (*commons.PaginatedList, error) {

	return s.Repository.FindAll(input)
}

func (s *CurrencyService) FindBySymbol(symbol string) (*models.Currency, error) {

	return s.Repository.FindBySymbol(symbol)
}
