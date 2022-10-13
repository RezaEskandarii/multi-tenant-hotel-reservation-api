package repositories

import (
	"errors"
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/message_keys"
	"reservation-api/internal/models"
	"reservation-api/internal/utils"
	"reservation-api/pkg/database/connection_resolver"
)

var (
	RoomTypeHasRoomErr = errors.New(message_keys.RoomTypeHasRoomErr)
)

type RoomTypeRepository struct {
	ConnectionResolver *connection_resolver.TenantConnectionResolver
}

func NewRoomTypeRepository(r *connection_resolver.TenantConnectionResolver) *RoomTypeRepository {
	return &RoomTypeRepository{ConnectionResolver: r}
}

func (r *RoomTypeRepository) Create(roomType *models.RoomType, tenantID uint64) (*models.RoomType, error) {

	db := r.ConnectionResolver.GetDB(tenantID)

	if tx := db.Create(&roomType); tx.Error != nil {
		return nil, tx.Error
	}

	return roomType, nil
}

func (r *RoomTypeRepository) Update(roomType *models.RoomType, tenantID uint64) (*models.RoomType, error) {

	db := r.ConnectionResolver.GetDB(tenantID)

	if tx := db.Updates(&roomType); tx.Error != nil {
		return nil, tx.Error
	}

	return roomType, nil
}

func (r *RoomTypeRepository) Find(id uint64, tenantID uint64) (*models.RoomType, error) {

	db := r.ConnectionResolver.GetDB(tenantID)
	model := models.RoomType{}

	if tx := db.Where("id=?", id).Find(&model); tx.Error != nil {
		return nil, tx.Error
	}

	if model.Id == 0 {
		return nil, nil
	}

	return &model, nil
}

func (r *RoomTypeRepository) FindAll(input *dto.PaginationFilter) (*commons.PaginatedResult, error) {

	db := r.ConnectionResolver.GetDB(input.TenantID)
	return paginatedList(&models.RoomType{}, db, input)
}

func (r RoomTypeRepository) Delete(id uint64, tenantID uint64) error {

	var count int64 = 0
	db := r.ConnectionResolver.GetDB(tenantID)

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

func (r *RoomTypeRepository) Seed(jsonFilePath string, tenantID uint64) error {

	db := r.ConnectionResolver.GetDB(tenantID)

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
