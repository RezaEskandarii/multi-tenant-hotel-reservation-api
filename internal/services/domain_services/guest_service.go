package domain_services

import (
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
	"reservation-api/internal/repositories"
)

type GuestService struct {
	Repository *repositories.GuestRepository
}

// NewGuestService returns new GuestService
func NewGuestService(r *repositories.GuestRepository) *GuestService {
	return &GuestService{Repository: r}
}

// Create creates new Guest.
func (s *GuestService) Create(guest *models.Guest, tenantID uint64) (*models.Guest, error) {

	return s.Repository.Create(guest, tenantID)
}

// Update updates Guest.
func (s *GuestService) Update(guest *models.Guest, tenantID uint64) (*models.Guest, error) {

	return s.Repository.Update(guest, tenantID)
}

// Find returns Guest and if it does not find the Guest, it returns nil.
func (s *GuestService) Find(id uint64, tenantID uint64) (*models.Guest, error) {

	return s.Repository.Find(id, tenantID)
}

// FindAll returns paginates list of cities.
func (s *GuestService) FindAll(input *dto.PaginationFilter) (*commons.PaginatedResult, error) {

	return s.Repository.FindAll(input)
}

// ReservationsCount returns guest reserves count
func (s *GuestService) ReservationsCount(guestId uint64, tenantID uint64) (error, uint64) {

	return s.Repository.ReservationsCount(guestId, tenantID)
}
