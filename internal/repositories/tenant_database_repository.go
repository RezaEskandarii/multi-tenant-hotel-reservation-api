package repositories

import (
	"reservation-api/internal/models"
	"reservation-api/pkg/database/tenant_database_resolver"
)

type TenantRepository struct {
	DatabaseResolver *tenant_database_resolver.TenantDatabaseResolver
}

// NewTenantDatabaseRepository returns new TenantRepository.
func NewTenantDatabaseRepository(resolver *tenant_database_resolver.TenantDatabaseResolver) *TenantRepository {

	return &TenantRepository{DatabaseResolver: resolver}
}

func (r *TenantRepository) Create(tenant *models.Tenant) (*models.Tenant, error) {

	publicDB := r.DatabaseResolver.GetTenantDB(0).Debug()

	if err := publicDB.AutoMigrate(&models.Tenant{}); err != nil {
		return nil, err
	}

	if tx := publicDB.Create(&tenant); tx.Error != nil {
		return nil, tx.Error
	}

	resolver := tenant_database_resolver.NewTenantDatabaseResolver()
	tenantDB := resolver.GetTenantDB(tenant.Id).Debug()

	resolver.CreateDbForTenant(publicDB, tenant.Id)
	resolver.Migrate(tenantDB, tenant.Id)

	userRepository := NewUserRepository(resolver)
	roomTypeRepository := NewRoomTypeRepository(resolver)
	currencyRepository := NewCurrencyRepository(resolver)

	// seed users
	if err := userRepository.Seed("./data/seed/users.json", tenant.Id); err != nil {
		panic(err)
	}
	// seed roomTypes
	if err := roomTypeRepository.Seed("./data/seed/room_types.json", tenant.Id); err != nil {
		panic(err)
	}

	// seed currencies
	if err := currencyRepository.Seed("./data/seed/currencies.json", tenant.Id); err != nil {
		panic(err)
	}

	return tenant, nil
}

func (r *TenantRepository) FindByTenantID(tenantID uint64) (*models.Tenant, error) {

	entity := models.Tenant{}
	db := r.DatabaseResolver.GetTenantDB(tenantID)

	if tx := db.Where("tenant_id=?", tenantID).Find(&entity); tx.Error != nil {
		return nil, tx.Error
	}
	return &entity, nil
}
