package handlers

import (
	"github.com/labstack/echo/v4"
	middlewares2 "hotel-reservation/api/middlewares"
	"hotel-reservation/internal/commons"
	"hotel-reservation/internal/dto"
	"hotel-reservation/internal/message_keys"
	"hotel-reservation/internal/models"
	"hotel-reservation/internal/services"
	"hotel-reservation/internal/utils"
	"hotel-reservation/pkg/applogger"
	"hotel-reservation/pkg/translator"
	"net/http"
)

// RateGroupHandler RateGroup endpoint handler
type RateGroupHandler struct {
	Router     *echo.Group
	Service    *services.RateGroupService
	translator *translator.Translator
	logger     applogger.Logger
}

func (handler *RateGroupHandler) Register(input *dto.HandlerInput, service *services.RateGroupService) {
	handler.Router = input.Router
	handler.Service = service
	handler.translator = input.Translator
	handler.logger = input.Logger

	routeGroup := handler.Router.Group("/rate-groups")

	routeGroup.POST("", handler.create)
	routeGroup.PUT("/:id", handler.update)
	routeGroup.GET("/:id", handler.find)
	routeGroup.DELETE("/:id", handler.delete)
	routeGroup.GET("", handler.findAll, middlewares2.PaginationMiddleware)
}

func (handler *RateGroupHandler) create(c echo.Context) error {

	model := &models.RateGroup{}

	lang := c.Request().Header.Get(acceptLanguage)

	if err := c.Bind(&model); err != nil {

		handler.logger.LogError(err.Error())

		return c.JSON(http.StatusBadRequest,
			commons.ApiResponse{
				Data:         nil,
				ResponseCode: http.StatusBadRequest,
				Message:      handler.translator.Localize(lang, message_keys.BadRequest),
			})
	}

	output, err := handler.Service.Create(model)

	if err != nil {

		handler.logger.LogError(err.Error())

		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			ResponseCode: http.StatusBadRequest,
			Message:      err.Error(),
		})
	}

	return c.JSON(http.StatusBadRequest, commons.ApiResponse{
		ResponseCode: http.StatusOK,
		Message:      handler.translator.Localize(lang, message_keys.Created),
		Data:         output,
	})
}

func (handler *RateGroupHandler) update(c echo.Context) error {

	id, err := utils.ConvertToUint(c.Param("id"))
	if err != nil {

		handler.logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}
	model, err := handler.Service.Find(id)

	lang := c.Request().Header.Get(acceptLanguage)

	if err != nil {

		handler.logger.LogError(err.Error())

		return c.JSON(http.StatusInternalServerError, commons.ApiResponse{
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

		handler.logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)

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

func (handler *RateGroupHandler) find(c echo.Context) error {
	id, err := utils.ConvertToUint(c.Param("id"))
	if err != nil {
		handler.logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}
	model, err := handler.Service.Find(id)
	lang := c.Request().Header.Get(acceptLanguage)

	if err != nil {

		handler.logger.LogError(err.Error())

		return c.JSON(http.StatusInternalServerError, commons.ApiResponse{
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

func (handler *RateGroupHandler) findAll(c echo.Context) error {

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

func (handler *RateGroupHandler) delete(c echo.Context) error {

	id, err := utils.ConvertToUint(c.Param("id"))
	lang := c.Request().Header.Get(acceptLanguage)

	if err != nil {

		handler.logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			ResponseCode: http.StatusBadRequest,
			Message:      handler.translator.Localize(lang, message_keys.BadRequest),
		})
	}

	err = handler.Service.Delete(id)

	if err != nil {

		handler.logger.LogError(err.Error())
		return c.JSON(http.StatusConflict, commons.ApiResponse{
			ResponseCode: http.StatusConflict,
			Message:      handler.translator.Localize(lang, err.Error()),
		})
	}

	return c.JSON(http.StatusOK, commons.ApiResponse{
		ResponseCode: http.StatusOK,
		Message:      handler.translator.Localize(lang, message_keys.Deleted),
	})
}
