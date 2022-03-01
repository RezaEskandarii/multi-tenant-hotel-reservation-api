package utils

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// ConvertToUint gives interface as a input and converts interface to uint.
func ConvertToUint(input interface{}) (uint64, error) {

	if input == nil {
		return 0, nil
	}

	switch input.(type) {
	case string:
		returnValue, err := strconv.ParseUint(input.(string), 10, 64)
		if err != nil {
			return 0, err
		}

		return returnValue, nil
	}

	return 0, nil
}

// GenerateCacheKey returns string to use as a cache key.
func GenerateCacheKey(keys ...interface{}) string {
	strBuilder := strings.Builder{}
	for _, str := range keys {
		str := fmt.Sprintf("%v", str)
		strBuilder.Write([]byte(str))
	}
	return GenerateSHA256(strBuilder.String())
}

// ToJson converts given model to json and returns as a byte array.
func ToJson(model interface{}) []byte {
	result, err := json.Marshal(&model)

	if err == nil {
		return result
	}
	return nil
}
