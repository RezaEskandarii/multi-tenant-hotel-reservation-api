package repositories

import (
	"gorm.io/gorm"
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
)

type AuditRepository struct {
	DB *gorm.DB
}

func NewAuditRepository(Db *gorm.DB) *AuditRepository {
	return &AuditRepository{
		DB: Db,
	}
}

func (r *AuditRepository) Create(model *models.Audit) (*models.Audit, error) {
	if err := r.DB.Create(&model).Error; err != nil {
		return nil, err
	}
	return model, nil
}

func (r *AuditRepository) FindAll(input *dto.PaginationFilter) (*commons.PaginatedResult, error) {

	return paginatedList(&models.City{}, r.DB, input)
}
