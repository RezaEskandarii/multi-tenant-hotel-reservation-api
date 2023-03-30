package repositories

import (
	"context"
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
	"reservation-api/pkg/multi_tenancy_database/tenant_database_resolver"
)

type RateGroupRepository struct {
	DbResolver *tenant_database_resolver.TenantDatabaseResolver
}

// NewRateGroupRepository returns new RateGroupRepository.
func NewRateGroupRepository(r *tenant_database_resolver.TenantDatabaseResolver) *RateGroupRepository {

	return &RateGroupRepository{DbResolver: r}
}

func (r *RateGroupRepository) Create(ctx context.Context, model *models.RateGroup) (*models.RateGroup, error) {

	db := r.DbResolver.GetTenantDB(ctx)

	if tx := db.Create(&model); tx.Error != nil {
		return nil, tx.Error
	}

	return model, nil
}

func (r *RateGroupRepository) Update(ctx context.Context, model *models.RateGroup) (*models.RateGroup, error) {

	db := r.DbResolver.GetTenantDB(ctx)

	if tx := db.Updates(&model); tx.Error != nil {
		return nil, tx.Error
	}

	return model, nil
}

func (r *RateGroupRepository) Find(ctx context.Context, id uint64) (*models.RateGroup, error) {

	model := models.RateGroup{}
	db := r.DbResolver.GetTenantDB(ctx)

	if tx := db.Where("id=?", id).Find(&model); tx.Error != nil {
		return nil, tx.Error
	}

	if model.Id == 0 {
		return nil, nil
	}

	return &model, nil
}

func (r *RateGroupRepository) FindAll(ctx context.Context, input *dto.PaginationFilter) (*commons.PaginatedResult, error) {

	db := r.DbResolver.GetTenantDB(ctx)
	return paginatedList(&models.RateGroup{}, db, input)
}

func (r RateGroupRepository) Delete(ctx context.Context, id uint64) error {

	db := r.DbResolver.GetTenantDB(ctx)

	if query := db.Model(&models.RateGroup{}).Where("id=?", id).Delete(&models.RateGroup{}); query.Error != nil {
		return query.Error
	}

	return nil
}
