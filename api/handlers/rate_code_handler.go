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

// RateCodeHandler RateCode endpoint handler
type RateCodeHandler struct {
	handlerBase
	Service               *domain_services.RateCodeService
	RateCodeDetailService *domain_services.RateCodeDetailService
}

// Register RateCodeHandler
// this method registers all routes,routeGroups and passes RateCodeHandler's related dependencies
func (handler *RateCodeHandler) Register(config *dto.HandlerConfig, service *domain_services.RateCodeService, rateCodeDetailService *domain_services.RateCodeDetailService) {
	handler.Service = service
	handler.RateCodeDetailService = rateCodeDetailService
	handler.Router = config.Router
	handler.Logger = config.Logger
	handler.registerRoutes()
}

// @Tags RateCode
// @Accept json
// @Param X-Tenant-ID header int true "X-Tenant-ID"
// @Produce json
// @Param  RateCode body  models.RateCode true "RateCode"
// @Success 200 {object} models.RateCode
// @Router /rate-codes/{id} [post]
func (handler *RateCodeHandler) create(c echo.Context) error {

	rateCode := &models.RateCode{}
	user := currentUser(c)

	if err := c.Bind(&rateCode); err != nil {

		handler.Logger.LogError(err.Error())

		return c.JSON(http.StatusBadRequest,
			commons.ApiResponse{
				ResponseCode: http.StatusBadRequest,
				Message:      translator.Localize(c.Request().Context(), message_keys.BadRequest),
			})
	}

	rateCode.SetAudit(user)
	result, err := handler.Service.Create(tenantContext(c), rateCode)

	if err != nil {
		handler.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			ResponseCode: http.StatusBadRequest,
			Message:      err.Error(),
		})
	}

	return c.JSON(http.StatusBadRequest, commons.ApiResponse{
		ResponseCode: http.StatusOK,
		Message:      translator.Localize(c.Request().Context(), message_keys.Created),
		Data:         result,
	})
}

// @Tags RateCode
// @Accept json
// @Param X-Tenant-ID header int true "X-Tenant-ID"
// @Param Id path int true "Id"
// @Produce json
// @Param  RateCode body  models.RateCode true "RateCode"
// @Success 200 {object} models.RateCode
// @Router /rate-codes/{id} [put]
func (handler *RateCodeHandler) update(c echo.Context) error {

	id, err := utils.ConvertToUint(c.Param("id"))
	user := currentUser(c)

	if err != nil {
		handler.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}

	rateCode, err := handler.Service.Find(tenantContext(c), id)

	if err != nil {
		handler.Logger.LogError(err.Error())
		return c.JSON(http.StatusInternalServerError, commons.ApiResponse{
			ResponseCode: http.StatusInternalServerError,
			Message:      translator.Localize(c.Request().Context(), message_keys.InternalServerError),
		})
	}

	if rateCode == nil {
		return c.JSON(http.StatusNotFound, commons.ApiResponse{
			ResponseCode: http.StatusNotFound,
			Message:      translator.Localize(c.Request().Context(), message_keys.NotFound),
		})
	}
	rateCode.SetUpdatedBy(user)
	if err := c.Bind(&rateCode); err != nil {

		handler.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)

	}

	if result, err := handler.Service.Update(tenantContext(c), rateCode); err == nil {

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

// @Tags RateCode
// @Accept json
// @Param X-Tenant-ID header int true "X-Tenant-ID"
// @Param Id path int true "Id"
// @Produce json
// @Success 200 {array} models.RateCode
// @Router /rate-codes/{id} [get]
func (handler *RateCodeHandler) find(c echo.Context) error {
	id, err := utils.ConvertToUint(c.Param("id"))
	if err != nil {
		handler.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}
	rateCode, err := handler.Service.Find(tenantContext(c), id)

	if err != nil {

		handler.Logger.LogError(err.Error())

		return c.JSON(http.StatusInternalServerError, commons.ApiResponse{
			ResponseCode: http.StatusInternalServerError,
			Message:      translator.Localize(c.Request().Context(), message_keys.InternalServerError),
		})
	}

	if rateCode == nil {
		return c.JSON(http.StatusNotFound, commons.ApiResponse{
			ResponseCode: http.StatusNotFound,
			Message:      translator.Localize(c.Request().Context(), message_keys.NotFound),
		})
	}

	return c.JSON(http.StatusOK, commons.ApiResponse{
		Data:         rateCode,
		ResponseCode: http.StatusOK,
	})
}

// @Tags RateCode
// @Accept json
// @Param X-Tenant-ID header int true "X-Tenant-ID"
// @Produce json
// @Success 200 {array} models.RateCode
// @Router /rate-codes [get]
func (handler *RateCodeHandler) findAll(c echo.Context) error {

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

// @Tags RateCode
// @Accept json
// @Param X-Tenant-ID header int true "X-Tenant-ID"
// @Param Id path int true "Id"
// @Produce json
// @Success 200 {array} models.RateCode
// @Router /rate-codes [delete]
func (handler *RateCodeHandler) delete(c echo.Context) error {

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

// @Tags RateCode
// @Accept json
// @Param X-Tenant-ID header int true "X-Tenant-ID"
// @Param Id path int true "Id"
// @Param RateCode body models.RateCodeDetail true "RateCode"
// @Produce json
// @Param  RateCode body  models.RateCodeDetail true "RateCode"
// @Success 200 {object} models.RateCodeDetail
// @Router /rate-codes/add-details/{id} [post]
func (handler *RateCodeHandler) addDetails(c echo.Context) error {

	id, err := utils.ConvertToUint(c.Get("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			ResponseCode: http.StatusBadRequest,
		})
	}
	requestBody := models.RateCodeDetail{RateCodeId: id}
	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			ResponseCode: http.StatusBadRequest,
		})
	}
	// create new rateCodeDetail.
	result, err := handler.RateCodeDetailService.Create(tenantContext(c), &requestBody)
	if err != nil {
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
func (handler *RateCodeHandler) registerRoutes() {
	routeGroup := handler.Router.Group("/rate-codes")
	routeGroup.POST("", handler.create)
	routeGroup.PUT("/:id", handler.update)
	routeGroup.GET("/:id", handler.find)
	routeGroup.DELETE("/:id", handler.delete)
	routeGroup.GET("", handler.findAll, middlewares2.PaginationMiddleware)
	routeGroup.POST("/add-details/:id", handler.addDetails)
}
