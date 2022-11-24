package repositories

import (
	"context"
	"reservation-api/internal/models"
	"reservation-api/internal/tenant_resolver"
	"reservation-api/pkg/database/tenant_database_resolver"
)

// PaymentRepository
type PaymentRepository struct {
	ConnectionResolver *tenant_database_resolver.TenantDatabaseResolver
}

func NewPaymentRepository(r *tenant_database_resolver.TenantDatabaseResolver) *PaymentRepository {
	return &PaymentRepository{r}
}

func (p *PaymentRepository) Create(ctx context.Context, payment *models.Payment) (*models.Payment, error) {

	db := p.ConnectionResolver.GetTenantDB(tenant_resolver.GetCurrentTenant(ctx))

	if err := db.Create(payment).Error; err != nil {
		return nil, err
	}

	return payment, nil
}

func (p *PaymentRepository) Find(ctx context.Context, id uint64) (*models.Payment, error) {

	db := p.ConnectionResolver.GetTenantDB(tenant_resolver.GetCurrentTenant(ctx))
	result := models.Payment{}

	if err := db.Where("id=?").Find(&result).
		Preload("Reservation").Preload("Payer").Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (p *PaymentRepository) Delete(ctx context.Context, id uint64) error {

	db := p.ConnectionResolver.GetTenantDB(tenant_resolver.GetCurrentTenant(ctx))

	if err := db.Model(&models.Payment{}).Where("id=?", id).Delete(&models.Payment{}).Error; err != nil {
		return err
	}

	return nil
}

func (p *PaymentRepository) GetListByReservationID(ctx context.Context, reservationID uint64, paymentType *models.PaymentType) ([]*models.Payment, error) {

	result := make([]*models.Payment, 0)
	db := p.ConnectionResolver.GetTenantDB(tenant_resolver.GetCurrentTenant(ctx))
	query := db.Model(&models.Payment{}).Where("reservation_id=?", reservationID)

	if paymentType != nil {
		query = query.Where("payment_type=?", paymentType)
	}

	if err := query.Scan(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (p *PaymentRepository) GetBalance(ctx context.Context, reservationID uint64, paymentType *models.PaymentType) (float64, error) {

	var result float64
	db := p.ConnectionResolver.GetTenantDB(tenant_resolver.GetCurrentTenant(ctx))

	query := db.Model(&models.Payment{}).Where("reservation_id=?", reservationID)

	if paymentType != nil {
		query = query.Where("payment_type=?", paymentType)
	}

	if err := query.Select("SUM(amount)").Scan(&result).Error; err != nil {
		return 0, err
	}

	return result, nil
}
