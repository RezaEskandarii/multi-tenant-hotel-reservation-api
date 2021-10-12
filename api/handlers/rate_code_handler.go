package handlers

import (
	"github.com/labstack/echo/v4"
	"hotel-reservation/internal/dto"
	"hotel-reservation/internal/services"
	"hotel-reservation/pkg/applogger"
	"hotel-reservation/pkg/translator"
)

// RateCodeHandler RateCode endpoint handler
type RateCodeHandler struct {
	Router     *echo.Group
	Service    *services.RateCodeService
	translator *translator.Translator
	logger     applogger.Logger
}

func (handler *RateCodeHandler) Register(input *dto.HandlerInput, service *services.RateCodeService) {
	handler.Router = input.Router
	handler.Service = service
	handler.translator = input.Translator
	handler.logger = input.Logger

	//routeCode := handler.Router.Group("/rate-Codes")

}
