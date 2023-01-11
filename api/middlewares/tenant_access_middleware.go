package middlewares

import (
	"github.com/labstack/echo/v4"
)

func TenantAccessMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		return next(c)
		//claims := c.Get("user_claims").(*domain_services.Claims)
		//if utils.Decrypt(claims.TenantID) == utils.Decrypt(fmt.Sprintf("%s", c.Get(global_variables.TenantIDKey))) {
		//	return next(c)
		//}
		//
		//return echo.NewHTTPError(http.StatusBadRequest, "invalid tenant id")
	}
}
