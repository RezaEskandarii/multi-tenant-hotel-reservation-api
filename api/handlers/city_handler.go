package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	m "reservation-api/api/middlewares"
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/message_keys"
	"reservation-api/internal/models"
	"reservation-api/internal/services/domain_services"
	"reservation-api/internal/utils"
)

// CityHandler City endpoint handler
type CityHandler struct {
	Service *domain_services.CityService
	Config  *dto.HandlerConfig
}

func (handler *CityHandler) Register(config *dto.HandlerConfig, service *domain_services.CityService) {
	handler.Service = service
	handler.Config = config
	routeGroup := handler.Config.Router.Group("/cities")
	routeGroup.POST("", handler.create)
	routeGroup.PUT("/:id", handler.update)
	routeGroup.GET("/:id", handler.find)
	routeGroup.GET("", handler.findAll, m.PaginationMiddleware)
}

// create new city
func (handler *CityHandler) create(c echo.Context) error {

	currentUser := getCurrentUser(c)
	city := models.City{}
	city.SetAudit(currentUser)
	lang := getAcceptLanguage(c)

	if err := c.Bind(&city); err != nil {
		return c.JSON(http.StatusBadRequest,
			commons.ApiResponse{
				Data:         nil,
				ResponseCode: http.StatusBadRequest,
				Message:      handler.Config.Translator.Localize(lang, message_keys.BadRequest),
			})
	}

	if _, err := handler.Service.Create(getCurrentTenantContext(c), &city); err == nil {

		return c.JSON(http.StatusBadRequest,
			commons.ApiResponse{
				Data:         city,
				ResponseCode: http.StatusOK,
				Message:      handler.Config.Translator.Localize(lang, message_keys.Created),
			})
	} else {

		handler.Config.Logger.LogError(err.Error())
		return c.JSON(http.StatusInternalServerError,
			commons.ApiResponse{
				Data:         nil,
				ResponseCode: http.StatusInternalServerError,
				Message:      handler.Config.Translator.Localize(lang, message_keys.InternalServerError),
			})
	}

}

/*====================================================================================*/
func (handler *CityHandler) update(c echo.Context) error {

	id, err := utils.ConvertToUint(c.Param("id"))

	if err != nil {
		handler.Config.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}

	lang := getAcceptLanguage(c)
	currentUser := getCurrentUser(c)
	model, err := handler.Service.Find(getCurrentTenantContext(c), id)

	if err != nil {

		handler.Config.Logger.LogError(err.Error())

		return c.JSON(http.StatusInternalServerError, commons.ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusInternalServerError,
			Message:      handler.Config.Translator.Localize(lang, message_keys.InternalServerError),
		})
	}

	if model == nil {
		return c.JSON(http.StatusNotFound, commons.ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusNotFound,
			Message:      handler.Config.Translator.Localize(lang, message_keys.NotFound),
		})
	}

	if err := c.Bind(&model); err != nil {
		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusBadRequest,
			Message:      handler.Config.Translator.Localize(lang, message_keys.BadRequest),
		})
	}
	model.SetUpdatedBy(currentUser)
	if output, err := handler.Service.Update(getCurrentTenantContext(c), model); err == nil {

		return c.JSON(http.StatusOK, commons.ApiResponse{
			Data:         output,
			ResponseCode: http.StatusOK,
			Message:      handler.Config.Translator.Localize(lang, message_keys.Updated),
		})
	} else {
		handler.Config.Logger.LogError(err.Error())
		return c.JSON(http.StatusInternalServerError, nil)
	}
}

/*====================================================================================*/
func (handler *CityHandler) find(c echo.Context) error {

	id, err := utils.ConvertToUint(c.Param("id"))
	if err != nil {
		handler.Config.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}

	model, err := handler.Service.Find(getCurrentTenantContext(c), id)
	lang := getAcceptLanguage(c)

	if err != nil {
		handler.Config.Logger.LogError(err.Error())
		return c.JSON(http.StatusInternalServerError, commons.ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusInternalServerError,
			Message:      handler.Config.Translator.Localize(lang, message_keys.InternalServerError),
		})
	}

	if model == nil {
		return c.JSON(http.StatusNotFound, commons.ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusNotFound,
			Message:      handler.Config.Translator.Localize(lang, message_keys.NotFound),
		})
	}

	return c.JSON(http.StatusOK, commons.ApiResponse{
		Data:         model,
		ResponseCode: http.StatusOK,
		Message:      "",
	})
}

/*====================================================================================*/
func (handler *CityHandler) findAll(c echo.Context) error {

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
