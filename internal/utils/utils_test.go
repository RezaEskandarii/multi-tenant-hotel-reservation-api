package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestFileExists(t *testing.T) {
	assert.True(t, FileExists("./mapper.go"))
	assert.False(t, FileExists("fff00x.go"))
}

func TestGenerateSHA256(t *testing.T) {
	strToHash1 := "this sample text"
	strToHash2 := "this sample text ."

	result1 := GenerateSHA256(strToHash1)
	result2 := GenerateSHA256(strToHash1)
	result3 := GenerateSHA256(strToHash2)

	assert.NotNil(t, result1)
	assert.NotNil(t, result2)
	assert.Equal(t, result1, result2)
	assert.NotEqual(t, result1, result3)
	assert.NotEqual(t, strToHash1, result1)
	assert.NotEqual(t, strToHash1, result2)
	assert.NotEqual(t, strToHash2, result3)
}

func TestConvertToInterfaceSlice(t *testing.T) {
	type testCase struct {
		key   int
		value interface{}
	}

	testCases := []testCase{
		{
			key:   1,
			value: 10,
		},
		{
			key:   2,
			value: "test",
		}, {
			key:   3,
			value: "my test",
		},
	}

	result, err := ConvertToInterfaceSlice(testCases)
	assert.Nil(t, err)
	assert.Equal(t, len(result), len(testCases))

	testCaseStr := fmt.Sprintf("%v", testCases)
	resultStr := fmt.Sprintf("%v", result)

	assert.Equal(t, testCaseStr, resultStr)
	assert.True(t, strings.Contains(testCaseStr, "["))
}

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

	t.Run("test_can_return_error_for_invalid_strings", func(t *testing.T) {

		numbers := []string{"-1", "a", "b", "cc"}

		for _, number := range numbers {

			result, err := ConvertToUint(number)

			assert.NotNil(t, err)
			assert.Equal(t, uint64(0), result)
			assert.Error(t, err)
		}
	})
}

func TestToJson(t *testing.T) {

	type testCase struct {
		Code int
	}

	testCases := []struct {
		input  testCase
		result string
	}{
		{
			input:  testCase{Code: 1},
			result: `{"Code":1}`,
		},
		{
			input:  testCase{Code: 2},
			result: `{"Code":2}`,
		},
	}

	for _, item := range testCases {

		result := ToJson(item.input)
		assert.NotNil(t, result)
		assert.True(t, len(result) > 0)
		assert.Equal(t, string(result), item.result)
	}
}
