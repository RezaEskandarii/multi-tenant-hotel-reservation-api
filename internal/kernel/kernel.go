package kernel

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"os"
	"reservation-api/cmd"
	"reservation-api/internal/config"
	"reservation-api/internal/registery"
	"reservation-api/pkg/applogger"
	"reservation-api/pkg/database"
)

var (
	db, dbErr = database.GetDb(false)
	logger    = applogger.New(nil)
)

func SetCommands() {
	commands := make([]cmd.Runner, 0)

	migrateCmd := cmd.NewCommand(cmd.Migrate, "false", "migrate structs to database", func() error {
		return database.Migrate(db)
	})

	commands = append(commands, migrateCmd)

	for _, arg := range os.Args {

		for _, c := range commands {
			if c.Name() == arg {
				c.Run()
			}
		}
	}

}

// Run run application
func Run(port int) error {

	SetCommands()

	if dbErr != nil {
		return dbErr
	}

	logger.LogInfo("application started ...")

	defer func() {
		if r := recover(); r != nil {
			logger.LogError(r)
			return
		}
	}()

	cfg, err := config.NewConfig()

	if err != nil {
		return err
	}

	if err != nil {
		logger.LogError(err.Error())
		return err
	}

	if cfg.Application.SqlDebug {
		db = db.Debug()
	}

	portStr := fmt.Sprintf(":%d", port)

	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

	corsCfg := middleware.CORSConfig{
		Skipper: func(context echo.Context) bool {
			return true
		},
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}

	e.Use(middleware.CORSWithConfig(corsCfg))

	router := e.Group("/v1")

	registery.RegisterServices(db, router)

	if err := registery.ApplySeed(db); err != nil {
		logger.LogError(err.Error())
		return err
	}

	e.Logger.Fatal(e.Start(portStr))

	return nil
}
