package repositories

import (
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
	"reservation-api/pkg/database/tenant_database_resolver"
)

type RateCodeDetailRepository struct {
	ConnectionResolver *tenant_database_resolver.TenantDatabaseResolver
}

// NewRateCodeDetailRepository returns new RateCodeDetailRepository.
func NewRateCodeDetailRepository(r *tenant_database_resolver.TenantDatabaseResolver) *RateCodeDetailRepository {

	return &RateCodeDetailRepository{ConnectionResolver: r}
}

func (r *RateCodeDetailRepository) Create(model *models.RateCodeDetail, tenantID uint64) (*models.RateCodeDetail, error) {

	db := r.ConnectionResolver.GetTenantDB(tenantID)

	if tx := db.Create(&model); tx.Error != nil {
		return nil, tx.Error
	}
	return model, nil
}

func (r *RateCodeDetailRepository) Update(model *models.RateCodeDetail, tenantID uint64) (*models.RateCodeDetail, error) {

	db := r.ConnectionResolver.GetTenantDB(tenantID)
	tx := db.Begin()

	// remove old price.
	if err := tx.Where("rate_code_detail_id=?", model.Id).Delete(&models.RateCodeDetailPrice{}).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	// update
	if err := tx.Updates(&model).Error; err != nil {
		return nil, err
	}

	tx.Commit()

	return model, nil
}

func (r *RateCodeDetailRepository) FindPrice(id uint64, tenantID uint64) (*models.RateCodeDetailPrice, error) {

	model := models.RateCodeDetailPrice{}
	db := r.ConnectionResolver.GetTenantDB(tenantID)

	if err := db.Model(models.RateCodeDetailPrice{}).Where("id=?", id).Find(&model).Error; err != nil {
		return nil, err
	}

	if model.Id == 0 {
		return nil, nil
	}

	return &model, nil
}

func (r *RateCodeDetailRepository) Find(id uint64, tenantID uint64) (*models.RateCodeDetail, error) {

	model := models.RateCodeDetail{}
	db := r.ConnectionResolver.GetTenantDB(tenantID)

	if tx := db.Where("id=?", id).Find(&model).Preload("RateCodeDetailPrice"); tx.Error != nil {
		return nil, tx.Error
	}

	if model.Id == 0 {
		return nil, nil
	}
	return &model, nil
}

func (r *RateCodeDetailRepository) FindAll(input *dto.PaginationFilter) (*commons.PaginatedResult, error) {
	db := r.ConnectionResolver.GetTenantDB(input.TenantID)
	return paginatedList(&models.RateCodeDetail{}, db, input)
}

func (r RateCodeDetailRepository) Delete(id uint64, tenantID uint64) error {

	db := r.ConnectionResolver.GetTenantDB(tenantID)

	if query := db.Model(&models.RateCodeDetail{}).Where("id=?", id).Delete(&models.RateCodeDetail{}); query.Error != nil {
		return query.Error
	}
	return nil
}
