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

func (s *PaymentService) Create(payment *models.Payment, tenantID uint64) (*models.Payment, error) {

	return s.Repository.Create(payment, tenantID)
}

func (s *PaymentService) Find(id uint64, tenantID uint64) (*models.Payment, error) {

	return s.Repository.Find(id, tenantID)
}

func (s *PaymentService) Delete(id uint64, tenantID uint64) error {

	return s.Repository.Delete(id, tenantID)
}

func (s *PaymentService) GetListByReservationID(reservationID uint64, paymentType *models.PaymentType, tenantID uint64) ([]*models.Payment, error) {

	return s.Repository.GetListByReservationID(reservationID, paymentType, tenantID)
}

func (s *PaymentService) GetBalance(reservationID uint64, paymentType *models.PaymentType, tenantID uint64) (float64, error) {

	return s.Repository.GetBalance(reservationID, paymentType, tenantID)
}
