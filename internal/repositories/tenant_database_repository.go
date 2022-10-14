package repositories

import (
	"gorm.io/gorm"
	"reservation-api/internal/models"
	"reservation-api/pkg/database/connection_resolver"
)

type TenantRepository struct {
	DB *gorm.DB
}

// NewTenantDatabaseRepository returns new TenantRepository.
func NewTenantDatabaseRepository(db *gorm.DB) *TenantRepository {

	return &TenantRepository{DB: db}
}

func (r *TenantRepository) Create(tenant *models.Tenant) (*models.Tenant, error) {

	if tx := r.DB.Create(&tenant); tx.Error != nil {
		return nil, tx.Error
	}

	resolver := connection_resolver.NewTenantConnectionResolver()
	db1 := resolver.GetDB(0)
	tenantDB := resolver.GetDB(tenant.Id)

	resolver.CreateDbForTenant(db1, tenant.Id)
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
	if tx := r.DB.Where("tenant_id=?", tenantID).Find(&entity); tx.Error != nil {
		return nil, tx.Error
	}
	return &entity, nil
}
