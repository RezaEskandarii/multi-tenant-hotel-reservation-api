package repositories

import (
	"context"
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
	"reservation-api/internal/tenant_resolver"
	"reservation-api/pkg/multi_tenancy_database/tenant_database_resolver"
)

type ProvinceRepository struct {
	ConnectionResolver *tenant_database_resolver.TenantDatabaseResolver
}

func NewProvinceRepository(r *tenant_database_resolver.TenantDatabaseResolver) *ProvinceRepository {
	return &ProvinceRepository{
		ConnectionResolver: r,
	}
}

func (r *ProvinceRepository) Create(ctx context.Context, province *models.Province) (*models.Province, error) {

	db := r.ConnectionResolver.GetTenantDB(tenant_resolver.GetCurrentTenant(ctx))

	if tx := db.Create(&province); tx.Error != nil {
		return nil, tx.Error
	}

	return province, nil
}

func (r *ProvinceRepository) Update(ctx context.Context, province *models.Province) (*models.Province, error) {

	db := r.ConnectionResolver.GetTenantDB(tenant_resolver.GetCurrentTenant(ctx))

	if tx := db.Updates(&province); tx.Error != nil {
		return nil, tx.Error
	}

	return province, nil
}

func (r *ProvinceRepository) Find(ctx context.Context, id uint64) (*models.Province, error) {

	model := models.Province{}
	db := r.ConnectionResolver.GetTenantDB(tenant_resolver.GetCurrentTenant(ctx))

	if tx := db.Where("id=?", id).Find(&model); tx.Error != nil {
		return nil, tx.Error
	}

	if model.Id == 0 {
		return nil, nil
	}

	return &model, nil
}

func (r *ProvinceRepository) FindAll(ctx context.Context, input *dto.PaginationFilter) (*commons.PaginatedResult, error) {

	db := r.ConnectionResolver.GetTenantDB(tenant_resolver.GetCurrentTenant(ctx))
	return paginatedList(&models.Province{}, db, input)
}

func (r *ProvinceRepository) GetCities(ctx context.Context, provinceId uint64) ([]*models.City, error) {

	var result []*models.City
	db := r.ConnectionResolver.GetTenantDB(tenant_resolver.GetCurrentTenant(ctx))

	query := db.Model(&models.City{}).
		Where("province_id=?", provinceId).Find(&result)

	if query.Error != nil {
		return nil, query.Error
	}

	return result, nil
}
