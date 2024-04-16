package utils

import (
	"reflect"
	"testing"
)

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
