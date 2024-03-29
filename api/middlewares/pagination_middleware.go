package middlewares

import (
	"github.com/labstack/echo/v4"
	"reservation-api/internal/dto"
	"strconv"
)

// PaginationMiddleware set pagination global object
func PaginationMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		page := 1
		pageVar, err := strconv.ParseUint(c.QueryParam("page"), 10, 64)
		if err == nil {
			page = int(pageVar)
		}

		var perPage = 20
		pageSize, err := strconv.ParseUint(c.QueryParam("page_size"), 10, 64)

		if err == nil {
			perPage = int(pageSize)
		}

		input := dto.NewPaginatedInput(page, perPage)

		//tenantID, _ := utils.ConvertToUint(c.Get(config.TenantIDKey))
		//input.TenantID = tenantID

		c.Set("paginationInput", input)

		return next(c)
	}
}
