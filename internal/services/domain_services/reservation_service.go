package domain_services

import (
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
	"reservation-api/internal/repositories"
)

type ReservationService struct {
	Repository *repositories.ReservationRepository
}

// NewReservationService returns new ReservationService
func NewReservationService() *ReservationService {
	return &ReservationService{}
}

// Create creates new Reservation.
func (s *ReservationService) Create(model *models.Reservation) (*models.Reservation, error) {

	return s.Repository.Create(model)
}

// Update updates Reservation.
func (s *ReservationService) Update(model *models.Reservation) (*models.Reservation, error) {

	return s.Repository.Update(model)
}

// CreateReservationRequest creates reservation request for given room to prevent concurrent request for specific room.
func (s *ReservationService) CreateReservationRequest(dto *dto.RoomRequestDto) (*models.ReservationRequest, error) {

	return s.Repository.CreateReservationRequest(dto)
}

func (s *ReservationService) HasConflict(request *dto.RoomRequestDto) (bool, error) {
	return s.Repository.HasConflict(request)
}

// CancelReservationRequest this function remove reservation request bt given requestKey param.
func (s *ReservationService) CancelReservationRequest(requestKey string) error {
	return s.Repository.CancelReservationRequest(requestKey)
}
