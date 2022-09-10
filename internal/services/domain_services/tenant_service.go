package domain_services

import (
	"reservation-api/internal/models"
	"reservation-api/internal/repositories"
)

type TenantService struct {
	Repository *repositories.TenantRepository
}

func NewTenantService(repository *repositories.TenantRepository) *TenantService {
	return &TenantService{Repository: repository}
}

func (s *TenantService) Create(model *models.Tenant) (*models.Tenant, error) {

	return s.Repository.Create(model)
}
