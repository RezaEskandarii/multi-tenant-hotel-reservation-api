package repositories

import (
	"errors"
	"gorm.io/gorm"
	"hotel-reservation/internal/commons"
	"hotel-reservation/internal/dto"
	"hotel-reservation/internal/message_keys"
	"hotel-reservation/internal/models"
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

func (r *RoomTypeRepository) FindAll(input *dto.PaginationInput) (*commons.PaginatedList, error) {

	return finAll(&models.RoomType{}, r.DB, input)
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
