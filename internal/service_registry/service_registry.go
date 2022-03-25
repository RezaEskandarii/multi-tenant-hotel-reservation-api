package service_registry

import (
	"context"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"reservation-api/api/handlers"
	"reservation-api/internal/config"
	"reservation-api/internal/dto"
	"reservation-api/internal/repositories"
	"reservation-api/internal/services/common_services"
	"reservation-api/internal/services/domain_services"
	"reservation-api/pkg/applogger"
	"reservation-api/pkg/message_broker"
	"reservation-api/pkg/translator"
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
	metricHandler      = handlers.MetricHandler{}
)

// RegisterServicesAndRoutes register dependencies for services and handlers
func RegisterServicesAndRoutes(db *gorm.DB, router *echo.Group, cfg *config.Config) {

	// logger
	logger := applogger.New(nil)
	// i18n translator
	i18nTranslator := translator.New()

	// fill handlers shared dependencies in handlerInput struct and pass this
	// struct to handlers inserted of pass many duplicated objects
	handlerInput := &dto.HandlerInput{
		Router:     router,
		Translator: i18nTranslator,
		Logger:     logger,
	}

	/****************************************** register services ***********************************************************/
	ctx := context.Background()

	// gomail implementation
	emailService := common_services.NewEmailService(cfg.Smtp.Host,
		cfg.Smtp.Username, cfg.Smtp.Password, cfg.Smtp.Port,
	)
	// rabbitmq implementation
	rabbitMqManager := message_broker.New(cfg.MessageBroker.Url, logger)
	// minio implementation
	fileService := common_services.NewFileTransferService(cfg.Minio.Endpoint, cfg.Minio.AccessKeyID,
		cfg.Minio.SecretAccessKey, cfg.Minio.UseSSL, ctx)
	// redis implementation
	cacheService := common_services.NewCacheService(cfg.Redis.Addr, cfg.Redis.Password, cfg.Redis.CacheDB, ctx)

	eventService := common_services.NewEventService(rabbitMqManager, emailService)
	var (
		countryService    = domain_services.NewCountryService(repositories.NewCountryRepository(db))
		provinceService   = domain_services.NewProvinceService(repositories.NewProvinceRepository(db))
		cityService       = domain_services.NewCityService(repositories.NewCityRepository(db), cacheService)
		currencyService   = domain_services.NewCurrencyService(repositories.NewCurrencyRepository(db))
		userService       = domain_services.NewUserService(repositories.NewUserRepository(db))
		hotelTypeService  = domain_services.NewHotelTypeService(repositories.NewHotelTypeRepository(db))
		hotelGradeService = domain_services.NewHotelGradeService(repositories.NewHotelGradeRepository(db))
		hotelService      = domain_services.NewHotelService(repositories.NewHotelRepository(db, fileService))
		roomTypeService   = domain_services.NewRoomTypeService(repositories.NewRoomTypeRepository(db))
		roomService       = domain_services.NewRoomService(repositories.NewRoomRepository(db))
		guestService      = domain_services.NewGuestService(repositories.NewGuestRepository(db))
		rateGroupService  = domain_services.NewRateGroupService(repositories.NewRateGroupRepository(db))
		rateCodeService   = domain_services.NewRateCodeService(repositories.NewRateCodeRepository(db))
		//auditService          = domain_services.NewAuditService(repositories.NewAuditRepository(db))
		rateCodeDetailService = domain_services.NewRateCodeDetailService(repositories.NewRateCodeDetailRepository(db))

		reservationRepository = repositories.NewReservationRepository(db, rateCodeDetailService.Repository)
		reservationService    = domain_services.NewReservationService(reservationRepository, rabbitMqManager)
		paymentService        = domain_services.NewPaymentService()
		authService           = domain_services.NewAuthService(userService, cfg)
	)

	/****************************************** end register services ***********************************************************/

	// authHandler does bot need to authMiddleware.
	authHandler.Register(handlerInput, userService, authService)

	// add authentication middleware to all routes.
	router.Use( /**middlewares.JWTAuthMiddleware, */ )

	metricHandler.Register(cfg)

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

	// schedule to remove expired reservation requests.
	scheduleRemoveExpiredReservationRequests(reservationService, logger)

	// listen to message broker on reservation event and send email in background.
	go eventService.SendEmailToGuestOnReservation()

}
