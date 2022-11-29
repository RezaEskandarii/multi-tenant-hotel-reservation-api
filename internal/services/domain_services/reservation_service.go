package domain_services

import (
	"context"
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/global_variables"
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
func (s *ReservationService) Create(ctx context.Context, model *models.Reservation) (*models.Reservation, error) {

	result, err := s.Repository.Create(ctx, model)
	if err != nil {
		s.MessageBrokerManager.PublishMessage(global_variables.ReservationQueueName, utils.ToJson(model))
	}
	return result, nil
}

// ChangeStatus changes the reservation check status.
func (s *ReservationService) ChangeStatus(ctx context.Context, id uint64, status models.ReservationCheckStatus) (*models.Reservation, error) {

	return s.Repository.ChangeStatus(ctx, id, status)
}

// Update updates Reservation.
func (s *ReservationService) Update(ctx context.Context, id uint64, model *models.Reservation) (*models.Reservation, error) {

	return s.Repository.Update(ctx, id, model)
}

// CreateReservationRequest creates reservation request for given room to prevent concurrent request for specific room.
func (s *ReservationService) CreateReservationRequest(ctx context.Context, requestDto *dto.RoomRequestDto) (*models.ReservationRequest, error) {

	return s.Repository.CreateReservationRequest(ctx, requestDto)
}

func (s *ReservationService) HasConflict(ctx context.Context, request *dto.RoomRequestDto, reservation *models.Reservation) (bool, error) {
	return s.Repository.HasConflict(ctx, request, reservation)
}

func (s *ReservationService) HasReservationConflict(ctx context.Context, checkInDate *time.Time, checkOutDate *time.Time, roomId uint64) (bool, error) {

	return s.Repository.HasReservationConflict(ctx, checkInDate, checkOutDate, roomId)
}

// RemoveReservationRequest this function remove reservation request bt given requestKey param.
func (s *ReservationService) RemoveReservationRequest(ctx context.Context, requestKey string) error {
	return s.Repository.DeleteReservationRequest(ctx, requestKey)
}

// GetRecommendedRateCodes returns list of recommended rateCodeDetails price per reservation condition.
func (s *ReservationService) GetRecommendedRateCodes(ctx context.Context, priceDto *dto.GetRatePriceDto) ([]*dto.RateCodePricesDto, error) {
	return s.Repository.GetRecommendedRateCodes(ctx, priceDto)
}

// Find find and returns reservation by id.
func (s *ReservationService) Find(ctx context.Context, id uint64) (*models.Reservation, error) {
	return s.Repository.Find(ctx, id)
}

// FindReservationRequest find and returns reservationRequest by  given roomId and requestKey.
func (s *ReservationService) FindReservationRequest(ctx context.Context, requestKey string) (*models.ReservationRequest, error) {
	return s.Repository.FindReservationRequest(ctx, requestKey)
}

// RemoveExpiredReservationRequests removes expired reservation requests.
func (s *ReservationService) RemoveExpiredReservationRequests(ctx context.Context) error {
	return s.Repository.RemoveExpiredReservationRequests(ctx)
}

// FindAll returns paginated list of reservations with some filters like
// guestname, fromDate,toDate etc.
func (s *ReservationService) FindAll(ctx context.Context, filter *dto.ReservationFilter) (error, *commons.PaginatedResult) {
	return s.Repository.FindAll(ctx, filter)
}
