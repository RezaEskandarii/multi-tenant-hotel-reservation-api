package handlers

import (
	"github.com/labstack/echo/v4"
	. "hotel-reservation/internal/commons"
	"hotel-reservation/internal/models"
	"hotel-reservation/internal/services"
	"net/http"
)

// CountryHandler country endpoint handler
type CountryHandler struct {
	Router  *echo.Group
	Service services.CountryService
}

func (handler *CountryHandler) Register(router *echo.Group, service services.CountryService) {
	handler.Router = router
	handler.Service = service
	handler.Router.POST("", handler.create)
}

func (handler *CountryHandler) create(c echo.Context) error {

	model := models.Country{}

	if err := c.Bind(&model); err != nil {
		return c.JSON(http.StatusBadRequest,
			ApiResponse{
				Data:         nil,
				ResponseCode: BadRequest,
				Message:      "BadRequest",
			})
	}

	if err, _ := handler.Service.Create(&model); err == nil {
		return c.JSON(http.StatusBadRequest,
			ApiResponse{
				Data:         model,
				ResponseCode: Ok,
				Message:      "Ok",
			})
	}

	return c.JSON(http.StatusInternalServerError,
		ApiResponse{
			Data:         nil,
			ResponseCode: InternalServerError,
			Message:      "BadRequest",
		})

}
