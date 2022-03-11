package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	middlewares2 "reservation-api/api/middlewares"
	. "reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/message_keys"
	"reservation-api/internal/models"
	"reservation-api/internal/services/domain_services"
	"reservation-api/internal/utils"
)

// CurrencyHandler Currency endpoint handler
type CurrencyHandler struct {
	Service *domain_services.CurrencyService
	Input   *dto.HandlerInput
}

func (handler *CurrencyHandler) Register(input *dto.HandlerInput, service *domain_services.CurrencyService) {
	handler.Service = service
	handler.Input = input
	routeGroup := handler.Input.Router.Group("/currencies")
	routeGroup.POST("", handler.create)
	routeGroup.PUT("/:id", handler.update)
	routeGroup.GET("/:id", handler.find)
	routeGroup.GET("", handler.findAll, middlewares2.PaginationMiddleware)
}

/*====================================================================================*/
func (handler *CurrencyHandler) create(c echo.Context) error {

	model := &models.Currency{}
	lang := getAcceptLanguage(c)
	user := getCurrentUser(c)

	if err := c.Bind(&model); err != nil {
		handler.Input.Logger.LogError(err.Error())

		return c.JSON(http.StatusBadRequest,
			ApiResponse{
				ResponseCode: http.StatusInternalServerError,
				Message:      handler.Input.Translator.Localize(lang, message_keys.BadRequest),
			})
	}

	model.SetAudit(user)
	if result, err := handler.Service.Create(model); err == nil {

		return c.JSON(http.StatusBadRequest,
			ApiResponse{
				Data:         result,
				ResponseCode: http.StatusOK,
				Message:      handler.Input.Translator.Localize(lang, message_keys.Created),
			})
	} else {

		handler.Input.Logger.LogError(err.Error())

		return c.JSON(http.StatusInternalServerError,
			ApiResponse{
				Data:         nil,
				ResponseCode: http.StatusInternalServerError,
				Message:      "",
			})

	}
}

/*====================================================================================*/
func (handler *CurrencyHandler) update(c echo.Context) error {

	id, err := utils.ConvertToUint(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	user := getCurrentUser(c)
	model, err := handler.Service.Find(id)
	lang := getAcceptLanguage(c)

	if err != nil {
		handler.Input.Logger.LogError(err.Error())
		return c.JSON(http.StatusInternalServerError, ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusInternalServerError,
			Message:      "",
		})
	}

	if model == nil {
		return c.JSON(http.StatusNotFound, ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusNotFound,
			Message:      handler.Input.Translator.Localize(lang, message_keys.NotFound),
		})
	}

	model.SetUpdatedBy(user)
	if err := c.Bind(&model); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	if result, err := handler.Service.Update(model); err == nil {

		return c.JSON(http.StatusOK, ApiResponse{
			Data:         result,
			ResponseCode: http.StatusOK,
			Message:      handler.Input.Translator.Localize(lang, message_keys.Updated),
		})
	} else {

		handler.Input.Logger.LogError(err.Error())
		return c.JSON(http.StatusInternalServerError, nil)
	}
}

/*====================================================================================*/
func (handler *CurrencyHandler) find(c echo.Context) error {
	id, err := utils.ConvertToUint(c.Param("id"))
	if err != nil {

		handler.Input.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}

	model, err := handler.Service.Find(id)
	lang := getAcceptLanguage(c)

	if err != nil {

		handler.Input.Logger.LogError(err.Error())

		return c.JSON(http.StatusInternalServerError, ApiResponse{
			ResponseCode: http.StatusInternalServerError,
			Message:      "",
		})
	}

	if model == nil {
		return c.JSON(http.StatusNotFound, ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusNotFound,
			Message:      handler.Input.Translator.Localize(lang, message_keys.NotFound),
		})
	}

	return c.JSON(http.StatusOK, ApiResponse{
		Data:         model,
		ResponseCode: http.StatusOK,
		Message:      "",
	})
}

/*====================================================================================*/
func (handler *CurrencyHandler) findAll(c echo.Context) error {

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
