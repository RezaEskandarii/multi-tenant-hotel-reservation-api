// Package tenant_database_resolver
// this package provides special connection string and multi_tenancy_database object per given tenant ID.
///**/
package tenant_database_resolver

import (
	"fmt"
	"gorm.io/gorm"
	"reservation-api/pkg/tenant_connection_string_resolver"
	"sync"
)

// TenantDatabaseResolver resolves multi tenancy connection strings.
type TenantDatabaseResolver struct {
	cache map[uint64]*gorm.DB
	Mutex sync.Mutex
}

// NewTenantDatabaseResolver returns new instance of TenantDatabaseResolver
func NewTenantDatabaseResolver() *TenantDatabaseResolver {
	return &TenantDatabaseResolver{
		cache: make(map[uint64]*gorm.DB),
	}
}

// GetTenantDB returns unique gorm DB object per given tenantID
// multi tenancy policy is unique multi_tenancy_database per Tenant
func (c *TenantDatabaseResolver) GetTenantDB(tenantID uint64) *gorm.DB {

	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	dbName := ""
	if tenantID != 0 {
		dbName = fmt.Sprintf("hotel_reservation_%d", tenantID)
	} else {
		// this name is public multi_tenancy_database's name that contains tenants information
		dbName = "hotel_reservation"
	}

	if c.cache[tenantID] == nil {
		cn, err := tenant_connection_string_resolver.ResolveDB(false, dbName)
		if err != nil {
			panic(err.Error())
		}

		c.cache[tenantID] = cn
		return cn
	}

	return c.cache[tenantID]

}

func (c *TenantDatabaseResolver) CreateDbForTenant(db *gorm.DB, tenantId uint64) {

	dbName := fmt.Sprintf("hotel_reservation_%d", tenantId)
	db.Exec(fmt.Sprintf("CREATE DATABASE %s ENCODING 'UTF8';", dbName))
}

// Migrate migrate all entities
func (c *TenantDatabaseResolver) Migrate(db *gorm.DB, tenantId uint64) {

	for _, entity := range tenant_connection_string_resolver.GetEntities() {
		if err := db.AutoMigrate(entity); err != nil {
			panic(err.Error())
		}
	}
}
