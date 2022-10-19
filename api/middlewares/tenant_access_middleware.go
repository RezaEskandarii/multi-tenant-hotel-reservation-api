package middlewares

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"reservation-api/internal/config"
	"reservation-api/internal/services/domain_services"
	"reservation-api/internal/utils"
)

func TenantAccessMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		claims := c.Get("user_claims").(*domain_services.Claims)
		fmt.Println(c.Get(config.TenantIDKey))
		if utils.Decrypt(claims.TenantID) == utils.Decrypt(fmt.Sprintf("%s", c.Get(config.TenantIDKey))) {
			return next(c)
		}

		return echo.NewHTTPError(http.StatusBadRequest, "invalid tenant id")
	}
}
