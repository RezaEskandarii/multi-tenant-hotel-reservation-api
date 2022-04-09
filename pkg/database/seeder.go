package database

import (
	"gorm.io/gorm"
	"reservation-api/internal/repositories"
	"reservation-api/internal/services/domain_services"
)

// ApplySeed seeds given json file to database.
func ApplySeed(db *gorm.DB) error {

	userService := domain_services.NewUserService(repositories.NewUserRepository(db))
	roomTypeService := domain_services.NewRoomTypeService(repositories.NewRoomTypeRepository(db))
	currencyService := domain_services.NewCurrencyService(repositories.NewCurrencyRepository(db))

	// seed users
	if err := userService.Seed("./data/seed/users.json"); err != nil {
		return err
	}
	// seed roomTypes
	if err := roomTypeService.Seed("./data/seed/room_types.json"); err != nil {
		return err
	}

	// seed currencies
	if err := currencyService.Seed("./data/seed/currencies.json"); err != nil {
		return err
	}

	return nil
}
