package httphelper

import (
	"context"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strings"
	"testing"
)

func TestCanSendHttpRequest(t *testing.T) {

	requestHelper := New()
	ctx := context.Background()

	t.Run("test_can_send_get_method", func(t *testing.T) {

		err, bodyDump, resp := requestHelper.Get(Request{Ctx: ctx, Url: "https://google.com"})
		assert.Nil(t, err)
		assert.NotNil(t, bodyDump)
		assert.NotNil(t, resp)
		assert.Equal(t, resp.StatusCode, http.StatusOK)
		assert.True(t, strings.Contains(bodyDump, "google"))
	})
}
