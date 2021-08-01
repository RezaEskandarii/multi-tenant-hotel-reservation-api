package repositories

import (
	"errors"
	"gorm.io/gorm"
	"hotel-reservation/internal/commons"
	"hotel-reservation/internal/dto"
	"hotel-reservation/internal/message_keys"
	"hotel-reservation/internal/models"
)

type ResidenceRepository struct {
	DB *gorm.DB
}

func NewResidenceRepository(db *gorm.DB) *ResidenceRepository {

	return &ResidenceRepository{DB: db}
}

func (r *ResidenceRepository) Create(residence *models.Residence) (*models.Residence, error) {

	if tx := r.DB.Create(&residence); tx.Error != nil {

		return nil, tx.Error
	}

	return residence, nil
}

func (r *ResidenceRepository) Update(residence *models.Residence) (*models.Residence, error) {

	if tx := r.DB.Updates(&residence); tx.Error != nil {

		return nil, tx.Error
	}

	return residence, nil
}

func (r *ResidenceRepository) Find(id uint64) (*models.Residence, error) {
	model := models.Residence{}

	if tx := r.DB.Where("id=?", id).Preload("Grades").Find(&model); tx.Error != nil {

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
		return query.Error
	}

	return nil
}

func (r *ResidenceRepository) hasRepeatData(residence *models.Residence) error {
	var countByName int64 = 0

	if tx := *r.DB.Model(&models.Residence{}).Where(&models.Residence{Name: residence.Name}).Count(&countByName); tx.Error != nil {
		return tx.Error
	}

	if countByName > 0 {

		return errors.New(message_keys.ResidenceRepeatPostalCode)
	}
	return nil
}
