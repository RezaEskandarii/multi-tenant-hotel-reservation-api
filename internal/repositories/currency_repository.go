package repositories

import (
	"gorm.io/gorm"
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
	"reservation-api/internal/utils"
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

		return nil, tx.Error
	}

	return currency, nil
}

func (r *CurrencyRepository) Update(currency *models.Currency) (*models.Currency, error) {

	if tx := r.DB.Updates(&currency); tx.Error != nil {

		return nil, tx.Error
	}

	return currency, nil
}

func (r *CurrencyRepository) Find(id uint64) (*models.Currency, error) {

	model := models.Currency{}
	if tx := r.DB.Where("id=?", id).Find(&model); tx.Error != nil {

		return nil, tx.Error
	}

	if model.Id == 0 {
		return nil, nil
	}

	return &model, nil
}

func (r *CurrencyRepository) FindAll(input *dto.PaginationFilter) (*commons.PaginatedList, error) {

	return paginatedList(&models.Currency{}, r.DB, input)
}

func (r *CurrencyRepository) FindBySymbol(symbol string) (*models.Currency, error) {
	model := models.Currency{}

	if tx := r.DB.Where("symbol=?", symbol).Find(&model); tx.Error != nil {

		return nil, tx.Error
	}

	if model.Id == 0 {
		return nil, nil
	}

	return &model, nil
}

func (r *CurrencyRepository) Seed(jsonFilePath string) error {

	currencies := make([]models.Currency, 0)
	if err := utils.CastJsonFileToStruct(jsonFilePath, &currencies); err == nil {
		for _, currency := range currencies {
			var count int64 = 0
			if err := r.DB.Model(models.Currency{}).Where("symbol", currency.Symbol).Count(&count).Error; err != nil {
				return err
			} else {
				if count == 0 {
					if err := r.DB.Create(&currency).Error; err != nil {
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
