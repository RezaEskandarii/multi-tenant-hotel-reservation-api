package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	middlewares2 "reservation-api/api/middlewares"
	. "reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/message_keys"
	"reservation-api/internal/models"

	_ "reservation-api/internal/models"
	"reservation-api/internal/services/domain_services"
	"reservation-api/internal/utils"
)

// CountryHandler country endpoint handler
type CountryHandler struct {
	Service *domain_services.CountryService
	Config  *dto.HandlerConfig
}

func (handler *CountryHandler) Register(config *dto.HandlerConfig, service *domain_services.CountryService) {
	handler.Service = service
	handler.Config = config
	routeGroup := handler.Config.Router.Group("/countries")
	routeGroup.POST("", handler.create)
	routeGroup.PUT("/:id", handler.update)
	routeGroup.GET("/:id", handler.find)
	routeGroup.GET("/:id/provinces", handler.provinces)
	routeGroup.GET("", handler.findAll, middlewares2.PaginationMiddleware)
}

// @Summary create new Country
// @Tags Country
// @Accept json
// @Produce json
// @Param X-TenantID header int true "X-TenantID"
// @Param  country body  models.Country true "Country"
// @Success 200 {object} models.Country
// @Router /countries [post]
func (handler *CountryHandler) create(c echo.Context) error {

	model := &models.CountryCreateUpdate{}
	lang := getAcceptLanguage(c)

	if err := c.Bind(&model); err != nil {
		handler.Config.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest,
			ApiResponse{
				Data:         nil,
				ResponseCode: http.StatusBadRequest,
				Message:      handler.Config.Translator.Localize(lang, message_keys.BadRequest),
			})
	}

	if ok, err := model.Validate(); err != nil && ok == false {
		return c.JSON(http.StatusBadRequest, ApiResponse{Message: err.Error()})
	}

	setCreatedByUpdatedBy(model.BaseModel, getCurrentUser(c))
	output, err := handler.Service.Create(tenantContext(c), model)

	if err != nil {
		handler.Config.Logger.LogError(err.Error())

		return c.JSON(http.StatusBadRequest, ApiResponse{
			ResponseCode: http.StatusBadRequest,
			Message:      err.Error(),
		})
	}

	return c.JSON(http.StatusOK, ApiResponse{
		ResponseCode: http.StatusOK,
		Message:      handler.Config.Translator.Localize(lang, message_keys.Created),
		Data:         output,
	})
}

// @Summary update Country
// @Tags Country
// @Accept json
// @Param X-TenantID header int true "X-TenantID"
// @Param Id path int true "Id"
// @Param Country body models.Country true "Country"
// @Produce json
// @Param  country body  models.Country true "Country"
// @Success 200 {object} models.Country
// @Router /countries/{id} [put]
func (handler *CountryHandler) update(c echo.Context) error {

	id, err := utils.ConvertToUint(c.Param("id"))
	if err != nil {

		handler.Config.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}

	user := getCurrentUser(c)
	model, err := handler.Service.Find(tenantContext(c), id)
	lang := getAcceptLanguage(c)

	if err != nil {
		handler.Config.Logger.LogError(err.Error())
		return c.JSON(http.StatusInternalServerError, ApiResponse{
			ResponseCode: http.StatusInternalServerError,
			Message:      handler.Config.Translator.Localize(lang, message_keys.InternalServerError),
		})
	}

	if model == nil {
		return c.JSON(http.StatusNotFound, ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusNotFound,
			Message:      handler.Config.Translator.Localize(lang, message_keys.NotFound),
		})
	}

	if err := c.Bind(&model); err != nil {
		handler.Config.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}

	model.SetUpdatedBy(user)
	if output, err := handler.Service.Update(tenantContext(c), model); err == nil {

		return c.JSON(http.StatusOK, ApiResponse{
			Data:         output,
			ResponseCode: http.StatusOK,
			Message:      handler.Config.Translator.Localize(lang, message_keys.Updated),
		})
	} else {
		handler.Config.Logger.LogError(err.Error())
		return c.JSON(http.StatusInternalServerError, nil)
	}
}

// @Summary find Country by id
// @Tags Country
// @Accept json
// @Param X-TenantID header int true "X-TenantID"
// @Param Id path int true "Id"
// @Produce json
// @Success 200 {object} models.Country
// @Router /countries/{id} [get]
func (handler *CountryHandler) find(c echo.Context) error {

	id, err := utils.ConvertToUint(c.Param("id"))
	if err != nil {
		handler.Config.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}

	model, err := handler.Service.Find(tenantContext(c), id)
	lang := getAcceptLanguage(c)

	if err != nil {
		handler.Config.Logger.LogError(err.Error())
		return c.JSON(http.StatusInternalServerError, ApiResponse{
			ResponseCode: http.StatusInternalServerError,
			Message:      handler.Config.Translator.Localize(lang, message_keys.InternalServerError),
		})
	}

	if model == nil {
		return c.JSON(http.StatusNotFound, ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusNotFound,
			Message:      handler.Config.Translator.Localize(lang, message_keys.NotFound),
		})
	}

	return c.JSON(http.StatusOK, ApiResponse{
		Data:         model,
		ResponseCode: http.StatusOK,
		Message:      "",
	})
}

// @Summary findAll Countries
// @Tags Country
// @Accept json
// @Param X-TenantID header int true "X-TenantID"
// @Produce json
// @Success 200 {array} models.Country
// @Router /countries [get]
func (handler *CountryHandler) findAll(c echo.Context) error {

	paginationInput := c.Get(paginationInput).(*dto.PaginationFilter)

	list, err := handler.Service.FindAll(tenantContext(c), paginationInput)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, ApiResponse{
		Data:         list,
		ResponseCode: http.StatusOK,
		Message:      "",
	})
}

// @Summary find Provinces by country ID
// @Tags Country
// @Accept json
// @Param X-TenantID header int true "X-TenantID"
// @Param Id path int true "Id"
// @Produce json
// @Success 200 {array} models.Province
// @Router /countries/{id}/provinces [get]
func (handler *CountryHandler) provinces(c echo.Context) error {

	id, err := utils.ConvertToUint(c.Param("id"))

	if err != nil {
		handler.Config.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}

	provinces, err := handler.Service.GetProvinces(tenantContext(c), id)

	if err != nil {

		handler.Config.Logger.LogError(err.Error())

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
