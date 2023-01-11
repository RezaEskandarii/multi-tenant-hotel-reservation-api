package repositories

import (
	"context"
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
	"reservation-api/internal/tenant_resolver"
	"reservation-api/internal_errors"
	"reservation-api/pkg/multi_tenancy_database/tenant_database_resolver"
)

type HotelTypeRepository struct {
	DbResolver *tenant_database_resolver.TenantDatabaseResolver
}

func NewHotelTypeRepository(r *tenant_database_resolver.TenantDatabaseResolver) *HotelTypeRepository {
	return &HotelTypeRepository{r}
}

func (r *HotelTypeRepository) Create(ctx context.Context, hotelType *models.HotelType) (*models.HotelType, error) {

	db := r.DbResolver.GetTenantDB(tenant_resolver.GetCurrentTenant(ctx))

	if tx := db.Create(&hotelType); tx.Error != nil {
		return nil, tx.Error
	}

	return hotelType, nil
}

func (r *HotelTypeRepository) Update(ctx context.Context, hotelType *models.HotelType) (*models.HotelType, error) {

	db := r.DbResolver.GetTenantDB(tenant_resolver.GetCurrentTenant(ctx))

	if tx := db.Updates(&hotelType); tx.Error != nil {
		return nil, tx.Error
	}

	return hotelType, nil
}

func (r *HotelTypeRepository) Find(ctx context.Context, id uint64) (*models.HotelType, error) {

	model := models.HotelType{}
	db := r.DbResolver.GetTenantDB(tenant_resolver.GetCurrentTenant(ctx))

	if tx := db.Where("id=?", id).Preload("Grades").Find(&model); tx.Error != nil {
		return nil, tx.Error
	}

	if model.Id == 0 {
		return nil, nil
	}

	return &model, nil
}

func (r *HotelTypeRepository) FindAll(ctx context.Context, input *dto.PaginationFilter) (*commons.PaginatedResult, error) {

	db := r.DbResolver.GetTenantDB(tenant_resolver.GetCurrentTenant(ctx))
	return paginatedList(&models.HotelType{}, db, input)
}

func (r HotelTypeRepository) Delete(ctx context.Context, id uint64) error {

	var count int64 = 0
	db := r.DbResolver.GetTenantDB(tenant_resolver.GetCurrentTenant(ctx))

	if query := db.Model(&models.Hotel{}).Where(&models.Hotel{HotelTypeId: id}).Count(&count); query.Error != nil {
		return query.Error
	}

	if count > 0 {
		return internal_errors.TypeHasHotelError
	}

	if query := db.Model(&models.HotelType{}).Where("id=?", id).Delete(&models.HotelType{}); query.Error != nil {
		return query.Error

	} else {
		query = db.Model(&models.HotelGrade{}).Where(&models.HotelGrade{HotelTypeId: id}).Delete(&models.HotelGrade{})
		if query.Error != nil {

			return query.Error
		}
	}

	return nil
}
