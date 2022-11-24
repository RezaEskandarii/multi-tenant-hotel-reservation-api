package domain_services

import (
	"context"
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
	"reservation-api/internal/repositories"
)

type CurrencyService struct {
	Repository *repositories.CurrencyRepository
}

// NewCurrencyService returns new CurrencyService
func NewCurrencyService(r *repositories.CurrencyRepository) *CurrencyService {
	return &CurrencyService{Repository: r}
}

// Create creates new currency.
func (s *CurrencyService) Create(ctx context.Context, currency *models.Currency) (*models.Currency, error) {

	return s.Repository.Create(ctx, currency)
}

// Update updates currency.
func (s *CurrencyService) Update(ctx context.Context, currency *models.Currency) (*models.Currency, error) {

	return s.Repository.Update(ctx, currency)
}

// Find returns currency and if it does not find the currency, it returns nil.
func (s *CurrencyService) Find(ctx context.Context, id uint64) (*models.Currency, error) {

	return s.Repository.Find(ctx, id)
}

// FindAll returns paginates list of currencies
func (s *CurrencyService) FindAll(ctx context.Context, input *dto.PaginationFilter) (*commons.PaginatedResult, error) {

	return s.Repository.FindAll(ctx, input)
}

// FindBySymbol returns currency by alias name.
func (s *CurrencyService) FindBySymbol(ctx context.Context, symbol string) (*models.Currency, error) {

	return s.Repository.FindBySymbol(ctx, symbol)
}

// Seed
func (s *CurrencyService) Seed(ctx context.Context, jsonFilePath string) error {

	return s.Repository.Seed(ctx, jsonFilePath)
}
