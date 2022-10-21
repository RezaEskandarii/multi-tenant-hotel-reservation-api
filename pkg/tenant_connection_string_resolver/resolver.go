package tenant_connection_string_resolver

import (
	"fmt"
	config2 "reservation-api/internal/config"
	"strings"
)

func ResolveConnectionString(tenantDbName string) (string, error) {
	dbCfg, err := config2.NewConfig()

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
		return fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
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
