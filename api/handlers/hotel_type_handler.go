// Package handlers
// handles all http requests
///**/
package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	middlewares2 "reservation-api/api/middlewares"
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
	"reservation-api/internal/services/domain_services"
	"reservation-api/internal_errors/message_keys"
	"reservation-api/pkg/translator"
	"strconv"
)

// HotelTypeHandler Province endpoint handler
type HotelTypeHandler struct {
	handlerBase
	Service *domain_services.HotelTypeService
}

// Register HotelTypeHandler
// this method registers all routes,routeGroups and passes HotelTypeHandler's related dependencies
func (handler *HotelTypeHandler) Register(config *dto.HandlerConfig, service *domain_services.HotelTypeService) {
	handler.Service = service
	handler.Router = config.Router
	handler.Logger = config.Logger
	handler.registerRoutes()
}

// @Tags HotelType
// @Accept json
// @Produce json
// @Param X-Tenant-ID header int true "X-Tenant-ID"
// @Param  HotelType body  models.HotelType true "HotelType"
// @Success 200 {object} models.HotelType
// @Router /hotel-types [post]
func (handler *HotelTypeHandler) create(c echo.Context) error {

	hotelType := &models.HotelType{}
	user := currentUser(c)

	if err := c.Bind(&hotelType); err != nil {
		handler.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			ResponseCode: http.StatusBadRequest,
			Message:      message_keys.BadRequest,
		})
	}

	hotelType.SetAudit(user)
	result, err := handler.Service.Create(tenantContext(c), hotelType)

	if err != nil {
		handler.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			ResponseCode: http.StatusBadRequest,
		})
	}

	return c.JSON(http.StatusOK, commons.ApiResponse{
		Data:         result,
		ResponseCode: http.StatusOK,
		Message:      translator.Localize(c.Request().Context(), message_keys.Created),
	})
}

// @Tags HotelType
// @Accept json
// @Param X-Tenant-ID header int true "X-Tenant-ID"
// @Param Id path int true "Id"
// @Param HotelType body models.HotelType true "HotelType"
// @Produce json
// @Param  HotelType body  models.HotelType true "HotelType"
// @Success 200 {object} models.HotelType
// @Router /hotel-types/{id} [put]
func (handler *HotelTypeHandler) update(c echo.Context) error {

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	user := currentUser(c)
	result, err := handler.Service.Find(tenantContext(c), id)
	if err != nil {
		handler.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			ResponseCode: http.StatusBadRequest,
			Message:      translator.Localize(c.Request().Context(), message_keys.BadRequest),
		})
	}

	if result == nil || (result != nil && result.Id == 0) {
		return c.JSON(http.StatusNotFound, commons.ApiResponse{
			ResponseCode: http.StatusNotFound,
			Message:      translator.Localize(c.Request().Context(), message_keys.NotFound),
		})
	}

	// prevent to update other fields by client.
	name := c.FormValue("name")
	result.Name = name
	result.SetUpdatedBy(user)

	updatedModel, err := handler.Service.Update(tenantContext(c), result)
	if err != nil {
		handler.Logger.LogError(err.Error())
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, commons.ApiResponse{
		Data:         updatedModel,
		ResponseCode: http.StatusOK,
	})
}

// @Tags HotelType
// @Accept json
// @Param X-Tenant-ID header int true "X-Tenant-ID"
// @Param Id path int true "Id"
// @Produce json
// @Success 200 {object} models.HotelType
// @Router /hotel-types/{id} [get]
func (handler *HotelTypeHandler) find(c echo.Context) error {

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		handler.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}

	result, err := handler.Service.Find(tenantContext(c), id)
	if err != nil {
		handler.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			ResponseCode: http.StatusBadRequest,
			Message:      translator.Localize(c.Request().Context(), message_keys.BadRequest),
		})
	}

	if result == nil || (result != nil && result.Id == 0) {
		return c.JSON(http.StatusNotFound, commons.ApiResponse{
			ResponseCode: http.StatusNotFound,
			Message:      translator.Localize(c.Request().Context(), message_keys.NotFound),
		})
	}

	return c.JSON(http.StatusOK, commons.ApiResponse{
		Data:         result,
		ResponseCode: http.StatusOK,
	})
}

// @Tags HotelType
// @Accept json
// @Param X-Tenant-ID header int true "X-Tenant-ID"
// @Produce json
// @Success 200 {array} models.HotelType
// @Router /hotel-types [get]
func (handler *HotelTypeHandler) findAll(c echo.Context) error {

	paginationInput := c.Get(paginationInput).(*dto.PaginationFilter)
	list, err := handler.Service.FindAll(tenantContext(c), paginationInput)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, commons.ApiResponse{
		Data:         list,
		ResponseCode: http.StatusOK,
	})
}

// @Tags HotelType
// @Accept json
// @Param X-Tenant-ID header int true "X-Tenant-ID"
// @Produce json
// @Success 200 {array} models.HotelType
// @Router /hotel-types [delete]
func (handler *HotelTypeHandler) delete(c echo.Context) error {

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		handler.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			ResponseCode: http.StatusBadRequest,
			Message:      translator.Localize(c.Request().Context(), message_keys.BadRequest),
		})
	}

	err = handler.Service.Delete(tenantContext(c), id)

	if err != nil {
		handler.Logger.LogError(err.Error())
		return c.JSON(http.StatusConflict, commons.ApiResponse{
			ResponseCode: http.StatusConflict,
			Message:      translator.Localize(c.Request().Context(), err.Error()),
		})
	}

	return c.JSON(http.StatusOK, commons.ApiResponse{
		ResponseCode: http.StatusOK,
		Message:      translator.Localize(c.Request().Context(), message_keys.Deleted),
	})
}

// ============================= register routes ================================================== //
func (handler *HotelTypeHandler) registerRoutes() {
	routeGroup := handler.Router.Group("/hotel-types")
	routeGroup.POST("", handler.create)
	routeGroup.PUT("/:id", handler.update)
	routeGroup.GET("/:id", handler.find)
	routeGroup.DELETE("/:id", handler.delete)
	routeGroup.GET("", handler.findAll, middlewares2.PaginationMiddleware)
}
