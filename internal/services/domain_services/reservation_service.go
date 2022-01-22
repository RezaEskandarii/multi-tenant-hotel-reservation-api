package domain_services

import (
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
	"reservation-api/internal/repositories"
	"time"
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

// ChangeStatus changes the reservation check status.
func (s *ReservationService) ChangeStatus(id uint64, status models.ReservationCheckStatus) (*models.Reservation, error) {

	return s.Repository.ChangeStatus(id, status)
}

// Update updates Reservation.
func (s *ReservationService) Update(id uint64, model *models.Reservation) (*models.Reservation, error) {

	return s.Repository.Update(id, model)
}

// CreateReservationRequest creates reservation request for given room to prevent concurrent request for specific room.
func (s *ReservationService) CreateReservationRequest(dto *dto.RoomRequestDto) (*models.ReservationRequest, error) {

	return s.Repository.CreateReservationRequest(dto)
}

func (s *ReservationService) HasConflict(request *dto.RoomRequestDto) (bool, error) {
	return s.Repository.HasConflict(request)
}

func (s *ReservationService) HasReservationConflict(checkInDate *time.Time, checkOutDate *time.Time, roomId uint64) (bool, error) {
	return s.Repository.HasReservationConflict(checkInDate, checkOutDate, roomId)
}

// RemoveReservationRequest this function remove reservation request bt given requestKey param.
func (s *ReservationService) RemoveReservationRequest(requestKey string) error {
	return s.Repository.DeleteReservationRequest(requestKey)
}

// GetRecommendedRateCodes returns list of recommended rateCodeDetails price per reservation condition.
func (s *ReservationService) GetRecommendedRateCodes(priceDto *dto.GetRatePriceDto) ([]*dto.RateCodePricesDto, error) {
	return s.Repository.GetRecommendedRateCodes(priceDto)
}

// Find find and returns reservation by id.
func (s *ReservationService) Find(id uint64) (*models.Reservation, error) {
	return s.Repository.Find(id)
}

// FindReservationRequest find and returns reservationRequest by  given roomId and requestKey.
func (s *ReservationService) FindReservationRequest(requestKey string, roomId uint64) (*models.ReservationRequest, error) {
	return s.Repository.FindReservationRequest(requestKey, roomId)
}
