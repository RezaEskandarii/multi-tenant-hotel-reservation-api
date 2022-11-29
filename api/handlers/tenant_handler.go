package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/message_keys"
	"reservation-api/internal/models"
	"reservation-api/internal/services/domain_services"
	"reservation-api/pkg/translator"
)

type TenantHandler struct {
	TenantService *domain_services.TenantService
	Config        *dto.HandlerConfig
}

func (handler *TenantHandler) Register(config *dto.HandlerConfig, service *domain_services.TenantService) {
	handler.TenantService = service
	handler.Config = config
	routeGroup := config.Router.Group("/tenants")
	routeGroup.POST("", handler.create)
}

func (handler *TenantHandler) create(c echo.Context) error {

	model := &models.Tenant{}

	if err := c.Bind(&model); err != nil {
		handler.Config.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusBadRequest,
			Message:      message_keys.BadRequest,
		})
	}

	result, err := handler.TenantService.SetUp(tenantContext(c), model)

	if err != nil {
		handler.Config.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusBadRequest,
			Message:      err.Error(),
		})
	}

	return c.JSON(http.StatusOK, commons.ApiResponse{
		Data:         result,
		ResponseCode: http.StatusOK,
		Message:      translator.Localize(c.Request().Context(), message_keys.Created),
	})
}
