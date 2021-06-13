package registery

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"hotel-reservation/internal/handlers"
	"hotel-reservation/internal/services"
)

// services
var (
	countryService = *services.NewCountryService()
)

// handlers
var (
	countryHandler = handlers.CountryHandler{}
)

func RegisterServices(db *gorm.DB, router *echo.Group) {

	countriesRouter := router.Group("/countries")
	countryService.Repository.DB = db
	countryHandler.Register(countriesRouter, countryService)
}
