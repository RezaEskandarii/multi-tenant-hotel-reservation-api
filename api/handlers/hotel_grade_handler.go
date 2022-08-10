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

// HotelGradeHandler Province endpoint handler
type HotelGradeHandler struct {
	Service *domain_services.HotelGradeService
	Input   *dto.HandlersShared
}

func (handler *HotelGradeHandler) Register(input *dto.HandlersShared, service *domain_services.HotelGradeService) {
	handler.Service = service
	handler.Input = input
	routeGroup := input.Router.Group("/hotel-grades")
	routeGroup.POST("", handler.create)
	routeGroup.PUT("/:id", handler.update)
	routeGroup.GET("/:id", handler.find)
	routeGroup.DELETE("/:id", handler.delete)
	routeGroup.GET("", handler.findAll, middlewares2.PaginationMiddleware)
}

// @Summary create new HotelGrade
// @Tags HotelGrade
// @Accept json
// @Produce json
// @Param X-TenantID header int true "X-TenantID"
// @Param  HotelGrade body  models.HotelGrade true "HotelGrade"
// @Success 200 {object} models.HotelGrade
// @Router /hotel-grades [post]
func (handler *HotelGradeHandler) create(c echo.Context) error {

	model := &models.HotelGrade{}
	lang := c.Request().Header.Get(acceptLanguage)
	user := getCurrentUser(c)

	if err := c.Bind(&model); err != nil {

		handler.Input.Logger.LogError(err.Error())

		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusBadRequest,
			Message:      message_keys.BadRequest,
		})
	}

	model.SetAudit(user)
	result, err := handler.Service.Create(model)

	if err != nil {

		handler.Input.Logger.LogError(err.Error())

		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusBadRequest,
			Message:      "",
		})
	}

	return c.JSON(http.StatusOK, commons.ApiResponse{
		Data:         result,
		ResponseCode: http.StatusOK,
		Message:      handler.Input.Translator.Localize(lang, message_keys.Created),
	})
}

// @Summary update HotelGrade
// @Tags HotelGrade
// @Accept json
// @Param X-TenantID header int true "X-TenantID"
// @Param Id path int true "Id"
// @Param HotelGrade body models.HotelGrade true "HotelGrade"
// @Produce json
// @Param  HotelGrade body  models.HotelGrade true "HotelGrade"
// @Success 200 {object} models.HotelGrade
// @Router /hotel-grades/{id} [put]
func (handler *HotelGradeHandler) update(c echo.Context) error {

	lang := c.Request().Header.Get(acceptLanguage)
	id, err := utils.ConvertToUint(c.Param("id"))
	user := getCurrentUser(c)

	if err != nil {
		handler.Input.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}

	result, err := handler.Service.Find(id)
	if err != nil {
		handler.Input.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			ResponseCode: http.StatusBadRequest,
			Message:      handler.Input.Translator.Localize(lang, message_keys.BadRequest),
		})
	}

	// If not found any item, return 404.
	if result == nil || (result != nil && result.Id == 0) {
		return c.JSON(http.StatusNotFound, commons.ApiResponse{
			ResponseCode: http.StatusNotFound,
			Message:      handler.Input.Translator.Localize(lang, message_keys.NotFound),
		})
	}

	// bind request to tmp struct.
	tmpModel := models.HotelGrade{}
	err = c.Bind(&tmpModel)

	if err != nil {
		handler.Input.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			ResponseCode: http.StatusBadRequest,
			Message:      handler.Input.Translator.Localize(lang, message_keys.BadRequest),
		})
	}

	// prevent to edit other fields.
	result.Name = tmpModel.Name
	result.SetUpdatedBy(user)
	updatedModel, err := handler.Service.Update(result)

	if err != nil {
		handler.Input.Logger.LogError(err.Error())
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, commons.ApiResponse{
		Data:         updatedModel,
		ResponseCode: http.StatusOK,
	})
}

// @Summary find HotelGrade by id
// @Tags HotelGrade
// @Accept json
// @Param X-TenantID header int true "X-TenantID"
// @Param Id path int true "Id"
// @Produce json
// @Success 200 {object} models.HotelGrade
// @Router /hotel-grades/{id} [get]
func (handler *HotelGradeHandler) find(c echo.Context) error {

	lang := c.Request().Header.Get(acceptLanguage)
	id, err := utils.ConvertToUint(c.Param("id"))

	if err != nil {
		handler.Input.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}

	result, err := handler.Service.Find(id)
	if err != nil {

		handler.Input.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			ResponseCode: http.StatusBadRequest,
			Message:      handler.Input.Translator.Localize(lang, message_keys.BadRequest),
		})
	}

	if result == nil || (result != nil && result.Id == 0) {
		return c.JSON(http.StatusNotFound, commons.ApiResponse{
			ResponseCode: http.StatusNotFound,
			Message:      handler.Input.Translator.Localize(lang, message_keys.NotFound),
		})
	}

	return c.JSON(http.StatusOK, commons.ApiResponse{
		Data:         result,
		ResponseCode: http.StatusOK,
	})
}

// @Summary findAll HotelGrades
// @Tags HotelGrade
// @Accept json
// @Param X-TenantID header int true "X-TenantID"
// @Produce json
// @Success 200 {array} models.HotelGrade
// @Router /hotel-grades [get]
func (handler *HotelGradeHandler) findAll(c echo.Context) error {

	paginationInput := c.Get(paginationInput).(*dto.PaginationFilter)
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

// @Summary Delete HotelGrade
// @Tags HotelGrade
// @Accept json
// @Param X-TenantID header int true "X-TenantID"
// @Produce json
// @Success 200 {array} models.HotelGrade
// @Router /hotel-grades [delete]
func (handler *HotelGradeHandler) delete(c echo.Context) error {

	id, err := utils.ConvertToUint(c.Param("id"))
	lang := c.Request().Header.Get(acceptLanguage)

	if err != nil {

		handler.Input.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			ResponseCode: http.StatusBadRequest,
			Message:      handler.Input.Translator.Localize(lang, message_keys.BadRequest),
		})
	}

	err = handler.Service.Delete(id)

	if err != nil {
		return c.JSON(http.StatusConflict, commons.ApiResponse{
			ResponseCode: http.StatusConflict,
			Message:      handler.Input.Translator.Localize(lang, err.Error()),
		})
	}

	return c.JSON(http.StatusOK, commons.ApiResponse{
		ResponseCode: http.StatusOK,
		Message:      handler.Input.Translator.Localize(lang, message_keys.Deleted),
	})
}
