package middlewares

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"reservation-api/internal/config"
	"reservation-api/internal/services/domain_services"
)

func TenantAccessMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		claims := c.Get("user_claims").(*domain_services.Claims)
		if claims.TenantID == c.Get(config.TenantIDKey) { ///utils.Decrypt(claims.TenantID) == utils.Decrypt(fmt.Sprintf("%s", c.Get(config.TenantIDKey))) {
			return next(c)
		}

		return echo.NewHTTPError(http.StatusBadRequest, "invalid tenant id")
	}
}
