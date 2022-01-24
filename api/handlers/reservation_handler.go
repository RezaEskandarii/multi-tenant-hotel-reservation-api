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
	"reservation-api/internal/utils"
	"strings"
	"time"
)

type ReservationHandler struct {
	Service *domain_services.ReservationService
	Input   *dto.HandlerInput
	Router  *echo.Group
}

func (handler *ReservationHandler) Register(input *dto.HandlerInput, service *domain_services.ReservationService) {
	handler.Router = input.Router
	routerGroup := handler.Router.Group("/reservation")
	handler.Input = input
	handler.Service = service
	routerGroup.POST("/room-request", handler.createRequest)
	routerGroup.POST("", handler.create)
	routerGroup.DELETE("/cancel", handler.cancelRequest)
	routerGroup.POST("/recommend-rate-codes", handler.recommendRateCodes)
	routerGroup.GET("/:id", handler.find)
	routerGroup.PUT("/:id", handler.update)
	routerGroup.PUT("/change-status/:id", handler.changeStatus)
}

/*=====================================================================================================*/
func (handler *ReservationHandler) createRequest(c echo.Context) error {

	lang := getAcceptLanguage(c)
	reservationIdStr := c.QueryParam("reservationId")
	reservation := &models.Reservation{}

	// If the client requests to edit a reservation,
	// client must send the reservation ID to avoid conflicts with other reservations on this check-in and check-out date.
	if strings.TrimSpace(reservationIdStr) != "" {

		reservationId, _ := utils.ConvertToUint(reservationIdStr)
		reservationResult, err := handler.Service.Find(reservationId)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, nil)
		}

		reservation = reservationResult

	} else {
		reservation = nil
	}
	request := dto.RoomRequestDto{}
	if err := c.Bind(&request); err != nil {
		handler.Input.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}
	// Checks if there is another reservation request for this room on the same check-in and check-out date,
	// otherwise do not allow a booking request.
	hasConflict, err := handler.Service.HasConflict(&request, reservation)
	if err != nil {
		handler.Input.Logger.LogError(err.Error())
		return c.JSON(http.StatusConflict, commons.ApiResponse{
			Message: err.Error(),
		})
	}
	// If there is a simultaneous booking request, the booking request is not given.
	if hasConflict {
		message := fmt.Sprintf(handler.Input.Translator.Localize(lang, message_keys.RoomHasReservationRequest), request.CheckInDate, request.CheckOutDate)
		return c.JSON(http.StatusConflict, commons.ApiResponse{
			Message: message,
		})
	}

	// prevent to reserve room for past dates.
	if request.CheckInDate.Before(time.Now()) && request.RequestType == dto.CreateReservation {
		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			Message: handler.Input.Translator.Localize(lang, message_keys.ImpossibleReservationLatDateError),
		})
	}

	if request.CheckInDate == nil {
		return c.JSON(http.StatusBadRequest,
			commons.ApiResponse{
				Message: handler.Input.Translator.Localize(lang, message_keys.CheckInDateEmptyError)})
	}
	if request.CheckOutDate == nil {
		return c.JSON(http.StatusBadRequest,
			commons.ApiResponse{
				Message: handler.Input.Translator.Localize(lang, message_keys.CheckOutDateEmptyError)})
	}

	// create new reservation request for requested room.
	result, err := handler.Service.CreateReservationRequest(&request)
	if err != nil {
		handler.Input.Logger.LogError(err.Error())
		return c.JSON(http.StatusConflict, commons.ApiResponse{
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, commons.ApiResponse{
		Data:    result,
		Message: handler.Input.Translator.Localize(getAcceptLanguage(c), message_keys.Created),
	})
}

/*==========================================================================================================*/

func (handler *ReservationHandler) create(c echo.Context) error {

	reservation := models.Reservation{}
	if err := c.Bind(&reservation); err != nil {
		handler.Input.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}

	lang := getAcceptLanguage(c)
	invalidReservationRequestKeyErr := handler.Input.Translator.Localize(lang, message_keys.InvalidReservationRequestKey)
	if strings.TrimSpace(reservation.RequestKey) == "" {
		return c.JSON(http.StatusBadRequest,
			commons.ApiResponse{
				Message: invalidReservationRequestKeyErr,
			})
	}

	reservationRequest, err := handler.Service.FindReservationRequest(reservation.RequestKey)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	if reservationRequest == nil {
		return c.JSON(http.StatusBadRequest,
			commons.ApiResponse{
				Message: invalidReservationRequestKeyErr,
			})
	}

	if time.Now().After(reservationRequest.ExpireTime) {
		return c.JSON(http.StatusBadRequest,
			commons.ApiResponse{
				Message: invalidReservationRequestKeyErr,
			})
	}

	if len(reservation.Sharers) == 0 {
		return c.JSON(http.StatusBadRequest,
			commons.ApiResponse{
				Message: handler.Input.Translator.Localize(lang, message_keys.EmptySharerError),
			})
	}

	hasReservationConflict, err := handler.Service.HasReservationConflict(reservation.CheckinDate, reservation.CheckoutDate, reservation.RoomId)
	if err != nil {
		handler.Input.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}

	if hasReservationConflict {
		return c.JSON(http.StatusBadRequest,
			commons.ApiResponse{
				Message: handler.Input.Translator.Localize(lang, message_keys.ReservationConflictError),
			})
	}
	handler.setReservationFields(&reservation, reservationRequest)
	// create new reservation.
	result, err := handler.Service.Create(&reservation)
	if err != nil {
		handler.Input.Logger.LogError(err.Error())
		return c.JSON(http.StatusConflict, commons.ApiResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, commons.ApiResponse{
		Data:    result,
		Message: handler.Input.Translator.Localize(getAcceptLanguage(c), message_keys.Created),
	})
}

/*===================================================================================*/

func (handler *ReservationHandler) update(c echo.Context) error {

	id, err := utils.ConvertToUint(c.Param("id"))
	lang := c.Request().Header.Get(acceptLanguage)

	if err != nil {
		handler.Input.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			ResponseCode: http.StatusBadRequest,
		})
	}

	reservationModel, err := handler.Service.Find(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			ResponseCode: http.StatusInternalServerError,
		})
	}

	if reservationModel == nil {
		return c.JSON(http.StatusNotFound, commons.ApiResponse{
			ResponseCode: http.StatusNotFound,
		})
	}

	reservation := models.Reservation{}
	if err := c.Bind(&reservation); err != nil {
		handler.Input.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}

	invalidReservationRequestKeyErr := handler.Input.Translator.Localize(lang, message_keys.InvalidReservationRequestKey)
	if strings.TrimSpace(reservation.RequestKey) == "" {
		return c.JSON(http.StatusBadRequest,
			commons.ApiResponse{
				Message: invalidReservationRequestKeyErr,
			})
	}

	reservationRequest, err := handler.Service.FindReservationRequest(reservation.RequestKey)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	if reservationRequest == nil {
		return c.JSON(http.StatusBadRequest,
			commons.ApiResponse{
				Message: invalidReservationRequestKeyErr,
			})
	}

	if time.Now().After(reservationRequest.ExpireTime) {
		return c.JSON(http.StatusBadRequest,
			commons.ApiResponse{
				Message: invalidReservationRequestKeyErr,
			})
	}

	if len(reservation.Sharers) == 0 {
		return c.JSON(http.StatusBadRequest,
			commons.ApiResponse{
				Message: handler.Input.Translator.Localize(lang, message_keys.EmptySharerError),
			})
	}

	hasReservationConflict, err := handler.Service.HasReservationConflict(reservation.CheckinDate, reservation.CheckoutDate, reservation.RoomId)
	if err != nil {
		handler.Input.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}

	if hasReservationConflict {
		return c.JSON(http.StatusBadRequest,
			commons.ApiResponse{
				Message: handler.Input.Translator.Localize(lang, message_keys.ReservationConflictError),
			})
	}
	handler.setReservationFields(&reservation, reservationRequest)
	// create new reservation.
	result, err := handler.Service.Update(id, &reservation)
	if err != nil {
		handler.Input.Logger.LogError(err.Error())
		return c.JSON(http.StatusConflict, commons.ApiResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, commons.ApiResponse{
		Data:    result,
		Message: handler.Input.Translator.Localize(getAcceptLanguage(c), message_keys.Created),
	})
}

/*===================================================================================*/

// If the client cancels the reservation request, they can call this endpoint to delete the reservation request.
func (handler *ReservationHandler) cancelRequest(c echo.Context) error {
	requestKey := c.QueryParam("requestKey")
	if err := handler.Service.RemoveReservationRequest(requestKey); err != nil {
		handler.Input.Logger.LogError(err.Error())
	}
	return c.JSON(http.StatusOK, nil)
}

/*===================================================================================*/
func (handler *ReservationHandler) recommendRateCodes(c echo.Context) error {

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

	result, err := handler.Service.GetRecommendedRateCodes(&priceDto)
	if err != nil {
		handler.Input.Logger.LogError(err.Error())
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, commons.ApiResponse{
		Data: result,
	})
}

/*===================================================================================*/
func (handler *ReservationHandler) find(c echo.Context) error {
	id, err := utils.ConvertToUint(c.Param("id"))

	if err != nil {
		handler.Input.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}

	result, err := handler.Service.Find(id)
	if err != nil {
		handler.Input.Logger.LogError(err.Error())
		return c.JSON(http.StatusInternalServerError, nil)
	}

	if result == nil {
		return c.JSON(http.StatusNotFound, nil)
	}

	return c.JSON(http.StatusOK, commons.ApiResponse{
		Data: result,
	})
}

/*===================================================================================*/
func (handler *ReservationHandler) changeStatus(c echo.Context) error {
	id, err := utils.ConvertToUint(c.Param("id"))

	if err != nil {
		handler.Input.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}

	reservation, err := handler.Service.Find(id)
	if err != nil {
		handler.Input.Logger.LogError(err.Error())
		return c.JSON(http.StatusInternalServerError, nil)
	}

	if reservation == nil {
		return c.JSON(http.StatusNotFound, nil)
	}

	statusVal, err := utils.ConvertToUint(c.QueryParam("status"))
	status := models.ReservationCheckStatus(statusVal)

	if err != nil {
		handler.Input.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}

	if status == models.CheckIn || status == models.Checkout {
		reservation.CheckStatus = status
		_, err := handler.Service.ChangeStatus(id, status)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, nil)
		}

	} else {
		return c.JSON(http.StatusBadRequest, nil)
	}

	return c.JSON(http.StatusOK, commons.ApiResponse{
		Data: reservation,
	})
}

/*====================================================================================*/
func (handler *ReservationHandler) setReservationFields(reservation *models.Reservation, reservationRequest *models.ReservationRequest) {
	reservation.CheckinDate = reservationRequest.CheckInDate
	reservation.CheckoutDate = reservationRequest.CheckOutDate
	reservation.RoomId = reservationRequest.RoomId
}
