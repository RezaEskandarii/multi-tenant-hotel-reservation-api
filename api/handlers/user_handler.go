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

// UserHandler User endpoint handler
type UserHandler struct {
	handlerBase
	Service *domain_services.UserService
}

func (handler *UserHandler) Register(config *dto.HandlerConfig, service *domain_services.UserService) {
	handler.Service = service
	handler.Router = config.Router
	handler.Logger = config.Logger

	routeGroup := handler.Router.Group("/users")
	routeGroup.POST("", handler.create)
	routeGroup.PUT("/:id", handler.update)
	routeGroup.GET("/:id", handler.find)
	routeGroup.GET("/by-username/:username", handler.findByUsername)
	routeGroup.GET("", handler.findAll, middlewares2.PaginationMiddleware)
}

// @Summary crete User
// @Tags User
// @Accept json
// @Param X-TenantID header int true "X-TenantID"
// @Produce json
// @Param  User body  models.User true "User"
// @Success 200 {object} models.User
// @Router /users [post]
func (handler *UserHandler) create(c echo.Context) error {

	model := models.User{}
	user := currentUser(c)

	if err := c.Bind(&model); err != nil {
		handler.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest,
			commons.ApiResponse{
				Data:         nil,
				ResponseCode: http.StatusBadRequest,
				Message:      translator.Localize(c.Request().Context(), message_keys.BadRequest),
			})
	}

	err, oldUser := handler.Service.FindByUsername(tenantContext(c), model.Username)

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

	model.SetAudit(user)
	if result, err := handler.Service.Create(tenantContext(c), &model); err == nil {

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

// @Summary update User
// @Tags User
// @Accept json
// @Param X-TenantID header int true "X-TenantID"
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

	if err := c.Bind(&model); err != nil {
		handler.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusBadRequest,
			Message:      translator.Localize(c.Request().Context(), message_keys.BadRequest),
		})
	}

	model.SetUpdatedBy(user)
	if result, err := handler.Service.Update(tenantContext(c), model); err == nil {

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

// @Summary find User by id
// @Tags User
// @Accept json
// @Param X-TenantID header int true "X-TenantID"
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

// @Summary findAll Users
// @Tags User
// @Accept json
// @Param X-TenantID header int true "X-TenantID"
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

/*====================================================================================*/
func (handler *UserHandler) findByUsername(c echo.Context) error {

	username := c.Param("username")
	err, model := handler.Service.FindByUsername(tenantContext(c), username)

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
