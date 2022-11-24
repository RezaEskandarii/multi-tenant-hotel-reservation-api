package handlers

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"reflect"
	"reservation-api/internal/config"
	"strings"
	"time"
)

var (
	paginationInput = "paginationInput"
	acceptLanguage  = "Accept-Language" // accept language header.
	tenantId        = "X-TenantIDKey"   // this value uses in http header request as a selected business id.
	EXCEL           = "excel"
	EXCEL_OUTPUT    = "xlsx"
	PDF             = "pdf"
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

// tenantContext function reads the tenant from the echo context and returns it
func tenantContext(c echo.Context) context.Context {
	return c.Get(config.TenantIDCtx).(context.Context)
}

// returns authenticated user from http Context
func getCurrentUser(c echo.Context) string {
	return fmt.Sprintf("%s", c.Get(config.ClaimsKey))
}

// getOutputQueryParamVal returns query param with "output" key to generate pdf or excel outputs.
func getOutputQueryParamVal(c echo.Context) string {
	return strings.TrimSpace(c.QueryParam("output"))
}

// writeBinaryHeaders method uses to sent file to client,
// where the headers for writing the file are set
//
func writeBinaryHeaders(context echo.Context, fName string, format string) {

	fileName := fmt.Sprintf("report-%s-%s.%s", fName, time.Now().Format("2006-01-02"), format)

	context.Response().Header().Set(echo.HeaderContentType, echo.MIMEOctetStream)
	context.Response().Header().Set(echo.HeaderContentDisposition, "attachment; filename="+fileName)
	context.Response().Header().Set("Content-Transfer-Encoding", "binary")
	context.Response().Header().Set("Expires", "0")
	context.Response().WriteHeader(http.StatusOK)

}

// setCreatedByUpdatedBy fills CreatedBy and UpdatedBy fields.
func setCreatedByUpdatedBy(entity interface{}, audit string) {
	val := reflect.Indirect(reflect.ValueOf(entity))
	val.FieldByName("CreatedBy").SetString(audit)
	val.FieldByName("UpdatedBy").SetString(audit)
}
