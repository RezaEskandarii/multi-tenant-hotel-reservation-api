package utils

import (
	"errors"
	"reflect"
)

var (
	NotStructErr = errors.New("given model or return model is not type of struct")
)

// ConvertToInterfaceSlice converts given input to slice of interface.
func ConvertToInterfaceSlice(input interface{}) ([]interface{}, error) {

	value := reflect.ValueOf(input)
	if value.Kind() != reflect.Slice {
		panic("InterfaceSlice() given a non-input type")
	}
	if value.IsNil() {
		return nil, errors.New("input is not slice")
	}

	result := make([]interface{}, value.Len())

	for i := 0; i < value.Len(); i++ {
		result[i] = value.Index(i).Interface()
	}

	return result, nil
}
