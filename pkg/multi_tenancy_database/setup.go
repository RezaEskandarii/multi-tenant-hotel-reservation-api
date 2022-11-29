package multi_tenancy_database

import (
	"gorm.io/gorm"
	"reservation-api/internal/models"
	"reservation-api/internal/repositories"
	"reservation-api/internal/services/domain_services"
	"reservation-api/pkg/multi_tenancy_database/tenant_database_resolver"
	"reservation-api/pkg/tenant_connection_string_resolver"
	"sync"
)

var (
	resolver = tenant_database_resolver.NewTenantDatabaseResolver()
	service  = domain_services.NewTenantService(&repositories.TenantRepository{
		DatabaseResolver: resolver,
	})
)

// SetUp application multi_tenancy_database
func SetUp() error {

	_, err := service.SetUp(nil, &models.Tenant{
		Name:        "fist tenant",
		Description: "first tenant",
	})

	return err
}

// Migrate migrate tables
func Migrate() error {

	publicDB := resolver.GetTenantDB(0).Debug()

	if err := publicDB.AutoMigrate(&models.Tenant{}); err != nil {
		return err
	}

	tenants := make([]models.Tenant, 0)
	wg := sync.WaitGroup{}

	publicDB.FindInBatches(&tenants, 100, func(tx *gorm.DB, batch int) error {
		for _, tenant := range tenants {
			wg.Add(1)
			tenantDB := resolver.GetTenantDB(tenant.Id).Debug()

			go func() {
				for _, entity := range tenant_connection_string_resolver.GetEntities() {
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
