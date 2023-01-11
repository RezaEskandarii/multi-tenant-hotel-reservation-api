package repositories

import (
	"context"
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
	"reservation-api/internal/tenant_resolver"
	"reservation-api/internal/utils"
	"reservation-api/pkg/multi_tenancy_database/tenant_database_resolver"
)

//type CountryRepository interface {
//	SetUp(country *models.Country) (*models.Country, error)
//	Update(country *models.Country) (*models.Country, error)
//	Find(id uint64) (*models.Country, error)
//	FindAll(input *dto.PaginationFilter) (*commons.PaginatedResult, error)
//	GetProvinces(countryId uint64) ([]*models.Province, error)
//}

type CountryRepository struct {
	DbResolver *tenant_database_resolver.TenantDatabaseResolver
}

func NewCountryRepository(connectionResolver *tenant_database_resolver.TenantDatabaseResolver) *CountryRepository {
	return &CountryRepository{DbResolver: connectionResolver}
}

func (r *CountryRepository) Create(ctx context.Context, country *models.Country) (*models.Country, error) {

	db := r.DbResolver.GetTenantDB(tenant_resolver.GetCurrentTenant(ctx))

	if tx := db.Create(&country); tx.Error != nil {
		return nil, tx.Error
	}

	return country, nil
}

func (r *CountryRepository) Update(ctx context.Context, country *models.Country) (*models.Country, error) {

	db := r.DbResolver.GetTenantDB(tenant_resolver.GetCurrentTenant(ctx))

	if tx := db.Updates(&country); tx.Error != nil {
		return nil, tx.Error
	}

	return country, nil
}

func (r *CountryRepository) Find(ctx context.Context, id uint64) (*models.Country, error) {

	model := models.Country{}
	db := r.DbResolver.GetTenantDB(tenant_resolver.GetCurrentTenant(ctx))

	if tx := db.Where("id=?", id).Preload("Provinces").Find(&model); tx.Error != nil {

		return nil, tx.Error
	}

	if model.Id == 0 {
		return nil, nil
	}

	return &model, nil
}

func (r *CountryRepository) FindAll(ctx context.Context, input *dto.PaginationFilter) (*commons.PaginatedResult, error) {

	db := r.DbResolver.GetTenantDB(tenant_resolver.GetCurrentTenant(ctx))
	return paginatedList(&models.Country{}, db, input)
}

func (r *CountryRepository) GetProvinces(ctx context.Context, countryId uint64) ([]*models.Province, error) {

	var result []*models.Province
	db := r.DbResolver.GetTenantDB(tenant_resolver.GetCurrentTenant(ctx))

	query := db.Model(&models.Province{}).
		Where("country_id=?", countryId).Find(&result)

	if query.Error != nil {
		return nil, query.Error
	}

	return result, nil
}

func (r *CountryRepository) Seed(ctx context.Context, jsonFilePath string) error {

	db := r.DbResolver.GetTenantDB(tenant_resolver.GetCurrentTenant(ctx))

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
