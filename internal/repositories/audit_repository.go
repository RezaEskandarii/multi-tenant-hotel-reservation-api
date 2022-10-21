package repositories

import (
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
	"reservation-api/pkg/database/tenant_database_resolver"
)

type AuditRepository struct {
	ConnectionResolver *tenant_database_resolver.TenantDatabaseResolver
}

func NewAuditRepository(r *tenant_database_resolver.TenantDatabaseResolver) *AuditRepository {
	return &AuditRepository{ConnectionResolver: r}
}

func (r *AuditRepository) Create(model *models.Audit, tenantID uint64) (*models.Audit, error) {

	db := r.ConnectionResolver.GetTenantDB(tenantID)

	if err := db.Create(&model).Error; err != nil {
		return nil, err
	}
	return model, nil
}

func (r *AuditRepository) FindAll(input *dto.PaginationFilter) (*commons.PaginatedResult, error) {
	db := r.ConnectionResolver.GetTenantDB(input.TenantID)
	return paginatedList(&models.City{}, db, input)
}
