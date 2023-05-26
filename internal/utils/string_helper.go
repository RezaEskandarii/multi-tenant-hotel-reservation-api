package utils

import (
	"encoding/json"
	"fmt"
	"strings"
)

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

	if model == nil {
		return nil
	}

	result, err := json.Marshal(&model)

	if err == nil {
		return result
	}
	return nil
}
