package middlewares

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"reservation-api/pkg/applogger"
)

func PanicRecoveryMiddleware(logger applogger.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			defer func() {
				if r := recover(); r != nil {

					logger.LogDebug(r)
					response := c.Response()
					responseBody, _ := json.Marshal(map[string]string{
						"error": "There was an internal server error",
					})
					response.Header().Set("Content-Type", "application/json")
					response.WriteHeader(http.StatusInternalServerError)
					response.Write(responseBody)

					fmt.Println(r)
				}
			}()

			return next(c)
		}
	}
}
