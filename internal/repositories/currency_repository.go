package repositories

import (
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
	"reservation-api/internal/utils"
	"reservation-api/pkg/database/connection_resolver"
)

type CurrencyRepository struct {
	ConnectionResolver *connection_resolver.ConnectionResolver
}

func NewCurrencyRepository(r *connection_resolver.ConnectionResolver) *CurrencyRepository {
	return &CurrencyRepository{
		ConnectionResolver: r,
	}
}

func (r *CurrencyRepository) Create(currency *models.Currency) (*models.Currency, error) {

	db := r.ConnectionResolver.GetDB(currency.TenantId)
	if tx := db.Create(&currency); tx.Error != nil {

		return nil, tx.Error
	}

	return currency, nil
}

func (r *CurrencyRepository) Update(currency *models.Currency) (*models.Currency, error) {

	db := r.ConnectionResolver.GetDB(currency.TenantId)
	if tx := db.Updates(&currency); tx.Error != nil {
		return nil, tx.Error
	}

	return currency, nil
}

func (r *CurrencyRepository) Find(id uint64, tenantID uint64) (*models.Currency, error) {

	model := models.Currency{}
	db := r.ConnectionResolver.GetDB(tenantID)
	if tx := db.Where("id=?", id).Find(&model); tx.Error != nil {

		return nil, tx.Error
	}

	if model.Id == 0 {
		return nil, nil
	}

	return &model, nil
}

func (r *CurrencyRepository) FindAll(input *dto.PaginationFilter) (*commons.PaginatedResult, error) {
	db := r.ConnectionResolver.GetDB(input.TenantID)
	return paginatedList(&models.Currency{}, db, input)
}

func (r *CurrencyRepository) FindBySymbol(symbol string, tenantID uint64) (*models.Currency, error) {

	db := r.ConnectionResolver.GetDB(tenantID)
	model := models.Currency{}

	if tx := db.Where("symbol=?", symbol).Find(&model); tx.Error != nil {
		return nil, tx.Error
	}

	if model.Id == 0 {
		return nil, nil
	}

	return &model, nil
}

func (r *CurrencyRepository) Seed(jsonFilePath string, tenantID uint64) error {

	db := r.ConnectionResolver.GetDB(tenantID)
	currencies := make([]models.Currency, 0)

	if err := utils.CastJsonFileToStruct(jsonFilePath, &currencies); err == nil {
		for _, currency := range currencies {
			var count int64 = 0
			if err := db.Model(models.Currency{}).Where("symbol", currency.Symbol).Count(&count).Error; err != nil {
				return err
			} else {
				if count == 0 {
					if err := db.Create(&currency).Error; err != nil {
						return err
					}
				}
			}
		}

	} else {

		return err
	}
	return nil
}
