package dto

import (
	"github.com/labstack/echo/v4"
	"reservation-api/pkg/applogger"
)

// HandlerConfig contains shared handler dependencies.
type HandlerConfig struct {
	Router *echo.Group
	Logger applogger.Logger
	///ReportService *common_services.ReportService
}
