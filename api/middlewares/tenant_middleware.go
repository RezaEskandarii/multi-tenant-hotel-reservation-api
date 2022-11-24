package middlewares

import (
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
	"reservation-api/internal/config"
	"strings"
)

func TenantMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		var tenant = strings.TrimSpace(c.Request().Header.Get("X-Tenant-ID"))

		if tenant == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "X-Tenant-ID header is null")
		}

		c.Set(config.TenantIDKey, tenant)
		ctx := context.WithValue(c.Request().Context(), "TenantID", tenant)
		c.Set(config.TenantIDCtx, ctx)

		return next(c)
	}
}
