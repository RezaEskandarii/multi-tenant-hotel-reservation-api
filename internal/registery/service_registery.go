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
	countryService        = *services.NewCountryService()
	provinceService       = *services.NewProvinceService()
	cityService           = *services.NewCityService()
	currencyService       = *services.NewCurrencyService()
	userService           = services.NewUserService()
	residenceTypeService  = *services.NewResidenceTypeService()
	residenceGradeService = *services.NewResidenceGradeService()
)

// handlers
var (
	countryHandler        = handlers.CountryHandler{}
	provinceHandler       = handlers.ProvinceHandler{}
	cityHandler           = handlers.CityHandler{}
	currencyHandler       = handlers.CurrencyHandler{}
	usersHandler          = handlers.UserHandler{}
	residenceTypeHandler  = handlers.ResidenceTypeHandler{}
	residenceGradeHandler = handlers.ResidenceGradeHandler{}
)

// pckgs
var (
	i18nTranslator = translator.New()
)

// RegisterServices register dependencies for services and handlers
func RegisterServices(db *gorm.DB, router *echo.Group) {

	setRepositoriesDb(db)

	countriesRouter := router.Group("/countries")
	countryHandler.Register(countriesRouter, countryService, i18nTranslator)

	provinceRouter := router.Group("/provinces")
	provinceHandler.Register(provinceRouter, provinceService, i18nTranslator)

	citiesRouter := router.Group("/cities")
	cityHandler.Register(citiesRouter, cityService, i18nTranslator)

	currencyRouter := router.Group("/currencies")
	currencyHandler.Register(currencyRouter, currencyService, i18nTranslator)

	usersRouter := router.Group("/users")
	usersHandler.Register(usersRouter, userService, i18nTranslator)

	residenceTypeRouter := router.Group("/residence-type")
	residenceTypeHandler.Register(residenceTypeRouter, residenceTypeService, i18nTranslator)

	residenceGradeRouter := router.Group("/residence-grade")
	residenceGradeHandler.Register(residenceGradeRouter, &residenceGradeService, i18nTranslator)
}

func setRepositoriesDb(db *gorm.DB) {
	countryService.Repository = repositories.NewCountryRepository(db)
	provinceService.Repository = repositories.NewProvinceRepository(db)
	cityService.Repository = repositories.NewCityRepository(db)
	currencyService.Repository = repositories.NewCurrencyRepository(db)
	userService.Repository = repositories.NewUserRepository(db)
	residenceTypeService.Repository = repositories.NewResidenceTypeRepository(db)
	residenceGradeService.Repository = repositories.NewResidenceGradeRepository(db)
}
