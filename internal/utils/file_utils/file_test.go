package file_utils

import (
	"reflect"
	"testing"
)

func TestIsFileExists(t *testing.T) {
	// Test with an existing file
	existingFile := "./file.go"
	if !FileExists(existingFile) {
		t.Errorf("FileExists(%s) returned false, expected true", existingFile)
	}

	// Test with a non-existing file
	nonExistingFile := "testdata/non-existing.txt"
	if FileExists(nonExistingFile) {
		t.Errorf("FileExists(%s) returned true, expected false", nonExistingFile)
	}

	// Test with an empty filename
	emptyFilename := ""
	if FileExists(emptyFilename) {
		t.Errorf("FileExists(%s) returned true, expected false", emptyFilename)
	}
}

func TestCastJsonFileToStruct(t *testing.T) {
	// Test with an existing file and valid struct model
	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	testFile := "testdata/test.json"
	expectedPerson := Person{Name: "John", Age: 25}
	var actualPerson Person
	err := CastJsonFileToStruct(testFile, &actualPerson)
	if err != nil {
		t.Errorf("CastJsonFileToStruct(%s, &actualPerson) returned error: %s", testFile, err.Error())
	}
	if !reflect.DeepEqual(expectedPerson, actualPerson) {
		t.Errorf("CastJsonFileToStruct(%s, &actualPerson) returned %v, expected %v", testFile, actualPerson, expectedPerson)
	}

	// Test with a non-existing file
	nonExistingFile := "testdata/non-existing.json"
	err = CastJsonFileToStruct(nonExistingFile, &actualPerson)
	if err == nil {
		t.Errorf("CastJsonFileToStruct(%s, &actualPerson) returned nil error, expected non-nil error", nonExistingFile)
	}

	// Test with an empty file path
	emptyPath := ""
	err = CastJsonFileToStruct(emptyPath, &actualPerson)
	if err == nil {
		t.Errorf("CastJsonFileToStruct(%s, &actualPerson) returned nil error, expected non-nil error", emptyPath)
	}

	// Test with a nil struct pointer
	var nilPointer *Person
	err = CastJsonFileToStruct(testFile, nilPointer)
	if err == nil {
		t.Errorf("CastJsonFileToStruct(%s, nil) returned nil error, expected non-nil error", testFile)
	}

	// Test with a struct model that is not a pointer
	notPointerModel := Person{}
	err = CastJsonFileToStruct(testFile, notPointerModel)
	if err == nil {
		t.Errorf("CastJsonFileToStruct(%s, notPointerModel) returned nil error, expected non-nil error", testFile)
	}
}
