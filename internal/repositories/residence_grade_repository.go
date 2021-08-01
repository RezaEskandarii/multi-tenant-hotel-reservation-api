package repositories

import (
	"gorm.io/gorm"
	"hotel-reservation/internal/commons"
	"hotel-reservation/internal/dto"
	"hotel-reservation/internal/models"
)

type ResidenceGradeRepository struct {
	DB *gorm.DB
}

func NewResidenceGradeRepository(db *gorm.DB) *ResidenceGradeRepository {
	return &ResidenceGradeRepository{DB: db}
}

func (r *ResidenceGradeRepository) Create(residenceGrade *models.ResidenceGrade) (*models.ResidenceGrade, error) {

	if tx := r.DB.Create(&residenceGrade); tx.Error != nil {

		return nil, tx.Error
	}

	return residenceGrade, nil
}

func (r *ResidenceGradeRepository) Update(residenceGrade *models.ResidenceGrade) (*models.ResidenceGrade, error) {

	if tx := r.DB.Updates(&residenceGrade); tx.Error != nil {

		return nil, tx.Error
	}

	return residenceGrade, nil
}

func (r *ResidenceGradeRepository) Find(id uint64) (*models.ResidenceGrade, error) {

	model := models.ResidenceGrade{}

	if tx := r.DB.Where("id=?", id).Preload("ResidenceType").Find(&model); tx.Error != nil {

		return nil, tx.Error
	}

	if model.Id == 0 {
		return nil, nil
	}

	return &model, nil
}

func (r *ResidenceGradeRepository) FindAll(input *dto.PaginationInput) (*commons.PaginatedList, error) {

	return finAll(&models.ResidenceGrade{}, r.DB, input)
}

func (r ResidenceGradeRepository) Delete(id uint64) error {

	var count int64 = 0

	if query := r.DB.Model(&models.Residence{}).Where(&models.Residence{ResidenceGradeId: id}).Count(&count); query.Error != nil {

		return query.Error
	}

	if count > 0 {
		return GradeHasResidence
	}

	if query := r.DB.Model(&models.ResidenceGrade{}).Where("id=?", id).Delete(&models.ResidenceGrade{}); query.Error != nil {

		return query.Error
	}

	return nil
}
