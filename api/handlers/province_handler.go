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
	"reservation-api/pkg/validator"
)

// ProvinceHandler Province endpoint handler
type ProvinceHandler struct {
	handlerBase
	Service *domain_services.ProvinceService
}

// Register ProvinceHandler
// this method registers all routes,routeGroups and passes ProvinceHandler's related dependencies
func (handler *ProvinceHandler) Register(config *dto.HandlerConfig, service *domain_services.ProvinceService) {
	handler.Service = service
	handler.Router = config.Router
	handler.Logger = config.Logger
	handler.registerRoutes()
}

// @Tags Province
// @Accept json
// @Param X-Tenant-ID header int true "X-Tenant-ID"
// @Param Id path int true "Id"
// @Param Province body models.Province true "Province"
// @Produce json
// @Param  Province body  models.Province true "Province"
// @Success 200 {object} models.Province
// @Router /provinces/{id} [put]
func (handler *ProvinceHandler) create(c echo.Context) error {

	province := &models.Province{}
	user := currentUser(c)

	if err := c.Bind(&province); err != nil {
		handler.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest,
			commons.ApiResponse{
				ResponseCode: http.StatusBadRequest,
				Message:      translator.Localize(c.Request().Context(), message_keys.BadRequest),
			})
	}

	if err, messages := validator.Validate(province); err != nil {
		return c.JSON(http.StatusBadRequest,
			commons.ApiResponse{
				ResponseCode: http.StatusBadRequest,
				Errors:       messages,
			})
	}

	province.SetAudit(user)
	if result, err := handler.Service.Create(tenantContext(c), province); err == nil {

		return c.JSON(http.StatusBadRequest,
			commons.ApiResponse{
				Data:         result,
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

// @Tags Province
// @Accept json
// @Param X-Tenant-ID header int true "X-Tenant-ID"
// @Param Id path int true "Id"
// @Param Province body models.Province true "Province"
// @Produce json
// @Param  Province body  models.Province true "Province"
// @Success 200 {object} models.Province
// @Router /provinces/{id} [put]
func (handler *ProvinceHandler) update(c echo.Context) error {

	id, err := utils.ConvertToUint(c.Param("id"))

	if err != nil {
		handler.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}

	user := currentUser(c)
	province, err := handler.Service.Find(tenantContext(c), id)

	if err != nil {

		handler.Logger.LogError(err.Error())
		return c.JSON(http.StatusInternalServerError, commons.ApiResponse{
			ResponseCode: http.StatusInternalServerError,
			Message:      translator.Localize(c.Request().Context(), message_keys.InternalServerError),
		})

	}

	if province == nil {
		return c.JSON(http.StatusNotFound, commons.ApiResponse{
			ResponseCode: http.StatusNotFound,
			Message:      translator.Localize(c.Request().Context(), message_keys.NotFound),
		})
	}

	if err := c.Bind(&province); err != nil {
		handler.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}

	if err, messages := validator.Validate(province); err != nil {
		return c.JSON(http.StatusBadRequest,
			commons.ApiResponse{
				ResponseCode: http.StatusBadRequest,
				Errors:       messages,
			})
	}
	province.SetUpdatedBy(user)
	if result, err := handler.Service.Update(tenantContext(c), province); err == nil {

		return c.JSON(http.StatusOK, commons.ApiResponse{
			Data:         result,
			ResponseCode: http.StatusOK,
			Message:      translator.Localize(c.Request().Context(), message_keys.Updated),
		})
	} else {

		handler.Logger.LogError(err.Error())
		return c.JSON(http.StatusInternalServerError, nil)
	}
}

// @Tags Province
// @Accept json
// @Param X-Tenant-ID header int true "X-Tenant-ID"
// @Param Id path int true "Id"
// @Produce json
// @Success 200 {object} models.Province
// @Router /provinces/{id} [get]
func (handler *ProvinceHandler) find(c echo.Context) error {

	id, err := utils.ConvertToUint(c.Param("id"))
	if err != nil {
		handler.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}

	model, err := handler.Service.Find(tenantContext(c), id)

	if err != nil {
		handler.Logger.LogError(err.Error())
		return c.JSON(http.StatusInternalServerError, commons.ApiResponse{
			ResponseCode: http.StatusInternalServerError,
		})
	}

	if model == nil {
		return c.JSON(http.StatusNotFound, commons.ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusNotFound,
			Message:      translator.Localize(c.Request().Context(), message_keys.NotFound),
		})
	}

	return c.JSON(http.StatusOK, commons.ApiResponse{
		Data:         model,
		ResponseCode: http.StatusOK,
	})
}

// @Tags Province
// @Accept json
// @Param X-Tenant-ID header int true "X-Tenant-ID"
// @Produce json
// @Success 200 {array} models.Province
// @Router /provinces [get]
func (handler *ProvinceHandler) findAll(c echo.Context) error {

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

// @Tags Province
// @Accept json
// @Param X-Tenant-ID header int true "X-Tenant-ID"
// @Param Id path int true "Id"
// @Produce json
// @Success 200 {array} models.Province
// @Router /provinces/{id}/cities [get]
func (handler *ProvinceHandler) cities(c echo.Context) error {

	id, err := utils.ConvertToUint(c.Param("id"))
	if err != nil {
		handler.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}

	cities, err := handler.Service.GetCities(tenantContext(c), id)
	if err != nil {

		handler.Logger.LogError(err.Error())
		return c.JSON(http.StatusInternalServerError, commons.NewApiResponse().
			SetResponseCode(http.StatusInternalServerError).SetMessage("bad request"),
		)

	}

	return c.JSON(http.StatusOK, commons.ApiResponse{
		Data:         cities,
		ResponseCode: http.StatusOK,
	})
}

// ============================= register routes ================================================== //
func (handler *ProvinceHandler) registerRoutes() {
	routeGroup := handler.Router.Group("/provinces")
	routeGroup.POST("", handler.create)
	routeGroup.PUT("/:id", handler.update)
	routeGroup.GET("/:id", handler.find)
	routeGroup.GET("/:id/cities", handler.cities)
	routeGroup.GET("", handler.findAll, middlewares2.PaginationMiddleware)
}
