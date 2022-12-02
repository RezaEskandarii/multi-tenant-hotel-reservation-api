package repositories

import (
	"context"
	"errors"
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/message_keys"
	"reservation-api/internal/models"
	"reservation-api/internal/tenant_resolver"
	"reservation-api/pkg/multi_tenancy_database/tenant_database_resolver"
)

type HotelRepository struct {
	ConnectionResolver *tenant_database_resolver.TenantDatabaseResolver
}

func NewHotelRepository(r *tenant_database_resolver.TenantDatabaseResolver) *HotelRepository {

	return &HotelRepository{
		ConnectionResolver: r,
	}
}

func (r *HotelRepository) Create(ctx context.Context, hotel *models.Hotel) (*models.Hotel, error) {

	db := r.ConnectionResolver.GetTenantDB(tenant_resolver.GetCurrentTenant(ctx))

	if tx := db.Create(&hotel); tx.Error != nil {
		return nil, tx.Error
	}

	return hotel, nil
}

func (r *HotelRepository) Update(ctx context.Context, hotel *models.Hotel) (*models.Hotel, error) {

	db := r.ConnectionResolver.GetTenantDB(tenant_resolver.GetCurrentTenant(ctx))

	if tx := db.Updates(&hotel); tx.Error != nil {
		return nil, tx.Error
	}

	return hotel, nil
}

func (r *HotelRepository) SetExtraData(ctx context.Context, id uint64, data string) error {

	db := r.ConnectionResolver.GetTenantDB(tenant_resolver.GetCurrentTenant(ctx))

	if err := db.Model(&models.Hotel{}).Update("extra_data", data).Where("id=?", id).Error; err != nil {
		return err
	}
	return nil
}

func (r *HotelRepository) Find(ctx context.Context, id uint64) (*models.Hotel, error) {

	model := models.Hotel{}
	db := r.ConnectionResolver.GetTenantDB(tenant_resolver.GetCurrentTenant(ctx))

	if tx := db.Where("id=?", id).Preload("Grades").Find(&model); tx.Error != nil {
		return nil, tx.Error
	}

	if model.Id == 0 {
		return nil, nil
	}

	return &model, nil
}

func (r *HotelRepository) FindAll(ctx context.Context, input *dto.PaginationFilter) (*commons.PaginatedResult, error) {
	db := r.ConnectionResolver.GetTenantDB(tenant_resolver.GetCurrentTenant(ctx))
	return paginatedList(&models.Hotel{}, db, input)
}

func (r HotelRepository) Delete(ctx context.Context, id uint64) error {

	db := r.ConnectionResolver.GetTenantDB(tenant_resolver.GetCurrentTenant(ctx))

	if query := db.Model(&models.Hotel{}).Where("id=?", id).Delete(&models.Hotel{}); query.Error != nil {
		return query.Error
	}

	return nil
}

func (r *HotelRepository) hasRepeatData(ctx context.Context, hotel *models.Hotel) error {

	var countByName int64 = 0
	db := r.ConnectionResolver.GetTenantDB(tenant_resolver.GetCurrentTenant(ctx))

	if tx := *db.Model(&models.Hotel{}).Where(&models.Hotel{Name: hotel.Name}).Count(&countByName); tx.Error != nil {
		return tx.Error
	}

	if countByName > 0 {

		return errors.New(message_keys.HotelRepeatPostalCode)
	}
	return nil
}
