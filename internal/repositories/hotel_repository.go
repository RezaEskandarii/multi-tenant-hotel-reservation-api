package repositories

import (
	"context"
	"errors"
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/global_variables"
	"reservation-api/internal/message_keys"
	"reservation-api/internal/models"
	"reservation-api/internal/services/common_services"
	"reservation-api/internal/tenant_resolver"
	"reservation-api/pkg/multi_tenancy_database/tenant_database_resolver"
	"sync"
)

type HotelRepository struct {
	ConnectionResolver  *tenant_database_resolver.TenantDatabaseResolver
	FileTransferService common_services.FileTransferer
}

func NewHotelRepository(r *tenant_database_resolver.TenantDatabaseResolver, fileTransferService common_services.FileTransferer) *HotelRepository {

	return &HotelRepository{
		ConnectionResolver:  r,
		FileTransferService: fileTransferService,
	}
}

func (r *HotelRepository) Create(ctx context.Context, hotel *models.Hotel) (*models.Hotel, error) {

	db := r.ConnectionResolver.GetTenantDB(tenant_resolver.GetCurrentTenant(ctx))

	if tx := db.Create(&hotel); tx.Error != nil {
		return nil, tx.Error
	}

	if hotel.Thumbnails != nil && len(hotel.Thumbnails) > 0 {

		var wg sync.WaitGroup
		errorsCh := make(chan error, 0)

		for _, file := range hotel.Thumbnails {
			if file != nil {
				wg.Add(1)
				go func() {
					result, err := r.FileTransferService.Upload(global_variables.HotelsBucketName, "", file, &wg)
					if err != nil {
						errorsCh <- err
						return
					}
					thumbnail := models.Thumbnail{
						VersionID:  result.VersionID,
						HotelId:    hotel.Id,
						BucketName: result.BucketName,
						FileName:   result.FileName,
						FileSize:   result.FileSize,
					}

					if err := db.Create(&thumbnail).Error; err != nil {
						errorsCh <- err
					}
				}()
			}
		}
		select {
		case err := <-errorsCh:
			return nil, err
		default:

		}
		wg.Wait()
		close(errorsCh)
	}

	return hotel, nil
}

func (r *HotelRepository) Update(ctx context.Context, hotel *models.Hotel) (*models.Hotel, error) {

	db := r.ConnectionResolver.GetTenantDB(tenant_resolver.GetCurrentTenant(ctx))

	if tx := db.Updates(&hotel); tx.Error != nil {
		return nil, tx.Error
	}

	return hotel, nil
}

func (r *HotelRepository) Find(ctx context.Context, id uint64) (*models.Hotel, error) {

	model := models.Hotel{}
	db := r.ConnectionResolver.GetTenantDB(tenant_resolver.GetCurrentTenant(ctx))

	if tx := db.Where("id=?", id).Preload("Grades").Find(&model); tx.Error != nil {
		return nil, tx.Error
	}

	if model.Id == 0 {
		return nil, nil
	}

	return &model, nil
}

func (r *HotelRepository) FindAll(ctx context.Context, input *dto.PaginationFilter) (*commons.PaginatedResult, error) {
	db := r.ConnectionResolver.GetTenantDB(tenant_resolver.GetCurrentTenant(ctx))
	return paginatedList(&models.Hotel{}, db, input)
}

func (r HotelRepository) Delete(ctx context.Context, id uint64) error {

	db := r.ConnectionResolver.GetTenantDB(tenant_resolver.GetCurrentTenant(ctx))

	if query := db.Model(&models.Hotel{}).Where("id=?", id).Delete(&models.Hotel{}); query.Error != nil {
		return query.Error
	}

	return nil
}

func (r *HotelRepository) hasRepeatData(ctx context.Context, hotel *models.Hotel) error {

	var countByName int64 = 0
	db := r.ConnectionResolver.GetTenantDB(tenant_resolver.GetCurrentTenant(ctx))

	if tx := *db.Model(&models.Hotel{}).Where(&models.Hotel{Name: hotel.Name}).Count(&countByName); tx.Error != nil {
		return tx.Error
	}

	if countByName > 0 {

		return errors.New(message_keys.HotelRepeatPostalCode)
	}
	return nil
}
