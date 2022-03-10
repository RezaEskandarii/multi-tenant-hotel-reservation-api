package repositories

import (
	"gorm.io/gorm"
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
	"reservation-api/internal/utils"
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

	valid, err := country.Validate()

	if err != nil {
		return nil, err
	}

	if valid == false {
		return nil, nil
	}

	if tx := r.DB.Create(&country); tx.Error != nil {

		return nil, tx.Error
	}

	return country, nil
}

func (r *CountryRepository) Update(country *models.Country) (*models.Country, error) {

	if tx := r.DB.Updates(&country); tx.Error != nil {

		return nil, tx.Error
	}

	return country, nil
}

func (r *CountryRepository) Find(id uint64) (*models.Country, error) {

	model := models.Country{}
	if tx := r.DB.Where("id=?", id).Preload("Provinces").Find(&model); tx.Error != nil {

		return nil, tx.Error
	}

	if model.Id == 0 {
		return nil, nil
	}

	return &model, nil
}

func (r *CountryRepository) FindAll(input *dto.PaginationInput) (*commons.PaginatedList, error) {

	return finAll(&models.Country{}, r.DB, input)
}

func (r *CountryRepository) GetProvinces(countryId uint64) ([]*models.Province, error) {
	var result []*models.Province

	query := r.DB.Model(&models.Province{}).
		Where("country_id=?", countryId).Find(&result)

	if query.Error != nil {
		return nil, query.Error
	}

	return result, nil
}

func (r *CountryRepository) Seed(jsonFilePath string) error {

	countries := make([]models.Country, 0)
	if err := utils.CastJsonFileToStruct(jsonFilePath, &countries); err == nil {
		for _, country := range countries {
			var count int64 = 0
			if err := r.DB.Model(models.Country{}).Where("name", country.Name).Count(&count).Error; err != nil {
				return err
			} else {
				if count == 0 {
					if err := r.DB.Create(&country).Error; err != nil {
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
