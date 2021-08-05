package registery

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	handlers2 "hotel-reservation/api/handlers"
	"hotel-reservation/internal/repositories"
	"hotel-reservation/internal/services"
	"hotel-reservation/pkg/applogger"
	"hotel-reservation/pkg/translator"
)

// services
var (
	countryService        = services.NewCountryService()
	provinceService       = services.NewProvinceService()
	cityService           = services.NewCityService()
	currencyService       = services.NewCurrencyService()
	userService           = services.NewUserService()
	residenceTypeService  = services.NewResidenceTypeService()
	residenceGradeService = services.NewResidenceGradeService()
	residenceService      = services.NewResidenceService()
	roomTypeService       = services.NewRoomTypeService()
)

// handlers
var (
	countryHandler        = handlers2.CountryHandler{}
	provinceHandler       = handlers2.ProvinceHandler{}
	cityHandler           = handlers2.CityHandler{}
	currencyHandler       = handlers2.CurrencyHandler{}
	usersHandler          = handlers2.UserHandler{}
	residenceTypeHandler  = handlers2.ResidenceTypeHandler{}
	residenceGradeHandler = handlers2.ResidenceGradeHandler{}
	residenceHandler      = handlers2.ResidenceHandler{}
	roomTypeHandler       = handlers2.RoomTypeHandler{}
)

// pckgs
var (
	i18nTranslator = translator.New()
)

// RegisterServices register dependencies for services and handlers
func RegisterServices(db *gorm.DB, router *echo.Group) {

	setServicesRepository(db)

	logger := applogger.New()

	countriesRouter := router.Group("/countries")
	countryHandler.Register(countriesRouter, countryService, i18nTranslator, logger)

	provinceRouter := router.Group("/provinces")
	provinceHandler.Register(provinceRouter, provinceService, i18nTranslator, logger)

	citiesRouter := router.Group("/cities")
	cityHandler.Register(citiesRouter, cityService, i18nTranslator, logger)

	currencyRouter := router.Group("/currencies")
	currencyHandler.Register(currencyRouter, currencyService, i18nTranslator, logger)

	usersRouter := router.Group("/users")
	usersHandler.Register(usersRouter, userService, i18nTranslator, logger)

	residenceTypeRouter := router.Group("/residence-type")
	residenceTypeHandler.Register(residenceTypeRouter, residenceTypeService, i18nTranslator, logger)

	residenceGradeRouter := router.Group("/residence-grade")
	residenceGradeHandler.Register(residenceGradeRouter, residenceGradeService, i18nTranslator, logger)

	residenceRouteGroup := router.Group("/residence")
	residenceHandler.Register(residenceRouteGroup, residenceService, i18nTranslator, logger)

	roomTypeRouteGroup := router.Group("/room-type")
	roomTypeHandler.Register(roomTypeRouteGroup, roomTypeService, i18nTranslator, logger)
}

func setServicesRepository(db *gorm.DB) {
	countryService.Repository = repositories.NewCountryRepository(db)
	provinceService.Repository = repositories.NewProvinceRepository(db)
	cityService.Repository = repositories.NewCityRepository(db)
	currencyService.Repository = repositories.NewCurrencyRepository(db)
	userService.Repository = repositories.NewUserRepository(db)
	residenceTypeService.Repository = repositories.NewResidenceTypeRepository(db)
	residenceGradeService.Repository = repositories.NewResidenceGradeRepository(db)
	residenceService.Repository = repositories.NewResidenceRepository(db)
	roomTypeService.Repository = repositories.NewRoomTypeRepository(db)
}
