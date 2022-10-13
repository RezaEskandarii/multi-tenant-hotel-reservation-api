package repositories

import (
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
	"reservation-api/pkg/database/connection_resolver"
)

type RateCodeRepository struct {
	ConnectionResolver *connection_resolver.TenantConnectionResolver
}

// NewRateCodeRepository returns new RateCodeRepository.
func NewRateCodeRepository(r *connection_resolver.TenantConnectionResolver) *RateCodeRepository {

	return &RateCodeRepository{ConnectionResolver: r}
}

func (r *RateCodeRepository) Create(model *models.RateCode, tenantID uint64) (*models.RateCode, error) {

	db := r.ConnectionResolver.GetDB(tenantID)

	if tx := db.Create(&model); tx.Error != nil {
		return nil, tx.Error
	}
	return model, nil
}

func (r *RateCodeRepository) Update(model *models.RateCode, tenantID uint64) (*models.RateCode, error) {

	db := r.ConnectionResolver.GetDB(tenantID)

	if tx := db.Updates(&model); tx.Error != nil {
		return nil, tx.Error
	}
	return model, nil
}

func (r *RateCodeRepository) Find(id uint64, tenantID uint64) (*models.RateCode, error) {

	model := models.RateCode{}
	db := r.ConnectionResolver.GetDB(tenantID)

	if tx := db.Where("id=?", id).Find(&model); tx.Error != nil {
		return nil, tx.Error
	}

	if model.Id == 0 {
		return nil, nil
	}
	return &model, nil
}

func (r *RateCodeRepository) FindAll(input *dto.PaginationFilter) (*commons.PaginatedResult, error) {

	db := r.ConnectionResolver.GetDB(input.TenantID)
	return paginatedList(&models.RateCode{}, db, input)
}

func (r RateCodeRepository) Delete(id uint64, tenantID uint64) error {

	db := r.ConnectionResolver.GetDB(tenantID)

	if query := db.Model(&models.RateCode{}).Where("id=?", id).Delete(&models.RateCode{}); query.Error != nil {
		return query.Error
	}
	return nil
}
