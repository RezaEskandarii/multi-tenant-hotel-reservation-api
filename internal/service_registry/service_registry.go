// Package service_registry
// register all handlers,services and dependencies
///**/
package service_registry

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"net/http"
	"reservation-api/api/handlers"
	"reservation-api/api/middlewares"
	"reservation-api/internal/appconfig"
	"reservation-api/internal/dto"
	"reservation-api/internal/repositories"
	"reservation-api/internal/services/common_services"
	"reservation-api/internal/services/domain_services"
	"reservation-api/pkg/applogger"
	"reservation-api/pkg/message_broker"
	"reservation-api/pkg/multi_tenancy_database/tenant_database_resolver"
)

// RegisterServicesAndRoutes register dependencies for services and handlers
func RegisterServicesAndRoutes(router *echo.Group) error {

	router.Use(middleware.Gzip())
	appConfig, err := appconfig.New()
	if err != nil {
		return err
	}

	addCorsMiddleware(router, appConfig)

	var (
		// ================================================================================================================
		logger = applogger.New(nil)
		ctx    = context.Background()
		// ================================= handlers =====================================================================
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
		// ================================================================================================================

		// ================================== common services =============================================================
		reportService = common_services.NewReportService()
		emailService  = common_services.NewEmailService(appConfig.Smtp.Host,
			appConfig.Smtp.Username, appConfig.Smtp.Password, appConfig.Smtp.Port,
		)

		rabbitMqManager = message_broker.New(appConfig.MessageBroker.Url, logger)
		fileService     = common_services.NewFileTransferService(appConfig.Minio.Endpoint, appConfig.Minio.AccessKeyID,
			appConfig.Minio.SecretAccessKey, appConfig.Minio.UseSSL, ctx)

		cacheService       = common_services.NewCacheService(appConfig.Redis.Addr, appConfig.Redis.Password, appConfig.Redis.CacheDB, ctx)
		eventService       = common_services.NewEventService(rabbitMqManager, emailService)
		connectionResolver = tenant_database_resolver.NewTenantDatabaseResolver()

		// =============================== domain services ===============================================================
		countryService        = domain_services.NewCountryService(repositories.NewCountryRepository(connectionResolver))
		provinceService       = domain_services.NewProvinceService(repositories.NewProvinceRepository(connectionResolver))
		cityService           = domain_services.NewCityService(repositories.NewCityRepository(connectionResolver), cacheService)
		currencyService       = domain_services.NewCurrencyService(repositories.NewCurrencyRepository(connectionResolver))
		userService           = domain_services.NewUserService(repositories.NewUserRepository(connectionResolver))
		hotelTypeService      = domain_services.NewHotelTypeService(repositories.NewHotelTypeRepository(connectionResolver))
		hotelGradeService     = domain_services.NewHotelGradeService(repositories.NewHotelGradeRepository(connectionResolver))
		hotelService          = domain_services.NewHotelService(repositories.NewHotelRepository(connectionResolver), fileService)
		roomTypeService       = domain_services.NewRoomTypeService(repositories.NewRoomTypeRepository(connectionResolver))
		roomService           = domain_services.NewRoomService(repositories.NewRoomRepository(connectionResolver))
		guestService          = domain_services.NewGuestService(repositories.NewGuestRepository(connectionResolver))
		rateGroupService      = domain_services.NewRateGroupService(repositories.NewRateGroupRepository(connectionResolver))
		rateCodeService       = domain_services.NewRateCodeService(repositories.NewRateCodeRepository(connectionResolver))
		rateCodeDetailService = domain_services.NewRateCodeDetailService(repositories.NewRateCodeDetailRepository(connectionResolver))
		reservationRepository = repositories.NewReservationRepository(connectionResolver, rateCodeDetailService.Repository)
		reservationService    = domain_services.NewReservationService(reservationRepository, rabbitMqManager)
		paymentService        = domain_services.NewPaymentService(repositories.NewPaymentRepository(connectionResolver))
		authService           = domain_services.NewAuthService(userService, appConfig)
		tenantService         = domain_services.NewTenantService(repositories.NewTenantDatabaseRepository(connectionResolver))
	)
	// ======================================================================================================================

	// fill handlers shared dependencies in handlerConf struct and pass this
	// struct to handlers inserted of pass many duplicated objects
	handlerConf := &dto.HandlerConfig{
		Router: router,
		Logger: logger,
	}

	// map swagger route
	router.GET("/swagger/*", echoSwagger.WrapHandler)

	// register tenant handler
	tenantHandler.Register(handlerConf, tenantService)

	// authHandler does bot need to authMiddleware.
	router.Use(middlewares.PanicRecoveryMiddleware(logger), middlewares.LoggerMiddleware(logger), middlewares.TenantMiddleware)

	// register auth handler
	authHandler.Register(handlerConf, userService, authService)

	// other handlers needs to this middlewares
	router.Use(middlewares.MetricsMiddleware, middlewares.JWTAuthMiddleware(authService), middlewares.TenantAccessMiddleware)

	// register all handlers
	metricHandler.Register(appConfig)
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
	scheduleRemoveExpiredReservationRequests(reservationService, logger, tenantService)

	// listen to message broker on reservation event and send email in background.
	go eventService.SendEmailToGuestOnReservation()

	return nil
}

func addCorsMiddleware(router *echo.Group, config *appconfig.Config) {
	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{config.Application.AllowedOrigins},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch,
			http.MethodPost, http.MethodDelete},
	}))
}
