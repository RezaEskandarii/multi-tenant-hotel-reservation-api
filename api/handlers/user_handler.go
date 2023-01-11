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

// UserHandler User endpoint handler
type UserHandler struct {
	handlerBase
	Service *domain_services.UserService
}

// Register UserHandler
// this method registers all routes,routeGroups and passes UserHandler's related dependencies
func (handler *UserHandler) Register(config *dto.HandlerConfig, service *domain_services.UserService) {
	handler.Service = service
	handler.Router = config.Router
	handler.Logger = config.Logger
	handler.registerRoutes()
}

// @Summary crete User
// @Tags User
// @Accept json
// @Param X-Tenant-ID header int true "X-Tenant-ID"
// @Produce json
// @Param  User body  models.User true "User"
// @Success 200 {object} models.User
// @Router /users [post]
func (handler *UserHandler) create(c echo.Context) error {

	userModel := models.User{}
	user := currentUser(c)

	if err := c.Bind(&userModel); err != nil {
		handler.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest,
			commons.ApiResponse{
				Data:         nil,
				ResponseCode: http.StatusBadRequest,
				Message:      translator.Localize(c.Request().Context(), message_keys.BadRequest),
			})
	}

	oldUser, err := handler.Service.FindByUsername(tenantContext(c), userModel.Username)

	if err != nil {

		handler.Logger.LogError(err.Error())

		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			ResponseCode: http.StatusBadRequest,
			Message:      translator.Localize(c.Request().Context(), err.Error()),
		})
	}

	// prevent to create repeat user.
	if oldUser != nil && oldUser.Id > 0 {

		return c.JSON(http.StatusConflict, commons.ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusConflict,
			Message:      translator.Localize(c.Request().Context(), message_keys.UsernameDuplicated),
		})
	}

	userModel.SetAudit(user)
	if result, err := handler.Service.Create(tenantContext(c), &userModel); err == nil {

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
				Data:         nil,
				ResponseCode: http.StatusInternalServerError,
				Message:      translator.Localize(c.Request().Context(), message_keys.InternalServerError),
			})
	}

}

// @Summary update
// @Tags User
// @Accept json
// @Param X-Tenant-ID header int true "X-Tenant-ID"
// @Param Id path int true "Id"
// @Produce json
// @Param  User body  models.User true "User"
// @Success 200 {object} models.User
// @Router /users/{id} [put]
func (handler *UserHandler) update(c echo.Context) error {

	id, err := utils.ConvertToUint(c.Param("id"))
	user := currentUser(c)

	if err != nil {
		handler.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}

	userModel, err := handler.Service.Find(tenantContext(c), id)

	if err != nil {
		handler.Logger.LogError(err.Error())
		return c.JSON(http.StatusInternalServerError, commons.ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusInternalServerError,
			Message:      translator.Localize(c.Request().Context(), message_keys.InternalServerError),
		})
	}

	if userModel == nil {
		return c.JSON(http.StatusNotFound, commons.ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusNotFound,
			Message:      translator.Localize(c.Request().Context(), message_keys.NotFound),
		})
	}

	if err := c.Bind(&userModel); err != nil {
		handler.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusBadRequest,
			Message:      translator.Localize(c.Request().Context(), message_keys.BadRequest),
		})
	}

	userModel.SetUpdatedBy(user)
	if result, err := handler.Service.Update(tenantContext(c), userModel); err == nil {

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

// @Summary findById
// @Tags User
// @Accept json
// @Param X-Tenant-ID header int true "X-Tenant-ID"
// @Param Id path int true "Id"
// @Produce json
// @Success 200 {object} models.User
// @Router /users/{id} [get]
func (handler *UserHandler) find(c echo.Context) error {
	id, err := utils.ConvertToUint(c.Param("id"))

	if err != nil {
		handler.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}

	model, err := handler.Service.Find(tenantContext(c), id)

	if err != nil {
		handler.Logger.LogError(err.Error())
		return c.JSON(http.StatusInternalServerError, commons.ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusInternalServerError,
			Message:      translator.Localize(c.Request().Context(), message_keys.InternalServerError),
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
		Message:      "",
	})
}

// @Summary findAll
// @Tags User
// @Accept json
// @Param X-Tenant-ID header int true "X-Tenant-ID"
// @Produce json
// @Success 200 {array} models.User
// @Router /users [get]
func (handler *UserHandler) findAll(c echo.Context) error {

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

// @Summary findByUsername
// @Tags User
// @Accept json
// @Param X-Tenant-ID header int true "X-Tenant-ID"
// @Param username query string true "username"
// @Produce json
// @Success 200 {object} models.User
// @Router /users [get]
func (handler *UserHandler) findByUsername(c echo.Context) error {

	username := c.Param("username")
	user, err := handler.Service.FindByUsername(tenantContext(c), username)

	if err != nil {
		handler.Logger.LogError(err.Error())
		return c.JSON(http.StatusInternalServerError, commons.ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusInternalServerError,
			Message:      translator.Localize(c.Request().Context(), message_keys.InternalServerError),
		})
	}

	if user == nil {
		return c.JSON(http.StatusNotFound, commons.ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusNotFound,
			Message:      translator.Localize(c.Request().Context(), message_keys.NotFound),
		})
	}

	return c.JSON(http.StatusOK, commons.ApiResponse{
		Data:         user,
		ResponseCode: http.StatusOK,
		Message:      "",
	})
}

// ============================= register routes ================================================== //
func (handler *UserHandler) registerRoutes() {
	routeGroup := handler.Router.Group("/users")
	routeGroup.POST("", handler.create)
	routeGroup.PUT("/:id", handler.update)
	routeGroup.GET("/:id", handler.find)
	routeGroup.GET("/by-username/:username", handler.findByUsername)
	routeGroup.GET("", handler.findAll, middlewares2.PaginationMiddleware)
}
