package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/message_keys"
	"reservation-api/internal/models"
	"reservation-api/internal/services/domain_services"
)

type TenantHandler struct {
	TenantService *domain_services.TenantService
	Input         *dto.HandlersShared
}

func (handler *TenantHandler) Register(input *dto.HandlersShared, service *domain_services.TenantService) {
	handler.TenantService = service
	handler.Input = input
	routeGroup := input.Router.Group("/tenants")
	routeGroup.POST("", handler.create)
}

func (handler *TenantHandler) create(c echo.Context) error {

	model := &models.Tenant{}
	lang := c.Request().Header.Get(acceptLanguage)

	if err := c.Bind(&model); err != nil {
		handler.Input.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusBadRequest,
			Message:      message_keys.BadRequest,
		})
	}

	result, err := handler.TenantService.Create(model)

	if err != nil {
		handler.Input.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusBadRequest,
			Message:      err.Error(),
		})
	}

	return c.JSON(http.StatusOK, commons.ApiResponse{
		Data:         result,
		ResponseCode: http.StatusOK,
		Message:      handler.Input.Translator.Localize(lang, message_keys.Created),
	})
}
