package middlewares

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"strings"
)

func JWTAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		authHeader := strings.Split(c.Request().Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			return c.JSON(http.StatusBadRequest, "invalid token")
		} else {
			jwtToken := authHeader[1]
			token, _ := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {

				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}

				SECRETKEY, _ := os.LookupEnv("JWT_KEY")
				return []byte(SECRETKEY), nil
			})

			if token == nil {
				return c.JSON(http.StatusBadRequest, "")
			}

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				c.Set("claims", claims)
				return next(c)
			} else {
				return c.JSON(http.StatusUnauthorized, "")
			}
		}
	}
}
