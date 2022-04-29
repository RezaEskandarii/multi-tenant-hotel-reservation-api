package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/message_keys"
	"reservation-api/internal/models"
	"reservation-api/internal/services/domain_services"
	"reservation-api/internal/utils"
)

// GuestHandler  Currency endpoint handler
type GuestHandler struct {
	Service *domain_services.GuestService
	Input   *dto.HandlersSharedObjects
}

func (handler *GuestHandler) Register(input *dto.HandlersSharedObjects, service *domain_services.GuestService) {
	handler.Input = input
	handler.Service = service
	routeGroup := handler.Input.Router.Group("/guests")
	routeGroup.POST("", handler.create)
	routeGroup.GET("/:id", handler.find)
	routeGroup.GET("", handler.findAll)
	routeGroup.PUT("", handler.update)
	//routeGroup.DELETE("", handler)
}

/*====================================================================================*/
func (handler *GuestHandler) create(c echo.Context) error {

	model := models.Guest{}
	lang := c.Request().Header.Get(acceptLanguage)
	user := getCurrentUser(c)

	if err := c.Bind(&model); err != nil {

		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			ResponseCode: http.StatusBadRequest,
			Message:      handler.Input.Translator.Localize(c.Request().Header.Get(acceptLanguage), message_keys.BadRequest),
		})
	}

	model.SetAudit(user)
	if _, err := handler.Service.Create(&model); err != nil {

		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			Message: handler.Input.Translator.Localize(lang, err.Error()),
		})
	}

	return c.JSON(http.StatusOK, commons.ApiResponse{
		Data:    model,
		Message: handler.Input.Translator.Localize(lang, message_keys.Created),
	})
}

/*====================================================================================*/
func (handler *GuestHandler) update(c echo.Context) error {

	model := models.Guest{}
	lang := c.Request().Header.Get(acceptLanguage)
	user := getCurrentUser(c)
	id, _ := utils.ConvertToUint(c.Get("id"))

	guest, _ := handler.Service.Find(id)

	if guest == nil || (guest != nil && guest.Id == 0) {

		return c.JSON(http.StatusNotFound, commons.ApiResponse{
			Message: handler.Input.Translator.Localize(lang, message_keys.NotFound),
		})
	}

	if err := c.Bind(&model); err != nil {

		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			ResponseCode: http.StatusBadRequest,
			Message:      handler.Input.Translator.Localize(c.Request().Header.Get(acceptLanguage), message_keys.BadRequest),
		})
	}

	model.SetUpdatedBy(user)
	if _, err := handler.Service.Update(&model); err != nil {

		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			Message: handler.Input.Translator.Localize(lang, err.Error()),
		})
	}

	return c.JSON(http.StatusOK, commons.ApiResponse{
		Data:    model,
		Message: handler.Input.Translator.Localize(lang, message_keys.Updated),
	})
}

/*====================================================================================*/
func (handler *GuestHandler) find(c echo.Context) error {

	lang := c.Request().Header.Get(acceptLanguage)
	id, _ := utils.ConvertToUint(c.Get("id"))

	guest, _ := handler.Service.Find(id)

	if guest == nil || (guest != nil && guest.Id == 0) {

		return c.JSON(http.StatusNotFound, commons.ApiResponse{
			Message: handler.Input.Translator.Localize(lang, message_keys.NotFound),
		})
	}

	return c.JSON(http.StatusOK, commons.ApiResponse{
		Data: guest,
	})
}

/*====================================================================================*/
func (handler *GuestHandler) findAll(c echo.Context) error {

	lang := c.Request().Header.Get(acceptLanguage)

	page, _ := utils.ConvertToUint(c.Param("page"))
	perPage, _ := utils.ConvertToUint(c.Param("perPage"))

	input := &dto.PaginationFilter{
		Page:    int(page),
		PerPage: int(perPage),
	}

	result, err := handler.Service.FindAll(input)

	if err != nil {
		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			Message: handler.Input.Translator.Localize(lang, err.Error()),
		})
	}

	return c.JSON(http.StatusOK, commons.ApiResponse{
		Data: result,
	})
}
