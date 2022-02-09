package middlewares

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"reservation-api/internal/services/domain_services"
	"strings"
)

func JWTAuthMiddleware(s *domain_services.UserService) echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			authHeader := strings.Split(c.Request().Header.Get("Authorization"), "Bearer ")
			if len(authHeader) != 2 {
				return echo.NewHTTPError(http.StatusBadRequest, "invalid token")
			} else {
				jwtToken := authHeader[1]

				token, _ := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {

					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("uexpected signing method: %v", token.Header["alg"])
					}

					secretkey, _ := os.LookupEnv("JWT_KEY")
					return []byte(secretkey), nil
				})

				if token == nil {
					return echo.NewHTTPError(http.StatusBadRequest, "")
				}

				if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

					username := claims["username"]
					c.Set("claims", claims)
					c.Set("username", username)
					user, _ := s.FindByUsername(fmt.Sprintf("%s", username))

					if user == nil {
						return echo.NewHTTPError(http.StatusUnauthorized, "invalid user")
					}

					return next(c)
				} else {
					return echo.NewHTTPError(http.StatusUnauthorized, "")
				}
			}
		}
	}

}
