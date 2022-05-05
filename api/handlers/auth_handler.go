package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/services/domain_services"
	"reservation-api/pkg/applogger"
	"reservation-api/pkg/translator"
	"strings"
)

type AuthHandler struct {
	Router      *echo.Group
	Service     *domain_services.UserService
	AuthService *domain_services.AuthService
	translator  *translator.Translator
	logger      applogger.Logger
}

func (handler *AuthHandler) Register(input *dto.HandlersShared, service *domain_services.UserService,
	authService *domain_services.AuthService) {

	handler.Router = input.Router
	handler.Service = service
	handler.translator = input.Translator
	handler.logger = input.Logger
	handler.AuthService = authService
	routeGroup := handler.Router.Group("/auth")
	routeGroup.POST("/signin", handler.signin)
	routeGroup.POST("/refresh-token", handler.refreshToken)
}

func (handler *AuthHandler) signin(c echo.Context) error {

	cerds := domain_services.Credentials{}
	if err := c.Bind(&cerds); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	if err, token := handler.AuthService.SignIn(cerds.Username, cerds.Password); err != nil {

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
