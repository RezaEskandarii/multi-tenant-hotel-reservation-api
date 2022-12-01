package service_registry

import (
	"context"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"reservation-api/api/handlers"
	"reservation-api/api/middlewares"
	"reservation-api/internal/config"
	"reservation-api/internal/dto"
	"reservation-api/internal/repositories"
	"reservation-api/internal/services/common_services"
	"reservation-api/internal/services/domain_services"
	"reservation-api/pkg/applogger"
	"reservation-api/pkg/message_broker"
	"reservation-api/pkg/multi_tenancy_database/tenant_database_resolver"
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
	tenantHandler      = handlers.TenantHandler{}
	metricHandler      = handlers.MetricHandler{}
)

// RegisterServicesAndRoutes register dependencies for services and handlers
func RegisterServicesAndRoutes(router *echo.Group, cfg *config.Config) {

	logger := applogger.New(nil)

	// fill handlers shared dependencies in handlerConf struct and pass this
	// struct to handlers inserted of pass many duplicated objects
	handlerConf := &dto.HandlerConfig{
		Router: router,
		Logger: logger,
	}

	router.GET("/swagger/*", echoSwagger.WrapHandler)

	reportService := common_services.NewReportService()
	ctx := context.Background()

	emailService := common_services.NewEmailService(cfg.Smtp.Host,
		cfg.Smtp.Username, cfg.Smtp.Password, cfg.Smtp.Port,
	)

	rabbitMqManager := message_broker.New(cfg.MessageBroker.Url, logger)

	fileService := common_services.NewFileTransferService(cfg.Minio.Endpoint, cfg.Minio.AccessKeyID,
		cfg.Minio.SecretAccessKey, cfg.Minio.UseSSL, ctx)

	cacheService := common_services.NewCacheService(cfg.Redis.Addr, cfg.Redis.Password, cfg.Redis.CacheDB, ctx)
	eventService := common_services.NewEventService(rabbitMqManager, emailService)
	connectionResolver := tenant_database_resolver.NewTenantDatabaseResolver()

	var (
		countryService        = domain_services.NewCountryService(repositories.NewCountryRepository(connectionResolver))
		provinceService       = domain_services.NewProvinceService(repositories.NewProvinceRepository(connectionResolver))
		cityService           = domain_services.NewCityService(repositories.NewCityRepository(connectionResolver), cacheService)
		currencyService       = domain_services.NewCurrencyService(repositories.NewCurrencyRepository(connectionResolver))
		userService           = domain_services.NewUserService(repositories.NewUserRepository(connectionResolver))
		hotelTypeService      = domain_services.NewHotelTypeService(repositories.NewHotelTypeRepository(connectionResolver))
		hotelGradeService     = domain_services.NewHotelGradeService(repositories.NewHotelGradeRepository(connectionResolver))
		hotelService          = domain_services.NewHotelService(repositories.NewHotelRepository(connectionResolver, fileService))
		roomTypeService       = domain_services.NewRoomTypeService(repositories.NewRoomTypeRepository(connectionResolver))
		roomService           = domain_services.NewRoomService(repositories.NewRoomRepository(connectionResolver))
		guestService          = domain_services.NewGuestService(repositories.NewGuestRepository(connectionResolver))
		rateGroupService      = domain_services.NewRateGroupService(repositories.NewRateGroupRepository(connectionResolver))
		rateCodeService       = domain_services.NewRateCodeService(repositories.NewRateCodeRepository(connectionResolver))
		rateCodeDetailService = domain_services.NewRateCodeDetailService(repositories.NewRateCodeDetailRepository(connectionResolver))
		reservationRepository = repositories.NewReservationRepository(connectionResolver, rateCodeDetailService.Repository)
		reservationService    = domain_services.NewReservationService(reservationRepository, rabbitMqManager)
		paymentService        = domain_services.NewPaymentService(repositories.NewPaymentRepository(connectionResolver))
		authService           = domain_services.NewAuthService(userService, cfg)
		tenantService         = domain_services.NewTenantService(repositories.NewTenantDatabaseRepository(connectionResolver))
		//auditService          = domain_services.NewAuditService(repositories.NewAuditRepository(connectionResolver))
	)

	tenantHandler.Register(handlerConf, tenantService)
	// authHandler does bot need to authMiddleware.
	router.Use(middlewares.PanicRecoveryMiddleware(logger), middlewares.LoggerMiddleware(logger), middlewares.TenantMiddleware)

	authHandler.Register(handlerConf, userService, authService)

	router.Use(middlewares.MetricsMiddleware, middlewares.JWTAuthMiddleware(authService),
		middlewares.TenantAccessMiddleware)

	// add authentication middleware to all routes.

	metricHandler.Register(cfg)
	countryHandler.Register(handlerConf, countryService)

	provinceHandler.Register(handlerConf, provinceService)

	cityHandler.Register(handlerConf, cityService)

	currencyHandler.Register(handlerConf, currencyService)

	usersHandler.Register(handlerConf, userService)

	hotelTypeHandler.Register(handlerConf, hotelTypeService)

	hotelGradeHandler.Register(handlerConf, hotelGradeService)

	hotelHandler.Register(handlerConf, hotelService)

	roomTypeHandler.Register(handlerConf, roomTypeService)

	roomHandler.Register(handlerConf, roomService)

	guestHandler.Register(handlerConf, guestService, reportService)

	rateGroupHandler.Register(handlerConf, rateGroupService)

	rateCodeHandler.Register(handlerConf, rateCodeService, rateCodeDetailService)

	reservationHandler.Register(handlerConf, reservationService, reportService)

	paymentHandler.Register(handlerConf, paymentService)

	// schedule to remove expired reservation requests.
	scheduleRemoveExpiredReservationRequests(reservationService, logger)

	// listen to message broker on reservation event and send email in background.
	go eventService.SendEmailToGuestOnReservation()

}
