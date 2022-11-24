package domain_services

import (
	"context"
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
	"reservation-api/internal/repositories"
)

type AuditService struct {
	Repository *repositories.AuditRepository
}

func NewAuditService(repository *repositories.AuditRepository) *AuditService {
	return &AuditService{Repository: repository}
}

// Save creates new audit.
func (s *AuditService) Save(ctx context.Context, model *models.Audit) (*models.Audit, error) {
	return s.Repository.Create(ctx, model)
}

//FindAll returns paginated list of audits.
func (s *AuditService) FindAll(ctx context.Context, input *dto.PaginationFilter) (*commons.PaginatedResult, error) {

	return s.Repository.FindAll(ctx, input)
}
