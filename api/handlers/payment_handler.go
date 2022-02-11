package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/services/domain_services"
)

type PaymentHandler struct {
	Input          *dto.HandlerInput
	PaymentService *domain_services.PaymentService
}

func (handler *PaymentHandler) Register(input *dto.HandlerInput, service *domain_services.PaymentService) {
	handler.Input = input
	routeGroup := handler.Input.Router.Group("/payment")
	handler.PaymentService = service
	routeGroup.POST("", handler.pay)
}

func (handler *PaymentHandler) pay(c echo.Context) error {

	b := handler.PaymentService.Pay()

	return c.JSON(http.StatusBadRequest, commons.ApiResponse{
		ResponseCode: http.StatusBadRequest,
		Message:      "",
		Data:         b,
	})
}
