package kernel

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"hotel-reservation/internal/config"
	"hotel-reservation/internal/registery"
	"hotel-reservation/pkg/application_loger"
	"hotel-reservation/pkg/database"
	"net/http"
)

// Run run application
func Run(port int) error {
	fmt.Println("application started ...")

	defer func() {
		if r := recover(); r != nil {
			application_loger.LogError(r)
			return
		}
	}()

	cfg, err := config.NewConfig()

	if err != nil {
		return err
	}

	db, err := database.GetDb(false)
	if err != nil {
		return err
	}

	if cfg.Application.SqlDebug {
		db = db.Debug()
	}

	if cfg.Application.IgnoreMigration == false {
		err = database.Migrate(db)
		if err != nil {
			return err
		}
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

	e.Logger.Fatal(e.Start(portStr))

	return nil
}
