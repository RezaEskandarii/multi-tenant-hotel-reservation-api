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
	cityService    = *services.NewCityService()
)

// handlers
var (
	countryHandler = handlers.CountryHandler{}
	cityHandler    = handlers.CityHandler{}
)

func RegisterServices(db *gorm.DB, router *echo.Group) {

	countriesRouter := router.Group("/countries")
	countryService.Repository.DB = db
	countryHandler.Register(countriesRouter, countryService)

	citiesRouter := router.Group("/cities")
	cityService.Repository.DB = db
	cityHandler.Register(citiesRouter, cityService)

}
