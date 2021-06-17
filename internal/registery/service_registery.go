package registery

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"hotel-reservation/internal/handlers"
	"hotel-reservation/internal/services"
)

// services
var (
	countryService  = *services.NewCountryService()
	provinceService = *services.NewProvinceService()
	cityService     = *services.NewCityService()
)

// handlers
var (
	countryHandler  = handlers.CountryHandler{}
	provinceHandler = handlers.ProvinceHandler{}
	cityHandler     = handlers.CityHandler{}
)

func RegisterServices(db *gorm.DB, router *echo.Group) {

	setRepositoriesDb(db)

	countriesRouter := router.Group("/countries")
	countryHandler.Register(countriesRouter, countryService)

	provinceRouter := router.Group("/provinces")
	provinceHandler.Register(provinceRouter, provinceService)

	citiesRouter := router.Group("/cities")
	cityHandler.Register(citiesRouter, cityService)

}

func setRepositoriesDb(db *gorm.DB) {
	countryService.Repository.DB = db
	provinceService.Repository.DB = db
	cityService.Repository.DB = db
}
