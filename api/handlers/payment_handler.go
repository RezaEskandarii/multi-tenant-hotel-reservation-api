// Package handlers
// handles all http requests
///**/
package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
	"reservation-api/internal/services/domain_services"
	"strconv"
)

type PaymentHandler struct {
	handlerBase
	PaymentService *domain_services.PaymentService
}

// Register PaymentHandler
// this method registers all routes,routeGroups and passes PaymentHandler's related dependencies
func (handler *PaymentHandler) Register(config *dto.HandlerConfig, service *domain_services.PaymentService) {
	handler.Router = config.Router
	handler.Logger = config.Logger
	handler.PaymentService = service
	handler.registerRoutes()
}

// @Tags Payment
// @Accept json
// @Produce json
// @Param X-Tenant-ID header int true "X-Tenant-ID"
// @Param  Payment body  models.Payment true "Payment"
// @Success 200 {object} models.Payment
// @Router /payments [post]
func (handler *PaymentHandler) create(c echo.Context) error {

	payment := models.Payment{}
	if err := c.Bind(&payment); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	if result, err := handler.PaymentService.Create(tenantContext(c), &payment); err != nil {
		handler.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	} else {
		return c.JSON(http.StatusOK, result)
	}
}

// @Tags Payment
// @Accept json
// @Param X-Tenant-ID header int true "X-Tenant-ID"
// @Produce json
// @Success 200 {array} models.Payment
// @Router /payments [delete]
func (handler *PaymentHandler) delete(c echo.Context) error {

	idStr := fmt.Sprintf("%s", c.Get("id"))
	id, _ := strconv.ParseUint(idStr, 10, 64)

	if err := handler.PaymentService.Delete(tenantContext(c), id); err != nil {
		handler.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}
	return c.JSON(http.StatusOK, nil)

}

// ============================= register routes ================================================== //
func (handler *PaymentHandler) registerRoutes() {
	routeGroup := handler.Router.Group("/payment")
	routeGroup.POST("", handler.create)
	routeGroup.DELETE("/:id", handler.delete)
}
