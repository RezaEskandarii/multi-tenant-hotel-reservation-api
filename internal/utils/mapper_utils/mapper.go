package mapper_utils

import (
	"encoding/json"
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

func Map[T any](source interface{}, dest T) T {

	var data []byte
	var err error

	if reflect.ValueOf(source).Kind() == reflect.Pointer {
		data, err = json.Marshal(source)
	} else {
		data, err = json.Marshal(source)
	}

	if err != nil {
		panic(err.Error())
	}

	if reflect.ValueOf(dest).Kind() == reflect.Pointer {
		err = json.Unmarshal(data, dest)
	} else {
		err = json.Unmarshal(data, &dest)
	}

	if err != nil {
		panic(err.Error())
	}

	return dest
}
