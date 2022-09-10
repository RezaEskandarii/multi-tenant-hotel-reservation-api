package connection_resolver

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	config2 "reservation-api/internal/config"
	"reservation-api/internal/models"
	"strings"
	"time"
)

type ConnectionResolver struct {
	cache map[uint64]*gorm.DB
}

func NewConnectionResolver() *ConnectionResolver {
	return &ConnectionResolver{
		cache: make(map[uint64]*gorm.DB),
	}
}

func (c *ConnectionResolver) Resolve(tenantID uint64) *gorm.DB {

	dbName := ""
	if tenantID != 0 {
		dbName = fmt.Sprintf("hotel_reservation_%d", tenantID)
	}

	if c.cache[tenantID] == nil {
		cn, err := getDB(false, dbName)
		if err != nil {
			panic(err.Error())
		}

		c.cache[tenantID] = cn
		return cn
	}

	return c.cache[tenantID]

}

func (c *ConnectionResolver) CreateDbForTenant(db *gorm.DB, tenantId uint64) {

	dbName := fmt.Sprintf("hotel_reservation_%d", tenantId)

	db.Exec(fmt.Sprintf("CREATE DATABASE %s ENCODING 'UTF8';", dbName))
}

func (c *ConnectionResolver) Migrate(db *gorm.DB, tenantId uint64) {

	var (
		entities = []interface{}{
			models.Country{},
			models.City{},
			models.Province{},
			models.Currency{},
			models.User{},
			models.Hotel{},
			models.Room{},
			models.RoomType{},
			models.Guest{},
			models.RateGroup{},
			models.RateCode{},
			models.HotelGrade{},
			models.HotelType{},
			models.ReservationRequest{},
			models.Reservation{},
			models.Audit{},
			models.RateCodeDetail{},
			models.RateCodeDetailPrice{},
			models.Sharer{},
			models.Thumbnail{},
		}
	)

	for _, entity := range entities {
		if err := db.AutoMigrate(entity); err != nil {
			panic(err.Error())
		}
	}
}

/*==========================================================================================*/
var dbLogger = logger.New(
	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	logger.Config{
		SlowThreshold: time.Second,   // Slow SQL threshold
		LogLevel:      logger.Silent, // Log level
		Colorful:      true,          // Disable color
	},
)

/*============================================================================================*/
// generate database connection string
func getDSN(tenantDbName string) (string, error) {
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

/*========================================================================================*/
func getDB(usesInTestEnv bool, tenantDbName string) (*gorm.DB, error) {

	if usesInTestEnv {
		os.Setenv("CONFIG_PATH", "../resources/config.yml")
	}

	connectionString, err := getDSN(tenantDbName)

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
		return nil, err
	}

	return connection, nil
}
