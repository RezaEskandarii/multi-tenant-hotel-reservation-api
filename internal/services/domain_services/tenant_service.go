package domain_services

import (
	"context"
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

	return s.Repository.Create(ctx, model)
}
