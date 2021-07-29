package utils

import (
	"fmt"
	"reflect"
)

func Map(givenModel interface{}, returnModel interface{}) (interface{}, error) {

	returnModelVal := reflect.ValueOf(returnModel)
	givenModelVal := reflect.ValueOf(givenModel)
	for i := 0; i < returnModelVal.NumField(); i++ {

		for j := 0; j < givenModelVal.NumField(); j++ {
			if returnModelVal.Type().Field(i).Name == givenModelVal.Type().Field(j).Name {

				valueFieldGiven := givenModelVal.Field(j)
				typeFieldGiven := givenModelVal.Type().Field(j)

				//valueFieldReturn := givenModelVal.Field()
				//typeFieldReturn := givenModelVal.Type().Field(j)

				f := valueFieldGiven.Interface()
				val := reflect.ValueOf(f)
				fmt.Println(typeFieldGiven.Type)
				fmt.Println(val)

				//fieldName := returnModelVal.Type().Field(i).Name

				switch returnModelVal.Field(i).Kind() {
				case reflect.String:
					val.Field(i).SetString("strconv.Atoi(s[val.Field(i).Name])")
				}
			}
		}
	}

	return returnModel, nil
}
