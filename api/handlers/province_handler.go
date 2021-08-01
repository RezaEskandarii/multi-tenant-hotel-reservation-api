package handlers

import (
	"github.com/labstack/echo/v4"
	middlewares2 "hotel-reservation/api/middlewares"
	"hotel-reservation/internal/commons"
	"hotel-reservation/internal/dto"
	"hotel-reservation/internal/message_keys"
	"hotel-reservation/internal/models"
	"hotel-reservation/internal/services"
	"hotel-reservation/internal/utils"
	"hotel-reservation/pkg/applogger"
	"hotel-reservation/pkg/translator"
	"net/http"
)

// ProvinceHandler Province endpoint handler
type ProvinceHandler struct {
	Router     *echo.Group
	Service    *services.ProvinceService
	translator *translator.Translator
}

func (handler *ProvinceHandler) Register(router *echo.Group, service *services.ProvinceService, translator *translator.Translator) {
	handler.Router = router
	handler.Service = service
	handler.translator = translator

	handler.Router.POST("", handler.create)
	handler.Router.PUT("/:id", handler.update)
	handler.Router.GET("/:id", handler.find)
	handler.Router.GET("/:id/cities", handler.cities)
	handler.Router.GET("", handler.findAll, middlewares2.PaginationMiddleware)
}

func (handler *ProvinceHandler) create(c echo.Context) error {

	model := &models.Province{}
	lang := c.Request().Header.Get(acceptLanguage)

	if err := c.Bind(&model); err != nil {

		applogger.LogError(err.Error())

		return c.JSON(http.StatusBadRequest,
			commons.ApiResponse{
				Data:         nil,
				ResponseCode: http.StatusBadRequest,
				Message:      handler.translator.Localize(lang, message_keys.BadRequest),
			})
	}

	if _, err := handler.Service.Create(model); err == nil {
		return c.JSON(http.StatusBadRequest,
			commons.ApiResponse{
				Data:         model,
				ResponseCode: http.StatusOK,
				Message:      handler.translator.Localize(lang, message_keys.Created),
			})
	} else {

		applogger.LogError(err.Error())
		return c.JSON(http.StatusInternalServerError,
			commons.ApiResponse{
				Data:         nil,
				ResponseCode: http.StatusInternalServerError,
				Message:      handler.translator.Localize(lang, message_keys.InternalServerError),
			})
	}

}

func (handler *ProvinceHandler) update(c echo.Context) error {

	id, err := utils.ConvertToUint(c.Param("id"))
	if err != nil {

		applogger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}
	model, err := handler.Service.Find(id)
	lang := c.Request().Header.Get(acceptLanguage)
	if err != nil {

		applogger.LogError(err.Error())

		return c.JSON(http.StatusInternalServerError, commons.ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusInternalServerError,
			Message:      handler.translator.Localize(lang, message_keys.InternalServerError),
		})

	}

	if model == nil {
		return c.JSON(http.StatusNotFound, commons.ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusNotFound,
			Message:      handler.translator.Localize(lang, message_keys.NotFound),
		})
	}

	if err := c.Bind(&model); err != nil {

		applogger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}

	if output, err := handler.Service.Update(model); err == nil {
		return c.JSON(http.StatusOK, commons.ApiResponse{
			Data:         output,
			ResponseCode: http.StatusOK,
			Message:      handler.translator.Localize(lang, message_keys.Updated),
		})
	} else {

		applogger.LogError(err.Error())
		return c.JSON(http.StatusInternalServerError, nil)
	}
}

func (handler *ProvinceHandler) find(c echo.Context) error {
	id, err := utils.ConvertToUint(c.Param("id"))
	if err != nil {

		applogger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}
	model, err := handler.Service.Find(id)
	lang := c.Request().Header.Get(acceptLanguage)

	if err != nil {

		applogger.LogError(err.Error())
		return c.JSON(http.StatusInternalServerError, commons.ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusInternalServerError,
			Message:      "",
		})
	}

	if model == nil {
		return c.JSON(http.StatusNotFound, commons.ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusNotFound,
			Message:      handler.translator.Localize(lang, message_keys.NotFound),
		})
	}

	return c.JSON(http.StatusOK, commons.ApiResponse{
		Data:         model,
		ResponseCode: http.StatusOK,
		Message:      "",
	})
}

func (handler *ProvinceHandler) findAll(c echo.Context) error {

	paginationInput := c.Get(paginationInput).(*dto.PaginationInput)

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

func (handler *ProvinceHandler) cities(c echo.Context) error {

	id, err := utils.ConvertToUint(c.Param("id"))
	if err != nil {

		applogger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}
	cities, err := handler.Service.GetCities(id)

	if err != nil {

		applogger.LogError(err.Error())
		return c.JSON(http.StatusInternalServerError, commons.ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusInternalServerError,
			Message:      "server error",
		})

	}

	return c.JSON(http.StatusOK, commons.ApiResponse{
		Data:         cities,
		ResponseCode: http.StatusOK,
		Message:      "",
	})
}
