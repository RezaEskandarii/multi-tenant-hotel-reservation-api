package middlewares

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"reservation-api/internal/models"
	"reservation-api/internal/services/domain_services"
	"sync"
	"time"
)

func AuditMiddleware(userService *domain_services.UserService, auditService *domain_services.AuditService, ch chan interface{}) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if c.Request().Method == http.MethodPost || c.Request().Method == http.MethodPut {
				user, err := userService.FindByUsername(c.Get("username").(string))
				if err != nil {
					return echo.NewHTTPError(http.StatusUnauthorized, "")
				}
				once := sync.Once{}
				once.Do(func() {
					go listenToAuditChannel(auditService, user, c, ch)
				})
			}

			return next(c)
		}
	}
}

func listenToAuditChannel(service *domain_services.AuditService, user *models.User, c echo.Context, ch chan interface{}) {

	for {
		time.Sleep(1000 * time.Millisecond)
		select {
		case v := <-ch:
			data, _ := json.Marshal(v)
			service.Save(&models.Audit{
				UserId:     user.Id,
				Username:   user.Username,
				HttpMethod: c.Request().Method,
				Url:        c.Request().URL.Path,
				Data:       fmt.Sprintf("%s", data),
			})

			return
		}
	}

}
