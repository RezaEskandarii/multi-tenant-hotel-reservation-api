package repositories

import (
	"gorm.io/gorm"
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
)

type RateCodeDetailRepository struct {
	DB *gorm.DB
}

// NewRateCodeDetailRepository returns new RateCodeDetailRepository.
func NewRateCodeDetailRepository(db *gorm.DB) *RateCodeDetailRepository {

	return &RateCodeDetailRepository{DB: db}
}

func (r *RateCodeDetailRepository) Create(model *models.RateCodeDetail) (*models.RateCodeDetail, error) {

	if tx := r.DB.Create(&model); tx.Error != nil {
		return nil, tx.Error
	}
	return model, nil
}

func (r *RateCodeDetailRepository) Update(model *models.RateCodeDetail) (*models.RateCodeDetail, error) {

	tx := r.DB.Begin()

	if err := tx.Model(&models.RateCodeDetailPrice{}).Where("rate_code_detail_id=?", model.Id).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Updates(&model).Error; err != nil {
		return nil, err
	}

	tx.Commit()

	return model, nil
}

func (r *RateCodeDetailRepository) Find(id uint64) (*models.RateCodeDetail, error) {

	model := models.RateCodeDetail{}

	if tx := r.DB.Where("id=?", id).Find(&model); tx.Error != nil {
		return nil, tx.Error
	}

	if model.Id == 0 {
		return nil, nil
	}
	return &model, nil
}

func (r *RateCodeDetailRepository) FindAll(input *dto.PaginationInput) (*commons.PaginatedList, error) {

	return finAll(&models.RateCodeDetail{}, r.DB, input)
}

func (r RateCodeDetailRepository) Delete(id uint64) error {

	if query := r.DB.Model(&models.RateCodeDetail{}).Where("id=?", id).Delete(&models.RateCodeDetail{}); query.Error != nil {
		return query.Error
	}
	return nil
}
