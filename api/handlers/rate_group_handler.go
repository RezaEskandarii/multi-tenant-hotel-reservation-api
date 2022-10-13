package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	middlewares2 "reservation-api/api/middlewares"
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/message_keys"
	"reservation-api/internal/models"
	"reservation-api/internal/services/domain_services"
	"reservation-api/internal/utils"
)

// RateGroupHandler RateGroup endpoint handler
type RateGroupHandler struct {
	Service *domain_services.RateGroupService
	Input   *dto.HandlersShared
}

func (handler *RateGroupHandler) Register(input *dto.HandlersShared, service *domain_services.RateGroupService) {
	handler.Service = service
	handler.Input = input
	routeGroup := handler.Input.Router.Group("/rate-groups")
	routeGroup.POST("", handler.create)
	routeGroup.PUT("/:id", handler.update)
	routeGroup.GET("/:id", handler.find)
	routeGroup.DELETE("/:id", handler.delete)
	routeGroup.GET("", handler.findAll, middlewares2.PaginationMiddleware)
}

// @Summary Create RateGroup
// @Tags RateGroup
// @Accept json
// @Param X-TenantID header int true "X-TenantID"
// @Produce json
// @Param  RateGroup body  models.RateGroup true "RateGroup"
// @Success 200 {object} models.RateGroup
// @Router /rate-groups/{id} [post]
func (handler *RateGroupHandler) create(c echo.Context) error {

	model := &models.RateGroup{}
	lang := c.Request().Header.Get(acceptLanguage)
	user := getCurrentUser(c)

	if err := c.Bind(&model); err != nil {
		handler.Input.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest,
			commons.ApiResponse{
				Data:         nil,
				ResponseCode: http.StatusBadRequest,
				Message:      handler.Input.Translator.Localize(lang, message_keys.BadRequest),
			})
	}

	model.SetUpdatedBy(user)
	result, err := handler.Service.Create(model, getCurrentTenant(c))

	if err != nil {
		handler.Input.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			ResponseCode: http.StatusBadRequest,
			Message:      err.Error(),
		})
	}

	return c.JSON(http.StatusBadRequest, commons.ApiResponse{
		ResponseCode: http.StatusOK,
		Message:      handler.Input.Translator.Localize(lang, message_keys.Created),
		Data:         result,
	})
}

// @Summary update RateGroup
// @Tags RateGroup
// @Accept json
// @Param X-TenantID header int true "X-TenantID"
// @Param Id path int true "Id"
// @Produce json
// @Param  RateGroup body  models.RateGroup true "RateGroup"
// @Success 200 {object} models.RateGroup
// @Router /rate-groups/{id} [put]
func (handler *RateGroupHandler) update(c echo.Context) error {

	id, err := utils.ConvertToUint(c.Param("id"))

	if err != nil {
		handler.Input.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}

	model, err := handler.Service.Find(id, getCurrentTenant(c))
	lang := c.Request().Header.Get(acceptLanguage)

	if err != nil {
		handler.Input.Logger.LogError(err.Error())
		return c.JSON(http.StatusInternalServerError, commons.ApiResponse{
			ResponseCode: http.StatusInternalServerError,
			Message:      handler.Input.Translator.Localize(lang, message_keys.InternalServerError),
		})
	}

	if model == nil {
		return c.JSON(http.StatusNotFound, commons.ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusNotFound,
			Message:      handler.Input.Translator.Localize(lang, message_keys.NotFound),
		})
	}

	if err := c.Bind(&model); err != nil {
		handler.Input.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}

	if result, err := handler.Service.Update(model, getCurrentTenant(c)); err == nil {

		return c.JSON(http.StatusOK, commons.ApiResponse{
			Data:         result,
			ResponseCode: http.StatusOK,
			Message:      handler.Input.Translator.Localize(lang, message_keys.Updated),
		})
	} else {

		handler.Input.Logger.LogError(err.Error())
		return c.JSON(http.StatusInternalServerError, nil)
	}
}

// @Summary find RateGroup
// @Tags RateGroup
// @Accept json
// @Param X-TenantID header int true "X-TenantID"
// @Param Id path int true "Id"
// @Produce json
// @Success 200 {array} models.RateGroup
// @Router /rate-groups/{id} [get]
func (handler *RateGroupHandler) find(c echo.Context) error {
	id, err := utils.ConvertToUint(c.Param("id"))

	if err != nil {
		handler.Input.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}

	model, err := handler.Service.Find(id, getCurrentTenant(c))
	lang := c.Request().Header.Get(acceptLanguage)

	if err != nil {
		handler.Input.Logger.LogError(err.Error())
		return c.JSON(http.StatusInternalServerError, commons.ApiResponse{
			ResponseCode: http.StatusInternalServerError,
			Message:      handler.Input.Translator.Localize(lang, message_keys.InternalServerError),
		})
	}

	if model == nil {
		return c.JSON(http.StatusNotFound, commons.ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusNotFound,
			Message:      handler.Input.Translator.Localize(lang, message_keys.NotFound),
		})
	}

	return c.JSON(http.StatusOK, commons.ApiResponse{
		Data:         model,
		ResponseCode: http.StatusOK,
		Message:      "",
	})
}

// @Summary findAll rate-codes
// @Tags RateGroup
// @Accept json
// @Param X-TenantID header int true "X-TenantID"
// @Produce json
// @Success 200 {array} models.RateGroup
// @Router /rate-groups [get]
func (handler *RateGroupHandler) findAll(c echo.Context) error {

	paginationInput := c.Get(paginationInput).(*dto.PaginationFilter)

	list, err := handler.Service.FindAll(paginationInput)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, commons.ApiResponse{
		Data:         list,
		ResponseCode: http.StatusOK,
		Message:      "",
	})
}

// @Summary Delete RateGroup
// @Tags RateGroup
// @Accept json
// @Param X-TenantID header int true "X-TenantID"
// @Param Id path int true "Id"
// @Produce json
// @Success 200 {array} models.RateGroup
// @Router /rate-groups [delete]
func (handler *RateGroupHandler) delete(c echo.Context) error {

	id, err := utils.ConvertToUint(c.Param("id"))
	lang := c.Request().Header.Get(acceptLanguage)

	if err != nil {

		handler.Input.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			ResponseCode: http.StatusBadRequest,
			Message:      handler.Input.Translator.Localize(lang, message_keys.BadRequest),
		})
	}

	err = handler.Service.Delete(id, getCurrentTenant(c))

	if err != nil {

		handler.Input.Logger.LogError(err.Error())
		return c.JSON(http.StatusConflict, commons.ApiResponse{
			ResponseCode: http.StatusConflict,
			Message:      handler.Input.Translator.Localize(lang, err.Error()),
		})
	}

	return c.JSON(http.StatusOK, commons.ApiResponse{
		ResponseCode: http.StatusOK,
		Message:      handler.Input.Translator.Localize(lang, message_keys.Deleted),
	})
}
