package string_utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertToUint(t *testing.T) {
	t.Run("test_can_convert_interface_to_uint", func(t *testing.T) {

		numbers := []string{"1", "2", "3", "4"}

		for _, number := range numbers {

			result, err := ConvertToUint(number)
			assert.Nil(t, err)
			assert.True(t, result > 0)
			assert.NotNil(t, result)

			numberStr1 := fmt.Sprintf("%s", number)
			numberStr2 := fmt.Sprintf("%d", result)
			assert.Equal(t, numberStr1, numberStr2)
		}
	})

}
