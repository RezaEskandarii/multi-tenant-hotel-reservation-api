package dto

import (
	"github.com/labstack/echo/v4"
	"hotel-reservation/pkg/applogger"
	"hotel-reservation/pkg/translator"
)

// HandlerInput contains shared handler dependencies.
type HandlerInput struct {
	Router     *echo.Group
	Translator *translator.Translator
	Logger     applogger.Logger
}
