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
func NewGuestService() *GuestService {
	return &GuestService{}
}

// Create creates new Guest.
func (s *GuestService) Create(guest *models.Guest) (*models.Guest, error) {

	return s.Repository.Create(guest)
}

// Update updates Guest.
func (s *GuestService) Update(guest *models.Guest) (*models.Guest, error) {

	return s.Repository.Update(guest)
}

// Find returns Guest and if it does not find the Guest, it returns nil.
func (s *GuestService) Find(id uint64) (*models.Guest, error) {

	return s.Repository.Find(id)
}

// FindAll returns paginates list of cities.
func (s *GuestService) FindAll(input *dto.PaginationInput) (*commons.PaginatedList, error) {

	return s.Repository.FindAll(input)
}

// ReservationsCount returns guest reserves count
func (s *GuestService) ReservationsCount(guestId uint64) (error, uint64) {

	return s.Repository.ReservationsCount(guestId)
}
