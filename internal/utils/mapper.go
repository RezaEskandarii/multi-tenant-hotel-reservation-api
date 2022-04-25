package utils

import (
	"errors"
	"fmt"
	"reflect"
)

var (
	NotStructErr = errors.New("given model or return model is not type of struct")
)

func Map(givenModel interface{}, returnModel interface{}) (interface{}, error) {

	if reflect.ValueOf(givenModel).Kind() == reflect.Struct && reflect.ValueOf(returnModel).Kind() == reflect.Struct {

		returnModelVal := reflect.ValueOf(returnModel)
		givenModelVal := reflect.ValueOf(givenModel)

		for i := 0; i < returnModelVal.NumField(); i++ {

			for j := 0; j < givenModelVal.NumField(); j++ {
				if returnModelVal.Type().Field(i).Name == givenModelVal.Type().Field(j).Name {

					valueFieldGiven := givenModelVal.Field(j)
					typeFieldGiven := givenModelVal.Type().Field(j)

					//valueFieldReturn := givenModelVal.Field()
					typeFieldReturn := givenModelVal.Type().Field(j)

					f := valueFieldGiven.Interface()
					val := reflect.ValueOf(f)
					fmt.Println(typeFieldGiven.Type)
					fmt.Println(val)

					//fieldName := returnModelVal.Type().Field(i).Name

					if typeFieldReturn.Type == typeFieldGiven.Type {
						val.Field(i).Set(val)
					}

				}
			}
		}

	} else {
		return nil, NotStructErr
	}

	return returnModel, nil
}

// ConvertToInterfaceSlice converts given input to slice of interface.
func ConvertToInterfaceSlice(input interface{}) (error, []interface{}) {

	value := reflect.ValueOf(input)
	if value.Kind() != reflect.Slice {
		panic("InterfaceSlice() given a non-input type")
	}
	if value.IsNil() {
		return errors.New("input is not slice"), nil
	}

	result := make([]interface{}, value.Len())

	for i := 0; i < value.Len(); i++ {
		result[i] = value.Index(i).Interface()
	}

	return nil, result
}
