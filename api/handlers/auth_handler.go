// Package handlers
// handles all http requests
///**/
package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/services/domain_services"
	"reservation-api/internal_errors/message_keys"
	"reservation-api/pkg/applogger"
	"reservation-api/pkg/translator"
	"reservation-api/pkg/validator"
	"strings"
)

type AuthHandler struct {
	Router      *echo.Group
	Service     *domain_services.UserService
	AuthService *domain_services.AuthService
	logger      applogger.Logger
}

// Register AuthHandler
// this method registers all routes,routeGroups and passes AuthHandler's related dependencies
func (handler *AuthHandler) Register(config *dto.HandlerConfig, service *domain_services.UserService,
	authService *domain_services.AuthService) {

	handler.Router = config.Router
	handler.Service = service
	handler.logger = config.Logger
	handler.AuthService = authService
	handler.registerRoutes()
}

func (handler *AuthHandler) signin(c echo.Context) error {

	cerds := domain_services.Credentials{}
	if err := c.Bind(&cerds); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	if err, messages := validator.Validate(cerds); err != nil {
		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			Errors:       messages,
			ResponseCode: http.StatusBadRequest,
		})
	}

	if err, messages := validator.Validate(cerds); err != nil {
		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			Errors:       messages,
			ResponseCode: http.StatusBadRequest,
		})
	}

	if user, _ := handler.Service.FindByUsername(tenantContext(c), cerds.Username); user == nil {

		return c.JSON(http.StatusNotFound, commons.ApiResponse{
			Errors:       translator.Localize(c.Request().Context(), message_keys.UserNotFound),
			ResponseCode: http.StatusNotFound,
		})

	} else if user.IsActive == false {

		return c.JSON(http.StatusForbidden, commons.ApiResponse{
			Errors:       translator.Localize(c.Request().Context(), message_keys.UserIsDeActive),
			ResponseCode: http.StatusForbidden,
		})
	}

	if err, token := handler.AuthService.SignIn(tenantContext(c), cerds.Username, cerds.Password); err != nil {

		return c.JSON(http.StatusBadRequest, nil)

	} else {

		return c.JSON(http.StatusOK, token)
	}
}

// refresh token.
func (handler *AuthHandler) refreshToken(c echo.Context) error {
	tokenStr := c.Request().Header.Get("Authorization")

	if tokenStr == "" {
		return c.JSON(http.StatusBadRequest, "Authorization header is empty.")
	}

	authToken := strings.Split(tokenStr, "Bearer ")
	tokenStr = authToken[1]

	if err, result := handler.AuthService.RefreshToken(tokenStr); err != nil {
		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			Message: err.Error(),
		})

	} else {
		return c.JSON(http.StatusOK, result)
	}
}

// ============================= register routes ================================================== //
func (handler *AuthHandler) registerRoutes() {
	routeGroup := handler.Router.Group("/auth")
	routeGroup.POST("/signin", handler.signin)
	routeGroup.POST("/refresh-token", handler.refreshToken)
}
