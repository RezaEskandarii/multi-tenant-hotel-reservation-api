package repositories

import (
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
	"reservation-api/internal/utils"
	"reservation-api/pkg/database/connection_resolver"
)

//type CountryRepository interface {
//	Create(country *models.Country) (*models.Country, error)
//	Update(country *models.Country) (*models.Country, error)
//	Find(id uint64) (*models.Country, error)
//	FindAll(input *dto.PaginationFilter) (*commons.PaginatedResult, error)
//	GetProvinces(countryId uint64) ([]*models.Province, error)
//}

type CountryRepository struct {
	ConnectionResolver *connection_resolver.TenantConnectionResolver
}

func NewCountryRepository(connectionResolver *connection_resolver.TenantConnectionResolver) *CountryRepository {
	return &CountryRepository{ConnectionResolver: connectionResolver}
}

func (r *CountryRepository) Create(country *models.Country, tenantID uint64) (*models.Country, error) {

	valid, err := country.Validate()
	db := r.ConnectionResolver.GetDB(tenantID)

	if err != nil {
		return nil, err
	}

	if valid == false {
		return nil, nil
	}

	if tx := db.Create(&country); tx.Error != nil {

		return nil, tx.Error
	}

	return country, nil
}

func (r *CountryRepository) Update(country *models.Country, tenantID uint64) (*models.Country, error) {

	db := r.ConnectionResolver.GetDB(tenantID)

	if tx := db.Updates(&country); tx.Error != nil {
		return nil, tx.Error
	}

	return country, nil
}

func (r *CountryRepository) Find(id uint64, tenantID uint64) (*models.Country, error) {

	model := models.Country{}
	db := r.ConnectionResolver.GetDB(tenantID)

	if tx := db.Where("id=?", id).Preload("Provinces").Find(&model); tx.Error != nil {

		return nil, tx.Error
	}

	if model.Id == 0 {
		return nil, nil
	}

	return &model, nil
}

func (r *CountryRepository) FindAll(input *dto.PaginationFilter) (*commons.PaginatedResult, error) {

	db := r.ConnectionResolver.GetDB(input.TenantID)
	return paginatedList(&models.Country{}, db, input)
}

func (r *CountryRepository) GetProvinces(countryId uint64, tenantID uint64) ([]*models.Province, error) {
	var result []*models.Province

	db := r.ConnectionResolver.GetDB(tenantID)

	query := db.Model(&models.Province{}).
		Where("country_id=?", countryId).Find(&result)

	if query.Error != nil {
		return nil, query.Error
	}

	return result, nil
}

func (r *CountryRepository) Seed(jsonFilePath string, tenantID uint64) error {

	db := r.ConnectionResolver.GetDB(tenantID)

	countries := make([]models.Country, 0)
	if err := utils.CastJsonFileToStruct(jsonFilePath, &countries); err == nil {
		for _, country := range countries {
			var count int64 = 0
			if err := db.Model(models.Country{}).Where("name", country.Name).Count(&count).Error; err != nil {
				return err
			} else {
				if count == 0 {
					if err := db.Create(&country).Error; err != nil {
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
