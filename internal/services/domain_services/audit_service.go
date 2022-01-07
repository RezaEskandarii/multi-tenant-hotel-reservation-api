package domain_services

import (
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
	"reservation-api/internal/repositories"
)

type AuditService struct {
	Repository *repositories.AuditRepository
}

func NewAuditService() *AuditService {
	return &AuditService{}
}

// Save creates new audit.
func (s *AuditService) Save(model *models.Audit) (*models.Audit, error) {
	return s.Repository.Create(model)
}

//FindAll returns paginated list of audits.
func (s *AuditService) FindAll(input *dto.PaginationInput) (*commons.PaginatedList, error) {

	return s.Repository.FindAll(input)
}
