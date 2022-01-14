package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/services/domain_services"
)

type ReservationHandler struct {
	Service *domain_services.ReservationService
	Input   *dto.HandlerInput
	Router  *echo.Group
}

func (r *ReservationHandler) Register(input *dto.HandlerInput, service *domain_services.ReservationService) {
	r.Router = input.Router
	routerGroup := r.Router.Group("reservation")

	r.Service = service
	routerGroup.POST("/room-request", r.createRequest)
}

func (r *ReservationHandler) createRequest(c echo.Context) error {

	request := dto.RoomRequestDto{}
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	result, err := r.Service.CreateReservationRequest(&request)
	if err != nil {
		return c.JSON(http.StatusConflict, commons.ApiResponse{
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, commons.ApiResponse{
		Data: result},
	)
}
