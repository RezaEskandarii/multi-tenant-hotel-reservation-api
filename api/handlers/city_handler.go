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
	"reservation-api/pkg/translator"
	"reservation-api/pkg/validator"
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

	currentUser := currentUser(c)
	city := models.City{}
	city.SetAudit(currentUser)

	if err := c.Bind(&city); err != nil {
		return c.JSON(http.StatusBadRequest,
			commons.ApiResponse{
				Data:         nil,
				ResponseCode: http.StatusBadRequest,
				Message:      translator.Localize(c.Request().Context(), message_keys.BadRequest),
				Errors:       nil,
			})
	}

	if err, messages := validator.Validate(city); err != nil {
		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			Errors:       messages,
			ResponseCode: http.StatusBadRequest,
		})
	}

	if _, err := handler.Service.Create(tenantContext(c), &city); err == nil {

		return c.JSON(http.StatusBadRequest,
			commons.ApiResponse{
				Data:         city,
				ResponseCode: http.StatusOK,
				Message:      translator.Localize(c.Request().Context(), message_keys.Created),
			})
	} else {

		handler.Config.Logger.LogError(err.Error())
		return c.JSON(http.StatusInternalServerError,
			commons.ApiResponse{
				Data:         nil,
				ResponseCode: http.StatusInternalServerError,
				Message:      translator.Localize(c.Request().Context(), message_keys.InternalServerError),
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

	currentUser := currentUser(c)
	model, err := handler.Service.Find(tenantContext(c), id)

	if err != nil {

		handler.Config.Logger.LogError(err.Error())

		return c.JSON(http.StatusInternalServerError, commons.ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusInternalServerError,
			Message:      translator.Localize(c.Request().Context(), message_keys.InternalServerError),
		})
	}

	if model == nil {
		return c.JSON(http.StatusNotFound, commons.ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusNotFound,
			Message:      translator.Localize(c.Request().Context(), message_keys.NotFound),
		})
	}

	if err := c.Bind(&model); err != nil {
		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusBadRequest,
			Message:      translator.Localize(c.Request().Context(), message_keys.BadRequest),
		})
	}

	if err, messages := validator.Validate(model); err != nil {
		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			Errors:       messages,
			ResponseCode: http.StatusBadRequest,
		})
	}

	model.SetUpdatedBy(currentUser)
	if output, err := handler.Service.Update(tenantContext(c), model); err == nil {

		return c.JSON(http.StatusOK, commons.ApiResponse{
			Data:         output,
			ResponseCode: http.StatusOK,
			Message:      translator.Localize(c.Request().Context(), message_keys.Updated),
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

	model, err := handler.Service.Find(tenantContext(c), id)

	if err != nil {
		handler.Config.Logger.LogError(err.Error())
		return c.JSON(http.StatusInternalServerError, commons.ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusInternalServerError,
			Message:      translator.Localize(c.Request().Context(), message_keys.InternalServerError),
		})
	}

	if model == nil {
		return c.JSON(http.StatusNotFound, commons.ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusNotFound,
			Message:      translator.Localize(c.Request().Context(), message_keys.NotFound),
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
