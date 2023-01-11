package repositories

import (
	"context"
	"errors"
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
	"reservation-api/internal/tenant_resolver"
	"reservation-api/internal/utils"
	"reservation-api/internal_errors/message_keys"
	"reservation-api/pkg/multi_tenancy_database/tenant_database_resolver"
)

var (
	RoomTypeHasRoomErr = errors.New(message_keys.RoomTypeHasRoomErr)
)

type RoomTypeRepository struct {
	DbResolver *tenant_database_resolver.TenantDatabaseResolver
}

func NewRoomTypeRepository(r *tenant_database_resolver.TenantDatabaseResolver) *RoomTypeRepository {
	return &RoomTypeRepository{DbResolver: r}
}

func (r *RoomTypeRepository) Create(ctx context.Context, roomType *models.RoomType) (*models.RoomType, error) {

	db := r.DbResolver.GetTenantDB(tenant_resolver.GetCurrentTenant(ctx))

	if tx := db.Create(&roomType); tx.Error != nil {
		return nil, tx.Error
	}

	return roomType, nil
}

func (r *RoomTypeRepository) Update(ctx context.Context, roomType *models.RoomType) (*models.RoomType, error) {

	db := r.DbResolver.GetTenantDB(tenant_resolver.GetCurrentTenant(ctx))

	if tx := db.Updates(&roomType); tx.Error != nil {
		return nil, tx.Error
	}

	return roomType, nil
}

func (r *RoomTypeRepository) Find(ctx context.Context, id uint64) (*models.RoomType, error) {

	db := r.DbResolver.GetTenantDB(tenant_resolver.GetCurrentTenant(ctx))
	model := models.RoomType{}

	if tx := db.Where("id=?", id).Find(&model); tx.Error != nil {
		return nil, tx.Error
	}

	if model.Id == 0 {
		return nil, nil
	}

	return &model, nil
}

func (r *RoomTypeRepository) FindAll(ctx context.Context, input *dto.PaginationFilter) (*commons.PaginatedResult, error) {

	db := r.DbResolver.GetTenantDB(tenant_resolver.GetCurrentTenant(ctx))
	return paginatedList(&models.RoomType{}, db, input)
}

func (r RoomTypeRepository) Delete(ctx context.Context, id uint64) error {

	var count int64 = 0
	db := r.DbResolver.GetTenantDB(tenant_resolver.GetCurrentTenant(ctx))

	if query := db.Model(&models.Room{}).Where(&models.Room{RoomTypeId: id}).Count(&count); query.Error != nil {

		return query.Error
	}

	if count > 0 {
		return RoomTypeHasRoomErr
	}

	if query := db.Model(&models.RoomType{}).Where("id=?", id).Delete(&models.RoomType{}); query.Error != nil {

		return query.Error
	}

	return nil
}

func (r *RoomTypeRepository) Seed(ctx context.Context, jsonFilePath string) error {

	db := r.DbResolver.GetTenantDB(tenant_resolver.GetCurrentTenant(ctx))

	roomTypes := make([]models.RoomType, 0)
	if err := utils.CastJsonFileToStruct(jsonFilePath, &roomTypes); err == nil {
		for _, roomType := range roomTypes {
			var count int64 = 0
			if err := db.Model(models.RoomType{}).Count(&count).Error; err != nil {
				return err
			} else {
				if count == 0 {
					if err := db.Create(&roomType).Error; err != nil {
						return err
					}
				}
			}
		}
	} else {
		return err
	}
	return nil
}
