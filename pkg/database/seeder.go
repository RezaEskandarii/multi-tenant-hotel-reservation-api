package database

import (
	"reservation-api/internal/repositories"
	"reservation-api/internal/services/domain_services"
	"reservation-api/pkg/database/connection_resolver"
)

// ApplySeed seeds given json file to database.
func ApplySeed(r *connection_resolver.TenantConnectionResolver, tenantID uint64) error {

	userService := domain_services.NewUserService(repositories.NewUserRepository(r))
	roomTypeService := domain_services.NewRoomTypeService(repositories.NewRoomTypeRepository(r))
	currencyService := domain_services.NewCurrencyService(repositories.NewCurrencyRepository(r))

	// seed users
	if err := userService.Seed("./data/seed/users.json", tenantID); err != nil {
		return err
	}
	// seed roomTypes
	if err := roomTypeService.Seed("./data/seed/room_types.json", tenantID); err != nil {
		return err
	}

	// seed currencies
	if err := currencyService.Seed("./data/seed/currencies.json", tenantID); err != nil {
		return err
	}

	return nil
}
