package query_fetch

import (
	"fmt"
	"strings"
	"testing"

	f "test_tools/fixtures"
	"test_tools/utils"
)

func TestDecode(t *testing.T) {
	setup := func() (payload *testQuery, responseBody []byte, expectedPayload *testQuery) {
		payload = &testQuery{}
		responseBody = []byte(fmt.Sprintf(`{
			"NextURL": null,
			"PreviousURL": null,
			"Results": [{"name": "%s"}]
		}`, f.StarterTown))
		expectedPayload = &testQuery{
			NextURL:     "",
			PreviousURL: "",
			Results:     []testResult{{Name: f.StarterTown}},
		}

		return payload, responseBody, expectedPayload
	}

	t.Run("should parse the response body correctly", func(t *testing.T) {
		t.Parallel()
		payload, body, expectedState := setup()
		if decode(body, payload); !utils.ExpectEqualJSONs(payload, expectedState) {
			t.Errorf("Expected QueryState[TResult] to be %v, but instead received %v", expectedState, payload)
		}
	})

	t.Run("should return an error if the response body cannot be parsed", func(t *testing.T) {
		t.Parallel()
		payload, _, _ := setup()

		err := decode(nil, payload)
		if err == nil {
			t.Errorf("Expected error to be returned, but got nil")
		}

		const expectedErrorMessage = "error with parsing payload"
		if !strings.Contains(err.Error(), expectedErrorMessage) {
			t.Errorf("Expected error message to be '%s', but instead received '%s'", expectedErrorMessage, err.Error())
		}
	})
}
