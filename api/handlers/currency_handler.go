package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	middlewares2 "reservation-api/api/middlewares"
	. "reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/message_keys"
	"reservation-api/internal/models"
	"reservation-api/internal/services/domain_services"
	"reservation-api/internal/utils"
	"reservation-api/pkg/translator"
)

// CurrencyHandler Currency endpoint handler
type CurrencyHandler struct {
	handlerBase
	Service *domain_services.CurrencyService
}

func (handler *CurrencyHandler) Register(config *dto.HandlerConfig, service *domain_services.CurrencyService) {
	handler.Service = service
	handler.Router = config.Router
	handler.Logger = config.Logger
	routeGroup := handler.Router.Group("/currencies")
	routeGroup.POST("", handler.create)
	routeGroup.PUT("/:id", handler.update)
	routeGroup.GET("/:id", handler.find)
	routeGroup.GET("", handler.findAll, middlewares2.PaginationMiddleware)
}

// @Summary create new Currency
// @Tags Currency
// @Accept json
// @Produce json
// @Param X-Tenant-ID header int true "X-Tenant-ID"
// @Param  Currency body  models.Currency true "Currency"
// @Success 200 {object} models.Currency
// @Router /currencies [post]
func (handler *CurrencyHandler) create(c echo.Context) error {

	model := &models.Currency{}
	user := currentUser(c)

	if err := c.Bind(&model); err != nil {
		handler.Logger.LogError(err.Error())

		return c.JSON(http.StatusBadRequest,
			ApiResponse{
				ResponseCode: http.StatusInternalServerError,
				Message:      translator.Localize(c.Request().Context(), message_keys.BadRequest),
			})
	}

	model.SetAudit(user)
	if result, err := handler.Service.Create(tenantContext(c), model); err == nil {

		return c.JSON(http.StatusBadRequest,
			ApiResponse{
				Data:         result,
				ResponseCode: http.StatusOK,
				Message:      translator.Localize(c.Request().Context(), message_keys.Created),
			})
	} else {

		handler.Logger.LogError(err.Error())

		return c.JSON(http.StatusInternalServerError,
			ApiResponse{
				Data:         nil,
				ResponseCode: http.StatusInternalServerError,
				Message:      "",
			})

	}
}

// @Summary update Currency
// @Tags Currency
// @Accept json
// @Param X-Tenant-ID header int true "X-Tenant-ID"
// @Param Id path int true "Id"
// @Param Currency body models.Currency true "Currency"
// @Produce json
// @Param  Currency body  models.Currency true "Currency"
// @Success 200 {object} models.Currency
// @Router /currencies/{id} [put]
func (handler *CurrencyHandler) update(c echo.Context) error {

	id, err := utils.ConvertToUint(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	user := currentUser(c)
	model, err := handler.Service.Find(tenantContext(c), id)

	if err != nil {
		handler.Logger.LogError(err.Error())
		return c.JSON(http.StatusInternalServerError, ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusInternalServerError,
			Message:      "",
		})
	}

	if model == nil {
		return c.JSON(http.StatusNotFound, ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusNotFound,
			Message:      translator.Localize(c.Request().Context(), message_keys.NotFound),
		})
	}

	model.SetUpdatedBy(user)
	if err := c.Bind(&model); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	if result, err := handler.Service.Update(tenantContext(c), model); err == nil {

		return c.JSON(http.StatusOK, ApiResponse{
			Data:         result,
			ResponseCode: http.StatusOK,
			Message:      translator.Localize(c.Request().Context(), message_keys.Updated),
		})
	} else {

		handler.Logger.LogError(err.Error())
		return c.JSON(http.StatusInternalServerError, nil)
	}
}

// @Summary find Currency by id
// @Tags Currency
// @Accept json
// @Param X-Tenant-ID header int true "X-Tenant-ID"
// @Param Id path int true "Id"
// @Produce json
// @Success 200 {object} models.Currency
// @Router /currencies/{id} [get]
func (handler *CurrencyHandler) find(c echo.Context) error {
	id, err := utils.ConvertToUint(c.Param("id"))
	if err != nil {

		handler.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}

	model, err := handler.Service.Find(tenantContext(c), id)

	if err != nil {

		handler.Logger.LogError(err.Error())

		return c.JSON(http.StatusInternalServerError, ApiResponse{
			ResponseCode: http.StatusInternalServerError,
			Message:      "",
		})
	}

	if model == nil {
		return c.JSON(http.StatusNotFound, ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusNotFound,
			Message:      translator.Localize(c.Request().Context(), message_keys.NotFound),
		})
	}

	return c.JSON(http.StatusOK, ApiResponse{
		Data:         model,
		ResponseCode: http.StatusOK,
		Message:      "",
	})
}

// @Summary findAll Currencies
// @Tags Currency
// @Accept json
// @Param X-Tenant-ID header int true "X-Tenant-ID"
// @Produce json
// @Success 200 {array} models.Currency
// @Router /currencies [get]
func (handler *CurrencyHandler) findAll(c echo.Context) error {

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
