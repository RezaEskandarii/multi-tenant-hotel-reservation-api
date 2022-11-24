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

// RoomHandler Province endpoint handler
type RoomHandler struct {
	Service *domain_services.RoomService
	Config  *dto.HandlerConfig
}

func (handler *RoomHandler) Register(config *dto.HandlerConfig, service *domain_services.RoomService) {
	handler.Service = service
	handler.Config = config
	routeGroup := config.Router.Group("/room")
	routeGroup.POST("", handler.create)
	routeGroup.PUT("/:id", handler.update)
	routeGroup.GET("/:id", handler.find)
	routeGroup.DELETE("/:id", handler.delete)
	routeGroup.GET("", handler.findAll, middlewares2.PaginationMiddleware)
}

/*====================================================================================*/
func (handler *RoomHandler) create(c echo.Context) error {

	model := &models.Room{}
	lang := getAcceptLanguage(c)
	user := getCurrentUser(c)

	if err := c.Bind(&model); err != nil {
		handler.Config.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusBadRequest,
			Message:      message_keys.BadRequest,
		})
	}

	model.SetAudit(user)
	result, err := handler.Service.Create(getCurrentTenantContext(c), model)

	if err != nil {
		handler.Config.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusBadRequest,
			Message:      err.Error(),
		})
	}

	return c.JSON(http.StatusOK, commons.ApiResponse{
		Data:         result,
		ResponseCode: http.StatusOK,
		Message:      handler.Config.Translator.Localize(lang, message_keys.Created),
	})
}

/*====================================================================================*/
func (handler *RoomHandler) update(c echo.Context) error {

	lang := c.Request().Header.Get(acceptLanguage)
	user := getCurrentUser(c)
	id, err := utils.ConvertToUint(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	result, err := handler.Service.Find(getCurrentTenantContext(c), id)
	if err != nil {
		handler.Config.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			ResponseCode: http.StatusBadRequest,
			Message:      handler.Config.Translator.Localize(lang, message_keys.BadRequest),
		})
	}

	if result == nil || (result != nil && result.Id == 0) {
		return c.JSON(http.StatusNotFound, commons.ApiResponse{
			ResponseCode: http.StatusNotFound,
			Message:      handler.Config.Translator.Localize(lang, message_keys.NotFound),
		})
	}

	name := c.FormValue("name")
	result.Name = name
	result.SetUpdatedBy(user)
	updatedModel, err := handler.Service.Update(getCurrentTenantContext(c), result)

	if err != nil {
		handler.Config.Logger.LogError(err.Error())
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, commons.ApiResponse{
		Data:         updatedModel,
		ResponseCode: http.StatusOK,
	})
}

/*====================================================================================*/
func (handler *RoomHandler) find(c echo.Context) error {

	lang := c.Request().Header.Get(acceptLanguage)
	id, err := utils.ConvertToUint(c.Param("id"))

	if err != nil {
		handler.Config.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}

	result, err := handler.Service.Find(getCurrentTenantContext(c), id)

	if err != nil {

		handler.Config.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			ResponseCode: http.StatusBadRequest,
			Message:      handler.Config.Translator.Localize(lang, message_keys.BadRequest),
		})
	}

	if result == nil || (result != nil && result.Id == 0) {
		return c.JSON(http.StatusNotFound, commons.ApiResponse{
			ResponseCode: http.StatusNotFound,
			Message:      handler.Config.Translator.Localize(lang, message_keys.NotFound),
		})
	}

	return c.JSON(http.StatusOK, commons.ApiResponse{
		Data:         result,
		ResponseCode: http.StatusOK,
	})
}

/*====================================================================================*/
func (handler *RoomHandler) findAll(c echo.Context) error {

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
func (handler *RoomHandler) delete(c echo.Context) error {

	id, err := utils.ConvertToUint(c.Param("id"))
	lang := c.Request().Header.Get(acceptLanguage)

	if err != nil {

		handler.Config.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			ResponseCode: http.StatusBadRequest,
			Message:      handler.Config.Translator.Localize(lang, message_keys.BadRequest),
		})
	}

	err = handler.Service.Delete(getCurrentTenantContext(c), id)

	if err != nil {

		handler.Config.Logger.LogError(err.Error())
		return c.JSON(http.StatusConflict, commons.ApiResponse{
			ResponseCode: http.StatusConflict,
			Message:      handler.Config.Translator.Localize(lang, err.Error()),
		})
	}

	return c.JSON(http.StatusOK, commons.ApiResponse{
		ResponseCode: http.StatusOK,
		Message:      handler.Config.Translator.Localize(lang, message_keys.Deleted),
	})
}
