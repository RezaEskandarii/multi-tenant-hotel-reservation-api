package repositories

import (
	"gorm.io/gorm"
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
)

type CityRepository struct {
	DB *gorm.DB
}

//type CityRepository interface {
//	Create(city *models.City) (*models.City, error)
//	Update(city *models.City) (*models.City, error)
//	Find(city *models.City) (*models.City, error)
//	FindAll(input *dto.PaginationInput) (commons.PaginatedList, error)
//	Delete(id uint64) error
//}

func NewCityRepository(db *gorm.DB) *CityRepository {
	return &CityRepository{
		DB: db,
	}
}

func (r *CityRepository) Create(city *models.City) (*models.City, error) {

	if tx := r.DB.Create(&city); tx.Error != nil {
		return nil, tx.Error
	}

	return city, nil
}

func (r *CityRepository) Update(city *models.City) (*models.City, error) {

	if tx := r.DB.Updates(&city); tx.Error != nil {

		return nil, tx.Error
	}

	return city, nil
}

func (r *CityRepository) Find(id uint64) (*models.City, error) {

	model := models.City{}
	if tx := r.DB.Where("id=?", id).Find(&model); tx.Error != nil {

		return nil, tx.Error
	}

	if model.Id == 0 {
		return nil, nil
	}

	return &model, nil
}

func (r *CityRepository) FindAll(input *dto.PaginationInput) (*commons.PaginatedList, error) {

	return finAll(&models.City{}, r.DB, input)
}
