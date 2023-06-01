package repositories

import (
	"context"
	"errors"
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
	"reservation-api/internal/utils/file_utils"
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

	db := r.DbResolver.GetTenantDB(ctx)

	if tx := db.Create(&roomType); tx.Error != nil {
		return nil, tx.Error
	}

	return roomType, nil
}

func (r *RoomTypeRepository) Update(ctx context.Context, roomType *models.RoomType) (*models.RoomType, error) {

	db := r.DbResolver.GetTenantDB(ctx)

	if tx := db.Updates(&roomType); tx.Error != nil {
		return nil, tx.Error
	}

	return roomType, nil
}

func (r *RoomTypeRepository) Find(ctx context.Context, id uint64) (*models.RoomType, error) {

	db := r.DbResolver.GetTenantDB(ctx)
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

	db := r.DbResolver.GetTenantDB(ctx)
	return paginatedList(&models.RoomType{}, db, input)
}

func (r RoomTypeRepository) Delete(ctx context.Context, id uint64) error {

	var count int64 = 0
	db := r.DbResolver.GetTenantDB(ctx)

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
	// Get a handle to the tenant database
	db := r.DbResolver.GetTenantDB(ctx)

	// Read the JSON file and convert its contents to a slice of RoomType structs
	roomTypes := make([]models.RoomType, 0)
	if err := file_utils.CastJsonFileToStruct(jsonFilePath, &roomTypes); err != nil {
		return err
	}

	// Iterate over each room type and check if it already exists in the database
	for _, roomType := range roomTypes {
		var count int64
		if err := db.Model(models.RoomType{}).Where("name = ?", roomType.Name).Count(&count).Error; err != nil {
			return err
		}

		// If the room type does not exist in the database, create a new record for it
		if count == 0 {
			if err := db.Create(&roomType).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
