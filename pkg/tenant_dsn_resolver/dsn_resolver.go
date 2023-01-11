package tenant_dsn_resolver

import (
	"fmt"
	config2 "reservation-api/internal/appconfig"
	"strings"
)

// ResolveConnectionString returns unique multi_tenancy_database connection string per given tenantID.
func ResolveConnectionString(tenantDbName string) (string, error) {

	// read configs
	dbCfg, err := config2.New()

	if err != nil {
		return "", nil
	}

	cfg := dbCfg.Database

	dbName := cfg.Name
	if tenantDbName != "" {
		dbName = tenantDbName
	}

	switch strings.TrimSpace(cfg.Engine) {

	case "postgres":
		return fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			cfg.Host, cfg.Port, cfg.Username, cfg.Password, dbName, cfg.SSLMode,
		), nil

	case "mysql":
		return fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4,utf8&parseTime=True",
			cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Name,
		), nil

	case "mssql":
		return fmt.Sprintf("sqlserver://%s:%s@%s:%s?multi_tenancy_database=%s",
			cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Name,
		), nil

	default:
		return fmt.Sprintf(
			//set default ConnectionString to postgresql
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.Name, cfg.SSLMode,
		), nil

	}
}
