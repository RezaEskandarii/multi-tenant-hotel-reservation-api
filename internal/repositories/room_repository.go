package repositories

import (
	"gorm.io/gorm"
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
)

type RoomRepository struct {
	DB *gorm.DB
}

func NewRoomRepository(db *gorm.DB) *RoomRepository {
	return &RoomRepository{DB: db}
}

func (r *RoomRepository) Create(room *models.Room) (*models.Room, error) {

	if tx := r.DB.Create(&room); tx.Error != nil {
		return nil, tx.Error
	}

	return room, nil
}

func (r *RoomRepository) Update(room *models.Room) (*models.Room, error) {

	if tx := r.DB.Updates(&room); tx.Error != nil {
		return nil, tx.Error
	}

	return room, nil
}

func (r *RoomRepository) Find(id uint64) (*models.Room, error) {

	model := models.Room{}

	if tx := r.DB.Where("id=?", id).Find(&model); tx.Error != nil {
		return nil, tx.Error
	}

	if model.Id == 0 {
		return nil, nil
	}

	return &model, nil
}

func (r *RoomRepository) FindAll(input *dto.PaginationFilter) (*commons.PaginatedList, error) {

	return paginate(&models.Room{}, r.DB, input)
}

func (r RoomRepository) Delete(id uint64) error {

	if query := r.DB.Model(&models.Room{}).Where("id=?", id).Delete(&models.Room{}); query.Error != nil {
		return query.Error
	}

	return nil
}
