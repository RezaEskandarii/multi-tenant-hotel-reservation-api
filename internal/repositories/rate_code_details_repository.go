package repositories

import (
	"context"
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
	"reservation-api/internal/tenant_resolver"
	"reservation-api/pkg/multi_tenancy_database/tenant_database_resolver"
)

type RateCodeDetailRepository struct {
	DbResolver *tenant_database_resolver.TenantDatabaseResolver
}

// NewRateCodeDetailRepository returns new RateCodeDetailRepository.
func NewRateCodeDetailRepository(r *tenant_database_resolver.TenantDatabaseResolver) *RateCodeDetailRepository {

	return &RateCodeDetailRepository{DbResolver: r}
}

func (r *RateCodeDetailRepository) Create(ctx context.Context, model *models.RateCodeDetail) (*models.RateCodeDetail, error) {

	db := r.DbResolver.GetTenantDB(tenant_resolver.GetCurrentTenant(ctx))

	if tx := db.Create(&model); tx.Error != nil {
		return nil, tx.Error
	}
	return model, nil
}

func (r *RateCodeDetailRepository) Update(ctx context.Context, model *models.RateCodeDetail) (*models.RateCodeDetail, error) {

	db := r.DbResolver.GetTenantDB(tenant_resolver.GetCurrentTenant(ctx))
	tx := db.Begin()

	// remove old price.
	if err := tx.Where("rate_code_detail_id=?", model.Id).Delete(&models.RateCodeDetailPrice{}).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	// update
	if err := tx.Updates(&model).Error; err != nil {
		return nil, err
	}

	tx.Commit()

	return model, nil
}

func (r *RateCodeDetailRepository) FindPrice(ctx context.Context, id uint64) (*models.RateCodeDetailPrice, error) {

	model := models.RateCodeDetailPrice{}
	db := r.DbResolver.GetTenantDB(tenant_resolver.GetCurrentTenant(ctx))

	if err := db.Model(models.RateCodeDetailPrice{}).Where("id=?", id).Find(&model).Error; err != nil {
		return nil, err
	}

	if model.Id == 0 {
		return nil, nil
	}

	return &model, nil
}

func (r *RateCodeDetailRepository) Find(ctx context.Context, id uint64) (*models.RateCodeDetail, error) {

	model := models.RateCodeDetail{}
	db := r.DbResolver.GetTenantDB(tenant_resolver.GetCurrentTenant(ctx))

	if tx := db.Where("id=?", id).Find(&model).Preload("RateCodeDetailPrice"); tx.Error != nil {
		return nil, tx.Error
	}

	if model.Id == 0 {
		return nil, nil
	}
	return &model, nil
}

func (r *RateCodeDetailRepository) FindAll(ctx context.Context, input *dto.PaginationFilter) (*commons.PaginatedResult, error) {
	db := r.DbResolver.GetTenantDB(tenant_resolver.GetCurrentTenant(ctx))
	return paginatedList(&models.RateCodeDetail{}, db, input)
}

func (r RateCodeDetailRepository) Delete(ctx context.Context, id uint64) error {

	db := r.DbResolver.GetTenantDB(tenant_resolver.GetCurrentTenant(ctx))

	if query := db.Model(&models.RateCodeDetail{}).Where("id=?", id).Delete(&models.RateCodeDetail{}); query.Error != nil {
		return query.Error
	}
	return nil
}
