package repositories

import (
	"gorm.io/gorm"
	"hotel-reservation/internal/commons"
	"hotel-reservation/internal/dto"
	"hotel-reservation/internal/models"
)

type GuestRepository struct {
	DB *gorm.DB
}

func NewGuestRepository(db *gorm.DB) *GuestRepository {
	return &GuestRepository{
		DB: db,
	}
}

func (r *GuestRepository) Create(guest *models.Guest) (*models.Guest, error) {

	if tx := r.DB.Create(&guest); tx.Error != nil {
		return nil, tx.Error
	}

	return guest, nil
}

func (r *GuestRepository) Update(guest *models.Guest) (*models.Guest, error) {

	if tx := r.DB.Updates(&guest); tx.Error != nil {

		return nil, tx.Error
	}

	return guest, nil
}

func (r *GuestRepository) Find(id uint64) (*models.Guest, error) {

	model := models.Guest{}
	if tx := r.DB.Where("id=?", id).Find(&model); tx.Error != nil {

		return nil, tx.Error
	}

	if model.Id == 0 {
		return nil, nil
	}

	return &model, nil
}

func (r *GuestRepository) FindByNationalId(id string) (*models.Guest, error) {

	model := models.Guest{}
	if tx := r.DB.Where(models.Guest{NationalId: id}).Find(&model); tx.Error != nil {

		return nil, tx.Error
	}

	if model.Id == 0 {
		return nil, nil
	}

	return &model, nil
}

func (r *GuestRepository) FindByPassportNumber(passNumber string) (*models.Guest, error) {

	model := models.Guest{}
	if tx := r.DB.Where(models.Guest{PassportNumber: passNumber}).Find(&model); tx.Error != nil {

		return nil, tx.Error
	}

	if model.Id == 0 {
		return nil, nil
	}

	return &model, nil
}

func (r *GuestRepository) ReservationsCount(guestId uint64) (error, uint64) {
	panic("not implemented")
}

func (r *GuestRepository) CheckIn(guestId uint64) error {
	panic("not implemented")
}

func (r *GuestRepository) CheckOut(guestId uint64) error {
	panic("not implemented")
}

func (r *GuestRepository) FindAll(input *dto.PaginationInput) (*commons.PaginatedList, error) {

	return finAll(&models.Guest{}, r.DB, input)
}
