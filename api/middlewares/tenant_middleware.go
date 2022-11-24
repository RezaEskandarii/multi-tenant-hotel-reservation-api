package middlewares

import (
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
	"reservation-api/internal/config"
	"reservation-api/internal/utils"
	"strings"
)

func TenantMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		var tenantStr = strings.TrimSpace(c.Request().Header.Get("X-Tenant-ID"))

		if tenantStr == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "X-Tenant-ID header is null")
		}

		tenantID, _ := utils.ConvertToUint(tenantStr)

		c.Set(config.TenantIDKey, tenantStr)
		ctx := context.WithValue(c.Request().Context(), "TenantID", tenantID)
		c.Set(config.TenantIDCtx, ctx)

		return next(c)
	}
}
