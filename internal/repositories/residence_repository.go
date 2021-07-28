package repositories

import (
	"gorm.io/gorm"
	"hotel-reservation/internal/commons"
	"hotel-reservation/internal/dto"
	"hotel-reservation/internal/models"
	"hotel-reservation/pkg/application_loger"
)

type ResidenceRepository struct {
	DB *gorm.DB
}

func NewResidenceRepository(db *gorm.DB) *ResidenceRepository {

	return &ResidenceRepository{DB: db}
}

func (r *ResidenceRepository) Create(residence *models.Residence) (*models.Residence, error) {

	if tx := r.DB.Create(&residence); tx.Error != nil {
		application_loger.LogError(tx.Error)
		return nil, tx.Error
	}

	return residence, nil
}

func (r *ResidenceRepository) Update(residence *models.Residence) (*models.Residence, error) {

	if tx := r.DB.Updates(&residence); tx.Error != nil {
		application_loger.LogError(tx.Error)
		return nil, tx.Error
	}

	return residence, nil
}

func (r *ResidenceRepository) Find(id uint64) (*models.Residence, error) {
	model := models.Residence{}

	if tx := r.DB.Where("id=?", id).Preload("Grades").Find(&model); tx.Error != nil {
		application_loger.LogError(tx.Error)
		return nil, tx.Error
	}

	if model.Id == 0 {
		return nil, nil
	}

	return &model, nil
}

func (r *ResidenceRepository) FindAll(input *dto.PaginationInput) (*commons.PaginatedList, error) {

	return finAll(&models.Residence{}, r.DB, input)
}

func (r ResidenceRepository) Delete(id uint64) error {
	if query := r.DB.Model(&models.Residence{}).Where("id=?", id).Delete(&models.Residence{}); query.Error != nil {

		application_loger.LogError(query.Error)
		return query.Error
	}

	return nil
}
