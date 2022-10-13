package domain_services

import (
	"reservation-api/internal/commons"
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
func NewReservationService(repository *repositories.ReservationRepository,
	messageBroker message_broker.MessageBrokerManager) *ReservationService {
	return &ReservationService{
		Repository:           repository,
		MessageBrokerManager: messageBroker,
	}
}

// Create creates new Reservation.
func (s *ReservationService) Create(model *models.Reservation, tenantID uint64) (*models.Reservation, error) {

	result, err := s.Repository.Create(model, tenantID)
	if err != nil {
		s.MessageBrokerManager.PublishMessage(config.ReservationQueueName, utils.ToJson(model))
	}
	return result, nil
}

// ChangeStatus changes the reservation check status.
func (s *ReservationService) ChangeStatus(id uint64, tenantID uint64, status models.ReservationCheckStatus) (*models.Reservation, error) {

	return s.Repository.ChangeStatus(id, tenantID, status)
}

// Update updates Reservation.
func (s *ReservationService) Update(id uint64, tenantID uint64, model *models.Reservation) (*models.Reservation, error) {

	return s.Repository.Update(id, tenantID, model)
}

// CreateReservationRequest creates reservation request for given room to prevent concurrent request for specific room.
func (s *ReservationService) CreateReservationRequest(requestDto *dto.RoomRequestDto, tenantID uint64) (*models.ReservationRequest, error) {

	return s.Repository.CreateReservationRequest(requestDto, tenantID)
}

func (s *ReservationService) HasConflict(request *dto.RoomRequestDto, reservation *models.Reservation, tenantID uint64) (bool, error) {
	return s.Repository.HasConflict(request, reservation, tenantID)
}

func (s *ReservationService) HasReservationConflict(checkInDate *time.Time, checkOutDate *time.Time, roomId uint64, tenantID uint64) (bool, error) {
	return s.Repository.HasReservationConflict(checkInDate, checkOutDate, roomId, tenantID)
}

// RemoveReservationRequest this function remove reservation request bt given requestKey param.
func (s *ReservationService) RemoveReservationRequest(requestKey string, tenantID uint64) error {
	return s.Repository.DeleteReservationRequest(requestKey, tenantID)
}

// GetRecommendedRateCodes returns list of recommended rateCodeDetails price per reservation condition.
func (s *ReservationService) GetRecommendedRateCodes(priceDto *dto.GetRatePriceDto, tenantID uint64) ([]*dto.RateCodePricesDto, error) {
	return s.Repository.GetRecommendedRateCodes(priceDto, tenantID)
}

// Find find and returns reservation by id.
func (s *ReservationService) Find(id uint64, tenantID uint64) (*models.Reservation, error) {
	return s.Repository.Find(id, tenantID)
}

// FindReservationRequest find and returns reservationRequest by  given roomId and requestKey.
func (s *ReservationService) FindReservationRequest(requestKey string, tenantID uint64) (*models.ReservationRequest, error) {
	return s.Repository.FindReservationRequest(requestKey, tenantID)
}

// RemoveExpiredReservationRequests removes expired reservation requests.
func (s *ReservationService) RemoveExpiredReservationRequests(tenantID uint64) error {
	return s.Repository.RemoveExpiredReservationRequests(tenantID)
}

// FindAll returns paginated list of reservations with some filters like
// guestname, fromDate,toDate etc.
func (s *ReservationService) FindAll(filter *dto.ReservationFilter) (error, *commons.PaginatedResult) {
	return s.Repository.FindAll(filter)
}
