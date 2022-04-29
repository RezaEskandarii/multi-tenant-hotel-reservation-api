package dto

import (
	"github.com/labstack/echo/v4"
	"reservation-api/pkg/applogger"
	"reservation-api/pkg/translator"
)

// HandlersSharedObjects contains shared handler dependencies.
type HandlersSharedObjects struct {
	Router     *echo.Group
	Translator *translator.Translator
	Logger     applogger.Logger
}
