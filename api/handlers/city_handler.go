// Package handlers
// handles all http requests
///**/
package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	m "reservation-api/api/middlewares"
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
	"reservation-api/internal/services/domain_services"
	"reservation-api/internal_errors/message_keys"
	"reservation-api/pkg/translator"
	"reservation-api/pkg/validator"
	"strconv"
)

// CityHandler City endpoint handler
type CityHandler struct {
	Service *domain_services.CityService
	handlerBase
}

// Register CityHandler
// this method registers all routes,routeGroups and passes CityHandler's related dependencies
func (handler *CityHandler) Register(config *dto.HandlerConfig, service *domain_services.CityService) {
	handler.Service = service
	handler.Router = config.Router
	handler.Logger = config.Logger
	handler.registerRoutes()
}

// @Tags city
// @Accept json
// @Produce json
// @Param X-Tenant-ID header int true "X-Tenant-ID"
// @Param  city body  models.City true "City"
// @Success 200 {object} models.City
// @Router /cities [post]
func (handler *CityHandler) create(c echo.Context) error {
	currentUser := currentUser(c)
	city := models.City{}
	city.SetAudit(currentUser)

	if err := c.Bind(&city); err != nil {
		return c.JSON(http.StatusBadRequest,
			commons.ApiResponse{
				ResponseCode: http.StatusBadRequest,
				Message:      translator.Localize(c.Request().Context(), message_keys.BadRequest),
			})
	}

	if err, messages := validator.Validate(city); err != nil {
		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			Errors:       messages,
			ResponseCode: http.StatusBadRequest,
		})
	}

	if _, err := handler.Service.Create(tenantContext(c), &city); err == nil {
		return c.JSON(http.StatusBadRequest,
			commons.ApiResponse{
				Data:         city,
				ResponseCode: http.StatusOK,
				Message:      translator.Localize(c.Request().Context(), message_keys.Created),
			})
	} else {
		handler.Logger.LogError(err.Error())
		return c.JSON(http.StatusInternalServerError,
			commons.ApiResponse{
				ResponseCode: http.StatusInternalServerError,
				Message:      translator.Localize(c.Request().Context(), message_keys.InternalServerError),
			})
	}

}

// @Tags City
// @Accept json
// @Param X-Tenant-ID header int true "X-Tenant-ID"
// @Param Id path int true "Id"
// @Param City body models.City true "City"
// @Produce json
// @Param  country body  models.City true "City"
// @Success 200 {object} models.City
// @Router /cities/{id} [put]
func (handler *CityHandler) update(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		handler.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}

	currentUser := currentUser(c)
	city, err := handler.Service.Find(tenantContext(c), id)

	if err != nil {
		handler.Logger.LogError(err.Error())
		return c.JSON(http.StatusInternalServerError, commons.ApiResponse{
			ResponseCode: http.StatusInternalServerError,
			Message:      translator.Localize(c.Request().Context(), message_keys.InternalServerError),
		})
	}

	if city == nil {
		return c.JSON(http.StatusNotFound, commons.ApiResponse{
			ResponseCode: http.StatusNotFound,
			Message:      translator.Localize(c.Request().Context(), message_keys.NotFound),
		})
	}

	if err := c.Bind(&city); err != nil {
		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			ResponseCode: http.StatusBadRequest,
			Message:      translator.Localize(c.Request().Context(), message_keys.BadRequest),
		})
	}

	if err, messages := validator.Validate(city); err != nil {
		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			Errors:       messages,
			ResponseCode: http.StatusBadRequest,
		})
	}

	city.SetUpdatedBy(currentUser)
	if output, err := handler.Service.Update(tenantContext(c), city); err == nil {

		return c.JSON(http.StatusOK, commons.ApiResponse{
			Data:         output,
			ResponseCode: http.StatusOK,
			Message:      translator.Localize(c.Request().Context(), message_keys.Updated),
		})
	} else {
		handler.Logger.LogError(err.Error())
		return c.JSON(http.StatusInternalServerError, nil)
	}
}

// @Tags City
// @Accept json
// @Param X-Tenant-ID header int true "X-Tenant-ID"
// @Param Id path int true "Id"
// @Produce json
// @Success 200 {object} models.City
// @Router /cities/{id} [get]
func (handler *CityHandler) find(c echo.Context) error {

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		handler.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}

	city, err := handler.Service.Find(tenantContext(c), id)
	if err != nil {
		handler.Logger.LogError(err.Error())
		return c.JSON(http.StatusInternalServerError, commons.ApiResponse{
			ResponseCode: http.StatusInternalServerError,
			Message:      translator.Localize(c.Request().Context(), message_keys.InternalServerError),
		})
	}

	if city == nil {
		return c.JSON(http.StatusNotFound, commons.ApiResponse{
			ResponseCode: http.StatusNotFound,
			Message:      translator.Localize(c.Request().Context(), message_keys.NotFound),
		})
	}

	return c.JSON(http.StatusOK, commons.ApiResponse{
		Data:         city,
		ResponseCode: http.StatusOK,
	})
}

// @Tags City
// @Accept json
// @Param X-Tenant-ID header int true "X-Tenant-ID"
// @Produce json
// @Success 200 {array} models.City
// @Router /cities [get]
func (handler *CityHandler) findAll(c echo.Context) error {

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

// ============================= register routes ================================================== //
func (handler *CityHandler) registerRoutes() {
	routeGroup := handler.Router.Group("/cities")
	routeGroup.POST("", handler.create)
	routeGroup.PUT("/:id", handler.update)
	routeGroup.GET("/:id", handler.find)
	routeGroup.GET("", handler.findAll, m.PaginationMiddleware)
}
