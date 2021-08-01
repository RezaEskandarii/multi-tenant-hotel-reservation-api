package repositories

import (
	"gorm.io/gorm"
	"hotel-reservation/internal/commons"
	"hotel-reservation/internal/dto"
	"hotel-reservation/internal/models"
)

type ProvinceRepository struct {
	DB *gorm.DB
}

func NewProvinceRepository(db *gorm.DB) *ProvinceRepository {
	return &ProvinceRepository{
		DB: db,
	}
}

func (r *ProvinceRepository) Create(province *models.Province) (*models.Province, error) {

	if tx := r.DB.Create(&province); tx.Error != nil {

		return nil, tx.Error
	}

	return province, nil
}

func (r *ProvinceRepository) Update(province *models.Province) (*models.Province, error) {

	if tx := r.DB.Updates(&province); tx.Error != nil {

		return nil, tx.Error
	}

	return province, nil
}

func (r *ProvinceRepository) Find(id uint64) (*models.Province, error) {

	model := models.Province{}
	if tx := r.DB.Where("id=?", id).Find(&model); tx.Error != nil {

		return nil, tx.Error
	}

	if model.Id == 0 {
		return nil, nil
	}

	return &model, nil
}

func (r *ProvinceRepository) FindAll(input *dto.PaginationInput) (*commons.PaginatedList, error) {

	return finAll(&models.Province{}, r.DB, input)
}

func (r *ProvinceRepository) GetCities(ProvinceId uint64) ([]*models.City, error) {
	var result []*models.City

	query := r.DB.Model(&models.City{}).
		Where("province_id=?", ProvinceId).Find(&result)

	if query.Error != nil {
		return nil, query.Error
	}

	return result, nil
}
