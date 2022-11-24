package repositories

import (
	"context"
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
	"reservation-api/internal/tenant_resolver"
	"reservation-api/pkg/database/tenant_database_resolver"
)

type GuestRepository struct {
	ConnectionResolver *tenant_database_resolver.TenantDatabaseResolver
}

func NewGuestRepository(r *tenant_database_resolver.TenantDatabaseResolver) *GuestRepository {
	return &GuestRepository{
		ConnectionResolver: r,
	}
}

func (r *GuestRepository) Create(ctx context.Context, guest *models.Guest) (*models.Guest, error) {

	db := r.ConnectionResolver.GetTenantDB(tenant_resolver.GetCurrentTenant(ctx))

	if tx := db.Create(&guest); tx.Error != nil {
		return nil, tx.Error
	}

	return guest, nil
}

func (r *GuestRepository) Update(ctx context.Context, guest *models.Guest) (*models.Guest, error) {

	db := r.ConnectionResolver.GetTenantDB(tenant_resolver.GetCurrentTenant(ctx))

	if tx := db.Updates(&guest); tx.Error != nil {
		return nil, tx.Error
	}

	return guest, nil
}

func (r *GuestRepository) Find(ctx context.Context, id uint64) (*models.Guest, error) {

	model := models.Guest{}
	db := r.ConnectionResolver.GetTenantDB(tenant_resolver.GetCurrentTenant(ctx))

	if tx := db.Where("id=?", id).Find(&model); tx.Error != nil {
		return nil, tx.Error
	}

	if model.Id == 0 {
		return nil, nil
	}

	return &model, nil
}

func (r *GuestRepository) FindByNationalId(ctx context.Context, id string) (*models.Guest, error) {

	model := models.Guest{}
	db := r.ConnectionResolver.GetTenantDB(tenant_resolver.GetCurrentTenant(ctx))

	if tx := db.Where(models.Guest{NationalId: id}).Find(&model); tx.Error != nil {
		return nil, tx.Error
	}

	if model.Id == 0 {
		return nil, nil
	}

	return &model, nil
}

func (r *GuestRepository) FindByPassportNumber(ctx context.Context, passNumber string) (*models.Guest, error) {

	model := models.Guest{}
	db := r.ConnectionResolver.GetTenantDB(tenant_resolver.GetCurrentTenant(ctx))

	if tx := db.Where(models.Guest{PassportNumber: passNumber}).Find(&model); tx.Error != nil {
		return nil, tx.Error
	}

	if model.Id == 0 {
		return nil, nil
	}

	return &model, nil
}

func (r *GuestRepository) ReservationsCount(ctx context.Context, guestId uint64) (error, uint64) {

	var count int64 = 0
	db := r.ConnectionResolver.GetTenantDB(tenant_resolver.GetCurrentTenant(ctx))

	if err := db.Model(&models.Reservation{}).Where("supervisor_id=?", guestId).Count(&count).Error; err != nil {
		return err, 0
	}

	return nil, uint64(count)
}

func (r *GuestRepository) FindAll(ctx context.Context, input *dto.PaginationFilter) (*commons.PaginatedResult, error) {

	db := r.ConnectionResolver.GetTenantDB(tenant_resolver.GetCurrentTenant(ctx))
	return paginatedList(&models.Guest{}, db, input)
}
