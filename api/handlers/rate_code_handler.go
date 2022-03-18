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

// RateCodeHandler RateCode endpoint handler
type RateCodeHandler struct {
	Service               *domain_services.RateCodeService
	RateCodeDetailService *domain_services.RateCodeDetailService
	Input                 *dto.HandlerInput
}

func (handler *RateCodeHandler) Register(input *dto.HandlerInput, service *domain_services.RateCodeService, rateCodeDetailService *domain_services.RateCodeDetailService) {
	handler.Service = service
	handler.RateCodeDetailService = rateCodeDetailService
	handler.Input = input
	routeGroup := handler.Input.Router.Group("/rate-codes")
	routeGroup.POST("", handler.create)
	routeGroup.PUT("/:id", handler.update)
	routeGroup.GET("/:id", handler.find)
	routeGroup.DELETE("/:id", handler.delete)
	routeGroup.GET("", handler.findAll, middlewares2.PaginationMiddleware)
	routeGroup.POST("/add-details/:id", handler.addDetails)
}

/*====================================================================================*/
func (handler *RateCodeHandler) create(c echo.Context) error {

	model := &models.RateCode{}
	user := getCurrentUser(c)
	lang := c.Request().Header.Get(acceptLanguage)

	if err := c.Bind(&model); err != nil {

		handler.Input.Logger.LogError(err.Error())

		return c.JSON(http.StatusBadRequest,
			commons.ApiResponse{
				Data:         nil,
				ResponseCode: http.StatusBadRequest,
				Message:      handler.Input.Translator.Localize(lang, message_keys.BadRequest),
			})
	}
	model.SetAudit(user)
	result, err := handler.Service.Create(model)

	if err != nil {
		handler.Input.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			ResponseCode: http.StatusBadRequest,
			Message:      err.Error(),
		})
	}

	return c.JSON(http.StatusBadRequest, commons.ApiResponse{
		ResponseCode: http.StatusOK,
		Message:      handler.Input.Translator.Localize(lang, message_keys.Created),
		Data:         result,
	})
}

/*====================================================================================*/
func (handler *RateCodeHandler) update(c echo.Context) error {

	id, err := utils.ConvertToUint(c.Param("id"))
	user := getCurrentUser(c)

	if err != nil {
		handler.Input.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}

	model, err := handler.Service.Find(id)
	lang := c.Request().Header.Get(acceptLanguage)

	if err != nil {
		handler.Input.Logger.LogError(err.Error())
		return c.JSON(http.StatusInternalServerError, commons.ApiResponse{
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
	model.SetUpdatedBy(user)
	if err := c.Bind(&model); err != nil {

		handler.Input.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)

	}

	if result, err := handler.Service.Update(model); err == nil {

		return c.JSON(http.StatusOK, commons.ApiResponse{
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
func (handler *RateCodeHandler) find(c echo.Context) error {
	id, err := utils.ConvertToUint(c.Param("id"))
	if err != nil {
		handler.Input.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}
	model, err := handler.Service.Find(id)
	lang := c.Request().Header.Get(acceptLanguage)

	if err != nil {

		handler.Input.Logger.LogError(err.Error())

		return c.JSON(http.StatusInternalServerError, commons.ApiResponse{
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

/*====================================================================================*/
func (handler *RateCodeHandler) findAll(c echo.Context) error {

	paginationInput := c.Get(paginationInput).(*dto.PaginationFilter)

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

/*====================================================================================*/
func (handler *RateCodeHandler) delete(c echo.Context) error {

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

/*====================================================================================*/
func (handler *RateCodeHandler) addDetails(c echo.Context) error {

	id, err := utils.ConvertToUint(c.Get("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			ResponseCode: http.StatusBadRequest,
		})
	}
	requestBody := models.RateCodeDetail{RateCodeId: id}
	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			ResponseCode: http.StatusBadRequest,
		})
	}
	// create new rateCodeDetail.
	result, err := handler.RateCodeDetailService.Create(&requestBody)
	if err != nil {
		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			ResponseCode: http.StatusBadRequest,
			Message:      err.Error(),
		})
	}

	return c.JSON(http.StatusOK, commons.ApiResponse{
		Data:         result,
		ResponseCode: http.StatusOK,
		Message:      handler.Input.Translator.Localize(c.Request().Header.Get(acceptLanguage), message_keys.Created),
	})
}
