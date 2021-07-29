package handlers

import (
	"github.com/labstack/echo/v4"
	. "hotel-reservation/internal/commons"
	"hotel-reservation/internal/dto"
	"hotel-reservation/internal/message_keys"
	"hotel-reservation/internal/middlewares"
	"hotel-reservation/internal/models"
	"hotel-reservation/internal/services"
	"hotel-reservation/internal/utils"
	"hotel-reservation/pkg/translator"
	"net/http"
)

// CountryHandler country endpoint handler
type CountryHandler struct {
	Router     *echo.Group
	Service    *services.CountryService
	translator *translator.Translator
}

func (handler *CountryHandler) Register(router *echo.Group, service *services.CountryService, translator *translator.Translator) {
	handler.Router = router
	handler.Service = service
	handler.translator = translator

	handler.Router.POST("", handler.create)
	handler.Router.PUT("/:id", handler.update)
	handler.Router.GET("/:id", handler.find)
	handler.Router.GET("/:id/provinces", handler.provinces)
	handler.Router.GET("", handler.findAll, middlewares.PaginationMiddleware)
}

func (handler *CountryHandler) create(c echo.Context) error {

	model := &models.Country{}

	lang := c.Request().Header.Get(acceptLanguage)

	if err := c.Bind(&model); err != nil {
		return c.JSON(http.StatusBadRequest,
			ApiResponse{
				Data:         nil,
				ResponseCode: http.StatusBadRequest,
				Message:      handler.translator.Localize(lang, message_keys.BadRequest),
			})
	}

	output, err := handler.Service.Create(model)

	if err != nil {

		return c.JSON(http.StatusBadRequest, ApiResponse{
			ResponseCode: http.StatusBadRequest,
			Message:      err.Error(),
		})
	}

	return c.JSON(http.StatusBadRequest, ApiResponse{
		ResponseCode: http.StatusOK,
		Message:      handler.translator.Localize(lang, message_keys.Created),
		Data:         output,
	})
}

func (handler *CountryHandler) update(c echo.Context) error {

	id, err := utils.ConvertToUint(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	model, err := handler.Service.Find(id)

	lang := c.Request().Header.Get(acceptLanguage)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, ApiResponse{
			ResponseCode: http.StatusInternalServerError,
			Message:      handler.translator.Localize(lang, message_keys.InternalServerError),
		})
	}

	if model == nil {
		return c.JSON(http.StatusNotFound, ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusNotFound,
			Message:      handler.translator.Localize(lang, message_keys.NotFound),
		})
	}

	if err := c.Bind(&model); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	if output, err := handler.Service.Update(model); err == nil {
		return c.JSON(http.StatusOK, ApiResponse{
			Data:         output,
			ResponseCode: http.StatusOK,
			Message:      handler.translator.Localize(lang, message_keys.Updated),
		})
	} else {
		return c.JSON(http.StatusInternalServerError, nil)
	}
}

func (handler *CountryHandler) find(c echo.Context) error {
	id, err := utils.ConvertToUint(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	model, err := handler.Service.Find(id)
	lang := c.Request().Header.Get(acceptLanguage)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, ApiResponse{
			ResponseCode: http.StatusInternalServerError,
			Message:      handler.translator.Localize(lang, message_keys.InternalServerError),
		})
	}

	if model == nil {
		return c.JSON(http.StatusNotFound, ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusNotFound,
			Message:      handler.translator.Localize(lang, message_keys.NotFound),
		})
	}

	return c.JSON(http.StatusOK, ApiResponse{
		Data:         model,
		ResponseCode: http.StatusOK,
		Message:      "",
	})
}

func (handler *CountryHandler) findAll(c echo.Context) error {

	paginationInput := c.Get(paginationInput).(*dto.PaginationInput)

	list, err := handler.Service.FindAll(paginationInput)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, ApiResponse{
		Data:         list,
		ResponseCode: http.StatusOK,
		Message:      "",
	})
}

func (handler *CountryHandler) provinces(c echo.Context) error {

	id, err := utils.ConvertToUint(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	provinces, err := handler.Service.GetProvinces(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, ApiResponse{
			ResponseCode: http.StatusInternalServerError,
			Message:      err.Error(),
		})
	}

	return c.JSON(http.StatusOK, ApiResponse{
		Data:         provinces,
		ResponseCode: http.StatusOK,
		Message:      "",
	})
}
