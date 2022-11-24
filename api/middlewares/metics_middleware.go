package middlewares

import (
	"github.com/labstack/echo/v4"
	"reservation-api/internal/commons/metrics"
)

func MetricsMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := metrics.Set(c.Request().Context())
		metrics.AddGoroutines(ctx)
		metrics.AddRequests(ctx)
		return next(c)
	}
}
