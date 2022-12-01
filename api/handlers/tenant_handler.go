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
	handlerBase
	TenantService *domain_services.TenantService
}

func (handler *TenantHandler) Register(config *dto.HandlerConfig, service *domain_services.TenantService) {
	handler.TenantService = service
	handler.Router = config.Router
	handler.Logger = config.Logger

	routeGroup := handler.Router.Group("/tenants")
	routeGroup.POST("", handler.create)
}

// @Summary crete RoomType
// @Tags RoomType
// @Accept json
// @Param X-TenantID header int true "X-TenantID"
// @Produce json
// @Param  RoomType body  models.RoomType true "RoomType"
// @Success 200 {object} models.RoomType
// @Router /room-types [post]
func (handler *TenantHandler) create(c echo.Context) error {

	model := &models.Tenant{}

	if err := c.Bind(&model); err != nil {
		handler.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusBadRequest,
			Message:      message_keys.BadRequest,
		})
	}

	result, err := handler.TenantService.SetUp(tenantContext(c), model)

	if err != nil {
		handler.Logger.LogError(err.Error())
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
