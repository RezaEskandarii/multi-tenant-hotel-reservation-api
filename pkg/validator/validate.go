// Package validator is package of validators and sanitizers for strings,
// structs and collections.
//
//**
package validator

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"strings"
)

// Validate the given struct by govalidator library,
// if there is an error, an error and a list of error messages are returned.
// Otherwise, nil is returned.
// Example of errors:
// "password: non zero value required",
// "username: non zero value required"
func Validate(s interface{}) (error, []string) {

	// validate given struct
	ok, err := govalidator.ValidateStruct(s)
	if !ok && err != nil {

		// split validations by ';'
		// Because govalidator returns the errors in a single string separated by ';'.
		validationErrors := strings.Split(fmt.Sprintf("%s", err.Error()), ";")

		// list of errors
		errorList := make([]string, len(validationErrors))

		for i, m := range validationErrors {
			errorList[i] = m
		}

		return err, errorList
	}

	return nil, nil
}
