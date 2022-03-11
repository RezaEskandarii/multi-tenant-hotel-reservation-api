package repositories

import (
	"gorm.io/gorm"
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
)

type HotelGradeRepository struct {
	DB *gorm.DB
}

func NewHotelGradeRepository(db *gorm.DB) *HotelGradeRepository {
	return &HotelGradeRepository{DB: db}
}

func (r *HotelGradeRepository) Create(hotelGrade *models.HotelGrade) (*models.HotelGrade, error) {

	if tx := r.DB.Create(&hotelGrade); tx.Error != nil {

		return nil, tx.Error
	}

	return hotelGrade, nil
}

func (r *HotelGradeRepository) Update(hotelGrade *models.HotelGrade) (*models.HotelGrade, error) {

	if tx := r.DB.Updates(&hotelGrade); tx.Error != nil {

		return nil, tx.Error
	}

	return hotelGrade, nil
}

func (r *HotelGradeRepository) Find(id uint64) (*models.HotelGrade, error) {

	model := models.HotelGrade{}

	if tx := r.DB.Where("id=?", id).Preload("HotelType").Find(&model); tx.Error != nil {

		return nil, tx.Error
	}

	if model.Id == 0 {
		return nil, nil
	}

	return &model, nil
}

func (r *HotelGradeRepository) FindAll(input *dto.PaginationInput) (*commons.PaginatedList, error) {

	return finAll(&models.HotelGrade{}, r.DB, input)
}

func (r HotelGradeRepository) Delete(id uint64) error {

	var count int64 = 0

	if query := r.DB.Model(&models.Hotel{}).Where(&models.Hotel{HotelGradeId: id}).Count(&count); query.Error != nil {

		return query.Error
	}

	if count > 0 {
		return GradeHasHotel
	}

	if query := r.DB.Model(&models.HotelGrade{}).Where("id=?", id).Delete(&models.HotelGrade{}); query.Error != nil {

		return query.Error
	}

	return nil
}
