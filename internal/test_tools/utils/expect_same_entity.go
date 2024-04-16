package utils

import (
	"reflect"
	"testing"
)

// ExpectSameEntity compares two values and fails the test if they are not the
// same entity, which is determined by their respective pointers.
//
// params:
// - t: the testing.T instance
// - actual: the actual value
// - expected: the expected value
// - fieldName: the name of the field being tested
//
// Example usage:
// utils.ExpectSameEntity(t, actualScanner, expectedScanner, "Scanner")
func ExpectSameEntity(t *testing.T, actual any, expected any, fieldName string) {
	if actualPointer, expectedPointer := reflect.ValueOf(actual).Pointer(), reflect.ValueOf(expected).Pointer(); actualPointer != expectedPointer {
		t.Errorf("Expected %s to be set with the argued entity, %v, but instead got %v\n", fieldName, expectedPointer, actualPointer)
	}
}
