package database

import (
	"gorm.io/gorm"
	"hotel-reservation/internal/models"
	"hotel-reservation/pkg/application_loger"
)

var (
	entities = []interface{}{
		models.Country{},
		models.City{},
		models.Province{},
		models.Currency{},
		models.User{},
	}
)

// Migrate migrate tables
func Migrate(db *gorm.DB) error {

	application_loger.LogInfo("migration started ...")

	err := db.AutoMigrate(&models.City{})
	if err != nil {
		application_loger.LogDebug(err.Error())
	}
	db = db.Debug()

	for _, entity := range entities {
		err = db.AutoMigrate(entity)

		if err != nil {
			application_loger.LogError(err.Error())
			return err
		}
	}

	return nil
}
