package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	config2 "hotel-reservation/internal/config"
	"hotel-reservation/pkg/application_loger"
	"log"
	"os"
	"strings"
	"time"
)

var dbLogger = logger.New(
	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	logger.Config{
		SlowThreshold: time.Second,   // Slow SQL threshold
		LogLevel:      logger.Silent, // Log level
		Colorful:      true,          // Disable color
	},
)

// generate database connection string
func getDSN() (string, error) {
	dbCfg, err := config2.NewConfig()

	if err != nil {
		return "", nil
	}

	cfg := dbCfg.Database

	switch strings.TrimSpace(cfg.Engine) {

	case "postgres":
		return fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.Name, cfg.SSLMode,
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

func GetDb(usesInTestEnv bool) (*gorm.DB, error) {

	if usesInTestEnv {
		os.Setenv("CONFIG_PATH", "../resources/config.yml")
	}

	connectionString, err := getDSN()

	if err != nil {
		return nil, err
	}

	connection, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  connectionString,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		Logger:                                   dbLogger,
		SkipDefaultTransaction:                   true,
		DisableForeignKeyConstraintWhenMigrating: true,
		PrepareStmt:                              false,
	})

	sqlDB, err := connection.DB()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(25)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(25)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	if err != nil {
		application_loger.LogError(err.Error())
		return nil, err
	}

	return connection, nil
}
