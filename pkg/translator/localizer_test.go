package translator

import (
	"github.com/stretchr/testify/assert"
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

	translationService := New()
	assert.NotNil(t, translationService)

	for _, item := range testCases {
		result := translationService.Localize("en", item.input)
		assert.NotNil(t, result)
		assert.NotEqual(t, result, "")
		assert.Equal(t, result, item.want)
	}

}
