package handlers

import (
	"github.com/labstack/echo/v4"
	"hotel-reservation/internal/commons"
	"hotel-reservation/internal/message_keys"
	"hotel-reservation/internal/middlewares"
	"hotel-reservation/internal/models"
	"hotel-reservation/internal/services"
	"hotel-reservation/internal/utils"
	"hotel-reservation/pkg/translator"
	"net/http"
)

// ResidenceGradeHandler Province endpoint handler
type ResidenceGradeHandler struct {
	Router     *echo.Group
	Service    services.ResidenceGradeService
	translator *translator.Translator
}

func (handler *ResidenceGradeHandler) Register(router *echo.Group, service services.ResidenceGradeService, translator *translator.Translator) {
	handler.Router = router
	handler.Service = service
	handler.translator = translator

	handler.Router.POST("", handler.create)
	handler.Router.PUT("/:id", handler.update)
	handler.Router.GET("/:id", handler.find)
	handler.Router.GET("/:id/cities", handler.cities)
	handler.Router.GET("", handler.findAll, middlewares.PaginationMiddleware)
}

func (handler *ResidenceGradeHandler) create(c echo.Context) error {

	model := &models.ResidenceGrade{}
	lang := c.Request().Header.Get(acceptLanguage)

	if err := c.Bind(&model); err != nil {
		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusBadRequest,
			Message:      message_keys.BadRequest,
		})
	}

	result, err := handler.Service.Create(model)

	if err != nil {
		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusBadRequest,
			Message:      "",
		})
	}

	return c.JSON(http.StatusOK, commons.ApiResponse{
		Data:         result,
		ResponseCode: http.StatusOK,
		Message:      handler.translator.Localize(lang, message_keys.Created),
	})
}

func (handler *ResidenceGradeHandler) update(c echo.Context) error {

	lang := c.Request().Header.Get(acceptLanguage)
	id, err := utils.ConvertToUint(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	result, err := handler.Service.Find(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			ResponseCode: http.StatusBadRequest,
			Message:      handler.translator.Localize(lang, message_keys.BadRequest),
		})
	}

	if result == nil || (result != nil && result.Id == 0) {
		return c.JSON(http.StatusNotFound, commons.ApiResponse{
			ResponseCode: http.StatusNotFound,
			Message:      handler.translator.Localize(lang, message_keys.NotFound),
		})
	}

	name := c.FormValue("name")
	result.Name = name

	updatedMode, err := handler.Service.Update(result)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, commons.ApiResponse{
		Data:         updatedMode,
		ResponseCode: http.StatusOK,
	})
}

func (handler *ResidenceGradeHandler) find(c echo.Context) error {

	lang := c.Request().Header.Get(acceptLanguage)
	id, err := utils.ConvertToUint(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	result, err := handler.Service.Find(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			ResponseCode: http.StatusBadRequest,
			Message:      handler.translator.Localize(lang, message_keys.BadRequest),
		})
	}

	if result == nil || (result != nil && result.Id == 0) {
		return c.JSON(http.StatusNotFound, commons.ApiResponse{
			ResponseCode: http.StatusNotFound,
			Message:      handler.translator.Localize(lang, message_keys.NotFound),
		})
	}

	return c.JSON(http.StatusOK, commons.ApiResponse{
		Data:         result,
		ResponseCode: http.StatusOK,
	})
}

func (handler *ResidenceGradeHandler) cities(c echo.Context) error {

	panic("not implemented")
}

func (handler *ResidenceGradeHandler) findAll(c echo.Context) error {

	panic("not implemented")
}
