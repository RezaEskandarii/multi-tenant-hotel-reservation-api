package middlewares

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"reservation-api/internal/commons"
	"reservation-api/pkg/applogger"
)

func LoggerMiddleware(logger applogger.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {

		return func(c echo.Context) error {

			var trace = commons.RequestTrace{
				TraceID:   uuid.NewString(),
				IPAddress: c.RealIP(),
				EndPoint:  c.Request().RequestURI,
				Data:      nil,
			}

			c.Response().Before(func() {
				trace.Tag = "BeforeResponse"
				logger.LogInfoJSON(trace)
			})

			c.Response().After(func() {
				trace.Tag = "AfterResponse"
				trace.ResponseCode = c.Response().Status
				logger.LogInfoJSON(trace)
			})

			return next(c)
		}
	}
}
