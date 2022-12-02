package service_registry

import (
	"context"
	"reservation-api/api/handlers"
	"reservation-api/internal/appconfig"
	"reservation-api/internal/repositories"
	"reservation-api/internal/services/common_services"
	"reservation-api/internal/services/domain_services"
	"reservation-api/pkg/applogger"
	"reservation-api/pkg/message_broker"
	"reservation-api/pkg/multi_tenancy_database/tenant_database_resolver"
)

var (
	// ================================================================================================================
	logger             = applogger.New(nil)
	appConfig, confErr = appconfig.New()
	ctx                = context.Background()
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
