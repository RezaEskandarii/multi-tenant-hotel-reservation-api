package repositories

import (
	"gorm.io/gorm"
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
)

type RateGroupRepository struct {
	DB *gorm.DB
}

// NewRateGroupRepository returns new RateGroupRepository.
func NewRateGroupRepository(db *gorm.DB) *RateGroupRepository {

	return &RateGroupRepository{DB: db}
}

func (r *RateGroupRepository) Create(model *models.RateGroup) (*models.RateGroup, error) {

	if tx := r.DB.Create(&model); tx.Error != nil {
		return nil, tx.Error
	}

	return model, nil
}

func (r *RateGroupRepository) Update(model *models.RateGroup) (*models.RateGroup, error) {

	if tx := r.DB.Updates(&model); tx.Error != nil {
		return nil, tx.Error
	}

	return model, nil
}

func (r *RateGroupRepository) Find(id uint64) (*models.RateGroup, error) {

	model := models.RateGroup{}

	if tx := r.DB.Where("id=?", id).Find(&model); tx.Error != nil {
		return nil, tx.Error
	}

	if model.Id == 0 {
		return nil, nil
	}

	return &model, nil
}

func (r *RateGroupRepository) FindAll(input *dto.PaginationInput) (*commons.PaginatedList, error) {

	return finAll(&models.RateGroup{}, r.DB, input)
}

func (r RateGroupRepository) Delete(id uint64) error {

	if query := r.DB.Model(&models.RateGroup{}).Where("id=?", id).Delete(&models.RateGroup{}); query.Error != nil {

		return query.Error
	}

	return nil
}
