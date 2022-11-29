// Package translator localize messages /**/
package translator

import (
	"context"
	"github.com/stretchr/testify/assert"
	"reservation-api/internal/global_variables"
	"testing"
)

type TranslationTestCase struct {
	input string
	want  string
}

func TestCanLocalizeMessage(t *testing.T) {

	testCases := []TranslationTestCase{
		{
			input: "AppName",
			want:  "reservation-api",
		},
		{
			input: "errors.BadRequest",
			want:  "BadRequest",
		},
	}

	ctx := context.WithValue(context.Background(), global_variables.CurrentLang, "en")

	for _, item := range testCases {
		result := Localize(ctx, item.input)
		assert.NotNil(t, result)
		assert.NotEqual(t, result, "")
		assert.Equal(t, result, item.want)
	}

}
