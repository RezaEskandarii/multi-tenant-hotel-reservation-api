package database

import (
	"gorm.io/gorm"
	"hotel-reservation/internal/models"
	"hotel-reservation/pkg/applogger"
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
	}
)

// Migrate migrate tables
func Migrate(db *gorm.DB) error {

	applogger.LogInfo("migration started ...")

	err := db.AutoMigrate(&models.City{})
	if err != nil {
		applogger.LogDebug(err.Error())
	}

	for _, entity := range entities {

		err = db.AutoMigrate(entity)

		if err != nil {
			applogger.LogError(err.Error())
			return err
		}
	}

	return nil
}
