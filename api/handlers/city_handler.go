package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	middlewares2 "reservation-api/api/middlewares"
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/message_keys"
	"reservation-api/internal/models"
	"reservation-api/internal/services/domain_services"
	"reservation-api/internal/utils"
)

// CityHandler City endpoint handler
type CityHandler struct {
	Service *domain_services.CityService
	Input   *dto.HandlerInput
}

func (handler *CityHandler) Register(input *dto.HandlerInput, service *domain_services.CityService) {
	handler.Service = service
	handler.Input = input
	routeGroup := handler.Input.Router.Group("/cities")
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
				Message:      handler.Input.Translator.Localize(lang, message_keys.BadRequest),
			})
	}

	if _, err := handler.Service.Create(&model); err == nil {
		handler.Input.AuditChannel <- model

		return c.JSON(http.StatusBadRequest,
			commons.ApiResponse{
				Data:         model,
				ResponseCode: http.StatusOK,
				Message:      handler.Input.Translator.Localize(lang, message_keys.Created),
			})
	} else {

		handler.Input.Logger.LogError(err.Error())
		return c.JSON(http.StatusInternalServerError,
			commons.ApiResponse{
				Data:         nil,
				ResponseCode: http.StatusInternalServerError,
				Message:      handler.Input.Translator.Localize(lang, message_keys.InternalServerError),
			})
	}

}

func (handler *CityHandler) update(c echo.Context) error {

	id, err := utils.ConvertToUint(c.Param("id"))
	if err != nil {
		handler.Input.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}
	lang := getAcceptLanguage(c)
	model, err := handler.Service.Find(id)

	if err != nil {

		handler.Input.Logger.LogError(err.Error())

		return c.JSON(http.StatusInternalServerError, commons.ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusInternalServerError,
			Message:      handler.Input.Translator.Localize(lang, message_keys.InternalServerError),
		})
	}

	if model == nil {
		return c.JSON(http.StatusNotFound, commons.ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusNotFound,
			Message:      handler.Input.Translator.Localize(lang, message_keys.NotFound),
		})
	}

	if err := c.Bind(&model); err != nil {
		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusBadRequest,
			Message:      handler.Input.Translator.Localize(lang, message_keys.BadRequest),
		})
	}

	if output, err := handler.Service.Update(model); err == nil {
		handler.Input.AuditChannel <- output

		return c.JSON(http.StatusOK, commons.ApiResponse{
			Data:         output,
			ResponseCode: http.StatusOK,
			Message:      handler.Input.Translator.Localize(lang, message_keys.Updated),
		})
	} else {
		handler.Input.Logger.LogError(err.Error())
		return c.JSON(http.StatusInternalServerError, nil)
	}
}

func (handler *CityHandler) find(c echo.Context) error {

	id, err := utils.ConvertToUint(c.Param("id"))
	if err != nil {
		handler.Input.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}

	model, err := handler.Service.Find(id)
	lang := getAcceptLanguage(c)

	if err != nil {
		handler.Input.Logger.LogError(err.Error())
		return c.JSON(http.StatusInternalServerError, commons.ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusInternalServerError,
			Message:      handler.Input.Translator.Localize(lang, message_keys.InternalServerError),
		})
	}

	if model == nil {
		return c.JSON(http.StatusNotFound, commons.ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusNotFound,
			Message:      handler.Input.Translator.Localize(lang, message_keys.NotFound),
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
