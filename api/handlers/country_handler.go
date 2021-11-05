package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	middlewares2 "reservation-api/api/middlewares"
	. "reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/message_keys"
	"reservation-api/internal/models"
	"reservation-api/internal/services"
	"reservation-api/internal/utils"
	"reservation-api/pkg/applogger"
	"reservation-api/pkg/translator"
)

// CountryHandler country endpoint handler
type CountryHandler struct {
	Router     *echo.Group
	Service    *services.CountryService
	translator *translator.Translator
	logger     applogger.Logger
}

func (handler *CountryHandler) Register(input *dto.HandlerInput, service *services.CountryService) {
	handler.Router = input.Router
	handler.Service = service
	handler.translator = input.Translator
	handler.logger = input.Logger

	routeGroup := handler.Router.Group("/countries")
	routeGroup.POST("", handler.create)
	routeGroup.PUT("/:id", handler.update)
	routeGroup.GET("/:id", handler.find)
	routeGroup.GET("/:id/provinces", handler.provinces)
	routeGroup.GET("", handler.findAll, middlewares2.PaginationMiddleware)
}

func (handler *CountryHandler) create(c echo.Context) error {

	model := &models.Country{}
	lang := getAcceptLanguage(c)

	if err := c.Bind(&model); err != nil {
		handler.logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest,
			ApiResponse{
				Data:         nil,
				ResponseCode: http.StatusBadRequest,
				Message:      handler.translator.Localize(lang, message_keys.BadRequest),
			})
	}

	output, err := handler.Service.Create(model)
	if err != nil {
		handler.logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, ApiResponse{
			ResponseCode: http.StatusBadRequest,
			Message:      err.Error(),
		})
	}

	return c.JSON(http.StatusBadRequest, ApiResponse{
		ResponseCode: http.StatusOK,
		Message:      handler.translator.Localize(lang, message_keys.Created),
		Data:         output,
	})
}

func (handler *CountryHandler) update(c echo.Context) error {

	id, err := utils.ConvertToUint(c.Param("id"))
	if err != nil {

		handler.logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}

	model, err := handler.Service.Find(id)
	lang := getAcceptLanguage(c)

	if err != nil {
		handler.logger.LogError(err.Error())
		return c.JSON(http.StatusInternalServerError, ApiResponse{
			ResponseCode: http.StatusInternalServerError,
			Message:      handler.translator.Localize(lang, message_keys.InternalServerError),
		})
	}

	if model == nil {
		return c.JSON(http.StatusNotFound, ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusNotFound,
			Message:      handler.translator.Localize(lang, message_keys.NotFound),
		})
	}

	if err := c.Bind(&model); err != nil {

		handler.logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)

	}

	if output, err := handler.Service.Update(model); err == nil {

		return c.JSON(http.StatusOK, ApiResponse{
			Data:         output,
			ResponseCode: http.StatusOK,
			Message:      handler.translator.Localize(lang, message_keys.Updated),
		})
	} else {

		handler.logger.LogError(err.Error())
		return c.JSON(http.StatusInternalServerError, nil)
	}
}

func (handler *CountryHandler) find(c echo.Context) error {
	id, err := utils.ConvertToUint(c.Param("id"))
	if err != nil {
		handler.logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}

	model, err := handler.Service.Find(id)
	lang := getAcceptLanguage(c)

	if err != nil {

		handler.logger.LogError(err.Error())

		return c.JSON(http.StatusInternalServerError, ApiResponse{
			ResponseCode: http.StatusInternalServerError,
			Message:      handler.translator.Localize(lang, message_keys.InternalServerError),
		})
	}

	if model == nil {
		return c.JSON(http.StatusNotFound, ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusNotFound,
			Message:      handler.translator.Localize(lang, message_keys.NotFound),
		})
	}

	return c.JSON(http.StatusOK, ApiResponse{
		Data:         model,
		ResponseCode: http.StatusOK,
		Message:      "",
	})
}

func (handler *CountryHandler) findAll(c echo.Context) error {

	paginationInput := c.Get(paginationInput).(*dto.PaginationInput)

	list, err := handler.Service.FindAll(paginationInput)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, ApiResponse{
		Data:         list,
		ResponseCode: http.StatusOK,
		Message:      "",
	})
}

func (handler *CountryHandler) provinces(c echo.Context) error {

	id, err := utils.ConvertToUint(c.Param("id"))
	if err != nil {
		handler.logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}
	provinces, err := handler.Service.GetProvinces(id)

	if err != nil {

		handler.logger.LogError(err.Error())

		return c.JSON(http.StatusInternalServerError, ApiResponse{
			ResponseCode: http.StatusInternalServerError,
			Message:      err.Error(),
		})
	}

	return c.JSON(http.StatusOK, ApiResponse{
		Data:         provinces,
		ResponseCode: http.StatusOK,
		Message:      "",
	})
}
