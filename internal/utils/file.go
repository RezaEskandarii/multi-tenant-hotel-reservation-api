package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

// FileExists check if given file exists or no.
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func CastJsonFileToModel(path string, model interface{}) error {
	if FileExists(path) {
		file, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		} else {
			return json.Unmarshal(file, model)
		}
	} else {
		return errors.New(fmt.Sprintf("%s file not found", path))
	}
}
