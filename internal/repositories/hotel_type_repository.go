package repositories

import (
	"gorm.io/gorm"
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
)

type HotelTypeRepository struct {
	DB *gorm.DB
}

func NewHotelTypeRepository(db *gorm.DB) *HotelTypeRepository {
	return &HotelTypeRepository{DB: db}
}

func (r *HotelTypeRepository) Create(hotelType *models.HotelType) (*models.HotelType, error) {

	if tx := r.DB.Create(&hotelType); tx.Error != nil {
		return nil, tx.Error
	}

	return hotelType, nil
}

func (r *HotelTypeRepository) Update(hotelType *models.HotelType) (*models.HotelType, error) {

	if tx := r.DB.Updates(&hotelType); tx.Error != nil {
		return nil, tx.Error
	}

	return hotelType, nil
}

func (r *HotelTypeRepository) Find(id uint64) (*models.HotelType, error) {

	model := models.HotelType{}

	if tx := r.DB.Where("id=?", id).Preload("Grades").Find(&model); tx.Error != nil {
		return nil, tx.Error
	}

	if model.Id == 0 {
		return nil, nil
	}

	return &model, nil
}

func (r *HotelTypeRepository) FindAll(input *dto.PaginationFilter) (*commons.PaginatedList, error) {

	return paginate(&models.HotelType{}, r.DB, input)
}

func (r HotelTypeRepository) Delete(id uint64) error {

	var count int64 = 0

	if query := r.DB.Model(&models.Hotel{}).Where(&models.Hotel{HotelTypeId: id}).Count(&count); query.Error != nil {

		return query.Error
	}

	if count > 0 {
		return TypeHasHotelError
	}

	if query := r.DB.Model(&models.HotelType{}).Where("id=?", id).Delete(&models.HotelType{}); query.Error != nil {

		return query.Error
	} else {

		query = r.DB.Model(&models.HotelGrade{}).Where(&models.HotelGrade{HotelTypeId: id}).Delete(&models.HotelGrade{})

		if query.Error != nil {
			return query.Error
		}
	}

	return nil
}
