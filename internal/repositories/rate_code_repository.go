package repositories

import (
	"context"
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
	"reservation-api/pkg/multi_tenancy_database/tenant_database_resolver"
)

type RateCodeRepository struct {
	DbResolver *tenant_database_resolver.TenantDatabaseResolver
}

// NewRateCodeRepository returns new RateCodeRepository.
func NewRateCodeRepository(r *tenant_database_resolver.TenantDatabaseResolver) *RateCodeRepository {

	return &RateCodeRepository{DbResolver: r}
}

func (r *RateCodeRepository) Create(ctx context.Context, model *models.RateCode) (*models.RateCode, error) {

	db := r.DbResolver.GetTenantDB(ctx)

	if tx := db.Create(&model); tx.Error != nil {
		return nil, tx.Error
	}
	return model, nil
}

func (r *RateCodeRepository) Update(ctx context.Context, model *models.RateCode) (*models.RateCode, error) {

	db := r.DbResolver.GetTenantDB(ctx)

	if tx := db.Updates(&model); tx.Error != nil {
		return nil, tx.Error
	}
	return model, nil
}

func (r *RateCodeRepository) Find(ctx context.Context, id uint64) (*models.RateCode, error) {

	model := models.RateCode{}
	db := r.DbResolver.GetTenantDB(ctx)

	if tx := db.Where("id=?", id).Find(&model); tx.Error != nil {
		return nil, tx.Error
	}

	if model.Id == 0 {
		return nil, nil
	}
	return &model, nil
}

func (r *RateCodeRepository) FindAll(ctx context.Context, input *dto.PaginationFilter) (*commons.PaginatedResult, error) {

	db := r.DbResolver.GetTenantDB(ctx)
	return paginatedList(&models.RateCode{}, db, input)
}

func (r RateCodeRepository) Delete(ctx context.Context, id uint64) error {

	db := r.DbResolver.GetTenantDB(ctx)

	if query := db.Model(&models.RateCode{}).Where("id=?", id).Delete(&models.RateCode{}); query.Error != nil {
		return query.Error
	}
	return nil
}
