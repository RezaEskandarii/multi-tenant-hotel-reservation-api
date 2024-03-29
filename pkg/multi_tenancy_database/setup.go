package multi_tenancy_database

import (
	"context"
	"gorm.io/gorm"
	"reservation-api/internal/global_variables"
	"reservation-api/internal/models"
	"reservation-api/internal/repositories"
	"reservation-api/internal/services/domain_services"
	"reservation-api/pkg/multi_tenancy_database/tenant_database_resolver"
	"reservation-api/pkg/tenant_dsn_resolver"
	"sync"
)

var (
	resolver = tenant_database_resolver.NewTenantDatabaseResolver()
	service  = domain_services.NewTenantService(&repositories.TenantRepository{
		DbResolver: resolver,
	})
)

// ClientSetUp application multi_tenancy_database
func ClientSetUp() error {

	_, err := service.SetUp(nil, &models.Tenant{
		Name:        "fist tenant",
		Description: "first tenant",
	})

	return err
}

// SetUp application multi_tenancy_database for first time in command line
func SetUp() error {

	tenant := &models.Tenant{
		Name:        "fist tenant",
		Description: "first tenant",
	}
	tenant.Id = global_variables.DefaultTenantID
	_, err := service.SetUp(nil, tenant)

	return err
}

// Migrate migrate tables
func Migrate() error {

	parentCtx := context.Background()
	ctx := context.WithValue(parentCtx, global_variables.TenantIDKey, 0)
	publicDB := resolver.GetTenantDB(ctx).Debug()

	if err := publicDB.AutoMigrate(&models.Tenant{}); err != nil {
		return err
	}

	tenants := make([]models.Tenant, 0)
	wg := sync.WaitGroup{}

	publicDB.FindInBatches(&tenants, 100, func(tx *gorm.DB, batch int) error {
		for _, tenant := range tenants {

			wg.Add(1)

			ctx = context.WithValue(parentCtx, global_variables.TenantIDKey, tenant.Id)
			tenantDB := resolver.GetTenantDB(ctx).Debug()

			go func() {
				for _, entity := range tenant_dsn_resolver.GetEntities() {
					if err := tenantDB.AutoMigrate(entity); err != nil {
						panic(err.Error())
					}
				}

				wg.Done()
			}()

		}
		return nil
	})

	wg.Wait()

	return nil
}
