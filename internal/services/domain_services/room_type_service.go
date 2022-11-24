package domain_services

import (
	"context"
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
	"reservation-api/internal/repositories"
)

type RoomTypeService struct {
	Repository *repositories.RoomTypeRepository
}

func NewRoomTypeService(r *repositories.RoomTypeRepository) *RoomTypeService {
	return &RoomTypeService{Repository: r}
}

// Create creates new RoomType.
func (s *RoomTypeService) Create(ctx context.Context, roomType *models.RoomType) (*models.RoomType, error) {

	return s.Repository.Create(ctx, roomType)
}

// Update updates RoomType.
func (s *RoomTypeService) Update(ctx context.Context, roomType *models.RoomType) (*models.RoomType, error) {

	return s.Repository.Update(ctx, roomType)
}

// Find returns RoomType and if it does not find the RoomType, it returns nil.
func (s *RoomTypeService) Find(ctx context.Context, id uint64) (*models.RoomType, error) {

	return s.Repository.Find(ctx, id)
}

// FindAll returns paginates list of hotel types.
func (s *RoomTypeService) FindAll(ctx context.Context, filter *dto.PaginationFilter) (*commons.PaginatedResult, error) {

	return s.Repository.FindAll(ctx, filter)
}

// Delete removes hotel type by given id.
func (s *RoomTypeService) Delete(ctx context.Context, id uint64) error {

	return s.Repository.Delete(ctx, id)
}

// Seed seed given json file data to roomTypes.
func (s *RoomTypeService) Seed(ctx context.Context, jsonFilePath string) error {
	return s.Repository.Seed(ctx, jsonFilePath)
}
