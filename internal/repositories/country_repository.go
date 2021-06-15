package repositories

import (
	"gorm.io/gorm"
	"hotel-reservation/internal/commons"
	"hotel-reservation/internal/dto"
	"hotel-reservation/internal/models"
	"hotel-reservation/pkg/application_loger"
)

type CountryRepository struct {
	DB *gorm.DB
}

func NewCountryRepository(db *gorm.DB) *CountryRepository {
	return &CountryRepository{
		DB: db,
	}
}

func (r *CountryRepository) Create(country *models.Country) (*models.Country, error) {

	if tx := r.DB.Create(&country); tx.Error != nil {
		application_loger.LogError(tx.Error)
		return nil, tx.Error
	}

	return country, nil
}

func (r *CountryRepository) Update(country *models.Country) (*models.Country, error) {

	if tx := r.DB.Updates(&country); tx.Error != nil {
		application_loger.LogError(tx.Error)
		return nil, tx.Error
	}

	return country, nil
}

func (r *CountryRepository) Find(id uint64) (*models.Country, error) {

	model := models.Country{}
	if tx := r.DB.Where("id=?", id).Find(&model); tx.Error != nil {
		application_loger.LogError(tx.Error)
		return nil, tx.Error
	}

	if model.Id == 0 {
		return nil, nil
	}

	return &model, nil
}

func (r *CountryRepository) FindAll(input *dto.PaginationInput) (*commons.PaginatedList, error) {

	var list []models.Country
	var total int64

	query := r.DB.Model(&models.Country{})
	query.Count(&total)
	result := commons.NewPaginatedList(uint(total), uint(input.Page), uint(input.PerPage))
	query = query.Limit(int(result.PerPage)).Offset(int(result.Page)).Order("id desc").Find(&list)

	if query.Error != nil {
		application_loger.LogError(query.Error)
		return nil, query.Error
	}

	result.Data = list
	return result, nil
}

func (r *CountryRepository) GetCities(countryId uint64) ([]*models.City, error) {
	var result []*models.City

	query := r.DB.Model(&models.City{}).
		Where("country_id=?", countryId).Find(&result)

	if query.Error != nil {
		return nil, query.Error
	}

	return result, nil
}
