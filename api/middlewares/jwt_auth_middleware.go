package middlewares

import (
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
	"reservation-api/internal/global_variables"
	"reservation-api/internal/services/domain_services"
	"reservation-api/internal/utils"
	"strings"
)

func JWTAuthMiddleware(s *domain_services.AuthService) echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			authHeader := strings.Split(c.Request().Header.Get("Authorization"), "Bearer ")
			if len(authHeader) != 2 {

				return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")

			} else {
				jwtToken := authHeader[1]

				if jwtToken == "" {
					return echo.NewHTTPError(http.StatusUnauthorized, "")
				}

				tenantID, _ := utils.ConvertToUint(c.Get(global_variables.TenantIDKey))
				if err, ok := s.VerifyToken(c.Get(global_variables.TenantIDCtx).(context.Context), jwtToken, tenantID); err == nil && ok {
					claims := s.ParseClaims(jwtToken)
					c.Set(global_variables.UserClaims, claims)
					c.Set(global_variables.ClaimsKey, claims.Username)
					return next(c)
				} else {

					return echo.NewHTTPError(http.StatusUnauthorized, "")
				}
			}
		}
	}

}
