package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/message_keys"
	"reservation-api/internal/models"
	"reservation-api/internal/services/domain_services"
)

type ReservationHandler struct {
	Service *domain_services.ReservationService
	Input   *dto.HandlerInput
	Router  *echo.Group
}

func (r *ReservationHandler) Register(input *dto.HandlerInput, service *domain_services.ReservationService) {
	r.Router = input.Router
	routerGroup := r.Router.Group("/reservation")
	r.Input = input

	r.Service = service
	routerGroup.POST("/room-request", r.createRequest)
	routerGroup.POST("", r.create)
	routerGroup.DELETE("/cancel", r.cancelRequest)
}

func (r *ReservationHandler) createRequest(c echo.Context) error {

	request := dto.RoomRequestDto{}
	if err := c.Bind(&request); err != nil {
		r.Input.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}
	hasConflict, err := r.Service.HasConflict(&request)
	if err != nil {
		r.Input.Logger.LogError(err.Error())
		return c.JSON(http.StatusConflict, commons.ApiResponse{
			Message: err.Error(),
		})
	}
	if hasConflict {
		message := fmt.Sprintf(r.Input.Translator.Localize(getAcceptLanguage(c), message_keys.RoomHasReservationRequest), request.CheckInDate, request.CheckOutDate)
		return c.JSON(http.StatusConflict, commons.ApiResponse{
			Message: message,
		})
	}
	result, err := r.Service.CreateReservationRequest(&request)
	if err != nil {
		r.Input.Logger.LogError(err.Error())
		return c.JSON(http.StatusConflict, commons.ApiResponse{
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, commons.ApiResponse{
		Data:    result,
		Message: r.Input.Translator.Localize(getAcceptLanguage(c), message_keys.Created),
	})
}

func (r *ReservationHandler) create(c echo.Context) error {

	reservation := models.Reservation{}
	if err := c.Bind(&reservation); err != nil {
		r.Input.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}
	result, err := r.Service.Create(&reservation)
	if err != nil {
		r.Input.Logger.LogError(err.Error())
		return c.JSON(http.StatusConflict, commons.ApiResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, commons.ApiResponse{
		Data:    result,
		Message: r.Input.Translator.Localize(getAcceptLanguage(c), message_keys.Created),
	})
}

func (r *ReservationHandler) cancelRequest(c echo.Context) error {
	requestKey := c.QueryParam("requestKey")
	if err := r.Service.CancelReservationRequest(requestKey); err != nil {
		r.Input.Logger.LogError(err.Error())
	}
	return c.JSON(http.StatusOK, nil)
}
