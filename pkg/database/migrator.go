package database

import (
	"gorm.io/gorm"
	"reservation-api/internal/models"
	"reservation-api/pkg/applogger"
)

var (
	entities = []interface{}{
		models.Tenant{},
		models.Country{},
		models.City{},
		models.Province{},
		models.Currency{},
		models.User{},
		models.Hotel{},
		models.Room{},
		models.RoomType{},
		models.Guest{},
		models.RateGroup{},
		models.RateCode{},
		models.HotelGrade{},
		models.HotelType{},
		models.ReservationRequest{},
		models.Reservation{},
		models.Audit{},
		models.RateCodeDetail{},
		models.RateCodeDetailPrice{},
		models.Sharer{},
		models.Thumbnail{},
	}
)

// Migrate migrate tables
func Migrate(db *gorm.DB) error {
	logger := applogger.New(nil)
	logger.LogInfo("migration started ...")

	for _, entity := range entities {
		err := db.Debug().AutoMigrate(entity)

		if err != nil {
			logger.LogError(err.Error())
			return err
		}
	}

	return nil
}
