package querystate

import (
	"testing"

	f "testtools/fixtures"
)

func TestWithResult(t *testing.T) {
	runWithResultTest := func() (item *result, state *queryTestState) {
		item = &result{Name: f.StarterTown}
		state = &queryTestState{}

		WithResult(*item)(state)

		return item, state
	}

	t.Run("should add location to the QueryState[TResult].Results field.", func(t *testing.T) {
		if _, state := runWithResultTest(); len(state.Results) != 1 {
			t.Errorf("Expected QueryState[TResult].NextURL to have 1 item, but got %d", len(state.Results))
		}
	})

	t.Run("should add Location with unchanged values.", func(t *testing.T) {
		if item, state := runWithResultTest(); state.Results[0] != *item {
			t.Errorf("Expected QueryState[TResult].NextURL[0] to be %v, but got %v", item, state.Results[0])
		}
	})
}
