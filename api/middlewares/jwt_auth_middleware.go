package middlewares

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"reservation-api/internal/config"
	"reservation-api/internal/services/domain_services"
	"strings"
)

func JWTAuthMiddleware(s *domain_services.AuthService) echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			authHeader := strings.Split(c.Request().Header.Get("Authorization"), "Bearer ")
			if len(authHeader) != 2 {
				return echo.NewHTTPError(http.StatusBadRequest, "invalid token")
			} else {
				jwtToken := authHeader[1]

				if jwtToken == "" {
					return echo.NewHTTPError(http.StatusUnauthorized, "")
				}
				if err, claims := s.VerifyToken(jwtToken, c.Get(config.TenantID).(uint64)); err == nil && claims != nil {

					c.Set("claims", claims.Username)
					return next(c)
				} else {

					return echo.NewHTTPError(http.StatusUnauthorized, "")
				}
			}
		}
	}

}
