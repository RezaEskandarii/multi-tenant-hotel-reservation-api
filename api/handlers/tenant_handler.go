// Package handlers
// handles all http requests
///**/
package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
	"reservation-api/internal/services/domain_services"
	"reservation-api/internal_errors/message_keys"
	"reservation-api/pkg/translator"
)

type TenantHandler struct {
	handlerBase
	TenantService *domain_services.TenantService
}

// Register TenantHandler
// this method registers all routes,routeGroups and passes TenantHandler's related dependencies
func (handler *TenantHandler) Register(config *dto.HandlerConfig, service *domain_services.TenantService) {
	handler.TenantService = service
	handler.Router = config.Router
	handler.Logger = config.Logger
	handler.registerRoutes()
}

// @Tags RoomType
// @Accept json
// @Param X-Tenant-ID header int true "X-Tenant-ID"
// @Produce json
// @Param  RoomType body  models.RoomType true "RoomType"
// @Success 200 {object} models.RoomType
// @Router /room-types [post]
func (handler *TenantHandler) create(c echo.Context) error {

	tenant := &models.Tenant{}

	if err := c.Bind(&tenant); err != nil {
		handler.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, commons.ApiResponse{

			ResponseCode: http.StatusBadRequest,
			Message:      message_keys.BadRequest,
		})
	}

	result, err := handler.TenantService.SetUp(tenantContext(c), tenant)

	if err != nil {
		handler.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, commons.ApiResponse{

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

// ============================= register routes ================================================== //
func (handler *TenantHandler) registerRoutes() {
	routeGroup := handler.Router.Group("/tenants")
	routeGroup.POST("", handler.create)
}
