package dto

import (
	"github.com/labstack/echo/v4"
	"reservation-api/pkg/applogger"
	"reservation-api/pkg/translator"
)

// HandlersShared contains shared handler dependencies.
type HandlersShared struct {
	Router     *echo.Group
	Translator *translator.Translator
	Logger     applogger.Logger
	///ReportService *common_services.ReportService
}
