package utils

import (
	"reflect"
	"testing"
)

// @TODO: Add stronger generic typing.
// ExpectSameEntity compares two values and fails the test if they are not the
// same entity, which is determined by their respective pointers.
//
// params:
// - t: the testing.T instance
// - actual: the actual value
// - expected: the expected value
// - fieldName: the name of the field being tested
//
// returns: nothing
//
// Example usage:
// utils.ExpectSameEntity(t, actualScanner, expectedScanner, "Scanner")
func ExpectSameEntity(t *testing.T, actual any, expected any, fieldName string) {
	// argsType := reflect.TypeOf(actual).Kind()
	// if argsType == reflect.Func && reflect.ValueOf(actual).Pointer() == reflect.ValueOf(expected).Pointer() {

	// } else if argsType == reflect.Struct && actual == expected {

	// } else {

	// }
	if reflect.ValueOf(actual).Pointer() == reflect.ValueOf(expected).Pointer() {
		return
	}
	t.Errorf("Expected %s to be set with the argued function, %v, but instead got %v\n", fieldName, expected, actual)
}
