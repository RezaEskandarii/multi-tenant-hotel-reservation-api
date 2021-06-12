package database

import (
	"hotel-reservation/internal/models"
	"hotel-reservation/pkg/application_loger"
)

var (
	entities = []interface{}{
		models.Country{},
		models.City{},
	}
)

// Migrate migrate tables
func Migrate() error {

	application_loger.LogInfo("migration started ...")

	defer func() {
		if r := recover(); r != nil {
			application_loger.LogError(r)
			return
		}
	}()

	db, err := GetDb()

	if err != nil {
		return err
	}

	err = db.AutoMigrate(&models.City{})
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
