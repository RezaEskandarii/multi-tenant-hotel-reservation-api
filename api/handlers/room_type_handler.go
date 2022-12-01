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
	"reservation-api/pkg/translator"
)

// RoomTypeTypeHandler Province endpoint handler
type RoomTypeTypeHandler struct {
	handlerBase
	Service *domain_services.RoomTypeService
}

func (handler *RoomTypeTypeHandler) Register(config *dto.HandlerConfig, service *domain_services.RoomTypeService) {

	handler.Service = service
	handler.Router = config.Router
	handler.Logger = config.Logger

	// register RoomType routes
	routeGroup := handler.Router.Group("/room-types")
	routeGroup.POST("", handler.create)
	routeGroup.PUT("/:id", handler.update)
	routeGroup.GET("/:id", handler.find)
	routeGroup.DELETE("/:id", handler.delete)
	routeGroup.GET("", handler.findAll, middlewares2.PaginationMiddleware)

}

// @Summary crete RoomType
// @Tags RoomType
// @Accept json
// @Param X-TenantID header int true "X-TenantID"
// @Produce json
// @Param  RoomType body  models.RoomType true "RoomType"
// @Success 200 {object} models.RoomType
// @Router /room-types [post]
func (handler *RoomTypeTypeHandler) create(c echo.Context) error {

	model := &models.RoomType{}
	user := currentUser(c)

	if err := c.Bind(&model); err != nil {
		handler.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusBadRequest,
			Message:      message_keys.BadRequest,
		})
	}

	model.SetAudit(user)
	result, err := handler.Service.Create(tenantContext(c), model)

	if err != nil {
		handler.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			Data:         nil,
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

// @Summary update RoomType
// @Tags RoomType
// @Accept json
// @Param X-TenantID header int true "X-TenantID"
// @Param Id path int true "Id"
// @Produce json
// @Param  RoomType body  models.RoomType true "RoomType"
// @Success 200 {object} models.RoomType
// @Router /room-types/{id} [put]
func (handler *RoomTypeTypeHandler) update(c echo.Context) error {

	id, err := utils.ConvertToUint(c.Param("id"))
	user := currentUser(c)

	if err != nil {
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

	// prevent to edit other fields by client.
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

// @Summary find RoomType by id
// @Tags RoomType
// @Accept json
// @Param X-TenantID header int true "X-TenantID"
// @Param Id path int true "Id"
// @Produce json
// @Success 200 {object} models.RoomType
// @Router /room-types/{id} [get]
func (handler *RoomTypeTypeHandler) find(c echo.Context) error {

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

// @Summary findAll RoomTypes
// @Tags RoomType
// @Accept json
// @Param X-TenantID header int true "X-TenantID"
// @Produce json
// @Success 200 {array} models.RoomType
// @Router /room-types [get]
func (handler *RoomTypeTypeHandler) findAll(c echo.Context) error {

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

// @Summary Delete RoomType
// @Tags RoomType
// @Accept json
// @Param X-TenantID header int true "X-TenantID"
// @Param Id path int true "Id"
// @Produce json
// @Success 200
// @Router /room-types/{id} [delete]
func (handler *RoomTypeTypeHandler) delete(c echo.Context) error {

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
