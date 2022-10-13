package repositories

import (
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
	"reservation-api/pkg/database/connection_resolver"
)

type RateGroupRepository struct {
	ConnectionResolver *connection_resolver.TenantConnectionResolver
}

// NewRateGroupRepository returns new RateGroupRepository.
func NewRateGroupRepository(r *connection_resolver.TenantConnectionResolver) *RateGroupRepository {

	return &RateGroupRepository{ConnectionResolver: r}
}

func (r *RateGroupRepository) Create(model *models.RateGroup, tenantID uint64) (*models.RateGroup, error) {

	db := r.ConnectionResolver.GetDB(tenantID)

	if tx := db.Create(&model); tx.Error != nil {
		return nil, tx.Error
	}

	return model, nil
}

func (r *RateGroupRepository) Update(model *models.RateGroup, tenantID uint64) (*models.RateGroup, error) {

	db := r.ConnectionResolver.GetDB(tenantID)

	if tx := db.Updates(&model); tx.Error != nil {
		return nil, tx.Error
	}

	return model, nil
}

func (r *RateGroupRepository) Find(id uint64, tenantID uint64) (*models.RateGroup, error) {

	model := models.RateGroup{}
	db := r.ConnectionResolver.GetDB(tenantID)

	if tx := db.Where("id=?", id).Find(&model); tx.Error != nil {
		return nil, tx.Error
	}

	if model.Id == 0 {
		return nil, nil
	}

	return &model, nil
}

func (r *RateGroupRepository) FindAll(input *dto.PaginationFilter) (*commons.PaginatedResult, error) {

	db := r.ConnectionResolver.GetDB(input.TenantID)
	return paginatedList(&models.RateGroup{}, db, input)
}

func (r RateGroupRepository) Delete(id uint64, tenantID uint64) error {

	db := r.ConnectionResolver.GetDB(tenantID)

	if query := db.Model(&models.RateGroup{}).Where("id=?", id).Delete(&models.RateGroup{}); query.Error != nil {
		return query.Error
	}

	return nil
}
