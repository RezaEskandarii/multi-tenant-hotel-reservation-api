// Package handlers
// handles all http requests
///**/
package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	middlewares2 "reservation-api/api/middlewares"
	. "reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
	"reservation-api/internal_errors/message_keys"
	"reservation-api/pkg/translator"

	_ "reservation-api/internal/models"
	"reservation-api/internal/services/domain_services"
	"reservation-api/internal/utils"
)

// CountryHandler country endpoint handler
type CountryHandler struct {
	handlerBase
	Service *domain_services.CountryService
}

// Register CountryHandler
// this method registers all routes,routeGroups and passes CountryHandler's related dependencies
func (handler *CountryHandler) Register(config *dto.HandlerConfig, service *domain_services.CountryService) {
	handler.Service = service
	handler.Router = config.Router
	handler.Logger = config.Logger
	handler.registerRoutes()
}

// @Tags Country
// @Accept json
// @Produce json
// @Param X-Tenant-ID header int true "X-Tenant-ID"
// @Param  country body  models.Country true "Country"
// @Success 200 {object} models.Country
// @Router /countries [post]
func (handler *CountryHandler) create(c echo.Context) error {

	country := &models.CountryCreateUpdate{}

	if err := c.Bind(&country); err != nil {
		handler.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest,
			ApiResponse{
				Data:         nil,
				ResponseCode: http.StatusBadRequest,
				Message:      translator.Localize(c.Request().Context(), message_keys.BadRequest),
			})
	}

	if ok, err := country.Validate(); err != nil && ok == false {
		return c.JSON(http.StatusBadRequest, ApiResponse{Message: err.Error()})
	}

	setCreatedByUpdatedBy(country.BaseModel, currentUser(c))
	output, err := handler.Service.Create(tenantContext(c), country)

	if err != nil {
		handler.Logger.LogError(err.Error())

		return c.JSON(http.StatusBadRequest, ApiResponse{
			ResponseCode: http.StatusBadRequest,
			Message:      err.Error(),
		})
	}

	return c.JSON(http.StatusOK, ApiResponse{
		ResponseCode: http.StatusOK,
		Message:      translator.Localize(c.Request().Context(), message_keys.Created),
		Data:         output,
	})
}

// @Tags Country
// @Accept json
// @Param X-Tenant-ID header int true "X-Tenant-ID"
// @Param Id path int true "Id"
// @Param Country body models.Country true "Country"
// @Produce json
// @Param  country body  models.Country true "Country"
// @Success 200 {object} models.Country
// @Router /countries/{id} [put]
func (handler *CountryHandler) update(c echo.Context) error {

	id, err := utils.ConvertToUint(c.Param("id"))
	if err != nil {

		handler.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}

	user := currentUser(c)
	country, err := handler.Service.Find(tenantContext(c), id)

	if err != nil {
		handler.Logger.LogError(err.Error())
		return c.JSON(http.StatusInternalServerError, ApiResponse{
			ResponseCode: http.StatusInternalServerError,
			Message:      translator.Localize(c.Request().Context(), message_keys.InternalServerError),
		})
	}

	if country == nil {
		return c.JSON(http.StatusNotFound, ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusNotFound,
			Message:      translator.Localize(c.Request().Context(), message_keys.NotFound),
		})
	}

	if err := c.Bind(&country); err != nil {
		handler.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}

	country.SetUpdatedBy(user)
	if output, err := handler.Service.Update(tenantContext(c), country); err == nil {

		return c.JSON(http.StatusOK, ApiResponse{
			Data:         output,
			ResponseCode: http.StatusOK,
			Message:      translator.Localize(c.Request().Context(), message_keys.Updated),
		})
	} else {
		handler.Logger.LogError(err.Error())
		return c.JSON(http.StatusInternalServerError, nil)
	}
}

// @Tags Country
// @Accept json
// @Param X-Tenant-ID header int true "X-Tenant-ID"
// @Param Id path int true "Id"
// @Produce json
// @Success 200 {object} models.Country
// @Router /countries/{id} [get]
func (handler *CountryHandler) find(c echo.Context) error {

	id, err := utils.ConvertToUint(c.Param("id"))
	if err != nil {
		handler.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}

	country, err := handler.Service.Find(tenantContext(c), id)

	if err != nil {
		handler.Logger.LogError(err.Error())
		return c.JSON(http.StatusInternalServerError, ApiResponse{
			ResponseCode: http.StatusInternalServerError,
			Message:      translator.Localize(c.Request().Context(), message_keys.InternalServerError),
		})
	}

	if country == nil {
		return c.JSON(http.StatusNotFound, ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusNotFound,
			Message:      translator.Localize(c.Request().Context(), message_keys.NotFound),
		})
	}

	return c.JSON(http.StatusOK, ApiResponse{
		Data:         country,
		ResponseCode: http.StatusOK,
		Message:      "",
	})
}

// @Tags Country
// @Accept json
// @Param X-Tenant-ID header int true "X-Tenant-ID"
// @Produce json
// @Success 200 {array} models.Country
// @Router /countries [get]
func (handler *CountryHandler) findAll(c echo.Context) error {

	paginationInput := c.Get(paginationInput).(*dto.PaginationFilter)

	list, err := handler.Service.FindAll(tenantContext(c), paginationInput)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, ApiResponse{
		Data:         list,
		ResponseCode: http.StatusOK,
		Message:      "",
	})
}

// @Summary find Provinces by country ID
// @Tags Country
// @Accept json
// @Param X-Tenant-ID header int true "X-Tenant-ID"
// @Param Id path int true "Id"
// @Produce json
// @Success 200 {array} models.Province
// @Router /countries/{id}/provinces [get]
func (handler *CountryHandler) provinces(c echo.Context) error {

	id, err := utils.ConvertToUint(c.Param("id"))

	if err != nil {
		handler.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}

	provinces, err := handler.Service.GetProvinces(tenantContext(c), id)

	if err != nil {

		handler.Logger.LogError(err.Error())

		return c.JSON(http.StatusInternalServerError, ApiResponse{
			ResponseCode: http.StatusInternalServerError,
		})
	}

	return c.JSON(http.StatusOK, ApiResponse{
		Data:         provinces,
		ResponseCode: http.StatusOK,
	})
}

// ============================= register routes ================================================== //
func (handler *CountryHandler) registerRoutes() {
	routeGroup := handler.Router.Group("/countries")
	routeGroup.POST("", handler.create)
	routeGroup.PUT("/:id", handler.update)
	routeGroup.GET("/:id", handler.find)
	routeGroup.GET("/:id/provinces", handler.provinces)
	routeGroup.GET("", handler.findAll, middlewares2.PaginationMiddleware)
}
