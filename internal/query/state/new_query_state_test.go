package querystate

import (
	"testing"
)

func TestNewQueryState(t *testing.T) {
	expectURLToBeNil := func(url *string, label string) {
		if url != nil {
			t.Errorf("Expected %s to be nil, but got non-nil", label)
		}
	}

	t.Run("should create QueryState[TResult] with it's NextURL field set to nil by default.", func(t *testing.T) {
		state := NewQueryState[result]()
		expectURLToBeNil(state.NextURL, "QueryState[TResult].NextURL")
	})

	t.Run("should create QueryState[TResult] with it's PreviousURL field set to nil by default.", func(t *testing.T) {
		state := NewQueryState[result]()
		expectURLToBeNil(state.PreviousURL, "QueryState[TResult].PreviousURL")
	})

	t.Run("should create QueryState[TResult] with it's Results field set with an empty slice by default.", func(t *testing.T) {
		state := NewQueryState[result]()
		if totalResults := len(state.Results); totalResults != 0 {
			t.Errorf("Expected results to be empty, but got %d elements", totalResults)
		}
	})
}
