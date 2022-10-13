package repositories

import (
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
	"reservation-api/pkg/database/connection_resolver"
)

type HotelTypeRepository struct {
	ConnectionResolver *connection_resolver.TenantConnectionResolver
}

func NewHotelTypeRepository(r *connection_resolver.TenantConnectionResolver) *HotelTypeRepository {
	return &HotelTypeRepository{r}
}

func (r *HotelTypeRepository) Create(hotelType *models.HotelType, tenantID uint64) (*models.HotelType, error) {

	db := r.ConnectionResolver.GetDB(tenantID)

	if tx := db.Create(&hotelType); tx.Error != nil {
		return nil, tx.Error
	}

	return hotelType, nil
}

func (r *HotelTypeRepository) Update(hotelType *models.HotelType, tenantID uint64) (*models.HotelType, error) {

	db := r.ConnectionResolver.GetDB(tenantID)

	if tx := db.Updates(&hotelType); tx.Error != nil {
		return nil, tx.Error
	}

	return hotelType, nil
}

func (r *HotelTypeRepository) Find(id uint64, tenantID uint64) (*models.HotelType, error) {

	model := models.HotelType{}
	db := r.ConnectionResolver.GetDB(tenantID)

	if tx := db.Where("id=?", id).Preload("Grades").Find(&model); tx.Error != nil {
		return nil, tx.Error
	}

	if model.Id == 0 {
		return nil, nil
	}

	return &model, nil
}

func (r *HotelTypeRepository) FindAll(input *dto.PaginationFilter) (*commons.PaginatedResult, error) {

	db := r.ConnectionResolver.GetDB(input.TenantID)
	return paginatedList(&models.HotelType{}, db, input)
}

func (r HotelTypeRepository) Delete(id uint64, tenantID uint64) error {

	var count int64 = 0
	db := r.ConnectionResolver.GetDB(tenantID)

	if query := db.Model(&models.Hotel{}).Where(&models.Hotel{HotelTypeId: id}).Count(&count); query.Error != nil {
		return query.Error
	}

	if count > 0 {
		return TypeHasHotelError
	}

	if query := db.Model(&models.HotelType{}).Where("id=?", id).Delete(&models.HotelType{}); query.Error != nil {
		return query.Error

	} else {
		query = db.Model(&models.HotelGrade{}).Where(&models.HotelGrade{HotelTypeId: id}).Delete(&models.HotelGrade{})
		if query.Error != nil {

			return query.Error
		}
	}

	return nil
}
