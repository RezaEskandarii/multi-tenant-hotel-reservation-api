package repositories

import (
	"gorm.io/gorm"
	"reservation-api/internal/models"
)

// PaymentRepository
type PaymentRepository struct {
	DB *gorm.DB
}

func NewPaymentRepository(DB *gorm.DB) *PaymentRepository {
	return &PaymentRepository{DB: DB}
}

func (p *PaymentRepository) Create(payment *models.Payment) (*models.Payment, error) {

	if err := p.DB.Create(payment).Error; err != nil {
		return nil, err
	}
	return payment, nil
}

func (p *PaymentRepository) Find(id uint64) (*models.Payment, error) {

	result := models.Payment{}
	if err := p.DB.Where("id=?").Find(&result).
		Preload("Reservation").Preload("Payer").Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (p *PaymentRepository) Delete(id uint64) error {

	if err := p.DB.Model(&models.Payment{}).Where("id=?", id).Delete(&models.Payment{}).Error; err != nil {
		return err
	}
	return nil
}

func (p *PaymentRepository) GetListByReservationID(reservationID uint64, paymentType *models.PaymentType) ([]*models.Payment, error) {

	result := make([]*models.Payment, 0)
	query := p.DB.Model(&models.Payment{}).Where("reservation_id=?", reservationID)
	if paymentType != nil {
		query = query.Where("payment_type=?", paymentType)
	}
	if err := query.Scan(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (p *PaymentRepository) GetBalance(reservationID uint64, paymentType *models.PaymentType) (float64, error) {

	var result float64
	query := p.DB.Model(&models.Payment{}).Where("reservation_id=?", reservationID)

	if paymentType != nil {
		query = query.Where("payment_type=?", paymentType)
	}

	if err := query.Select("SUM(amount)").Scan(&result).Error; err != nil {
		return 0, err
	}
	return result, nil
}
