package repositories

import (
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
	"reservation-api/pkg/database/connection_resolver"
)

type CityRepository struct {
	ConnectionResolver *connection_resolver.ConnectionResolver
}

//type CityRepository interface {
//	Create(city *models.City) (*models.City, error)
//	Update(city *models.City) (*models.City, error)
//	Find(city *models.City) (*models.City, error)
//	FindAll(input *dto.PaginationFilter) (commons.PaginatedResult, error)
//	Delete(id uint64) error
//}

func NewCityRepository(connectionResolver *connection_resolver.ConnectionResolver) *CityRepository {
	return &CityRepository{
		ConnectionResolver: connectionResolver,
	}
}

func (r *CityRepository) Create(city *models.City) (*models.City, error) {
	db := r.ConnectionResolver.GetDB(city.TenantId)

	if tx := db.Create(&city); tx.Error != nil {
		return nil, tx.Error
	}

	return city, nil
}

func (r *CityRepository) Update(city *models.City) (*models.City, error) {

	db := r.ConnectionResolver.GetDB(city.TenantId)
	if tx := db.Updates(&city); tx.Error != nil {

		return nil, tx.Error
	}

	return city, nil
}

func (r *CityRepository) Delete(id uint64, tenantID uint64) error {

	db := r.ConnectionResolver.GetDB(tenantID)
	if err := db.Model(&models.City{}).Where("id=?", id).Delete(&models.City{}).Error; err != nil {

		return err
	}

	return nil
}

func (r *CityRepository) Find(id uint64, tenantID uint64) (*models.City, error) {

	model := models.City{}
	db := r.ConnectionResolver.GetDB(tenantID)

	if tx := db.Where("id=?", id).Find(&model); tx.Error != nil {

		return nil, tx.Error
	}

	if model.Id == 0 {
		return nil, nil
	}

	return &model, nil
}

func (r *CityRepository) FindAll(input *dto.PaginationFilter) (*commons.PaginatedResult, error) {
	db := r.ConnectionResolver.GetDB(input.TenantID)
	return paginatedList(&models.City{}, db, input)
}
