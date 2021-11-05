package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/message_keys"
	"reservation-api/internal/models"
	"reservation-api/internal/services"
	"reservation-api/internal/utils"
	"reservation-api/pkg/applogger"
	"reservation-api/pkg/translator"
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

	routeGroup := handler.Router.Group("/guests")
	routeGroup.POST("", handler.create)
	routeGroup.GET("/:id", handler.find)
	routeGroup.GET("", handler.findAll)
	routeGroup.PUT("", handler.update)
	//routeGroup.DELETE("", handler)
}

func (handler *GuestHandler) create(c echo.Context) error {

	model := models.Guest{}

	lang := c.Request().Header.Get(acceptLanguage)

	if err := c.Bind(&model); err != nil {

		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			ResponseCode: http.StatusBadRequest,
			Message:      handler.translator.Localize(c.Request().Header.Get(acceptLanguage), message_keys.BadRequest),
		})
	}

	if _, err := handler.Service.Create(&model); err != nil {

		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			Message: handler.translator.Localize(lang, err.Error()),
		})
	}

	return c.JSON(http.StatusOK, commons.ApiResponse{
		Data:    model,
		Message: handler.translator.Localize(lang, message_keys.Created),
	})
}

func (handler *GuestHandler) update(c echo.Context) error {

	model := models.Guest{}
	lang := c.Request().Header.Get(acceptLanguage)

	id, _ := utils.ConvertToUint(c.Get("id"))

	guest, _ := handler.Service.Find(id)

	if guest == nil || (guest != nil && guest.Id == 0) {

		return c.JSON(http.StatusNotFound, commons.ApiResponse{
			Message: handler.translator.Localize(lang, message_keys.NotFound),
		})
	}

	if err := c.Bind(&model); err != nil {

		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			ResponseCode: http.StatusBadRequest,
			Message:      handler.translator.Localize(c.Request().Header.Get(acceptLanguage), message_keys.BadRequest),
		})
	}

	if _, err := handler.Service.Update(&model); err != nil {

		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			Message: handler.translator.Localize(lang, err.Error()),
		})
	}

	return c.JSON(http.StatusOK, commons.ApiResponse{
		Data:    model,
		Message: handler.translator.Localize(lang, message_keys.Updated),
	})
}

func (handler *GuestHandler) find(c echo.Context) error {

	lang := c.Request().Header.Get(acceptLanguage)
	id, _ := utils.ConvertToUint(c.Get("id"))

	guest, _ := handler.Service.Find(id)

	if guest == nil || (guest != nil && guest.Id == 0) {

		return c.JSON(http.StatusNotFound, commons.ApiResponse{
			Message: handler.translator.Localize(lang, message_keys.NotFound),
		})
	}

	return c.JSON(http.StatusOK, commons.ApiResponse{
		Data: guest,
	})
}

func (handler *GuestHandler) findAll(c echo.Context) error {

	lang := c.Request().Header.Get(acceptLanguage)

	page, _ := utils.ConvertToUint(c.Param("page"))
	perPage, _ := utils.ConvertToUint(c.Param("perPage"))

	input := &dto.PaginationInput{
		Page:    int(page),
		PerPage: int(perPage),
	}

	result, err := handler.Service.FindAll(input)

	if err != nil {

		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			Message: handler.translator.Localize(lang, err.Error()),
		})
	}

	return c.JSON(http.StatusOK, commons.ApiResponse{
		Data: result,
	})
}
