package tenant_connection_string_resolver

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"time"
)

func ResolveDB(usesInTestEnv bool, tenantDbName string) (*gorm.DB, error) {

	if usesInTestEnv {
		// read configs from given path in tests.
		os.Setenv("CONFIG_PATH", "../resources/config.yml")
	}

	// get unique connection string per given tenantID
	connectionString, err := ResolveConnectionString(tenantDbName)

	if err != nil {
		return nil, err
	}

	connection, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  connectionString,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		Logger:                                   GetDbLogger(),
		SkipDefaultTransaction:                   true,
		DisableForeignKeyConstraintWhenMigrating: true,
		PrepareStmt:                              false,
	})

	sqlDB, err := connection.DB()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(25)

	// SetMaxOpenConns sets the maximum number of open connections to the multi_tenancy_database.
	sqlDB.SetMaxOpenConns(25)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	if err != nil {
		return nil, err
	}

	return connection, nil
}
