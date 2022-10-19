package middlewares

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"reservation-api/internal/config"
	"reservation-api/internal/services/domain_services"
	"reservation-api/internal/utils"
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

				tenantID, _ := utils.ConvertToUint(c.Get(config.TenantIDKey))
				if err, claims := s.VerifyToken(jwtToken, tenantID); err == nil && claims != nil {

					c.Set("user_claims", claims)
					c.Set(config.ClaimsKey, claims.Username)
					return next(c)
				} else {

					return echo.NewHTTPError(http.StatusUnauthorized, "")
				}
			}
		}
	}

}
