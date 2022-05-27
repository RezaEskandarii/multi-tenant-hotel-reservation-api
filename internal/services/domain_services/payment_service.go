package domain_services

import (
	"reservation-api/internal/models"
	"reservation-api/internal/repositories"
)

type PaymentService struct {
	Repository *repositories.PaymentRepository
}

func NewPaymentService(repository *repositories.PaymentRepository) *PaymentService {
	return &PaymentService{Repository: repository}
}

func (s *PaymentService) Create(payment *models.Payment) (*models.Payment, error) {

	return s.Repository.Create(payment)
}

func (s *PaymentService) Find(id uint64) (*models.Payment, error) {

	return s.Repository.Find(id)
}

func (s *PaymentService) Delete(id uint64) error {

	return s.Repository.Delete(id)
}

func (s *PaymentService) GetListByReservationID(reservationID uint64, paymentType *models.PaymentType) ([]*models.Payment, error) {

	return s.Repository.GetListByReservationID(reservationID, paymentType)
}

func (s *PaymentService) GetBalance(reservationID uint64, paymentType *models.PaymentType) (float64, error) {

	return s.Repository.GetBalance(reservationID, paymentType)
}
