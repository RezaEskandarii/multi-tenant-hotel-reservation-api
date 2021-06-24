package handlers

import (
	"github.com/labstack/echo/v4"
	. "hotel-reservation/internal/commons"
	"hotel-reservation/internal/dto"
	"hotel-reservation/internal/middlewares"
	"hotel-reservation/internal/models"
	"hotel-reservation/internal/services"
	"hotel-reservation/internal/utils"
	"net/http"
)

// CurrencyHandler Currency endpoint handler
type CurrencyHandler struct {
	Router  *echo.Group
	Service services.CurrencyService
}

func (handler *CurrencyHandler) Register(router *echo.Group, service services.CurrencyService) {
	handler.Router = router
	handler.Service = service
	handler.Router.POST("", handler.create)
	handler.Router.PUT("/:id", handler.update)
	handler.Router.GET("/:id", handler.find)
	handler.Router.GET("", handler.findAll, middlewares.PaginationMiddleware)
}

func (handler *CurrencyHandler) create(c echo.Context) error {

	model := &models.Currency{}

	if err := c.Bind(&model); err != nil {
		return c.JSON(http.StatusBadRequest,
			ApiResponse{
				Data:         nil,
				ResponseCode: BadRequest,
				Message:      "BadRequest",
			})
	}

	if _, err := handler.Service.Create(model); err == nil {
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

func (handler *CurrencyHandler) update(c echo.Context) error {

	id, err := utils.ConvertToUint(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	model, err := handler.Service.Find(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, ApiResponse{
			Data:         nil,
			ResponseCode: InternalServerError,
			Message:      "server error",
		})
	}

	if model == nil {
		return c.JSON(http.StatusNotFound, ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusNotFound,
			Message:      "Not Found",
		})
	}

	if err := c.Bind(&model); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	if output, err := handler.Service.Update(model); err == nil {
		return c.JSON(http.StatusOK, ApiResponse{
			Data:         output,
			ResponseCode: http.StatusOK,
			Message:      "Successfully updated",
		})
	} else {
		return c.JSON(http.StatusInternalServerError, nil)
	}
}

func (handler *CurrencyHandler) find(c echo.Context) error {
	id, err := utils.ConvertToUint(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	model, err := handler.Service.Find(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, ApiResponse{
			Data:         nil,
			ResponseCode: InternalServerError,
			Message:      "server error",
		})
	}

	if model == nil {
		return c.JSON(http.StatusNotFound, ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusNotFound,
			Message:      "Not Found",
		})
	}

	return c.JSON(http.StatusOK, ApiResponse{
		Data:         model,
		ResponseCode: http.StatusOK,
		Message:      "",
	})
}

func (handler *CurrencyHandler) findAll(c echo.Context) error {

	paginationInput := c.Get(paginationInput).(*dto.PaginationInput)

	list, err := handler.Service.FindAll(paginationInput)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, ApiResponse{
		Data:         list,
		ResponseCode: http.StatusOK,
		Message:      "",
	})
}
