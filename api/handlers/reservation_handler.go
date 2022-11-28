package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/message_keys"
	"reservation-api/internal/models"
	"reservation-api/internal/services/common_services"
	"reservation-api/internal/services/domain_services"
	"reservation-api/internal/utils"
	"strings"
	"time"
)

type ReservationHandler struct {
	Service       *domain_services.ReservationService
	Config        *dto.HandlerConfig
	Router        *echo.Group
	ReportService *common_services.ReportService
}

func (handler *ReservationHandler) Register(config *dto.HandlerConfig, service *domain_services.ReservationService,
	reportService *common_services.ReportService) {
	handler.Router = config.Router
	handler.ReportService = reportService
	routerGroup := handler.Router.Group("/reservation")
	handler.Config = config
	handler.Service = service
	routerGroup.POST("/room-request", handler.createRequest)
	routerGroup.POST("", handler.create)
	routerGroup.DELETE("/cancel", handler.cancelRequest)
	routerGroup.POST("/recommend-rate-codes", handler.recommendRateCodes)
	routerGroup.GET("/:id", handler.find)
	routerGroup.GET("", handler.findAll)
	routerGroup.PUT("/:id", handler.update)
	routerGroup.PUT("/change-status/:id", handler.changeStatus)

}

// @Summary SetUp Reservation
// @Tags Reservation
// @Accept json
// @Param X-TenantID header int true "X-TenantID"
// @Produce json
// @Param  Reservation body  models.Reservation true "Reservation"
// @Success 200 {object} models.Reservation
// @Router /reservations/{id} [post]
func (handler *ReservationHandler) createRequest(c echo.Context) error {

	lang := getAcceptLanguage(c)
	reservationIdStr := c.QueryParam("reservationId")
	reservation := &models.Reservation{}

	// If the client requests to edit a reservation,
	// client must send the reservation ID to avoid conflicts with other reservations on this check-in and check-out date.
	if strings.TrimSpace(reservationIdStr) != "" {

		reservationId, _ := utils.ConvertToUint(reservationIdStr)
		reservationResult, err := handler.Service.Find(tenantContext(c), reservationId)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, nil)
		}

		reservation = reservationResult

	} else {
		reservation = nil
	}
	request := dto.RoomRequestDto{}
	if err := c.Bind(&request); err != nil {
		handler.Config.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}
	// Checks if there is another reservation request for this room on the same check-in and check-out date,
	// otherwise do not allow a booking request.
	hasConflict, err := handler.Service.HasConflict(tenantContext(c), &request, reservation)
	if err != nil {
		handler.Config.Logger.LogError(err.Error())
		return c.JSON(http.StatusConflict, commons.ApiResponse{
			Message: err.Error(),
		})
	}
	// If there is a simultaneous booking request, the booking request is not given.
	if hasConflict {
		message := fmt.Sprintf(handler.Config.Translator.Localize(lang, message_keys.RoomHasReservationRequest), request.CheckInDate, request.CheckOutDate)
		return c.JSON(http.StatusConflict, commons.ApiResponse{
			Message: message,
		})
	}

	// prevent to reserve room for past dates.
	if request.CheckInDate.Before(time.Now()) && request.RequestType == dto.CreateReservation {
		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			Message: handler.Config.Translator.Localize(lang, message_keys.ImpossibleReservationLatDateError),
		})
	}

	if request.CheckInDate == nil {
		return c.JSON(http.StatusBadRequest,
			commons.ApiResponse{
				Message: handler.Config.Translator.Localize(lang, message_keys.CheckInDateEmptyError)})
	}
	if request.CheckOutDate == nil {
		return c.JSON(http.StatusBadRequest,
			commons.ApiResponse{
				Message: handler.Config.Translator.Localize(lang, message_keys.CheckOutDateEmptyError)})
	}

	// create new reservation request for requested room.
	result, err := handler.Service.CreateReservationRequest(tenantContext(c), &request)
	if err != nil {
		handler.Config.Logger.LogError(err.Error())
		return c.JSON(http.StatusConflict, commons.ApiResponse{
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, commons.ApiResponse{
		Data:    result,
		Message: handler.Config.Translator.Localize(getAcceptLanguage(c), message_keys.Created),
	})
}

// @Summary SetUp Reservation
// @Tags Reservation
// @Accept json
// @Param X-TenantID header int true "X-TenantID"
// @Produce json
// @Param  Reservation body  models.Reservation true "Reservation"
// @Success 200 {object} models.Reservation
// @Router /reservations/{id} [post]
func (handler *ReservationHandler) create(c echo.Context) error {

	reservation := models.Reservation{}
	if err := c.Bind(&reservation); err != nil {
		handler.Config.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}

	lang := getAcceptLanguage(c)
	user := currentUser(c)

	invalidReservationRequestKeyErr := handler.Config.Translator.Localize(lang, message_keys.InvalidReservationRequestKey)
	if strings.TrimSpace(reservation.RequestKey) == "" {
		return c.JSON(http.StatusBadRequest,
			commons.ApiResponse{
				Message: invalidReservationRequestKeyErr,
			})
	}

	reservationRequest, err := handler.Service.FindReservationRequest(tenantContext(c), reservation.RequestKey)
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
				Message: handler.Config.Translator.Localize(lang, message_keys.EmptySharerError),
			})
	}

	hasReservationConflict, err := handler.Service.HasReservationConflict(tenantContext(c), reservation.CheckinDate, reservation.CheckoutDate, reservation.RoomId)
	if err != nil {
		handler.Config.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}

	if hasReservationConflict {
		return c.JSON(http.StatusBadRequest,
			commons.ApiResponse{
				Message: handler.Config.Translator.Localize(lang, message_keys.ReservationConflictError),
			})
	}
	handler.setReservationFields(&reservation, reservationRequest)
	reservation.SetAudit(user)
	// create new reservation.
	result, err := handler.Service.Create(tenantContext(c), &reservation)
	if err != nil {
		handler.Config.Logger.LogError(err.Error())
		return c.JSON(http.StatusConflict, commons.ApiResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, commons.ApiResponse{
		Data:    result,
		Message: handler.Config.Translator.Localize(getAcceptLanguage(c), message_keys.Created),
	})
}

// @Summary update Reservation
// @Tags Reservation
// @Accept json
// @Param X-TenantID header int true "X-TenantID"
// @Param Id path int true "Id"
// @Produce json
// @Param  Reservation body  models.Reservation true "Reservation"
// @Success 200 {object} models.Reservation
// @Router /reservations/{id} [put]
func (handler *ReservationHandler) update(c echo.Context) error {

	id, err := utils.ConvertToUint(c.Param("id"))
	lang := c.Request().Header.Get(acceptLanguage)
	user := currentUser(c)

	if err != nil {
		handler.Config.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, commons.ApiResponse{
			ResponseCode: http.StatusBadRequest,
		})
	}

	reservationModel, err := handler.Service.Find(tenantContext(c), id)
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
		handler.Config.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}

	invalidReservationRequestKeyErr := handler.Config.Translator.Localize(lang, message_keys.InvalidReservationRequestKey)
	if strings.TrimSpace(reservation.RequestKey) == "" {
		return c.JSON(http.StatusBadRequest,
			commons.ApiResponse{
				Message: invalidReservationRequestKeyErr,
			})
	}

	reservationRequest, err := handler.Service.FindReservationRequest(tenantContext(c), reservation.RequestKey)
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
				Message: handler.Config.Translator.Localize(lang, message_keys.EmptySharerError),
			})
	}

	hasReservationConflict, err := handler.Service.HasReservationConflict(tenantContext(c), reservation.CheckinDate, reservation.CheckoutDate, reservation.RoomId)
	if err != nil {
		handler.Config.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}

	if hasReservationConflict {
		return c.JSON(http.StatusBadRequest,
			commons.ApiResponse{
				Message: handler.Config.Translator.Localize(lang, message_keys.ReservationConflictError),
			})
	}
	reservation.SetUpdatedBy(user)
	handler.setReservationFields(&reservation, reservationRequest)
	// create new reservation.
	result, err := handler.Service.Update(tenantContext(c), id, &reservation)
	if err != nil {
		handler.Config.Logger.LogError(err.Error())
		return c.JSON(http.StatusConflict, commons.ApiResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, commons.ApiResponse{
		Data:    result,
		Message: handler.Config.Translator.Localize(getAcceptLanguage(c), message_keys.Created),
	})
}

// If the client cancels the reservation request, they can call this endpoint to delete the reservation request.

// @Summary Delete Reservation
// @Tags Reservation
// @Accept json
// @Param X-TenantID header int true "X-TenantID"
// @Param Id path int true "Id"
// @Produce json
// @Success 200
// @Router /reservations/cancel/{id} [delete]
func (handler *ReservationHandler) cancelRequest(c echo.Context) error {
	requestKey := c.QueryParam("requestKey")
	if err := handler.Service.RemoveReservationRequest(tenantContext(c), requestKey); err != nil {
		handler.Config.Logger.LogError(err.Error())
	}
	return c.JSON(http.StatusOK, nil)
}

// @Summary SetUp Reservation
// @Tags Reservation
// @Accept json
// @Param X-TenantID header int true "X-TenantID"
// @Produce json
// @Param  GetRatePriceDto body  dto.GetRatePriceDto true "GetRatePriceDto"
// @Success 200 {object} dto.RateCodePricesDto
// @Router /reservations/recommend-rate-codes/{id} [post]
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

	result, err := handler.Service.GetRecommendedRateCodes(tenantContext(c), &priceDto)
	if err != nil {
		handler.Config.Logger.LogError(err.Error())
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, commons.ApiResponse{
		Data: result,
	})
}

// @Summary Delete Reservation
// @Tags Reservation
// @Accept json
// @Param X-TenantID header int true "X-TenantID"
// @Param Id path int true "Id"
// @Produce json
// @Success 200 {array} models.Reservation
// @Router /reservations/{id} [get]
func (handler *ReservationHandler) find(c echo.Context) error {
	id, err := utils.ConvertToUint(c.Param("id"))

	if err != nil {
		handler.Config.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}

	result, err := handler.Service.Find(tenantContext(c), id)
	if err != nil {
		handler.Config.Logger.LogError(err.Error())
		return c.JSON(http.StatusInternalServerError, nil)
	}

	if result == nil {
		return c.JSON(http.StatusNotFound, nil)
	}

	return c.JSON(http.StatusOK, commons.ApiResponse{
		Data: result,
	})
}

// @Summary update Reservation
// @Tags Reservation
// @Accept json
// @Param X-TenantID header int true "X-TenantID"
// @Param status query int true "status"
// @Produce json
// @Param  Reservation body  models.Reservation true "Reservation"
// @Success 200 {object} models.Reservation
// @Router /reservations/change-status/{id} [put]
func (handler *ReservationHandler) changeStatus(c echo.Context) error {
	id, err := utils.ConvertToUint(c.Param("id"))

	if err != nil {
		handler.Config.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}

	reservation, err := handler.Service.Find(tenantContext(c), id)
	if err != nil {
		handler.Config.Logger.LogError(err.Error())
		return c.JSON(http.StatusInternalServerError, nil)
	}

	if reservation == nil {
		return c.JSON(http.StatusNotFound, nil)
	}

	statusVal, err := utils.ConvertToUint(c.QueryParam("status"))
	status := models.ReservationCheckStatus(statusVal)

	if err != nil {
		handler.Config.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}

	if status == models.CheckIn || status == models.Checkout {
		reservation.CheckStatus = status
		_, err := handler.Service.ChangeStatus(tenantContext(c), id, status)

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

func (handler *ReservationHandler) findAll(c echo.Context) error {
	paginationInput := c.Get(paginationInput).(*dto.PaginationFilter)
	output := getOutputQueryParamVal(c)

	filter := dto.ReservationFilter{}
	filter.PaginationFilter = *paginationInput
	filter.IgnorePagination = output != ""

	if err := c.Bind(&filter); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	result, err := handler.Service.FindAll(tenantContext(c), &filter)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	if output != "" {
		if output == EXCEL {
			report, err := handler.ReportService.ExportToExcel(result, getAcceptLanguage(c))
			if err != nil {
				handler.Config.Logger.LogError(err.Error())
				return c.JSON(http.StatusInternalServerError, commons.ApiResponse{})
			}
			setBinaryHeaders(c, "reservations", EXCEL_OUTPUT)
			c.Response().Write(report)
			return nil
		}
	}

	return c.JSON(http.StatusOK, commons.ApiResponse{
		Data:         result,
		ResponseCode: http.StatusOK,
		Message:      "",
	})
}
