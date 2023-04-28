// Package handlers
// handles all http requests
///**/
package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/global_variables"
	"reservation-api/internal/models"
	"reservation-api/internal/services/common_services"
	"reservation-api/internal/services/domain_services"
	"reservation-api/internal/utils"
	"reservation-api/internal_errors/message_keys"
	"reservation-api/pkg/translator"
)

// GuestHandler  Guest endpoint handler
type GuestHandler struct {
	handlerBase
	Service       *domain_services.GuestService
	ReportService *common_services.ReportService
}

// Register GuestHandler
// this method registers all routes,routeGroups and passes GuestHandler's related dependencies
func (handler *GuestHandler) Register(config *dto.HandlerConfig,
	service *domain_services.GuestService, reportService *common_services.ReportService) {
	handler.ReportService = reportService
	handler.Router = config.Router
	handler.Logger = config.Logger
	handler.Service = service
	handler.registerRoutes()
}

// @Tags Guest
// @Accept json
// @Produce json
// @Param X-Tenant-ID header int true "X-Tenant-ID"
// @Param  Guest body  models.Guest true "Guest"
// @Success 200 {object} models.Guest
// @Router /guests [post]
func (handler *GuestHandler) create(c echo.Context) error {

	guest := models.Guest{}
	user := currentUser(c)

	if err := c.Bind(&guest); err != nil {
		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			ResponseCode: http.StatusBadRequest,
			Message:      translator.Localize(c.Request().Context(), message_keys.BadRequest),
		})
	}

	guest.SetAudit(user)
	if _, err := handler.Service.Create(tenantContext(c), &guest); err != nil {
		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			Message: translator.Localize(c.Request().Context(), err.Error()),
		})
	}

	return c.JSON(http.StatusOK, commons.ApiResponse{
		Data:    guest,
		Message: translator.Localize(c.Request().Context(), message_keys.Created),
	})
}

// @Tags Guest
// @Accept json
// @Param X-Tenant-ID header int true "X-Tenant-ID"
// @Param Id path int true "Id"
// @Param Guest body models.Guest true "Guest"
// @Produce json
// @Param  Guest body  models.Guest true "Guest"
// @Success 200 {object} models.Guest
// @Router /guests/{id} [put]
func (handler *GuestHandler) update(c echo.Context) error {

	model := models.Guest{}
	user := currentUser(c)
	id, _ := utils.ConvertToUint(c.Param("id"))
	guest, _ := handler.Service.Find(tenantContext(c), id)

	if guest == nil || (guest != nil && guest.Id == 0) {
		return c.JSON(http.StatusNotFound, commons.ApiResponse{
			Message: translator.Localize(c.Request().Context(), message_keys.NotFound),
		})
	}

	if err := c.Bind(&model); err != nil {
		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			ResponseCode: http.StatusBadRequest,
			Message:      translator.Localize(c.Request().Context(), message_keys.BadRequest),
		})
	}

	model.SetUpdatedBy(user)
	if _, err := handler.Service.Update(tenantContext(c), &model); err != nil {
		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			Message: translator.Localize(c.Request().Context(), err.Error()),
		})
	}

	return c.JSON(http.StatusOK, commons.ApiResponse{
		Data:    model,
		Message: translator.Localize(c.Request().Context(), message_keys.Updated),
	})
}

// @Tags Guest
// @Accept json
// @Param X-Tenant-ID header int true "X-Tenant-ID"
// @Param Id path int true "Id"
// @Produce json
// @Success 200 {object} models.Guest
// @Router /guests/{id} [get]
func (handler *GuestHandler) find(c echo.Context) error {

	id, _ := utils.ConvertToUint(c.Param("id"))
	guest, _ := handler.Service.Find(tenantContext(c), id)

	if guest == nil || (guest != nil && guest.Id == 0) {
		return c.JSON(http.StatusNotFound, commons.ApiResponse{
			Message: translator.Localize(c.Request().Context(), message_keys.NotFound),
		})
	}

	return c.JSON(http.StatusOK, commons.ApiResponse{
		Data: guest,
	})
}

// @Tags Guest
// @Accept json
// @Param X-Tenant-ID header int true "X-Tenant-ID"
// @Produce json
// @Success 200 {array} models.Guest
// @Router /guests [get]
func (handler *GuestHandler) findAll(c echo.Context) error {

	page, _ := utils.ConvertToUint(c.Param("page"))
	perPage, _ := utils.ConvertToUint(c.Param("perPage"))
	output := getOutputQueryParamVal(c)

	input := &dto.PaginationFilter{
		Page:     int(page),
		PageSize: int(perPage),
	}

	input.IgnorePagination = output != ""

	result, err := handler.Service.FindAll(tenantContext(c), input)

	if err != nil {
		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			Message: translator.Localize(c.Request().Context(), err.Error()),
		})
	}

	if output != "" {
		if output == EXCEL {
			report, err := handler.ReportService.ExportToExcel(result, c.Request().Context().Value(global_variables.CurrentLang).(string))
			if err != nil {
				handler.Logger.LogError(err.Error())
				return c.JSON(http.StatusInternalServerError, commons.ApiResponse{})
			}
			setBinaryHeaders(c, "guests", EXCEL_OUTPUT)
			c.Response().Write(report)
			return nil
		}
	}

	return c.JSON(http.StatusOK, commons.ApiResponse{
		Data: result,
	})
}

// ============================= register routes ================================================== //
func (handler *GuestHandler) registerRoutes() {
	routeGroup := handler.Router.Group("/guests")
	routeGroup.POST("", handler.create)
	routeGroup.GET("/:id", handler.find)
	routeGroup.GET("", handler.findAll)
	routeGroup.PUT("/:id", handler.update)
	//routeGroup.DELETE("", handler)
}
