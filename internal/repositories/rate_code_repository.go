package repositories

import (
	"gorm.io/gorm"
	"hotel-reservation/internal/commons"
	"hotel-reservation/internal/dto"
	"hotel-reservation/internal/models"
)

type RateCodeRepository struct {
	DB *gorm.DB
}

// NewRateCodeRepository returns new RateCodeRepository.
func NewRateCodeRepository(db *gorm.DB) *RateCodeRepository {

	return &RateCodeRepository{DB: db}
}

func (r *RateCodeRepository) Create(model *models.RateCode) (*models.RateCode, error) {

	if tx := r.DB.Create(&model); tx.Error != nil {
		return nil, tx.Error
	}
	return model, nil
}

func (r *RateCodeRepository) Update(model *models.RateCode) (*models.RateCode, error) {

	if tx := r.DB.Updates(&model); tx.Error != nil {
		return nil, tx.Error
	}
	return model, nil
}

func (r *RateCodeRepository) Find(id uint64) (*models.RateCode, error) {

	model := models.RateCode{}

	if tx := r.DB.Where("id=?", id).Find(&model); tx.Error != nil {
		return nil, tx.Error
	}

	if model.Id == 0 {
		return nil, nil
	}
	return &model, nil
}

func (r *RateCodeRepository) FindAll(input *dto.PaginationInput) (*commons.PaginatedList, error) {

	return finAll(&models.RateCode{}, r.DB, input)
}

func (r RateCodeRepository) Delete(id uint64) error {

	if query := r.DB.Model(&models.RateCode{}).Where("id=?", id).Delete(&models.RateCode{}); query.Error != nil {

		return query.Error
	}
	return nil
}
