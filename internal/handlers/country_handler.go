package handlers

import (
	"github.com/labstack/echo/v4"
	. "hotel-reservation/internal/commons"
	"hotel-reservation/internal/models"
	"net/http"
)

// CountryHandler country endpoint handler
type CountryHandler struct {
	router *echo.Group
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

	return c.JSON(http.StatusBadRequest,
		ApiResponse{
			Data:         model,
			ResponseCode: Ok,
			Message:      "Ok",
		})
}
