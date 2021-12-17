package database

import (
	"gorm.io/gorm"
	"reservation-api/internal/models"
	"reservation-api/pkg/applogger"
)

var (
	entities = []interface{}{
		models.Country{},
		models.City{},
		models.Province{},
		models.Currency{},
		models.User{},
		models.Residence{},
		models.Room{},
		models.RoomType{},
		models.Guest{},
		models.RateGroup{},
		models.RateCode{},
		models.ResidenceGrade{},
		models.ResidenceType{},
		models.ReservationRequest{},
		models.Reservation{},
		models.Audit{},
	}
)

// Migrate migrate tables
func Migrate(db *gorm.DB) error {

	logger := applogger.New(nil)
	logger.LogInfo("migration started ...")

	err := db.AutoMigrate(&models.City{})
	if err != nil {
		logger.LogDebug(err.Error())
	}

	for _, entity := range entities {

		err = db.AutoMigrate(entity)

		if err != nil {
			logger.LogError(err.Error())
			return err
		}
	}

	return nil
}
