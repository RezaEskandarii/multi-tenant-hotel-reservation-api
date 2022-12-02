// Package handlers
// handles all http requests
///**/
package handlers

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"reflect"
	"reservation-api/internal/global_variables"
	"strings"
	"time"
)

// setBinaryHeaders method uses to sent file to client,
// where the headers for writing the file are set
// The set headers are:
// *MIMEOctetStream*: A media type (also known as a Multipurpose Internet Mail Extensions or MIME type) indicates the nature and format of a document,
// file, or assortment of bytes. MIME types are defined and standardized in IETF's RFC 6838
// *HeaderContentDisposition*: The Content-Disposition header is defined in the larger context of MIME messages for e-mail,
// but only a subset of the possible parameters apply to HTTP forms and POST requests.
// Only the value form-data, as well as the optional directive name and filename, can be used in the HTTP context.
// *Content-Transfer-Encoding*: Each MIME part may contain a header that specifies whether the part was processed
// for transfer and how the body of the message part is currently represented. The field name of this header is Content-Transfer-Encoding.
// *Expires*: The Expires HTTP header contains the date/time after which the response is considered expired.
// Invalid expiration dates with value 0 represent a date in the past and mean that the resource is already expired.
func setBinaryHeaders(context echo.Context, fileName string, format string) {

	fileNameStr := fmt.Sprintf("report-%s-%s.%s", fileName, time.Now().Format("2006-01-02"), format)
	context.Response().Header().Set(echo.HeaderContentType, echo.MIMEOctetStream)
	context.Response().Header().Set(echo.HeaderContentDisposition, "attachment; filename="+fileNameStr)
	context.Response().Header().Set("Content-Transfer-Encoding", "binary")
	context.Response().Header().Set("Expires", "10000000")
	context.Response().WriteHeader(http.StatusOK)

}

// tenantContext function reads the tenant from the echo context and returns it.
// The tenant context is set in Tenant middleware.
func tenantContext(c echo.Context) context.Context {

	var tenantCtx = c.Get(global_variables.TenantIDCtx)
	if tenantCtx == nil {
		panic("tenant context in nil in function tenantContext")
	}
	return c.Get(global_variables.TenantIDCtx).(context.Context)
}

// returns authenticated user's username from echo context
// the authenticated user's username set in user middleware.
func currentUser(c echo.Context) string {
	return fmt.Sprintf("%s", c.Get(global_variables.ClaimsKey))
}

// getOutputQueryParamVal returns query param with "output" key to generate pdf or excel outputs.
func getOutputQueryParamVal(c echo.Context) string {
	return strings.TrimSpace(c.QueryParam("output"))
}

// setCreatedByUpdatedBy fills CreatedBy and UpdatedBy fields.
func setCreatedByUpdatedBy(entity interface{}, audit string) {
	val := reflect.Indirect(reflect.ValueOf(entity))
	val.FieldByName("CreatedBy").SetString(audit)
	val.FieldByName("UpdatedBy").SetString(audit)
}
