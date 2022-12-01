package handlers

import (
	"github.com/labstack/echo/v4"
	"reservation-api/pkg/applogger"
)

type handlerBase struct {
	Router *echo.Group
	Logger applogger.Logger
}
