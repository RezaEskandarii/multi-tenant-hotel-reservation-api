package repositories

import (
	"errors"
	"gorm.io/gorm"
	"reservation-api/internal/commons"
	"reservation-api/internal/config"
	"reservation-api/internal/dto"
	"reservation-api/internal/message_keys"
	"reservation-api/internal/models"
	"reservation-api/internal/services/common_services"
	"sync"
)

type HotelRepository struct {
	DB                  *gorm.DB
	FileTransferService common_services.FileTransferer
}

func NewHotelRepository(db *gorm.DB, fileTransferService common_services.FileTransferer) *HotelRepository {

	return &HotelRepository{
		DB:                  db,
		FileTransferService: fileTransferService,
	}
}

func (r *HotelRepository) Create(hotel *models.Hotel) (*models.Hotel, error) {

	if tx := r.DB.Create(&hotel); tx.Error != nil {

		return nil, tx.Error
	}

	if hotel.Thumbnails != nil && len(hotel.Thumbnails) > 0 {

		var wg sync.WaitGroup
		errorsCh := make(chan error, 0)

		for _, thumbnail := range hotel.Thumbnails {
			if thumbnail.File != nil {
				wg.Add(1)
				go func() {
					result, err := r.FileTransferService.Upload(config.HotelsBucketName, "", thumbnail.File, &wg)
					if err != nil {
						errorsCh <- err
						return
					}

					thumbnail.VersionID = result.VersionID
					thumbnail.HotelId = hotel.Id
					thumbnail.BucketName = result.BucketName
					thumbnail.FileName = result.FileName
					thumbnail.FileSize = result.FileSize

					if err := r.DB.Create(&thumbnail).Error; err != nil {
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

func (r *HotelRepository) Update(hotel *models.Hotel) (*models.Hotel, error) {

	if tx := r.DB.Updates(&hotel); tx.Error != nil {

		return nil, tx.Error
	}

	return hotel, nil
}

func (r *HotelRepository) Find(id uint64) (*models.Hotel, error) {
	model := models.Hotel{}

	if tx := r.DB.Where("id=?", id).Preload("Grades").Find(&model); tx.Error != nil {

		return nil, tx.Error
	}

	if model.Id == 0 {
		return nil, nil
	}

	return &model, nil
}

func (r *HotelRepository) FindAll(input *dto.PaginationInput) (*commons.PaginatedList, error) {

	return finAll(&models.Hotel{}, r.DB, input)
}

func (r HotelRepository) Delete(id uint64) error {
	if query := r.DB.Model(&models.Hotel{}).Where("id=?", id).Delete(&models.Hotel{}); query.Error != nil {
		return query.Error
	}

	return nil
}

func (r *HotelRepository) hasRepeatData(hotel *models.Hotel) error {
	var countByName int64 = 0

	if tx := *r.DB.Model(&models.Hotel{}).Where(&models.Hotel{Name: hotel.Name}).Count(&countByName); tx.Error != nil {
		return tx.Error
	}

	if countByName > 0 {

		return errors.New(message_keys.HotelRepeatPostalCode)
	}
	return nil
}
