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
	routerGroup.POST("/recommend-rate-codes", r.recommendRateCodes)
}

/*=====================================================================================================*/
func (r *ReservationHandler) createRequest(c echo.Context) error {

	request := dto.RoomRequestDto{}
	if err := c.Bind(&request); err != nil {
		r.Input.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}
	// Checks if there is another reservation request for this room on the same check-in and check-out date,
	// otherwise do not allow a booking request.
	hasConflict, err := r.Service.HasConflict(&request)
	if err != nil {
		r.Input.Logger.LogError(err.Error())
		return c.JSON(http.StatusConflict, commons.ApiResponse{
			Message: err.Error(),
		})
	}
	// If there is a simultaneous booking request, the booking request is not given.
	if hasConflict {
		message := fmt.Sprintf(r.Input.Translator.Localize(getAcceptLanguage(c), message_keys.RoomHasReservationRequest), request.CheckInDate, request.CheckOutDate)
		return c.JSON(http.StatusConflict, commons.ApiResponse{
			Message: message,
		})
	}
	// create new reservation request for requested room.
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

/*==========================================================================================================*/

func (r *ReservationHandler) create(c echo.Context) error {

	reservation := models.Reservation{}
	if err := c.Bind(&reservation); err != nil {
		r.Input.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}
	// create new reservation.
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

/*===================================================================================*/

// If the client cancels the reservation request, they can call this endpoint to delete the reservation request.
func (r *ReservationHandler) cancelRequest(c echo.Context) error {
	requestKey := c.QueryParam("requestKey")
	if err := r.Service.CancelReservationRequest(requestKey); err != nil {
		r.Input.Logger.LogError(err.Error())
	}
	return c.JSON(http.StatusOK, nil)
}

/*===================================================================================*/
func (r *ReservationHandler) recommendRateCodes(c echo.Context) error {

	priceDto := dto.GetRatePriceDto{}

	if err := c.Bind(&priceDto); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	isValid, err := priceDto.Validate()
	if isValid == false && err != nil {
		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			Data: err.Error(),
		})
	}

	result, err := r.Service.GetRecommendedRateCodes(&priceDto)
	if err != nil {
		r.Input.Logger.LogError(err.Error())
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, commons.ApiResponse{
		Data: result,
	})
}
