package repositories

import (
	"context"
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
	"reservation-api/pkg/multi_tenancy_database/tenant_database_resolver"
)

type RoomRepository struct {
	DbResolver *tenant_database_resolver.TenantDatabaseResolver
}

func NewRoomRepository(r *tenant_database_resolver.TenantDatabaseResolver) *RoomRepository {
	return &RoomRepository{r}
}

func (r *RoomRepository) Create(ctx context.Context, room *models.Room) (*models.Room, error) {
	db := r.DbResolver.GetTenantDB(ctx)
	if tx := db.Create(&room); tx.Error != nil {
		return nil, tx.Error
	}
	return room, nil
}

func (r *RoomRepository) Update(ctx context.Context, room *models.Room) (*models.Room, error) {

	db := r.DbResolver.GetTenantDB(ctx)

	if tx := db.Updates(&room); tx.Error != nil {
		return nil, tx.Error
	}

	return room, nil
}

func (r *RoomRepository) Find(ctx context.Context, id uint64) (*models.Room, error) {

	model := models.Room{}
	db := r.DbResolver.GetTenantDB(ctx)

	if tx := db.Where("id=?", id).Find(&model); tx.Error != nil {
		return nil, tx.Error
	}

	if model.Id == 0 {
		return nil, nil
	}

	return &model, nil
}

func (r *RoomRepository) FindAll(ctx context.Context, input *dto.PaginationFilter) (*commons.PaginatedResult, error) {

	return paginatedList(&models.Room{}, r.DbResolver.GetTenantDB(ctx), input)
}

func (r RoomRepository) Delete(ctx context.Context, id uint64) error {

	db := r.DbResolver.GetTenantDB(ctx)

	if query := db.Model(&models.Room{}).Where("id=?", id).Delete(&models.Room{}); query.Error != nil {
		return query.Error
	}

	return nil
}
