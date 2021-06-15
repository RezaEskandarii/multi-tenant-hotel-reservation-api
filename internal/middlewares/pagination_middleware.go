package middlewares

import (
	"github.com/labstack/echo/v4"
	"hotel-reservation/internal/dto"
	"hotel-reservation/internal/utils"
)

func PaginationMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		page := 1
		pageVar, err := utils.ConvertToUint(c.QueryParam("page"))
		if err == nil {
			page = int(pageVar)
		}
		var perPage int
		perPageVar, err := utils.ConvertToUint(c.QueryParam("perPage"))

		if err == nil {
			perPage = int(perPageVar)
		}

		input := dto.NewPaginatedInput(page, perPage)

		c.Set("paginationInput", input)

		return next(c)
	}
}