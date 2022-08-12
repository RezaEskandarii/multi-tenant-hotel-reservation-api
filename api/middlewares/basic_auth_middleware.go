package middlewares

import (
	"crypto/subtle"
	"github.com/labstack/echo/v4"
	"os"
)

func BasicAuth(username, password string, c echo.Context) (bool, error) {
	user, _ := os.LookupEnv("basic_auth_user")
	pass, _ := os.LookupEnv("basic_auth_pass")

	if subtle.ConstantTimeCompare([]byte(username), []byte(user)) == 1 &&
		subtle.ConstantTimeCompare([]byte(password), []byte(pass)) == 1 {
		return true, nil
	}
	return false, nil
}
