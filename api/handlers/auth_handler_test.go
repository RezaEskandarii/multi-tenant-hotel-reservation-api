package handlers

import (
	"context"
	"github.com/stretchr/testify/assert"
	"net/http"
	"reservation-api/internal/commons"
	"reservation-api/internal/commons/httphelper"
	"testing"
	"time"
)

var (
	httpReq = httphelper.New()
	baseUrl = "http://127.0.0.1:8080/api/v1"
)

func init() {
	setup()
}

func setup() {

	time.Sleep(10000)

}

func getHeaders() map[string]string {

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["X-Tenant-ID"] = "1"
	return headers
}

func TestAuthHandler(t *testing.T) {

	t.Run("test_can_signin", func(t *testing.T) {

		body := make(map[string]interface{})

		body["username"] = "reza"
		body["password"] = "Reza1234"

		o := commons.JWTTokenResponse{}

		err, _, resp := httpReq.Post(httphelper.Request{
			Ctx:     context.Background(),
			Url:     baseUrl + "/auth/signin",
			Body:    body,
			Out:     &o,
			Headers: getHeaders(),
		})

		assert.Nil(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, resp.StatusCode, http.StatusOK)
		assert.NotEqual(t, o.AccessToken, "")
		assert.True(t, len(o.AccessToken) > 0)
		assert.True(t, o.ExpireAt.After(time.Now()))

	})
}
