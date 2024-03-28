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
	if reflect.DeepEqual(actual, expected) {
		return
	}
	t.Errorf("Expected %s to be set with the argued function, but instead got %v\n", fieldName, actual)
}
