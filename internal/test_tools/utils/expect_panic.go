package utils

import (
	"reflect"
	"testing"
)

// ExpectPanic expects a panic to occur when running the subject function.
//
// Generic Constraints:
//   - TFunc: The type of the function to test.
//
// Parameters:
//   - t: The testing.T object to use for assertions.
//   - subjectFunc: The function to test.
//
// Example usage:
//
//	ExpectPanic(t, func() {
//		panic("expected panic")
//	})
func ExpectPanic[TFunc any](t *testing.T, subjectFunc TFunc) {
	if reflect.TypeOf(subjectFunc).Kind() != reflect.Func {
		t.Error("expected a function")
	}

	defer func() {
		if r := recover(); r == nil {
			t.Error("expected a panic to occur")
		}
	}()

	// Run the subject function for testing.
	var fn func()
	reflect.ValueOf(&fn).Elem().Set(reflect.ValueOf(subjectFunc))
	fn()
}
