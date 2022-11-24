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

// UserHandler User endpoint handler
type UserHandler struct {
	Service *domain_services.UserService
	Config  *dto.HandlerConfig
}

func (handler *UserHandler) Register(config *dto.HandlerConfig, service *domain_services.UserService) {
	handler.Service = service
	handler.Config = config
	routeGroup := config.Router.Group("/users")
	routeGroup.POST("", handler.create)
	routeGroup.PUT("/:id", handler.update)
	routeGroup.GET("/:id", handler.find)
	routeGroup.GET("/by-username/:username", handler.findByUsername)
	routeGroup.GET("", handler.findAll, middlewares2.PaginationMiddleware)
}

/*====================================================================================*/
func (handler *UserHandler) create(c echo.Context) error {

	model := models.User{}
	lang := c.Request().Header.Get(acceptLanguage)
	user := getCurrentUser(c)

	if err := c.Bind(&model); err != nil {
		handler.Config.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest,
			commons.ApiResponse{
				Data:         nil,
				ResponseCode: http.StatusBadRequest,
				Message:      handler.Config.Translator.Localize(lang, message_keys.BadRequest),
			})
	}

	oldUser, err := handler.Service.FindByUsername(getCurrentTenantContext(c), model.Username)

	if err != nil {

		handler.Config.Logger.LogError(err.Error())

		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			ResponseCode: http.StatusBadRequest,
			Message:      handler.Config.Translator.Localize(lang, err.Error()),
		})
	}

	// prevent to create repeat user.
	if oldUser != nil && oldUser.Id > 0 {

		return c.JSON(http.StatusConflict, commons.ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusConflict,
			Message:      handler.Config.Translator.Localize(lang, message_keys.UsernameDuplicated),
		})
	}

	model.SetAudit(user)
	if result, err := handler.Service.Create(getCurrentTenantContext(c), &model); err == nil {

		return c.JSON(http.StatusBadRequest,
			commons.ApiResponse{
				Data:         result,
				ResponseCode: http.StatusOK,
				Message:      handler.Config.Translator.Localize(lang, message_keys.Created),
			})
	} else {

		handler.Config.Logger.LogError(err.Error())
		return c.JSON(http.StatusInternalServerError,
			commons.ApiResponse{
				Data:         nil,
				ResponseCode: http.StatusInternalServerError,
				Message:      handler.Config.Translator.Localize(lang, message_keys.InternalServerError),
			})
	}

}

/*====================================================================================*/
func (handler *UserHandler) update(c echo.Context) error {

	id, err := utils.ConvertToUint(c.Param("id"))
	user := getCurrentUser(c)

	if err != nil {
		handler.Config.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}

	lang := c.Request().Header.Get(acceptLanguage)
	model, err := handler.Service.Find(getCurrentTenantContext(c), id)

	if err != nil {
		handler.Config.Logger.LogError(err.Error())
		return c.JSON(http.StatusInternalServerError, commons.ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusInternalServerError,
			Message:      handler.Config.Translator.Localize(lang, message_keys.InternalServerError),
		})
	}

	if model == nil {
		return c.JSON(http.StatusNotFound, commons.ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusNotFound,
			Message:      handler.Config.Translator.Localize(lang, message_keys.NotFound),
		})
	}

	if err := c.Bind(&model); err != nil {
		handler.Config.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusBadRequest,
			Message:      handler.Config.Translator.Localize(lang, message_keys.BadRequest),
		})
	}

	model.SetUpdatedBy(user)
	if result, err := handler.Service.Update(getCurrentTenantContext(c), model); err == nil {

		return c.JSON(http.StatusOK, commons.ApiResponse{
			Data:         result,
			ResponseCode: http.StatusOK,
			Message:      handler.Config.Translator.Localize(lang, message_keys.Updated),
		})
	} else {

		handler.Config.Logger.LogError(err.Error())
		return c.JSON(http.StatusInternalServerError, nil)
	}
}

/*====================================================================================*/
func (handler *UserHandler) find(c echo.Context) error {
	id, err := utils.ConvertToUint(c.Param("id"))

	if err != nil {
		handler.Config.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}

	model, err := handler.Service.Find(getCurrentTenantContext(c), id)
	lang := c.Request().Header.Get(acceptLanguage)

	if err != nil {
		handler.Config.Logger.LogError(err.Error())
		return c.JSON(http.StatusInternalServerError, commons.ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusInternalServerError,
			Message:      handler.Config.Translator.Localize(lang, message_keys.InternalServerError),
		})
	}

	if model == nil {
		return c.JSON(http.StatusNotFound, commons.ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusNotFound,
			Message:      handler.Config.Translator.Localize(lang, message_keys.NotFound),
		})
	}

	return c.JSON(http.StatusOK, commons.ApiResponse{
		Data:         model,
		ResponseCode: http.StatusOK,
		Message:      "",
	})
}

/*====================================================================================*/
func (handler *UserHandler) findAll(c echo.Context) error {

	paginationInput := c.Get(paginationInput).(*dto.PaginationFilter)
	list, err := handler.Service.FindAll(getCurrentTenantContext(c), paginationInput)

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
	model, err := handler.Service.FindByUsername(getCurrentTenantContext(c), username)
	lang := c.Request().Header.Get(acceptLanguage)

	if err != nil {
		handler.Config.Logger.LogError(err.Error())
		return c.JSON(http.StatusInternalServerError, commons.ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusInternalServerError,
			Message:      handler.Config.Translator.Localize(lang, message_keys.InternalServerError),
		})
	}

	if model == nil {
		return c.JSON(http.StatusNotFound, commons.ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusNotFound,
			Message:      handler.Config.Translator.Localize(lang, message_keys.NotFound),
		})
	}

	return c.JSON(http.StatusOK, commons.ApiResponse{
		Data:         model,
		ResponseCode: http.StatusOK,
		Message:      "",
	})
}
