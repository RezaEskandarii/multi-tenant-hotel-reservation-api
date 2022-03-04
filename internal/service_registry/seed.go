package service_registry

import "gorm.io/gorm"

// ApplySeed seeds given json file to database.
func ApplySeed(db *gorm.DB) error {

	setServicesDependencies(db, nil, nil, true)

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
