package repositories

import (
	"context"
	"reservation-api/internal/models"
	"reservation-api/pkg/multi_tenancy_database/tenant_database_resolver"
)

type TenantRepository struct {
	DbResolver *tenant_database_resolver.TenantDatabaseResolver
}

// NewTenantDatabaseRepository returns new TenantRepository.
func NewTenantDatabaseRepository(resolver *tenant_database_resolver.TenantDatabaseResolver) *TenantRepository {

	return &TenantRepository{DbResolver: resolver}
}

func (r *TenantRepository) Create(ctx context.Context, tenant *models.Tenant) (*models.Tenant, error) {

	publicDB := r.DbResolver.GetTenantDB(ctx).Debug()

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
	tenantDB := resolver.GetTenantDB(ctx).Debug()

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

func (r *TenantRepository) GetAll() ([]models.Tenant, error) {

	db := r.DbResolver.GetDefaultDB()
	var tenantsCount int64 = 0
	if err := db.Model(&models.Tenant{}).Count(&tenantsCount).Error; err != nil {
		return nil, err
	}

	tenants := make([]models.Tenant, tenantsCount)
	if err := db.Model(&models.Tenant{}).Scan(&tenants).Error; err != nil {
		return nil, err
	}
	return tenants, nil
}
