package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
)

// FileExists check if given file exists or no.
func FileExists(filename string) bool {
	if filename == "" {
		return false
	}
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// CastJsonFileToStruct Takes the address of the json file and  pointer of struct
// And Marshall json file data int structModel.
func CastJsonFileToStruct(path string, structModel interface{}) error {
	if path == "" {
		return errors.New("file path is empty")
	}
	if structModel == nil {
		return errors.New("struct pointer is nil")
	}
	if reflect.ValueOf(structModel).Type().Kind() != reflect.Ptr {
		return errors.New("given structModel type is not pointer")
	}
	if FileExists(path) {
		file, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		} else {
			return json.Unmarshal(file, structModel)
		}
	} else {
		return errors.New(fmt.Sprintf("%s file not found", path))
	}
}
