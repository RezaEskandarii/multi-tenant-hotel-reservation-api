package middlewares

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"reservation-api/internal/models"
	"reservation-api/internal/services"
)

func AuditMiddleware(userService *services.UserService, auditService services.AuditService, ch chan interface{}) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user, err := userService.FindByUsername(c.Get("username").(string))
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "")
			}
			data := <-ch
			if data != nil {
				auditService.Save(&models.Audit{
					UserId:     user.Id,
					HttpMethod: c.Request().Method,
					Path:       c.Path(),
					Data:       fmt.Sprintf("%v", data),
				})
			}
			return next(c)
		}
	}
}
