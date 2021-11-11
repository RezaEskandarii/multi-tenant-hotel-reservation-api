package repositories

import (
	"fmt"
	"gorm.io/gorm"
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
)

type AuditRepository struct {
	Completed chan bool
	Data      chan interface{}
	DB        *gorm.DB
}

func NewAuditRepository(Db *gorm.DB, completedCh chan bool, DataCh chan interface{}) *AuditRepository {
	return &AuditRepository{
		Completed: completedCh,
		Data:      DataCh,
		DB:        Db,
	}
}

func (r *AuditRepository) Create(model *models.Audit) (*models.Audit, error) {
	model.Data = fmt.Sprintf("%v", model.DataChannel)
	if err := r.DB.Save(&model).Error; err != nil {
		return nil, err
	}
	return model, nil
}

func (r *AuditRepository) FindAll(input *dto.PaginationInput) (*commons.PaginatedList, error) {

	return finAll(&models.City{}, r.DB, input)
}
