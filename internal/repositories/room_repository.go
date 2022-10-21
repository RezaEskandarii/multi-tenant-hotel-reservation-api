package repositories

import (
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
	"reservation-api/pkg/database/tenant_database_resolver"
)

type RoomRepository struct {
	ConnectionResolver *tenant_database_resolver.TenantDatabaseResolver
}

func NewRoomRepository(r *tenant_database_resolver.TenantDatabaseResolver) *RoomRepository {
	return &RoomRepository{r}
}

func (r *RoomRepository) Create(room *models.Room, tenantID uint64) (*models.Room, error) {

	db := r.ConnectionResolver.GetTenantDB(tenantID)

	if tx := db.Create(&room); tx.Error != nil {
		return nil, tx.Error
	}

	return room, nil
}

func (r *RoomRepository) Update(room *models.Room, tenantID uint64) (*models.Room, error) {

	db := r.ConnectionResolver.GetTenantDB(tenantID)

	if tx := db.Updates(&room); tx.Error != nil {
		return nil, tx.Error
	}

	return room, nil
}

func (r *RoomRepository) Find(id uint64, tenantID uint64) (*models.Room, error) {

	model := models.Room{}
	db := r.ConnectionResolver.GetTenantDB(tenantID)

	if tx := db.Where("id=?", id).Find(&model); tx.Error != nil {
		return nil, tx.Error
	}

	if model.Id == 0 {
		return nil, nil
	}

	return &model, nil
}

func (r *RoomRepository) FindAll(input *dto.PaginationFilter) (*commons.PaginatedResult, error) {

	return paginatedList(&models.Room{}, r.ConnectionResolver.GetTenantDB(input.TenantID), input)
}

func (r RoomRepository) Delete(id uint64, tenantID uint64) error {

	db := r.ConnectionResolver.GetTenantDB(tenantID)

	if query := db.Model(&models.Room{}).Where("id=?", id).Delete(&models.Room{}); query.Error != nil {
		return query.Error
	}

	return nil
}
