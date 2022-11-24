package domain_services

import (
	"context"
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
func (s *GuestService) Create(ctx context.Context, guest *models.Guest) (*models.Guest, error) {

	return s.Repository.Create(ctx, guest)
}

// Update updates Guest.
func (s *GuestService) Update(ctx context.Context, guest *models.Guest) (*models.Guest, error) {

	return s.Repository.Update(ctx, guest)
}

// Find returns Guest and if it does not find the Guest, it returns nil.
func (s *GuestService) Find(ctx context.Context, id uint64) (*models.Guest, error) {

	return s.Repository.Find(ctx, id)
}

// FindAll returns paginates list of cities.
func (s *GuestService) FindAll(ctx context.Context, filter *dto.PaginationFilter) (*commons.PaginatedResult, error) {

	return s.Repository.FindAll(ctx, filter)
}

// ReservationsCount returns guest reserves count
func (s *GuestService) ReservationsCount(ctx context.Context, guestId uint64) (error, uint64) {

	return s.Repository.ReservationsCount(ctx, guestId)
}
