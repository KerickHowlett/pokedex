package queryfetch

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	qs "query/state"
	f "test_tools/fixtures"
)

type MockedResult struct {
	Name string
}

type MockedQueryState = qs.QueryState[MockedResult]

func TestFetchQuery(t *testing.T) {
	const (
		httpGetError     = "http-get-error"               // Used to simulate an error when fetching the API.
		readPayloadError = "read-payload-error"           // Used to simulate an error when reading the response body.
		parseJSONError   = "parse-json-error"             // Used to simulate an error when parsing the JSON payload.
		successResponse  = "success-response"             // Used to simulate a successful response from the API.
		errorResponse    = "error-response"               // Used to simulate an error response from the API.
		errorStatus      = http.StatusInternalServerError // Used to simulate an error status code.
	)

	setup := func(responseType string) (queryState *MockedQueryState, err error) {
		var httpHandler http.HandlerFunc
		var useInvalidAPIEndpoint bool = false

		switch responseType {
		case httpGetError:
			useInvalidAPIEndpoint = true
			httpHandler = func(w http.ResponseWriter, r *http.Request) {}
		case readPayloadError:
			httpHandler = func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Length", "1")
				w.(http.Flusher).Flush()
			}
		case parseJSONError:
			httpHandler = func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprint(w, "invalid json")
			}
		case successResponse:
			httpHandler = func(w http.ResponseWriter, r *http.Request) {
				mockedPayload := fmt.Sprintf(`{"next": "%s", "previous": null, "results": []}`, f.APIEndpoint)
				fmt.Fprint(w, mockedPayload)
			}
		case errorResponse:
			httpHandler = func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(errorStatus)
				http.Error(w, "Internal Server Error", errorStatus)
			}
		default:
			panic(fmt.Sprintf("Unexpected responseType used for test setup: %s", responseType))
		}

		server := httptest.NewServer(http.HandlerFunc(httpHandler))
		defer server.Close()

		apiEndpointForTest := server.URL
		if useInvalidAPIEndpoint {
			apiEndpointForTest = "INVALID"
		}

		queryState = qs.NewQueryState[MockedResult]()
		err = QueryFetch(apiEndpointForTest, queryState)

		return queryState, err
	}

	t.Run("FetchQueryState Happy Paths", func(t *testing.T) {
		t.Run("should fetch the list of state from the API successfully.", func(t *testing.T) {
			if _, err := setup(successResponse); err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		})

		t.Run("QueryState struct payload", func(t *testing.T) {
			t.Run("should set the NextURL field with the same named field's value.", func(t *testing.T) {
				state, _ := setup(successResponse)
				if actual := *state.NextURL; actual != f.APIEndpoint {
					t.Errorf("Expected QueryState[TResult].NextURL to be %s, but fetched: %v", f.APIEndpoint, actual)
				}
			})

			t.Run("should set the PreviousURL with the same named field's value.", func(t *testing.T) {
				state, _ := setup(successResponse)
				if actual := state.PreviousURL; actual != nil {
					t.Errorf("Expected QueryState[TResult].PreviousURL to be nil, but fetched: %v", actual)
				}
			})

			t.Run("should set the Results field with the same named field's slice.", func(t *testing.T) {
				state, _ := setup(successResponse)
				if actual := len(state.Results); actual != 0 {
					t.Errorf("Expected QueryState[TResult].Results to be empty, but fetched %d elements", actual)
				}
			})
		})
	})

	t.Run("FetchQueryState Error Handling Paths", func(t *testing.T) {
		expectFetchError := func(err error, expectedErrorMessage string) {
			if err == nil {
				t.Error("Expected an error, but got nil")
			}

			if err.Error()[:len(expectedErrorMessage)] != expectedErrorMessage {
				t.Errorf("Expected error message to be '%s', but got: %s", expectedErrorMessage, err.Error())
			}
		}

		getStatusMessageByCode := func(statusCode int) string {
			statusMessage := http.StatusText(statusCode)
			return fmt.Sprintf("error with response: %s", statusMessage)
		}

		t.Run("should return an error if http.Get fails.", func(t *testing.T) {
			_, err := setup(httpGetError)
			expectFetchError(err, "error while attempting to fetch query")
		})

		t.Run("should return an error if HTTP Response Body cannot be read.", func(t *testing.T) {
			_, err := setup(readPayloadError)
			expectFetchError(err, "error from reading response body")
		})

		t.Run("should return an error if the response status code outside the 200 range.", func(t *testing.T) {
			_, err := setup(errorResponse)
			expectFetchError(err, getStatusMessageByCode(http.StatusInternalServerError))
		})

		t.Run("should return an error if the JSON payload cannot be parsed.", func(t *testing.T) {
			_, err := setup(parseJSONError)
			expectFetchError(err, "error with parsing payload")
		})
	})
}
