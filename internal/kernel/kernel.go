package kernel

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"net/http"
	"reservation-api/internal/config"
	"reservation-api/internal/service_registry"
	"reservation-api/pkg/applogger"
)

var (
	logger        = applogger.New(nil)
	httpRouter    = getHttpRouter()
	v1RouterGroup = httpRouter.Group("/api/v1")
)

// Run run application
func Run() error {

	loadFlags()

	//connectionResolver := database.NewConnectionResolver()
	//db := connectionResolver.GetTenantDB("")

	cfg, err := config.NewConfig()

	if err != nil {
		return err
	}

	service_registry.RegisterServicesAndRoutes(v1RouterGroup, cfg)
	httpRouter.Logger.Fatal(httpRouter.Start(fmt.Sprintf(":%s", cfg.Application.Port)))

	return nil
}

// return new instance of echo.
func getHttpRouter() *echo.Echo {
	e := echo.New()
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	return e
}
