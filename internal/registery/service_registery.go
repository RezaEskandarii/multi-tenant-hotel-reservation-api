package registery

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"hotel-reservation/internal/handlers"
	"hotel-reservation/internal/repositories"
	"hotel-reservation/internal/services"
	"hotel-reservation/pkg/translator"
)

// services
var (
	countryService  = *services.NewCountryService()
	provinceService = *services.NewProvinceService()
	cityService     = *services.NewCityService()
	currencyService = *services.NewCurrencyService()
)

// handlers
var (
	countryHandler  = handlers.CountryHandler{}
	provinceHandler = handlers.ProvinceHandler{}
	cityHandler     = handlers.CityHandler{}
	currencyHandler = handlers.CurrencyHandler{}
)

// pckgs

var (
	i18nTranslator = translator.New()
)

// RegisterServices register dependencies for services and handlers
func RegisterServices(db *gorm.DB, router *echo.Group) {

	setRepositoriesDb(db)

	countriesRouter := router.Group("/countries")
	countryHandler.Register(countriesRouter, countryService)

	provinceRouter := router.Group("/provinces")
	provinceHandler.Register(provinceRouter, provinceService)

	citiesRouter := router.Group("/cities")
	cityHandler.Register(citiesRouter, cityService, i18nTranslator)

	currencyRouter := router.Group("/currencies")
	currencyHandler.Register(currencyRouter, currencyService)
}

func setRepositoriesDb(db *gorm.DB) {
	countryService.Repository.DB = db
	provinceService.Repository.DB = db
	cityService.Repository.DB = db

	currencyService.Repository = repositories.NewCurrencyRepository(db)

}
