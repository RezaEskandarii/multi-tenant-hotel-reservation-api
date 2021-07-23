package handlers

import (
	"github.com/labstack/echo/v4"
	"hotel-reservation/internal/commons"
	"hotel-reservation/internal/message_keys"
	"hotel-reservation/internal/middlewares"
	"hotel-reservation/internal/models"
	"hotel-reservation/internal/services"
	"hotel-reservation/pkg/translator"
	"net/http"
)

// ResidenceTypeHandler Province endpoint handler
type ResidenceTypeHandler struct {
	Router     *echo.Group
	Service    services.ResidenceTypeService
	translator *translator.Translator
}

func (handler *ResidenceTypeHandler) Register(router *echo.Group, service services.ResidenceTypeService, translator *translator.Translator) {
	handler.Router = router
	handler.Service = service
	handler.translator = translator

	handler.Router.POST("", handler.create)
	handler.Router.PUT("/:id", handler.update)
	handler.Router.GET("/:id", handler.find)
	handler.Router.GET("/:id/cities", handler.cities)
	handler.Router.GET("", handler.findAll, middlewares.PaginationMiddleware)
}

func (handler *ResidenceTypeHandler) create(c echo.Context) error {

	model := &models.ResidenceType{}
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

func (handler *ResidenceTypeHandler) update(c echo.Context) error {

	panic("not implemented")
}

func (handler *ResidenceTypeHandler) find(c echo.Context) error {

	panic("not implemented")
}

func (handler *ResidenceTypeHandler) cities(c echo.Context) error {

	panic("not implemented")
}

func (handler *ResidenceTypeHandler) findAll(c echo.Context) error {

	panic("not implemented")
}
