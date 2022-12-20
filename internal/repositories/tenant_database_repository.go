package repositories

import (
	"context"
	"reservation-api/internal/models"
	"reservation-api/pkg/multi_tenancy_database/tenant_database_resolver"
)

type TenantRepository struct {
	DatabaseResolver *tenant_database_resolver.TenantDatabaseResolver
}

// NewTenantDatabaseRepository returns new TenantRepository.
func NewTenantDatabaseRepository(resolver *tenant_database_resolver.TenantDatabaseResolver) *TenantRepository {

	return &TenantRepository{DatabaseResolver: resolver}
}

func (r *TenantRepository) Create(ctx context.Context, tenant *models.Tenant) (*models.Tenant, error) {

	publicDB := r.DatabaseResolver.GetTenantDB(0).Debug()

	if err := publicDB.AutoMigrate(&models.Tenant{}); err != nil {
		return nil, err
	}

	// check if tenant exists
	// this code uses for prevent create tenant in restart application
	// because create tenant command calls in docker cmd
	if tenant.Id != 0 {
		var count int64 = 0
		if err := publicDB.Model(&models.Tenant{}).Where("id=?", tenant.Id).Count(&count).Error; err != nil {
			return nil, err
		}
		if count > 0 {
			return nil, nil
		}
	}
	if tx := publicDB.Create(&tenant); tx.Error != nil {
		return nil, tx.Error
	}

	ctx = context.WithValue(context.Background(), "TenantID", tenant.Id)

	resolver := tenant_database_resolver.NewTenantDatabaseResolver()
	tenantDB := resolver.GetTenantDB(tenant.Id).Debug()

	resolver.CreateDbForTenant(publicDB, tenant.Id)
	resolver.Migrate(tenantDB, tenant.Id)

	userRepository := NewUserRepository(resolver)
	roomTypeRepository := NewRoomTypeRepository(resolver)
	currencyRepository := NewCurrencyRepository(resolver)

	// seed users
	if err := userRepository.Seed(ctx, "./data/seed/users.json"); err != nil {
		panic(err)
	}
	// seed roomTypes
	if err := roomTypeRepository.Seed(ctx, "./data/seed/room_types.json"); err != nil {
		panic(err)
	}

	// seed currencies
	if err := currencyRepository.Seed(ctx, "./data/seed/currencies.json"); err != nil {
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
