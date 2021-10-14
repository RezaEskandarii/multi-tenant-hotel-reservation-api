package registery

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"hotel-reservation/api/handlers"
	"hotel-reservation/internal/dto"
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
	roomService           = services.NewRoomService()
	guestService          = services.NewGuestService()
	rateGroupService      = services.NewRateGroupService()
	rateCodeService       = services.NewRateCodeService()
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
	residenceHandler      = handlers.ResidenceHandler{}
	roomTypeHandler       = handlers.RoomTypeHandler{}
	roomHandler           = handlers.RoomHandler{}
	guestHandler          = handlers.GuestHandler{}
	rateGroupHandler      = handlers.RateGroupHandler{}
	rateCodeHandler       = handlers.RateCodeHandler{}
)

// RegisterServices register dependencies for services and handlers
func RegisterServices(db *gorm.DB, router *echo.Group) {

	setServicesRepository(db)

	logger := applogger.New(nil)
	i18nTranslator := translator.New()

	handlerInput := &dto.HandlerInput{
		Router:     router,
		Translator: i18nTranslator,
		Logger:     logger,
	}

	countryHandler.Register(handlerInput, countryService)

	provinceHandler.Register(handlerInput, provinceService)

	cityHandler.Register(handlerInput, cityService)

	currencyHandler.Register(handlerInput, currencyService)

	usersHandler.Register(handlerInput, userService)

	residenceTypeHandler.Register(handlerInput, residenceTypeService)

	residenceGradeHandler.Register(handlerInput, residenceGradeService)

	residenceHandler.Register(handlerInput, residenceService)

	roomTypeHandler.Register(handlerInput, roomTypeService)

	roomHandler.Register(handlerInput, roomService)

	guestHandler.Register(handlerInput, guestService)

	rateGroupHandler.Register(handlerInput, rateGroupService)

	rateCodeHandler.Register(handlerInput, rateCodeService)

}

// set repository dependency
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
	roomService.Repository = repositories.NewRoomRepository(db)
	guestService.Repository = repositories.NewGuestRepository(db)
	rateGroupService.Repository = repositories.NewRateGroupRepository(db)
	rateCodeService.Repository = repositories.NewRateCodeRepository(db)
}
