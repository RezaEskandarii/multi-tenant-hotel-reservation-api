package repositories

import (
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
	"reservation-api/pkg/database/connection_resolver"
)

type GuestRepository struct {
	ConnectionResolver *connection_resolver.TenantConnectionResolver
}

func NewGuestRepository(r *connection_resolver.TenantConnectionResolver) *GuestRepository {
	return &GuestRepository{
		ConnectionResolver: r,
	}
}

func (r *GuestRepository) Create(guest *models.Guest, tenantID uint64) (*models.Guest, error) {

	db := r.ConnectionResolver.GetDB(tenantID)

	if tx := db.Create(&guest); tx.Error != nil {
		return nil, tx.Error
	}

	return guest, nil
}

func (r *GuestRepository) Update(guest *models.Guest, tenantID uint64) (*models.Guest, error) {

	db := r.ConnectionResolver.GetDB(tenantID)

	if tx := db.Updates(&guest); tx.Error != nil {
		return nil, tx.Error
	}

	return guest, nil
}

func (r *GuestRepository) Find(id uint64, tenantID uint64) (*models.Guest, error) {

	model := models.Guest{}
	db := r.ConnectionResolver.GetDB(tenantID)

	if tx := db.Where("id=?", id).Find(&model); tx.Error != nil {
		return nil, tx.Error
	}

	if model.Id == 0 {
		return nil, nil
	}

	return &model, nil
}

func (r *GuestRepository) FindByNationalId(id string, tenantID uint64) (*models.Guest, error) {

	model := models.Guest{}
	db := r.ConnectionResolver.GetDB(tenantID)

	if tx := db.Where(models.Guest{NationalId: id}).Find(&model); tx.Error != nil {
		return nil, tx.Error
	}

	if model.Id == 0 {
		return nil, nil
	}

	return &model, nil
}

func (r *GuestRepository) FindByPassportNumber(passNumber string, tenantID uint64) (*models.Guest, error) {

	model := models.Guest{}
	db := r.ConnectionResolver.GetDB(tenantID)

	if tx := db.Where(models.Guest{PassportNumber: passNumber}).Find(&model); tx.Error != nil {
		return nil, tx.Error
	}

	if model.Id == 0 {
		return nil, nil
	}

	return &model, nil
}

func (r *GuestRepository) ReservationsCount(guestId uint64, tenantID uint64) (error, uint64) {

	var count int64 = 0
	db := r.ConnectionResolver.GetDB(tenantID)

	if err := db.Model(&models.Reservation{}).Where("supervisor_id=?", guestId).Count(&count).Error; err != nil {
		return err, 0
	}

	return nil, uint64(count)
}

func (r *GuestRepository) FindAll(input *dto.PaginationFilter) (*commons.PaginatedResult, error) {

	db := r.ConnectionResolver.GetDB(input.TenantID)
	return paginatedList(&models.Guest{}, db, input)
}
