package repositories

import (
	"context"
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
	"reservation-api/internal/utils"
	"reservation-api/pkg/multi_tenancy_database/tenant_database_resolver"
)

type CurrencyRepository struct {
	DbResolver *tenant_database_resolver.TenantDatabaseResolver
}

func NewCurrencyRepository(r *tenant_database_resolver.TenantDatabaseResolver) *CurrencyRepository {
	return &CurrencyRepository{
		DbResolver: r,
	}
}

func (r *CurrencyRepository) Create(ctx context.Context, currency *models.Currency) (*models.Currency, error) {

	db := r.DbResolver.GetTenantDB(ctx)
	if tx := db.Create(&currency); tx.Error != nil {
		return nil, tx.Error
	}

	return currency, nil
}

func (r *CurrencyRepository) Update(ctx context.Context, currency *models.Currency) (*models.Currency, error) {

	db := r.DbResolver.GetTenantDB(ctx)
	if tx := db.Updates(&currency); tx.Error != nil {
		return nil, tx.Error
	}

	return currency, nil
}

func (r *CurrencyRepository) Find(ctx context.Context, id uint64) (*models.Currency, error) {

	model := models.Currency{}
	db := r.DbResolver.GetTenantDB(ctx)

	if tx := db.Where("id=?", id).Find(&model); tx.Error != nil {

		return nil, tx.Error
	}
	if model.Id == 0 {
		return nil, nil
	}

	return &model, nil
}

func (r *CurrencyRepository) FindAll(ctx context.Context, input *dto.PaginationFilter) (*commons.PaginatedResult, error) {
	db := r.DbResolver.GetTenantDB(ctx)
	return paginatedList(&models.Currency{}, db, input)
}

func (r *CurrencyRepository) FindBySymbol(ctx context.Context, symbol string) (*models.Currency, error) {

	db := r.DbResolver.GetTenantDB(ctx)
	model := models.Currency{}

	if tx := db.Where("symbol=?", symbol).Find(&model); tx.Error != nil {
		return nil, tx.Error
	}
	if model.Id == 0 {
		return nil, nil
	}

	return &model, nil
}

func (r *CurrencyRepository) Seed(ctx context.Context, jsonFilePath string) error {
	// Get a handle to the tenant database
	db := r.DbResolver.GetTenantDB(ctx)

	// Read the JSON file and convert its contents to a slice of Currency structs
	currencies := make([]models.Currency, 0)
	if err := utils.CastJsonFileToStruct(jsonFilePath, &currencies); err != nil {
		return err
	}

	// Iterate over each currency and check if it already exists in the database
	for _, currency := range currencies {
		var count int64
		if err := db.Model(models.Currency{}).Where("symbol", currency.Symbol).Count(&count).Error; err != nil {
			return err
		}

		// If the currency does not exist in the database, create a new record for it
		if count == 0 {
			if err := db.Create(&currency).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
