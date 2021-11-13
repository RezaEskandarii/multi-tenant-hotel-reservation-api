package middlewares

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"reservation-api/internal/models"
	"reservation-api/internal/services"
	"time"
)

func AuditMiddleware(userService *services.UserService, auditService *services.AuditService, ch chan interface{}) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if c.Request().Method == http.MethodPost || c.Request().Method == http.MethodPut {
				user, err := userService.FindByUsername(c.Get("username").(string))

				if err != nil {
					return echo.NewHTTPError(http.StatusUnauthorized, "")
				}

				go func() {
					for {
						time.Sleep(1000 * time.Millisecond)
						select {
						case v := <-ch:
							data, _ := json.Marshal(v)
							auditService.Save(&models.Audit{
								UserId:     user.Id,
								HttpMethod: c.Request().Method,
								Path:       c.Request().URL.Path,
								Data:       fmt.Sprintf("%s", data),
							})
							return
						}
					}

				}()
			}

			return next(c)
		}
	}
}
