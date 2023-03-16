package app

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"reservation-api/internal/appconfig"
	"reservation-api/internal/service_registry"
	"reservation-api/pkg/applogger"
)

var (
	logger = applogger.New(nil)
)

// Run starts application
func Run() error {

	loadFlags()

	cfg, err := appconfig.New()

	if err != nil {
		return err
	}

	e := echo.New()
	v1RouterGroup := e.Group("/api/v1")

	// register all routes,handlers and dependencies
	if err := service_registry.RegisterServicesAndRoutes(v1RouterGroup); err != nil {
		return err
	} else {

		e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.Application.Port)))
	}

	return nil
}
