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

// ProvinceHandler Province endpoint handler
type ProvinceHandler struct {
	Service *domain_services.ProvinceService
	Config  *dto.HandlerConfig
}

func (handler *ProvinceHandler) Register(config *dto.HandlerConfig, service *domain_services.ProvinceService) {
	handler.Service = service
	handler.Config = config
	routeGroup := config.Router.Group("/provinces")
	routeGroup.POST("", handler.create)
	routeGroup.PUT("/:id", handler.update)
	routeGroup.GET("/:id", handler.find)
	routeGroup.GET("/:id/cities", handler.cities)
	routeGroup.GET("", handler.findAll, middlewares2.PaginationMiddleware)
}

// @Summary update Province
// @Tags Province
// @Accept json
// @Param X-TenantID header int true "X-TenantID"
// @Param Id path int true "Id"
// @Param Province body models.Province true "Province"
// @Produce json
// @Param  Province body  models.Province true "Province"
// @Success 200 {object} models.Province
// @Router /provinces/{id} [put]
func (handler *ProvinceHandler) create(c echo.Context) error {

	province := &models.Province{}
	user := currentUser(c)
	lang := c.Request().Header.Get(acceptLanguage)

	if err := c.Bind(&province); err != nil {

		handler.Config.Logger.LogError(err.Error())

		return c.JSON(http.StatusBadRequest,
			commons.ApiResponse{
				Data:         nil,
				ResponseCode: http.StatusBadRequest,
				Message:      handler.Config.Translator.Localize(lang, message_keys.BadRequest),
			})
	}

	if ok, err := province.Validate(); !ok {
		return c.JSON(http.StatusBadRequest,
			commons.ApiResponse{
				Data:         nil,
				ResponseCode: http.StatusBadRequest,
				Message:      err.Error(),
			})
	}

	province.SetAudit(user)
	if result, err := handler.Service.Create(tenantContext(c), province); err == nil {

		return c.JSON(http.StatusBadRequest,
			commons.ApiResponse{
				Data:         result,
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

// @Summary update Province
// @Tags Province
// @Accept json
// @Param X-TenantID header int true "X-TenantID"
// @Param Id path int true "Id"
// @Param Province body models.Province true "Province"
// @Produce json
// @Param  Province body  models.Province true "Province"
// @Success 200 {object} models.Province
// @Router /provinces/{id} [put]
func (handler *ProvinceHandler) update(c echo.Context) error {

	id, err := utils.ConvertToUint(c.Param("id"))

	if err != nil {
		handler.Config.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}

	user := currentUser(c)
	province, err := handler.Service.Find(tenantContext(c), id)
	lang := getAcceptLanguage(c)

	if err != nil {

		handler.Config.Logger.LogError(err.Error())
		return c.JSON(http.StatusInternalServerError, commons.ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusInternalServerError,
			Message:      handler.Config.Translator.Localize(lang, message_keys.InternalServerError),
		})

	}

	if province == nil {
		return c.JSON(http.StatusNotFound, commons.ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusNotFound,
			Message:      handler.Config.Translator.Localize(lang, message_keys.NotFound),
		})
	}

	if err := c.Bind(&province); err != nil {

		handler.Config.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}

	if ok, err := province.Validate(); !ok {
		return c.JSON(http.StatusBadRequest,
			commons.ApiResponse{
				Data:         nil,
				ResponseCode: http.StatusBadRequest,
				Message:      err.Error(),
			})
	}
	province.SetUpdatedBy(user)
	if result, err := handler.Service.Update(tenantContext(c), province); err == nil {

		return c.JSON(http.StatusOK, commons.ApiResponse{
			Data:         result,
			ResponseCode: http.StatusOK,
			Message:      handler.Config.Translator.Localize(lang, message_keys.Updated),
		})
	} else {

		handler.Config.Logger.LogError(err.Error())
		return c.JSON(http.StatusInternalServerError, nil)
	}
}

// @Summary find Province by id
// @Tags Province
// @Accept json
// @Param X-TenantID header int true "X-TenantID"
// @Param Id path int true "Id"
// @Produce json
// @Success 200 {object} models.Province
// @Router /provinces/{id} [get]
func (handler *ProvinceHandler) find(c echo.Context) error {
	id, err := utils.ConvertToUint(c.Param("id"))
	if err != nil {

		handler.Config.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}
	model, err := handler.Service.Find(tenantContext(c), id)
	lang := c.Request().Header.Get(acceptLanguage)

	if err != nil {

		handler.Config.Logger.LogError(err.Error())
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
			Message:      handler.Config.Translator.Localize(lang, message_keys.NotFound),
		})
	}

	return c.JSON(http.StatusOK, commons.ApiResponse{
		Data:         model,
		ResponseCode: http.StatusOK,
		Message:      "",
	})
}

// @Summary findAll Provinces
// @Tags Province
// @Accept json
// @Param X-TenantID header int true "X-TenantID"
// @Produce json
// @Success 200 {array} models.Province
// @Router /provinces [get]
func (handler *ProvinceHandler) findAll(c echo.Context) error {

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

// @Summary find Province cities by Province ID
// @Tags Province
// @Accept json
// @Param X-TenantID header int true "X-TenantID"
// @Param Id path int true "Id"
// @Produce json
// @Success 200 {array} models.Province
// @Router /provinces/{id}/cities [get]
func (handler *ProvinceHandler) cities(c echo.Context) error {

	id, err := utils.ConvertToUint(c.Param("id"))
	if err != nil {

		handler.Config.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}
	cities, err := handler.Service.GetCities(tenantContext(c), id)

	if err != nil {

		handler.Config.Logger.LogError(err.Error())
		return c.JSON(http.StatusInternalServerError, commons.NewApiResponse().
			SetResponseCode(http.StatusInternalServerError).SetMessage("bad request"),
		)

	}

	return c.JSON(http.StatusOK, commons.ApiResponse{
		Data:         cities,
		ResponseCode: http.StatusOK,
		Message:      "",
	})
}
