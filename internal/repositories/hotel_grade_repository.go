package repositories

import (
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
	"reservation-api/pkg/database/connection_resolver"
)

type HotelGradeRepository struct {
	ConnectionResolver *connection_resolver.TenantConnectionResolver
}

func NewHotelGradeRepository(r *connection_resolver.TenantConnectionResolver) *HotelGradeRepository {
	return &HotelGradeRepository{ConnectionResolver: r}
}

func (r *HotelGradeRepository) Create(hotelGrade *models.HotelGrade, tenantID uint64) (*models.HotelGrade, error) {

	db := r.ConnectionResolver.GetDB(tenantID)

	if tx := db.Create(&hotelGrade); tx.Error != nil {
		return nil, tx.Error
	}

	return hotelGrade, nil
}

func (r *HotelGradeRepository) Update(hotelGrade *models.HotelGrade, tenantID uint64) (*models.HotelGrade, error) {

	db := r.ConnectionResolver.GetDB(tenantID)

	if tx := db.Updates(&hotelGrade); tx.Error != nil {
		return nil, tx.Error
	}

	return hotelGrade, nil
}

func (r *HotelGradeRepository) Find(id uint64, tenantID uint64) (*models.HotelGrade, error) {

	model := models.HotelGrade{}
	db := r.ConnectionResolver.GetDB(tenantID)

	if tx := db.Where("id=?", id).Preload("HotelType").Find(&model); tx.Error != nil {
		return nil, tx.Error
	}

	if model.Id == 0 {
		return nil, nil
	}

	return &model, nil
}

func (r *HotelGradeRepository) FindAll(input *dto.PaginationFilter) (*commons.PaginatedResult, error) {

	db := r.ConnectionResolver.GetDB(input.TenantID)
	return paginatedList(&models.HotelGrade{}, db, input)
}

func (r HotelGradeRepository) Delete(id uint64, tenantID uint64) error {

	var count int64 = 0
	db := r.ConnectionResolver.GetDB(tenantID)

	if query := db.Model(&models.Hotel{}).Where(&models.Hotel{HotelGradeId: id}).Count(&count); query.Error != nil {
		return query.Error
	}

	if count > 0 {
		return GradeHasHotel
	}

	if query := db.Model(&models.HotelGrade{}).Where("id=?", id).Delete(&models.HotelGrade{}); query.Error != nil {
		return query.Error
	}

	return nil
}
