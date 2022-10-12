package repositories

import (
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
	"reservation-api/pkg/database/connection_resolver"
)

type AuditRepository struct {
	ConnectionResolver *connection_resolver.ConnectionResolver
}

func NewAuditRepository(r *connection_resolver.ConnectionResolver) *AuditRepository {
	return &AuditRepository{ConnectionResolver: r}
}

func (r *AuditRepository) Create(model *models.Audit) (*models.Audit, error) {
	db := r.ConnectionResolver.GetDB(model.TenantId)
	if err := db.Create(&model).Error; err != nil {
		return nil, err
	}
	return model, nil
}

func (r *AuditRepository) FindAll(input *dto.PaginationFilter) (*commons.PaginatedResult, error) {
	db := r.ConnectionResolver.GetDB(input.TenantID)
	return paginatedList(&models.City{}, db, input)
}
