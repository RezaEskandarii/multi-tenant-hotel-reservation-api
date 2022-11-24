package middlewares

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
)

func PanicRecoveryMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		defer func() {
			if r := recover(); r != nil {
				response := c.Response()
				responseBody, _ := json.Marshal(map[string]string{
					"error": "There was an internal server error",
				})
				response.Header().Set("Content-Type", "application/json")
				response.WriteHeader(http.StatusInternalServerError)
				response.Write(responseBody)
			}
		}()

		return next(c)
	}
}
