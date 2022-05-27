package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
	"reservation-api/internal/services/domain_services"
	"reservation-api/internal/utils"
)

type PaymentHandler struct {
	Input          *dto.HandlersShared
	PaymentService *domain_services.PaymentService
}

func (handler *PaymentHandler) Register(input *dto.HandlersShared, service *domain_services.PaymentService) {
	handler.Input = input
	routeGroup := handler.Input.Router.Group("/payment")
	handler.PaymentService = service
	routeGroup.POST("", handler.create)
	routeGroup.DELETE("/:id", handler.delete)
}

func (handler *PaymentHandler) create(c echo.Context) error {

	payment := models.Payment{}
	if err := c.Bind(&payment); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	if result, err := handler.PaymentService.Create(&payment); err != nil {
		handler.Input.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	} else {
		return c.JSON(http.StatusOK, result)
	}
}

func (handler *PaymentHandler) delete(c echo.Context) error {

	id, _ := utils.ConvertToUint(c.Get("id"))
	if err := handler.PaymentService.Delete(id); err != nil {
		handler.Input.Logger.LogError(err.Error())
		return c.JSON(http.StatusBadRequest, nil)
	}
	return c.JSON(http.StatusOK, nil)

}
