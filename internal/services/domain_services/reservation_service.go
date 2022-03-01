package domain_services

import (
	"fmt"
	"reservation-api/internal/config"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
	"reservation-api/internal/repositories"
	"reservation-api/internal/utils"
	"reservation-api/pkg/message_broker"
	"time"
)

type ReservationService struct {
	Repository           *repositories.ReservationRepository
	MessageBrokerManager message_broker.MessageBrokerManager
}

// NewReservationService returns new ReservationService
func NewReservationService() *ReservationService {
	return &ReservationService{}
}

// Create creates new Reservation.
func (s *ReservationService) Create(model *models.Reservation) (*models.Reservation, error) {

	result, err := s.Repository.Create(model)
	if err != nil {
		fmt.Println(utils.ToJson(result))
		s.MessageBrokerManager.PublishMessage(config.ReservationQueueName, utils.ToJson(result))
	}
	return result, err
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
func (s *ReservationService) CreateReservationRequest(requestDto *dto.RoomRequestDto) (*models.ReservationRequest, error) {

	return s.Repository.CreateReservationRequest(requestDto)
}

func (s *ReservationService) HasConflict(request *dto.RoomRequestDto, reservation *models.Reservation) (bool, error) {
	return s.Repository.HasConflict(request, reservation)
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
func (s *ReservationService) FindReservationRequest(requestKey string) (*models.ReservationRequest, error) {
	return s.Repository.FindReservationRequest(requestKey)
}

// RemoveExpiredReservationRequests removes expired reservation requests.
func (s *ReservationService) RemoveExpiredReservationRequests() error {
	return s.Repository.RemoveExpiredReservationRequests()
}
