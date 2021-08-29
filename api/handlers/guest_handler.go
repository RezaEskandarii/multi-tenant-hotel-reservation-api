package handlers

import (
	"github.com/labstack/echo/v4"
	"hotel-reservation/internal/commons"
	"hotel-reservation/internal/dto"
	"hotel-reservation/internal/message_keys"
	"hotel-reservation/internal/models"
	"hotel-reservation/internal/services"
	"hotel-reservation/pkg/applogger"
	"hotel-reservation/pkg/translator"
	"net/http"
)

// GuestHandler  Currency endpoint handler
type GuestHandler struct {
	Router     *echo.Group
	Service    *services.GuestService
	translator *translator.Translator
	logger     applogger.Logger
}

func (handler *GuestHandler) Register(input *dto.HandlerInput, service *services.GuestService) {
	handler.Router = input.Router
	handler.Service = service
	handler.translator = input.Translator
	handler.logger = input.Logger

	routeGroup := handler.Router.Group("/currencies")

	routeGroup.POST("", handler.create)
}

func (handler *GuestHandler) create(c echo.Context) error {

	model := models.Guest{}

	if err := c.Bind(&model); err != nil {

		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			ResponseCode: http.StatusBadRequest,
			Message:      handler.translator.Localize(c.Request().Header.Get(acceptLanguage), message_keys.BadRequest),
		})
	}

	panic("not implemetend !!!")
}
