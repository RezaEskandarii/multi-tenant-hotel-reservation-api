package handlers

import (
	"github.com/labstack/echo/v4"
	"hotel-reservation/internal/commons"
	"hotel-reservation/internal/dto"
	"hotel-reservation/internal/middlewares"
	"hotel-reservation/internal/models"
	"hotel-reservation/internal/services"
	"hotel-reservation/internal/utils"
	"net/http"
)

// ProvinceHandler Province endpoint handler
type ProvinceHandler struct {
	Router  *echo.Group
	Service services.ProvinceService
}

func (handler *ProvinceHandler) Register(router *echo.Group, service services.ProvinceService) {
	handler.Router = router
	handler.Service = service
	handler.Router.POST("", handler.create)
	handler.Router.PUT("/:id", handler.update)
	handler.Router.GET("/:id", handler.find)
	handler.Router.GET("/:id/cities", handler.cities)
	handler.Router.GET("", handler.findAll, middlewares.PaginationMiddleware)
}

func (handler *ProvinceHandler) create(c echo.Context) error {

	model := &models.Province{}

	if err := c.Bind(&model); err != nil {
		return c.JSON(http.StatusBadRequest,
			commons.ApiResponse{
				Data:         nil,
				ResponseCode: http.StatusBadRequest,
				Message:      "BadRequest",
			})
	}

	if _, err := handler.Service.Create(model); err == nil {
		return c.JSON(http.StatusBadRequest,
			commons.ApiResponse{
				Data:         model,
				ResponseCode: http.StatusOK,
				Message:      "Ok",
			})
	}

	return c.JSON(http.StatusInternalServerError,
		commons.ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusInternalServerError,
			Message:      "BadRequest",
		})

}

func (handler *ProvinceHandler) update(c echo.Context) error {

	id, err := utils.ConvertToUint(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	model, err := handler.Service.Find(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, commons.ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusInternalServerError,
			Message:      "server error",
		})
	}

	if model == nil {
		return c.JSON(http.StatusNotFound, commons.ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusNotFound,
			Message:      "Not Found",
		})
	}

	if err := c.Bind(&model); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	if output, err := handler.Service.Update(model); err == nil {
		return c.JSON(http.StatusOK, commons.ApiResponse{
			Data:         output,
			ResponseCode: http.StatusOK,
			Message:      "Successfully updated",
		})
	} else {
		return c.JSON(http.StatusInternalServerError, nil)
	}
}

func (handler *ProvinceHandler) find(c echo.Context) error {
	id, err := utils.ConvertToUint(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	model, err := handler.Service.Find(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, commons.ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusInternalServerError,
			Message:      "server error",
		})
	}

	if model == nil {
		return c.JSON(http.StatusNotFound, commons.ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusNotFound,
			Message:      "Not Found",
		})
	}

	return c.JSON(http.StatusOK, commons.ApiResponse{
		Data:         model,
		ResponseCode: http.StatusOK,
		Message:      "",
	})
}

func (handler *ProvinceHandler) findAll(c echo.Context) error {

	paginationInput := c.Get(paginationInput).(*dto.PaginationInput)

	list, err := handler.Service.FindAll(paginationInput)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, commons.ApiResponse{
		Data:         list,
		ResponseCode: http.StatusOK,
		Message:      "",
	})
}

func (handler *ProvinceHandler) cities(c echo.Context) error {

	id, err := utils.ConvertToUint(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	cities, err := handler.Service.GetCities(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, commons.ApiResponse{
			Data:         nil,
			ResponseCode: http.StatusInternalServerError,
			Message:      "server error",
		})
	}

	return c.JSON(http.StatusOK, commons.ApiResponse{
		Data:         cities,
		ResponseCode: http.StatusOK,
		Message:      "",
	})
}
