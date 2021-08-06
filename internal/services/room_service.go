package services

import (
	"hotel-reservation/internal/commons"
	"hotel-reservation/internal/dto"
	"hotel-reservation/internal/models"
	"hotel-reservation/internal/repositories"
)

type RoomService struct {
	Repository *repositories.RoomRepository
}

func NewRoomService() *RoomService {
	return &RoomService{}
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
func (s *RoomService) FindAll(input *dto.PaginationInput) (*commons.PaginatedList, error) {

	return s.Repository.FindAll(input)
}

// Delete removes room  by given id.
func (s *RoomService) Delete(id uint64) error {

	return s.Repository.Delete(id)
}
