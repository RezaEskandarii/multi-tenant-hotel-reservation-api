// Package service_registry
// register all handlers,services and dependencies
///**/
package service_registry

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"net/http"
	"reservation-api/api/middlewares"
	"reservation-api/internal/dto"
)

// RegisterServicesAndRoutes register dependencies for services and handlers
func RegisterServicesAndRoutes(router *echo.Group) error {

	if confErr != nil {
		return confErr
	}

	router.Use(middleware.Gzip())

	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{appConfig.Application.AllowedOrigins},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch,
			http.MethodPost, http.MethodDelete},
	}))

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
	scheduleRemoveExpiredReservationRequests(reservationService, logger)

	// listen to message broker on reservation event and send email in background.
	go eventService.SendEmailToGuestOnReservation()

	return nil
}
