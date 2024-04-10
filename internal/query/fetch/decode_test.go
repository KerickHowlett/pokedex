package query_fetch

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	qs "query/state"
	f "test_tools/fixtures"
)

type result struct {
	Name string `json:"name"`
}

func TestDecode(t *testing.T) {
	t.Run("should parse the response body correctly", func(t *testing.T) {
		body := []byte(fmt.Sprintf(`{
			"NextURL": null,
			"PreviousURL": null,
			"Results": [{"name": "%s"}]
		}`, f.StarterTown))
		state := qs.NewQueryState[result]()
		expectedState := qs.NewQueryState(
			qs.WithResult(result{Name: f.StarterTown}),
		)

		if decode(body, state); !expectEqualJSONs(state, expectedState) {
			t.Errorf("Expected QueryState[TResult] to be %v, but instead received %v", expectedState, state)
		}
	})

	t.Run("should return an error if the response body cannot be parsed", func(t *testing.T) {
		err := decode(nil, qs.NewQueryState[result]())

		if err == nil {
			t.Errorf("Expected error to be returned, but got nil")
		}

		const expectedErrorMessage = "error with parsing payload"
		if !strings.Contains(err.Error(), expectedErrorMessage) {
			t.Errorf("Expected error message to be '%s', but instead received '%s'", expectedErrorMessage, err.Error())
		}
	})
}

func expectEqualJSONs(a, b interface{}) bool {
	aJSON, _ := json.Marshal(a)
	bJSON, _ := json.Marshal(b)
	return string(aJSON) == string(bJSON)
}
