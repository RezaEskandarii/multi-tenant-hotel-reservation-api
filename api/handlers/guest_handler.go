package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/message_keys"
	"reservation-api/internal/models"
	"reservation-api/internal/services/common_services"
	"reservation-api/internal/services/domain_services"
	"reservation-api/internal/utils"
)

// GuestHandler  Guest endpoint handler
type GuestHandler struct {
	Service       *domain_services.GuestService
	Config        *dto.HandlerConfig
	ReportService *common_services.ReportService
}

func (handler *GuestHandler) Register(config *dto.HandlerConfig,
	service *domain_services.GuestService, reportService *common_services.ReportService) {
	handler.ReportService = reportService
	handler.Config = config
	handler.Service = service
	routeGroup := handler.Config.Router.Group("/guests")
	routeGroup.POST("", handler.create)
	routeGroup.GET("/:id", handler.find)
	routeGroup.GET("", handler.findAll)
	routeGroup.PUT("", handler.update)
	//routeGroup.DELETE("", handler)
}

// @Summary create new Guest
// @Tags Guest
// @Accept json
// @Produce json
// @Param X-TenantID header int true "X-TenantID"
// @Param  Guest body  models.Guest true "Guest"
// @Success 200 {object} models.Guest
// @Router /guests [post]
func (handler *GuestHandler) create(c echo.Context) error {

	model := models.Guest{}
	lang := c.Request().Header.Get(acceptLanguage)
	user := getCurrentUser(c)

	if err := c.Bind(&model); err != nil {

		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			ResponseCode: http.StatusBadRequest,
			Message:      handler.Config.Translator.Localize(c.Request().Header.Get(acceptLanguage), message_keys.BadRequest),
		})
	}

	model.SetAudit(user)
	if _, err := handler.Service.Create(getCurrentTenantContext(c), &model); err != nil {

		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			Message: handler.Config.Translator.Localize(lang, err.Error()),
		})
	}

	return c.JSON(http.StatusOK, commons.ApiResponse{
		Data:    model,
		Message: handler.Config.Translator.Localize(lang, message_keys.Created),
	})
}

// @Summary update Guest
// @Tags Guest
// @Accept json
// @Param X-TenantID header int true "X-TenantID"
// @Param Id path int true "Id"
// @Param Guest body models.Guest true "Guest"
// @Produce json
// @Param  Guest body  models.Guest true "Guest"
// @Success 200 {object} models.Guest
// @Router /guests/{id} [put]
func (handler *GuestHandler) update(c echo.Context) error {

	model := models.Guest{}
	lang := c.Request().Header.Get(acceptLanguage)
	user := getCurrentUser(c)
	id, _ := utils.ConvertToUint(c.Get("id"))

	guest, _ := handler.Service.Find(getCurrentTenantContext(c), id)

	if guest == nil || (guest != nil && guest.Id == 0) {

		return c.JSON(http.StatusNotFound, commons.ApiResponse{
			Message: handler.Config.Translator.Localize(lang, message_keys.NotFound),
		})
	}

	if err := c.Bind(&model); err != nil {

		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			ResponseCode: http.StatusBadRequest,
			Message:      handler.Config.Translator.Localize(c.Request().Header.Get(acceptLanguage), message_keys.BadRequest),
		})
	}

	model.SetUpdatedBy(user)
	if _, err := handler.Service.Update(getCurrentTenantContext(c), &model); err != nil {

		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			Message: handler.Config.Translator.Localize(lang, err.Error()),
		})
	}

	return c.JSON(http.StatusOK, commons.ApiResponse{
		Data:    model,
		Message: handler.Config.Translator.Localize(lang, message_keys.Updated),
	})
}

// @Summary find Guest by id
// @Tags Guest
// @Accept json
// @Param X-TenantID header int true "X-TenantID"
// @Param Id path int true "Id"
// @Produce json
// @Success 200 {object} models.Guest
// @Router /guests/{id} [get]
func (handler *GuestHandler) find(c echo.Context) error {

	lang := c.Request().Header.Get(acceptLanguage)
	id, _ := utils.ConvertToUint(c.Get("id"))

	guest, _ := handler.Service.Find(getCurrentTenantContext(c), id)

	if guest == nil || (guest != nil && guest.Id == 0) {

		return c.JSON(http.StatusNotFound, commons.ApiResponse{
			Message: handler.Config.Translator.Localize(lang, message_keys.NotFound),
		})
	}

	return c.JSON(http.StatusOK, commons.ApiResponse{
		Data: guest,
	})
}

// @Summary findAll Guests
// @Tags Guest
// @Accept json
// @Param X-TenantID header int true "X-TenantID"
// @Produce json
// @Success 200 {array} models.Guest
// @Router /guests [get]
func (handler *GuestHandler) findAll(c echo.Context) error {

	lang := c.Request().Header.Get(acceptLanguage)

	page, _ := utils.ConvertToUint(c.Param("page"))
	perPage, _ := utils.ConvertToUint(c.Param("perPage"))
	output := getOutputQueryParamVal(c)

	input := &dto.PaginationFilter{
		Page:    int(page),
		PerPage: int(perPage),
	}

	input.IgnorePagination = output != ""

	result, err := handler.Service.FindAll(getCurrentTenantContext(c), input)

	if err != nil {
		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			Message: handler.Config.Translator.Localize(lang, err.Error()),
		})
	}

	if output != "" {
		if output == EXCEL {
			report, err := handler.ReportService.ExportToExcel(result, getAcceptLanguage(c))
			if err != nil {
				handler.Config.Logger.LogError(err.Error())
				return c.JSON(http.StatusInternalServerError, commons.ApiResponse{})
			}
			writeBinaryHeaders(c, "guests", EXCEL_OUTPUT)
			c.Response().Write(report)
			return nil
		}
	}

	return c.JSON(http.StatusOK, commons.ApiResponse{
		Data: result,
	})
}
