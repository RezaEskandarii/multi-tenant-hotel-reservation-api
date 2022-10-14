package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"reservation-api/internal/config"
	"reservation-api/internal/utils"
	"strings"
	"time"
)

var (
	paginationInput = "paginationInput"
	acceptLanguage  = "Accept-Language" // accept language header.
	tenantId        = "X-TenantID"      // this value uses in http header request as a selected business id.

	EXCEL        = "excel"
	EXCEL_OUTPUT = "xlsx"
	PDF          = "pdf"
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

func getCurrentTenant(c echo.Context) uint64 {
	tenantID, err := utils.ConvertToUint(c.Get(config.TenantID))
	if err != nil {
		panic(err)
	}
	return tenantID
}

// returns authenticated user from http Context
func getCurrentUser(c echo.Context) string {
	return fmt.Sprintf("%s", c.Get("username"))
}

// getOutputQueryParamVal returns query param with "output" key to generate pdf or excel outputs.
func getOutputQueryParamVal(c echo.Context) string {
	return strings.TrimSpace(c.QueryParam("output"))
}

// writeBinaryHeaders
func writeBinaryHeaders(context echo.Context, fName string, format string) {
	fileName := fmt.Sprintf("report-%s-%s.%s", fName, time.Now().Format("2006-01-02"), format)
	context.Response().Header().Set(echo.HeaderContentType, echo.MIMEOctetStream)
	context.Response().Header().Set(echo.HeaderContentDisposition, "attachment; filename="+fileName)
	context.Response().Header().Set("Content-Transfer-Encoding", "binary")
	context.Response().Header().Set("Expires", "0")
	context.Response().WriteHeader(http.StatusOK)
}
