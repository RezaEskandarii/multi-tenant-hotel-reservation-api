package repositories

import (
	"context"
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
	"reservation-api/internal/tenant_resolver"
	"reservation-api/pkg/multi_tenancy_database/tenant_database_resolver"
)

type CityRepository struct {
	DbResolver *tenant_database_resolver.TenantDatabaseResolver
}

func NewCityRepository(connectionResolver *tenant_database_resolver.TenantDatabaseResolver) *CityRepository {
	return &CityRepository{
		DbResolver: connectionResolver,
	}
}

func (r *CityRepository) Create(ctx context.Context, city *models.City) (*models.City, error) {

	db := r.DbResolver.GetTenantDB(tenant_resolver.GetCurrentTenant(ctx))
	if tx := db.Create(&city); tx.Error != nil {
		return nil, tx.Error
	}

	return city, nil
}

func (r *CityRepository) Update(ctx context.Context, city *models.City) (*models.City, error) {

	db := r.DbResolver.GetTenantDB(tenant_resolver.GetCurrentTenant(ctx))

	if tx := db.Updates(&city); tx.Error != nil {
		return nil, tx.Error
	}

	return city, nil
}

func (r *CityRepository) Delete(ctx context.Context, id uint64) error {

	db := r.DbResolver.GetTenantDB(tenant_resolver.GetCurrentTenant(ctx))

	if err := db.Model(&models.City{}).Where("id=?", id).Delete(&models.City{}).Error; err != nil {
		return err
	}

	return nil
}

func (r *CityRepository) Find(ctx context.Context, id uint64) (*models.City, error) {

	model := models.City{}
	db := r.DbResolver.GetTenantDB(tenant_resolver.GetCurrentTenant(ctx))

	if tx := db.Where("id=?", id).Find(&model); tx.Error != nil {

		return nil, tx.Error
	}

	if model.Id == 0 {
		return nil, nil
	}

	return &model, nil
}

func (r *CityRepository) FindAll(ctx context.Context, input *dto.PaginationFilter) (*commons.PaginatedResult, error) {
	db := r.DbResolver.GetTenantDB(tenant_resolver.GetCurrentTenant(ctx))
	return paginatedList(&models.City{}, db, input)
}
