package repositories

import (
	"errors"
	"gorm.io/gorm"
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/message_keys"
	"reservation-api/internal/models"
	"reservation-api/internal/utils"
)

var (
	RoomTypeHasRoomErr = errors.New(message_keys.RoomTypeHasRoomErr)
)

type RoomTypeRepository struct {
	DB *gorm.DB
}

func NewRoomTypeRepository(db *gorm.DB) *RoomTypeRepository {
	return &RoomTypeRepository{DB: db}
}

func (r *RoomTypeRepository) Create(roomType *models.RoomType) (*models.RoomType, error) {

	if tx := r.DB.Create(&roomType); tx.Error != nil {
		return nil, tx.Error
	}

	return roomType, nil
}

func (r *RoomTypeRepository) Update(roomType *models.RoomType) (*models.RoomType, error) {

	if tx := r.DB.Updates(&roomType); tx.Error != nil {
		return nil, tx.Error
	}

	return roomType, nil
}

func (r *RoomTypeRepository) Find(id uint64) (*models.RoomType, error) {

	model := models.RoomType{}

	if tx := r.DB.Where("id=?", id).Find(&model); tx.Error != nil {
		return nil, tx.Error
	}

	if model.Id == 0 {
		return nil, nil
	}

	return &model, nil
}

func (r *RoomTypeRepository) FindAll(input *dto.PaginationFilter) (*commons.PaginatedList, error) {

	return paginatedList(&models.RoomType{}, r.DB, input)
}

func (r RoomTypeRepository) Delete(id uint64) error {

	var count int64 = 0

	if query := r.DB.Model(&models.Room{}).Where(&models.Room{RoomTypeId: id}).Count(&count); query.Error != nil {

		return query.Error
	}

	if count > 0 {
		return RoomTypeHasRoomErr
	}

	if query := r.DB.Model(&models.RoomType{}).Where("id=?", id).Delete(&models.RoomType{}); query.Error != nil {

		return query.Error
	}

	return nil
}

func (r *RoomTypeRepository) Seed(jsonFilePath string) error {

	roomTypes := make([]models.RoomType, 0)
	if err := utils.CastJsonFileToStruct(jsonFilePath, &roomTypes); err == nil {
		for _, roomType := range roomTypes {
			var count int64 = 0
			if err := r.DB.Model(models.RoomType{}).Count(&count).Error; err != nil {
				return err
			} else {
				if count == 0 {
					if err := r.DB.Create(&roomType).Error; err != nil {
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
