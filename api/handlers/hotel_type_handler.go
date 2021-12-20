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
)

// HotelTypeHandler Province endpoint handler
type HotelTypeHandler struct {
	Service *services.HotelTypeService
	Input   *dto.HandlerInput
}

func (handler *HotelTypeHandler) Register(input *dto.HandlerInput, service *services.HotelTypeService) {
	handler.Service = service
	handler.Input = input
	routeGroup := input.Router.Group("/hotel-types")

	routeGroup.POST("", handler.create)
	routeGroup.PUT("/:id", handler.update)
	routeGroup.GET("/:id", handler.find)
	routeGroup.DELETE("/:id", handler.delete)
	routeGroup.GET("", handler.findAll, middlewares2.PaginationMiddleware)
}

func (handler *HotelTypeHandler) create(c echo.Context) error {

	model := &models.HotelType{}
	lang := c.Request().Header.Get(acceptLanguage)

	if err := c.Bind(&model); err != nil {

		handler.Input.Logger.LogError(err.Error())

		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusBadRequest,
			Message:      message_keys.BadRequest,
		})
	}

	result, err := handler.Service.Create(model)

	if err != nil {

		handler.Input.Logger.LogError(err.Error())

		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusBadRequest,
			Message:      "",
		})
	}

	return c.JSON(http.StatusOK, commons.ApiResponse{
		Data:         result,
		ResponseCode: http.StatusOK,
		Message:      handler.Input.Translator.Localize(lang, message_keys.Created),
	})
}

func (handler *HotelTypeHandler) update(c echo.Context) error {

	lang := c.Request().Header.Get(acceptLanguage)
	id, err := utils.ConvertToUint(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	result, err := handler.Service.Find(id)
	if err != nil {

		handler.Input.Logger.LogError(err.Error())

		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			ResponseCode: http.StatusBadRequest,
			Message:      handler.Input.Translator.Localize(lang, message_keys.BadRequest),
		})
	}

	if result == nil || (result != nil && result.Id == 0) {
		return c.JSON(http.StatusNotFound, commons.ApiResponse{
			ResponseCode: http.StatusNotFound,
			Message:      handler.Input.Translator.Localize(lang, message_keys.NotFound),
		})
	}

	name := c.FormValue("name")
	result.Name = name

	updatedMode, err := handler.Service.Update(result)

	if err != nil {

		handler.Input.Logger.LogError(err.Error())
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, commons.ApiResponse{
		Data:         updatedMode,
		ResponseCode: http.StatusOK,
	})
}

func (handler *HotelTypeHandler) find(c echo.Context) error {

	lang := c.Request().Header.Get(acceptLanguage)
	id, err := utils.ConvertToUint(c.Param("id"))

	if err != nil {

		handler.Input.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}

	result, err := handler.Service.Find(id)

	if err != nil {

		handler.Input.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			ResponseCode: http.StatusBadRequest,
			Message:      handler.Input.Translator.Localize(lang, message_keys.BadRequest),
		})
	}

	if result == nil || (result != nil && result.Id == 0) {
		return c.JSON(http.StatusNotFound, commons.ApiResponse{
			ResponseCode: http.StatusNotFound,
			Message:      handler.Input.Translator.Localize(lang, message_keys.NotFound),
		})
	}

	return c.JSON(http.StatusOK, commons.ApiResponse{
		Data:         result,
		ResponseCode: http.StatusOK,
	})
}

func (handler *HotelTypeHandler) findAll(c echo.Context) error {

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

func (handler *HotelTypeHandler) delete(c echo.Context) error {

	id, err := utils.ConvertToUint(c.Param("id"))
	lang := c.Request().Header.Get(acceptLanguage)

	if err != nil {

		handler.Input.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			ResponseCode: http.StatusBadRequest,
			Message:      handler.Input.Translator.Localize(lang, message_keys.BadRequest),
		})
	}

	err = handler.Service.Delete(id)

	if err != nil {

		handler.Input.Logger.LogError(err.Error())
		return c.JSON(http.StatusConflict, commons.ApiResponse{
			ResponseCode: http.StatusConflict,
			Message:      handler.Input.Translator.Localize(lang, err.Error()),
		})
	}

	return c.JSON(http.StatusOK, commons.ApiResponse{
		ResponseCode: http.StatusOK,
		Message:      handler.Input.Translator.Localize(lang, message_keys.Deleted),
	})
}