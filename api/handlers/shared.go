package handlers

import "github.com/labstack/echo/v4"

var (
	paginationInput = "paginationInput"
	acceptLanguage  = "Accept-Language" // accept language header.
	tenantId        = "X-TenantID"      // this value uses in http header request as a selected business id.
)

// getAcceptLanguage returns Accept-Language header value.
// If the client returns the Accept-Language header, it returns it; otherwise, it returns en by default.
func getAcceptLanguage(c echo.Context) string {
	lang := c.Request().Header.Get(acceptLanguage)
	if lang == "" {
		lang = "en"
	}
	return lang
}
