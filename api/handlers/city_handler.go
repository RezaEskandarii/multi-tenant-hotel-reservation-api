package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	middlewares2 "reservation-api/api/middlewares"
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/message_keys"
	"reservation-api/internal/models"
	"reservation-api/internal/services"
	"reservation-api/internal/utils"
	"reservation-api/pkg/applogger"
	"reservation-api/pkg/translator"
)

// CityHandler City endpoint handler
type CityHandler struct {
	Router     *echo.Group
	Service    *services.CityService
	translator *translator.Translator
	logger     applogger.Logger
}

func (handler *CityHandler) Register(input *dto.HandlerInput, service *services.CityService) {
	handler.Router = input.Router
	handler.Service = service
	handler.translator = input.Translator
	handler.logger = input.Logger

	routeGroup := handler.Router.Group("/cities")
	routeGroup.POST("", handler.create)
	routeGroup.PUT("/:id", handler.update)
	routeGroup.GET("/:id", handler.find)
	routeGroup.GET("", handler.findAll, middlewares2.PaginationMiddleware)
}

func (handler *CityHandler) create(c echo.Context) error {

	model := models.City{}
	lang := getAcceptLanguage(c)

	if err := c.Bind(&model); err != nil {
		return c.JSON(http.StatusBadRequest,
			commons.ApiResponse{
				Data:         nil,
				ResponseCode: http.StatusBadRequest,
				Message:      handler.translator.Localize(lang, message_keys.BadRequest),
			})
	}

	if _, err := handler.Service.Create(&model); err == nil {
		return c.JSON(http.StatusBadRequest,
			commons.ApiResponse{
				Data:         model,
				ResponseCode: http.StatusOK,
				Message:      handler.translator.Localize(lang, message_keys.Created),
			})
	} else {

		handler.logger.LogError(err.Error())
		return c.JSON(http.StatusInternalServerError,
			commons.ApiResponse{
				Data:         nil,
				ResponseCode: http.StatusInternalServerError,
				Message:      handler.translator.Localize(lang, message_keys.InternalServerError),
			})
	}

}

func (handler *CityHandler) update(c echo.Context) error {

	id, err := utils.ConvertToUint(c.Param("id"))
	if err != nil {
		handler.logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}
	lang := getAcceptLanguage(c)
	model, err := handler.Service.Find(id)

	if err != nil {

		handler.logger.LogError(err.Error())

		return c.JSON(http.StatusInternalServerError, commons.ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusInternalServerError,
			Message:      handler.translator.Localize(lang, message_keys.InternalServerError),
		})
	}

	if model == nil {
		return c.JSON(http.StatusNotFound, commons.ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusNotFound,
			Message:      handler.translator.Localize(lang, message_keys.NotFound),
		})
	}

	if err := c.Bind(&model); err != nil {
		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusBadRequest,
			Message:      handler.translator.Localize(lang, message_keys.BadRequest),
		})
	}

	if output, err := handler.Service.Update(model); err == nil {
		return c.JSON(http.StatusOK, commons.ApiResponse{
			Data:         output,
			ResponseCode: http.StatusOK,
			Message:      handler.translator.Localize(lang, message_keys.Updated),
		})
	} else {
		handler.logger.LogError(err.Error())
		return c.JSON(http.StatusInternalServerError, nil)
	}
}

func (handler *CityHandler) find(c echo.Context) error {

	id, err := utils.ConvertToUint(c.Param("id"))
	if err != nil {
		handler.logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}

	model, err := handler.Service.Find(id)
	lang := getAcceptLanguage(c)

	if err != nil {
		handler.logger.LogError(err.Error())
		return c.JSON(http.StatusInternalServerError, commons.ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusInternalServerError,
			Message:      handler.translator.Localize(lang, message_keys.InternalServerError),
		})
	}

	if model == nil {
		return c.JSON(http.StatusNotFound, commons.ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusNotFound,
			Message:      handler.translator.Localize(lang, message_keys.NotFound),
		})
	}

	return c.JSON(http.StatusOK, commons.ApiResponse{
		Data:         model,
		ResponseCode: http.StatusOK,
		Message:      "",
	})
}

func (handler *CityHandler) findAll(c echo.Context) error {

	paginationInput := c.Get(paginationInput).(*dto.PaginationInput)
	list, err := handler.Service.FindAll(paginationInput)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, commons.ApiResponse{
		Data:         list,
		ResponseCode: http.StatusOK,
		Message:      "",
	})
}
