package domain_services

import (
	"context"
	"reservation-api/internal/models"
	"reservation-api/internal/repositories"
)

type PaymentService struct {
	Repository *repositories.PaymentRepository
}

func NewPaymentService(repository *repositories.PaymentRepository) *PaymentService {
	return &PaymentService{Repository: repository}
}

func (s *PaymentService) Create(ctx context.Context, payment *models.Payment) (*models.Payment, error) {

	return s.Repository.Create(ctx, payment)
}

func (s *PaymentService) Find(ctx context.Context, id uint64) (*models.Payment, error) {

	return s.Repository.Find(ctx, id)
}

func (s *PaymentService) Delete(ctx context.Context, id uint64) error {

	return s.Repository.Delete(ctx, id)
}

func (s *PaymentService) GetListByReservationID(ctx context.Context, reservationID uint64, paymentType *models.PaymentType) ([]*models.Payment, error) {

	return s.Repository.GetListByReservationID(ctx, reservationID, paymentType)
}

func (s *PaymentService) GetBalance(ctx context.Context, reservationID uint64, paymentType *models.PaymentType) (float64, error) {

	return s.Repository.GetBalance(ctx, reservationID, paymentType)
}
