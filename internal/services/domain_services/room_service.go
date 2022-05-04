package domain_services

import (
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
func (s *RoomService) Create(room *models.Room) (*models.Room, error) {

	return s.Repository.Create(room)
}

// Update updates Room.
func (s *RoomService) Update(room *models.Room) (*models.Room, error) {

	return s.Repository.Update(room)
}

// Find returns Room and if it does not find the Room, it returns nil.
func (s *RoomService) Find(id uint64) (*models.Room, error) {

	return s.Repository.Find(id)
}

// FindAll returns paginates list of rooms.
func (s *RoomService) FindAll(input *dto.PaginationFilter) (*commons.PaginatedResult, error) {

	return s.Repository.FindAll(input)
}

// Delete removes room  by given id.
func (s *RoomService) Delete(id uint64) error {

	return s.Repository.Delete(id)
}
