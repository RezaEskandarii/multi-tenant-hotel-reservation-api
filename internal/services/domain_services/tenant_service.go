package domain_services

import (
	"context"
	"reservation-api/internal/global_variables"
	"reservation-api/internal/models"
	"reservation-api/internal/repositories"
)

type TenantService struct {
	Repository *repositories.TenantRepository
}

func NewTenantService(repository *repositories.TenantRepository) *TenantService {
	return &TenantService{Repository: repository}
}

func (s *TenantService) SetUp(ctx context.Context, model *models.Tenant) (*models.Tenant, error) {

	dbCtx := context.WithValue(ctx, global_variables.TenantIDKey, 0)

	return s.Repository.Create(dbCtx, model)
}
