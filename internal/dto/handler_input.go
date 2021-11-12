package dto

import (
	"github.com/labstack/echo/v4"
	"reservation-api/pkg/applogger"
	"reservation-api/pkg/translator"
)

// HandlerInput contains shared handler dependencies.
type HandlerInput struct {
	Router       *echo.Group
	Translator   *translator.Translator
	Logger       applogger.Logger
	AuditChannel chan interface{}
}
