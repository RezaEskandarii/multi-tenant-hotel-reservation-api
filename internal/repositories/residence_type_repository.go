package repositories

import (
	"gorm.io/gorm"
	"hotel-reservation/internal/commons"
	"hotel-reservation/internal/dto"
	"hotel-reservation/internal/models"
)

type ResidenceTypeRepository struct {
	DB *gorm.DB
}

func NewResidenceTypeRepository(db *gorm.DB) *ResidenceTypeRepository {
	return &ResidenceTypeRepository{DB: db}
}

func (r *ResidenceTypeRepository) Create(residenceType *models.ResidenceType) (*models.ResidenceType, error) {

	if tx := r.DB.Create(&residenceType); tx.Error != nil {
		return nil, tx.Error
	}

	return residenceType, nil
}

func (r *ResidenceTypeRepository) Update(residenceType *models.ResidenceType) (*models.ResidenceType, error) {

	if tx := r.DB.Updates(&residenceType); tx.Error != nil {
		return nil, tx.Error
	}

	return residenceType, nil
}

func (r *ResidenceTypeRepository) Find(id uint64) (*models.ResidenceType, error) {

	model := models.ResidenceType{}

	if tx := r.DB.Where("id=?", id).Preload("Grades").Find(&model); tx.Error != nil {
		return nil, tx.Error
	}

	if model.Id == 0 {
		return nil, nil
	}

	return &model, nil
}

func (r *ResidenceTypeRepository) FindAll(input *dto.PaginationInput) (*commons.PaginatedList, error) {

	return finAll(&models.ResidenceType{}, r.DB, input)
}

func (r ResidenceTypeRepository) Delete(id uint64) error {

	var count int64 = 0

	if query := r.DB.Model(&models.Residence{}).Where(&models.Residence{ResidenceTypeId: id}).Count(&count); query.Error != nil {

		return query.Error
	}

	if count > 0 {
		return TypeHasResidenceError
	}

	if query := r.DB.Model(&models.ResidenceType{}).Where("id=?", id).Delete(&models.ResidenceType{}); query.Error != nil {

		return query.Error
	} else {

		query = r.DB.Model(&models.ResidenceGrade{}).Where(&models.ResidenceGrade{ResidenceTypeId: id}).Delete(&models.ResidenceGrade{})

		if query.Error != nil {
			return query.Error
		}
	}

	return nil
}
