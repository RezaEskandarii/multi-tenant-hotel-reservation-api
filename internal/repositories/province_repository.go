package repositories

import (
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
	"reservation-api/pkg/database/tenant_database_resolver"
)

type ProvinceRepository struct {
	ConnectionResolver *tenant_database_resolver.TenantDatabaseResolver
}

func NewProvinceRepository(r *tenant_database_resolver.TenantDatabaseResolver) *ProvinceRepository {
	return &ProvinceRepository{
		ConnectionResolver: r,
	}
}

func (r *ProvinceRepository) Create(province *models.Province, tenantID uint64) (*models.Province, error) {

	db := r.ConnectionResolver.GetTenantDB(tenantID)

	if tx := db.Create(&province); tx.Error != nil {
		return nil, tx.Error
	}

	return province, nil
}

func (r *ProvinceRepository) Update(province *models.Province, tenantID uint64) (*models.Province, error) {

	db := r.ConnectionResolver.GetTenantDB(tenantID)

	if tx := db.Updates(&province); tx.Error != nil {
		return nil, tx.Error
	}

	return province, nil
}

func (r *ProvinceRepository) Find(id uint64, tenantID uint64) (*models.Province, error) {

	model := models.Province{}
	db := r.ConnectionResolver.GetTenantDB(tenantID)

	if tx := db.Where("id=?", id).Find(&model); tx.Error != nil {
		return nil, tx.Error
	}

	if model.Id == 0 {
		return nil, nil
	}

	return &model, nil
}

func (r *ProvinceRepository) FindAll(input *dto.PaginationFilter) (*commons.PaginatedResult, error) {

	db := r.ConnectionResolver.GetTenantDB(input.TenantID)
	return paginatedList(&models.Province{}, db, input)
}

func (r *ProvinceRepository) GetCities(provinceId uint64, tenantID uint64) ([]*models.City, error) {

	var result []*models.City
	db := r.ConnectionResolver.GetTenantDB(tenantID)

	query := db.Model(&models.City{}).
		Where("province_id=?", provinceId).Find(&result)

	if query.Error != nil {
		return nil, query.Error
	}

	return result, nil
}
