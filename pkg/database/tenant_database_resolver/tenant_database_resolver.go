package tenant_database_resolver

import (
	"fmt"
	"gorm.io/gorm"
	"reservation-api/pkg/tenant_connection_string_resolver"
)

type TenantDatabaseResolver struct {
	cache map[uint64]*gorm.DB
}

func NewTenantDatabaseResolver() *TenantDatabaseResolver {
	return &TenantDatabaseResolver{
		cache: make(map[uint64]*gorm.DB),
	}
}

func (c *TenantDatabaseResolver) GetTenantDB(tenantID uint64) *gorm.DB {

	dbName := ""
	if tenantID != 0 {
		dbName = fmt.Sprintf("hotel_reservation_%d", tenantID)
	} else {
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

func (c *TenantDatabaseResolver) Migrate(db *gorm.DB, tenantId uint64) {

	for _, entity := range tenant_connection_string_resolver.GetEntities() {
		if err := db.AutoMigrate(entity); err != nil {
			panic(err.Error())
		}
	}
}
