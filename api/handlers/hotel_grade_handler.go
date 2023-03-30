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
	"reservation-api/internal/utils"
	"reservation-api/internal_errors/message_keys"
	"reservation-api/pkg/translator"
)

// HotelGradeHandler Province endpoint handler
type HotelGradeHandler struct {
	handlerBase
	Service *domain_services.HotelGradeService
}

// Register HotelGradeHandler
// this method registers all routes,routeGroups and passes HotelGradeHandler's related dependencies
func (handler *HotelGradeHandler) Register(config *dto.HandlerConfig, service *domain_services.HotelGradeService) {
	handler.Service = service
	handler.Router = config.Router
	handler.Logger = config.Logger
	handler.registerRoutes()
}

// @Tags HotelGrade
// @Accept json
// @Produce json
// @Param X-Tenant-ID header int true "X-Tenant-ID"
// @Param  HotelGrade body  models.HotelGrade true "HotelGrade"
// @Success 200 {object} models.HotelGrade
// @Router /hotel-grades [post]
func (handler *HotelGradeHandler) create(c echo.Context) error {

	hotel := &models.HotelGrade{}
	user := currentUser(c)

	if err := c.Bind(&hotel); err != nil {

		handler.Logger.LogError(err.Error())

		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusBadRequest,
			Message:      message_keys.BadRequest,
		})
	}

	hotel.SetAudit(user)
	result, err := handler.Service.Create(tenantContext(c), hotel)

	if err != nil {

		handler.Logger.LogError(err.Error())

		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusBadRequest,
			Message:      "",
		})
	}

	return c.JSON(http.StatusOK, commons.ApiResponse{
		Data:         result,
		ResponseCode: http.StatusOK,
		Message:      translator.Localize(c.Request().Context(), message_keys.Created),
	})
}

// @Tags HotelGrade
// @Accept json
// @Param X-Tenant-ID header int true "X-Tenant-ID"
// @Param Id path int true "Id"
// @Param HotelGrade body models.HotelGrade true "HotelGrade"
// @Produce json
// @Param  HotelGrade body  models.HotelGrade true "HotelGrade"
// @Success 200 {object} models.HotelGrade
// @Router /hotel-grades/{id} [put]
func (handler *HotelGradeHandler) update(c echo.Context) error {

	id, err := utils.ConvertToUint(c.Param("id"))
	user := currentUser(c)

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

	// If not found any item, return 404.
	if result == nil || (result != nil && result.Id == 0) {
		return c.JSON(http.StatusNotFound, commons.ApiResponse{
			ResponseCode: http.StatusNotFound,
			Message:      translator.Localize(c.Request().Context(), message_keys.NotFound),
		})
	}

	// bind request to tmp struct.
	tmpHotel := models.HotelGrade{}
	err = c.Bind(&tmpHotel)

	if err != nil {
		handler.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			ResponseCode: http.StatusBadRequest,
			Message:      translator.Localize(c.Request().Context(), message_keys.BadRequest),
		})
	}

	// prevent to edit other fields.
	result.Name = tmpHotel.Name
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

// @Tags HotelGrade
// @Accept json
// @Param X-Tenant-ID header int true "X-Tenant-ID"
// @Param Id path int true "Id"
// @Produce json
// @Success 200 {object} models.HotelGrade
// @Router /hotel-grades/{id} [get]
func (handler *HotelGradeHandler) find(c echo.Context) error {

	id, err := utils.ConvertToUint(c.Param("id"))

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

// @Tags HotelGrade
// @Accept json
// @Param X-Tenant-ID header int true "X-Tenant-ID"
// @Produce json
// @Success 200 {array} models.HotelGrade
// @Router /hotel-grades [get]
func (handler *HotelGradeHandler) findAll(c echo.Context) error {

	paginationInput := c.Get(paginationInput).(*dto.PaginationFilter)
	list, err := handler.Service.FindAll(tenantContext(c), paginationInput)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, commons.ApiResponse{
		Data:         list,
		ResponseCode: http.StatusOK,
		Message:      "",
	})
}

// @Tags HotelGrade
// @Accept json
// @Param X-Tenant-ID header int true "X-Tenant-ID"
// @Produce json
// @Success 200 {array} models.HotelGrade
// @Router /hotel-grades [delete]
func (handler *HotelGradeHandler) delete(c echo.Context) error {

	id, err := utils.ConvertToUint(c.Param("id"))

	if err != nil {

		handler.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			ResponseCode: http.StatusBadRequest,
			Message:      translator.Localize(c.Request().Context(), message_keys.BadRequest),
		})
	}

	err = handler.Service.Delete(tenantContext(c), id)

	if err != nil {
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
func (handler *HotelGradeHandler) registerRoutes() {
	routeGroup := handler.Router.Group("/hotel-grades")
	routeGroup.POST("", handler.create)
	routeGroup.PUT("/:id", handler.update)
	routeGroup.GET("/:id", handler.find)
	routeGroup.DELETE("/:id", handler.delete)
	routeGroup.GET("", handler.findAll, middlewares2.PaginationMiddleware)
}
