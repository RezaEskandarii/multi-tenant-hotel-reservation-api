package registery

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"reservation-api/api/handlers"
	"reservation-api/internal/config"
	"reservation-api/internal/dto"
	"reservation-api/internal/repositories"
	"reservation-api/internal/services/common_services"
	"reservation-api/internal/services/domain_services"
	"reservation-api/pkg/applogger"
	"reservation-api/pkg/cache"
	"reservation-api/pkg/rabbitmq"
	"reservation-api/pkg/translator"
)

// services
var (
	countryService        = domain_services.NewCountryService()
	provinceService       = domain_services.NewProvinceService()
	cityService           = domain_services.NewCityService()
	currencyService       = domain_services.NewCurrencyService()
	userService           = domain_services.NewUserService()
	hotelTypeService      = domain_services.NewHotelTypeService()
	hotelGradeService     = domain_services.NewHotelGradeService()
	hotelService          = domain_services.NewHotelService()
	roomTypeService       = domain_services.NewRoomTypeService()
	roomService           = domain_services.NewRoomService()
	guestService          = domain_services.NewGuestService()
	rateGroupService      = domain_services.NewRateGroupService()
	rateCodeService       = domain_services.NewRateCodeService()
	auditService          = domain_services.NewAuditService()
	rateCodeDetailService = domain_services.NewRateCodeDetailService()
	reservationService    = domain_services.NewReservationService()
	paymentService        = domain_services.NewPaymentService()
)

// handlers
var (
	countryHandler     = handlers.CountryHandler{}
	provinceHandler    = handlers.ProvinceHandler{}
	cityHandler        = handlers.CityHandler{}
	currencyHandler    = handlers.CurrencyHandler{}
	usersHandler       = handlers.UserHandler{}
	hotelTypeHandler   = handlers.HotelTypeHandler{}
	hotelGradeHandler  = handlers.HotelGradeHandler{}
	hotelHandler       = handlers.HotelHandler{}
	roomTypeHandler    = handlers.RoomTypeHandler{}
	roomHandler        = handlers.RoomHandler{}
	guestHandler       = handlers.GuestHandler{}
	rateGroupHandler   = handlers.RateGroupHandler{}
	rateCodeHandler    = handlers.RateCodeHandler{}
	authHandler        = handlers.AuthHandler{}
	reservationHandler = handlers.ReservationHandler{}
	paymentHandler     = handlers.PaymentHandler{}
)

// RegisterServices register dependencies for services and handlers
func RegisterServices(db *gorm.DB, router *echo.Group, cfg *config.Config) {

	// set service layer repository and database object.
	setServicesRepository(db, cfg)

	logger := applogger.New(nil)
	i18nTranslator := translator.New()

	handlerInput := &dto.HandlerInput{
		Router:     router,
		Translator: i18nTranslator,
		Logger:     logger,
	}

	q := rabbitmq.New(cfg.MessageBroker.Url, logger)

	h := func(s interface{}) {
		fmt.Println(fmt.Sprintf("%s", s))
	}

	go func() {
		er := q.Consume("email", h)
		fmt.Println(er)
	}()
	// authHandler does bot need to authMiddleware.
	authHandler.Register(handlerInput, userService)

	// add authentication middleware to all routes.
	router.Use( /**middlewares.JWTAuthMiddleware, */ )

	countryHandler.Register(handlerInput, countryService)

	provinceHandler.Register(handlerInput, provinceService)

	cityHandler.Register(handlerInput, cityService)

	currencyHandler.Register(handlerInput, currencyService)

	usersHandler.Register(handlerInput, userService)

	hotelTypeHandler.Register(handlerInput, hotelTypeService)

	hotelGradeHandler.Register(handlerInput, hotelGradeService)

	hotelHandler.Register(handlerInput, hotelService)

	roomTypeHandler.Register(handlerInput, roomTypeService)

	roomHandler.Register(handlerInput, roomService)

	guestHandler.Register(handlerInput, guestService)

	rateGroupHandler.Register(handlerInput, rateGroupService)

	rateCodeHandler.Register(handlerInput, rateCodeService, rateCodeDetailService)

	reservationHandler.Register(handlerInput, reservationService)

	paymentHandler.Register(handlerInput, paymentService)

	scheduleRemoveExpiredReservationRequests(reservationService, logger)
}

// set repository dependency
func setServicesRepository(db *gorm.DB, cfg *config.Config) {

	cacheManager := &cache.Manager{}
	fileService := common_services.FileTransferService{}
	// fileTransferService context for minio
	ctx := context.Background()
	fileTransferService := fileService.New(cfg.Minio.Endpoint, cfg.Minio.AccessKeyID, cfg.Minio.SecretAccessKey, cfg.Minio.UseSSL, ctx)

	countryService.Repository = repositories.NewCountryRepository(db)
	provinceService.Repository = repositories.NewProvinceRepository(db)
	cityService.Repository = repositories.NewCityRepository(db)
	cityService.CacheManager = cacheManager
	currencyService.Repository = repositories.NewCurrencyRepository(db)
	userService.Repository = repositories.NewUserRepository(db)
	hotelTypeService.Repository = repositories.NewHotelTypeRepository(db)
	hotelGradeService.Repository = repositories.NewHotelGradeRepository(db)
	hotelService.Repository = repositories.NewHotelRepository(db, fileTransferService)
	roomTypeService.Repository = repositories.NewRoomTypeRepository(db)
	roomService.Repository = repositories.NewRoomRepository(db)
	guestService.Repository = repositories.NewGuestRepository(db)
	rateGroupService.Repository = repositories.NewRateGroupRepository(db)
	rateCodeService.Repository = repositories.NewRateCodeRepository(db)
	auditService.Repository = repositories.NewAuditRepository(db)
	rateCodeDetailService.Repository = repositories.NewRateCodeDetailRepository(db)
	reservationService.Repository = repositories.NewReservationRepository(db, rateCodeDetailService.Repository)
}

// ApplySeed seeds given json file to database.
func ApplySeed(db *gorm.DB) error {

	setServicesRepository(db, nil)

	// seed users
	if err := userService.Seed("./data/seed/users.json"); err != nil {
		return err
	}
	// seed roomTypes
	if err := roomTypeService.Seed("./data/seed/room_types.json"); err != nil {
		return err
	}

	// seed currencies
	if err := currencyService.Seed("./data/seed/currencies.json"); err != nil {
		return err
	}

	return nil
}
