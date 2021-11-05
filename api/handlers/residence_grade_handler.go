package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	middlewares2 "reservation-api/api/middlewares"
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/message_keys"
	"reservation-api/internal/models"
	"reservation-api/internal/services"
	"reservation-api/internal/utils"
	"reservation-api/pkg/applogger"
	"reservation-api/pkg/translator"
)

// ResidenceGradeHandler Province endpoint handler
type ResidenceGradeHandler struct {
	Router     *echo.Group
	Service    *services.ResidenceGradeService
	translator *translator.Translator
	logger     applogger.Logger
}

func (handler *ResidenceGradeHandler) Register(input *dto.HandlerInput, service *services.ResidenceGradeService) {
	handler.Router = input.Router
	handler.Service = service
	handler.translator = input.Translator
	handler.logger = input.Logger

	routeGroup := input.Router.Group("/residence-grade")

	routeGroup.POST("", handler.create)
	routeGroup.PUT("/:id", handler.update)
	routeGroup.GET("/:id", handler.find)
	routeGroup.DELETE("/:id", handler.delete)
	routeGroup.GET("", handler.findAll, middlewares2.PaginationMiddleware)
}

func (handler *ResidenceGradeHandler) create(c echo.Context) error {

	model := &models.ResidenceGrade{}
	lang := c.Request().Header.Get(acceptLanguage)

	if err := c.Bind(&model); err != nil {

		handler.logger.LogError(err.Error())

		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusBadRequest,
			Message:      message_keys.BadRequest,
		})
	}

	result, err := handler.Service.Create(model)

	if err != nil {

		handler.logger.LogError(err.Error())

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

		handler.logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}

	result, err := handler.Service.Find(id)
	if err != nil {

		handler.logger.LogError(err.Error())
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

	tmpModel := models.ResidenceGrade{}

	err = c.Bind(&tmpModel)

	if err != nil {
		handler.logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			ResponseCode: http.StatusBadRequest,
			Message:      handler.translator.Localize(lang, message_keys.BadRequest),
		})
	}

	result.Name = tmpModel.Name

	updatedMode, err := handler.Service.Update(result)

	if err != nil {

		handler.logger.LogError(err.Error())
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

		handler.logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}

	result, err := handler.Service.Find(id)
	if err != nil {

		handler.logger.LogError(err.Error())
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

func (handler *ResidenceGradeHandler) findAll(c echo.Context) error {

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

func (handler *ResidenceGradeHandler) delete(c echo.Context) error {

	id, err := utils.ConvertToUint(c.Param("id"))
	lang := c.Request().Header.Get(acceptLanguage)

	if err != nil {

		handler.logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			ResponseCode: http.StatusBadRequest,
			Message:      handler.translator.Localize(lang, message_keys.BadRequest),
		})
	}

	err = handler.Service.Delete(id)

	if err != nil {

		return c.JSON(http.StatusConflict, commons.ApiResponse{
			ResponseCode: http.StatusConflict,
			Message:      handler.translator.Localize(lang, err.Error()),
		})
	}

	return c.JSON(http.StatusOK, commons.ApiResponse{
		ResponseCode: http.StatusOK,
		Message:      handler.translator.Localize(lang, message_keys.Deleted),
	})
}
