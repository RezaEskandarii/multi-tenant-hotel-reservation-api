package repositories

import (
	"context"
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
	"reservation-api/internal/tenant_resolver"
	"reservation-api/pkg/multi_tenancy_database/tenant_database_resolver"
)

type HotelGradeRepository struct {
	ConnectionResolver *tenant_database_resolver.TenantDatabaseResolver
}

func NewHotelGradeRepository(r *tenant_database_resolver.TenantDatabaseResolver) *HotelGradeRepository {
	return &HotelGradeRepository{ConnectionResolver: r}
}

func (r *HotelGradeRepository) Create(ctx context.Context, hotelGrade *models.HotelGrade) (*models.HotelGrade, error) {

	db := r.ConnectionResolver.GetTenantDB(tenant_resolver.GetCurrentTenant(ctx))

	if tx := db.Create(&hotelGrade); tx.Error != nil {
		return nil, tx.Error
	}

	return hotelGrade, nil
}

func (r *HotelGradeRepository) Update(ctx context.Context, hotelGrade *models.HotelGrade) (*models.HotelGrade, error) {

	db := r.ConnectionResolver.GetTenantDB(tenant_resolver.GetCurrentTenant(ctx))

	if tx := db.Updates(&hotelGrade); tx.Error != nil {
		return nil, tx.Error
	}

	return hotelGrade, nil
}

func (r *HotelGradeRepository) Find(ctx context.Context, id uint64) (*models.HotelGrade, error) {

	model := models.HotelGrade{}
	db := r.ConnectionResolver.GetTenantDB(tenant_resolver.GetCurrentTenant(ctx))

	if tx := db.Where("id=?", id).Preload("HotelType").Find(&model); tx.Error != nil {
		return nil, tx.Error
	}

	if model.Id == 0 {
		return nil, nil
	}

	return &model, nil
}

func (r *HotelGradeRepository) FindAll(ctx context.Context, input *dto.PaginationFilter) (*commons.PaginatedResult, error) {

	db := r.ConnectionResolver.GetTenantDB(tenant_resolver.GetCurrentTenant(ctx))
	return paginatedList(&models.HotelGrade{}, db, input)
}

func (r HotelGradeRepository) Delete(ctx context.Context, id uint64) error {

	var count int64 = 0
	db := r.ConnectionResolver.GetTenantDB(tenant_resolver.GetCurrentTenant(ctx))

	if query := db.Model(&models.Hotel{}).Where(&models.Hotel{HotelGradeId: id}).Count(&count); query.Error != nil {
		return query.Error
	}

	if count > 0 {
		return GradeHasHotel
	}

	if query := db.Model(&models.HotelGrade{}).Where("id=?", id).Delete(&models.HotelGrade{}); query.Error != nil {
		return query.Error
	}

	return nil
}
