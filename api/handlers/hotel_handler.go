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

// HotelHandler Province endpoint handler
type HotelHandler struct {
	Service *domain_services.HotelService
	Config  *dto.HandlerConfig
}

func (handler *HotelHandler) Register(config *dto.HandlerConfig, service *domain_services.HotelService) {
	handler.Service = service
	handler.Config = config
	routeGroup := config.Router.Group("/hotels")
	routeGroup.POST("", handler.create)
	routeGroup.PUT("/:id", handler.update)
	routeGroup.GET("/:id", handler.find)
	routeGroup.DELETE("/:id", handler.delete)
	routeGroup.GET("", handler.findAll, middlewares2.PaginationMiddleware)
}

// @Summary create new Hotel
// @Tags Hotel
// @Accept json
// @Produce json
// @Param X-TenantID header int true "X-TenantID"
// @Param  Hotel body  models.Hotel true "Hotel"
// @Success 200 {object} models.Hotel
// @Router /hotels [post]
func (handler *HotelHandler) create(c echo.Context) error {

	createDto := dto.HotelCreateDto{}
	lang := c.Request().Header.Get(acceptLanguage)
	user := currentUser(c)

	if err := c.Bind(&createDto); err != nil {
		handler.Config.Logger.LogError(err.Error())

		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusBadRequest,
			Message:      message_keys.BadRequest,
		})
	}
	model := createDto.Data
	///	model.Thumbnails = createDto.Thumbnails
	model.SetAudit(user)
	result, err := handler.Service.Create(tenantContext(c), &model)

	if err != nil {
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

// @Summary update Hotel
// @Tags Hotel
// @Accept json
// @Param X-TenantID header int true "X-TenantID"
// @Param Id path int true "Id"
// @Param Hotel body models.Hotel true "Hotel"
// @Produce json
// @Param  Hotel body  models.Hotel true "Hotel"
// @Success 200 {object} models.Hotel
// @Router /hotels/{id} [put]
func (handler *HotelHandler) update(c echo.Context) error {

	lang := c.Request().Header.Get(acceptLanguage)
	id, err := utils.ConvertToUint(c.Param("id"))
	user := currentUser(c)

	if err != nil {

		handler.Config.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}

	mainModel, err := handler.Service.Find(tenantContext(c), id)

	if err != nil {
		handler.Config.Logger.LogError(err.Error())

		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			ResponseCode: http.StatusBadRequest,
			Message:      handler.Config.Translator.Localize(lang, message_keys.BadRequest),
		})
	}

	if mainModel == nil || (mainModel != nil && mainModel.Id == 0) {
		return c.JSON(http.StatusNotFound, commons.ApiResponse{
			ResponseCode: http.StatusNotFound,
			Message:      handler.Config.Translator.Localize(lang, message_keys.NotFound),
		})
	}

	clientModel := models.Hotel{}
	err = c.Bind(&clientModel)

	if err != nil {
		handler.Config.Logger.LogError(err.Error())

		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			ResponseCode: http.StatusBadRequest,
			Message:      handler.Config.Translator.Localize(lang, message_keys.BadRequest),
		})
	}

	// map client request fields.
	modelToUpdate := handler.Service.Map(&clientModel, mainModel)
	modelToUpdate.SetUpdatedBy(user)
	updatedModel, err := handler.Service.Update(tenantContext(c), modelToUpdate)

	if err != nil {
		handler.Config.Logger.LogError(err.Error())
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, commons.ApiResponse{
		Data:         updatedModel,
		ResponseCode: http.StatusOK,
	})
}

// @Summary find Hotel by id
// @Tags Hotel
// @Accept json
// @Param X-TenantID header int true "X-TenantID"
// @Param Id path int true "Id"
// @Produce json
// @Success 200 {object} models.Hotel
// @Router /hotels/{id} [get]
func (handler *HotelHandler) find(c echo.Context) error {

	lang := c.Request().Header.Get(acceptLanguage)
	id, err := utils.ConvertToUint(c.Param("id"))

	if err != nil {

		handler.Config.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}

	result, err := handler.Service.Find(tenantContext(c), id)

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

// @Summary findAll Hotels
// @Tags Hotel
// @Accept json
// @Param X-TenantID header int true "X-TenantID"
// @Produce json
// @Success 200 {array} models.Hotel
// @Router /hotels [get]
func (handler *HotelHandler) findAll(c echo.Context) error {

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

// @Summary Delete Hotel
// @Tags Hotel
// @Accept json
// @Param X-TenantID header int true "X-TenantID"
// @Produce json
// @Success 200 {array} models.Hotel
// @Router /hotels [delete]
func (handler *HotelHandler) delete(c echo.Context) error {

	id, err := utils.ConvertToUint(c.Param("id"))
	lang := c.Request().Header.Get(acceptLanguage)

	if err != nil {
		handler.Config.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			ResponseCode: http.StatusBadRequest,
			Message:      handler.Config.Translator.Localize(lang, message_keys.BadRequest),
		})
	}

	err = handler.Service.Delete(tenantContext(c), id)

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
