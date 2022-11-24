package domain_services

import (
	"context"
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
	"reservation-api/internal/repositories"
)

type RoomService struct {
	Repository *repositories.RoomRepository
}

func NewRoomService(r *repositories.RoomRepository) *RoomService {
	return &RoomService{Repository: r}
}

// Create creates new Room.
func (s *RoomService) Create(ctx context.Context, room *models.Room) (*models.Room, error) {

	return s.Repository.Create(ctx, room)
}

// Update updates Room.
func (s *RoomService) Update(ctx context.Context, room *models.Room) (*models.Room, error) {

	return s.Repository.Update(ctx, room)
}

// Find returns Room and if it does not find the Room, it returns nil.
func (s *RoomService) Find(ctx context.Context, id uint64) (*models.Room, error) {

	return s.Repository.Find(ctx, id)
}

// FindAll returns paginates list of rooms.
func (s *RoomService) FindAll(ctx context.Context, filter *dto.PaginationFilter) (*commons.PaginatedResult, error) {

	return s.Repository.FindAll(ctx, filter)
}

// Delete removes room  by given id.
func (s *RoomService) Delete(ctx context.Context, id uint64) error {

	return s.Repository.Delete(ctx, id)
}
