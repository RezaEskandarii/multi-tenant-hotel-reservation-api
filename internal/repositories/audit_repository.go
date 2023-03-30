package repositories

import (
	"context"
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
	"reservation-api/pkg/multi_tenancy_database/tenant_database_resolver"
)

type AuditRepository struct {
	DbResolver *tenant_database_resolver.TenantDatabaseResolver
}

func NewAuditRepository(r *tenant_database_resolver.TenantDatabaseResolver) *AuditRepository {
	return &AuditRepository{DbResolver: r}
}

func (r *AuditRepository) Create(ctx context.Context, model *models.Audit) (*models.Audit, error) {

	db := r.DbResolver.GetTenantDB(ctx)

	if err := db.Create(&model).Error; err != nil {
		return nil, err
	}
	return model, nil
}

func (r *AuditRepository) FindAll(ctx context.Context, input *dto.PaginationFilter) (*commons.PaginatedResult, error) {
	db := r.DbResolver.GetTenantDB(ctx)
	return paginatedList(&models.City{}, db, input)
}
