package repositories

import (
	"gorm.io/gorm"
	"hotel-reservation/internal/commons"
	"hotel-reservation/internal/dto"
	"hotel-reservation/internal/models"
	"hotel-reservation/pkg/application_loger"
)

type CurrencyRepository struct {
	DB *gorm.DB
}

func NewCurrencyRepository(db *gorm.DB) *CurrencyRepository {
	return &CurrencyRepository{
		DB: db,
	}
}

func (r *CurrencyRepository) Create(currency *models.Currency) (*models.Currency, error) {

	if tx := r.DB.Create(&currency); tx.Error != nil {
		application_loger.LogError(tx.Error)
		return nil, tx.Error
	}

	return currency, nil
}

func (r *CurrencyRepository) Update(currency *models.Currency) (*models.Currency, error) {

	if tx := r.DB.Updates(&currency); tx.Error != nil {
		application_loger.LogError(tx.Error)
		return nil, tx.Error
	}

	return currency, nil
}

func (r *CurrencyRepository) Find(id uint64) (*models.Currency, error) {

	model := models.Currency{}
	if tx := r.DB.Where("id=?", id).Find(&model); tx.Error != nil {
		application_loger.LogError(tx.Error)
		return nil, tx.Error
	}

	if model.Id == 0 {
		return nil, nil
	}

	return &model, nil
}

func (r *CurrencyRepository) FindAll(input *dto.PaginationInput) (*commons.PaginatedList, error) {

	return finAll(&models.Currency{}, r.DB, input)
}

func (r *CurrencyRepository) FindBySymbol(symbol string) (*models.Currency, error) {
	model := models.Currency{}
	if tx := r.DB.Where("symbol=?", symbol).Find(&model); tx.Error != nil {
		application_loger.LogError(tx.Error)
		return nil, tx.Error
	}

	if model.Id == 0 {
		return nil, nil
	}

	return &model, nil
}
