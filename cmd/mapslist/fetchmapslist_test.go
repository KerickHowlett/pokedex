package mapslist

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	f "testtools/fixtures"
)

func TestFetchMapsList(t *testing.T) {
	const (
		HTTPGetError     = "http-get-error"     // Used to simulate an error when fetching the API.
		ReadPayloadError = "read-payload-error" // Used to simulate an error when reading the response body.
		ParseJSONError   = "parse-json-error"   // Used to simulate an error when parsing the JSON payload.
		SuccessResponse  = "success-response"   // Used to simulate a successful response from the API.
		ErrorResponse    = "error-response"     // Used to simulate an error response from the API.
	)

	setup := func(responseType string) (maps *MapsList, err error) {
		var httpHandler http.HandlerFunc
		var useInvalidAPIEndpoint bool = false

		switch responseType {
		case HTTPGetError:
			useInvalidAPIEndpoint = true
			httpHandler = func(w http.ResponseWriter, r *http.Request) {}
		case ReadPayloadError:
			httpHandler = func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Length", "1")
				w.(http.Flusher).Flush()
			}
		case ParseJSONError:
			httpHandler = func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprint(w, "invalid json")
			}
		case SuccessResponse:
			httpHandler = func(w http.ResponseWriter, r *http.Request) {
				mockedPayload := fmt.Sprintf(`{"next": "%s", "previous": null, "results": []}`, f.APIEndpoint)
				fmt.Fprint(w, mockedPayload)
			}
		case ErrorResponse:
			errorStatus := http.StatusInternalServerError
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

		maps = NewMapsList()
		err = maps.FetchMapsList(apiEndpointForTest)

		return maps, err
	}

	t.Run("FetchMapsList Happy Paths", func(t *testing.T) {
		t.Run("should fetch the list of maps from the API successfully.", func(t *testing.T) {
			if _, err := setup(SuccessResponse); err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		})

		t.Run("MapsList struct payload", func(t *testing.T) {
			t.Run("should set the NextURL field with the same named field's value.", func(t *testing.T) {
				maps, _ := setup(SuccessResponse)
				if actual := *maps.NextURL; actual != f.APIEndpoint {
					t.Errorf("Expected maps.NextURL to be %s, but fetched: %v", f.APIEndpoint, actual)
				}
			})

			t.Run("should set the PreviousURL with the same named field's value.", func(t *testing.T) {
				maps, _ := setup(SuccessResponse)
				if actual := maps.PreviousURL; actual != nil {
					t.Errorf("Expected maps.PreviousURL to be nil, but fetched: %v", actual)
				}
			})

			t.Run("should set the Locations field with the same named field's slice.", func(t *testing.T) {
				maps, _ := setup(SuccessResponse)
				if actual := len(maps.Locations); actual != 0 {
					t.Errorf("Expected maps.Locations to be empty, but fetched %d elements", actual)
				}
			})
		})
	})

	t.Run("FetchMapsList Error Handling Paths", func(t *testing.T) {
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
			_, err := setup(HTTPGetError)
			expectFetchError(err, "error fetching maps list")
		})

		t.Run("should return an error if HTTP Response Body cannot be read.", func(t *testing.T) {
			_, err := setup(ReadPayloadError)
			expectFetchError(err, "error reading response body")
		})

		t.Run("should return an error if the response status code outside the 200 range.", func(t *testing.T) {
			_, err := setup(ErrorResponse)
			expectFetchError(err, getStatusMessageByCode(http.StatusInternalServerError))
		})

		t.Run("should return an error if the JSON payload cannot be parsed.", func(t *testing.T) {
			_, err := setup(ParseJSONError)
			expectFetchError(err, "error with parsing payload")
		})
	})
}

func TestIsSuccessfulResponse(t *testing.T) {
	runIsSuccessfulResponseTest := func(statusCode int) bool {
		mapsList := NewMapsList()
		response := &http.Response{StatusCode: statusCode}

		return mapsList.isSuccessfulResponse(response)
	}

	t.Run("should return true if the response status code is in the 200 range.", func(t *testing.T) {
		if !runIsSuccessfulResponseTest(http.StatusOK) {
			t.Error("Expected response to be successful, but got false")
		}
	})

	t.Run("should return false if the response status code is outside the 200 range.", func(t *testing.T) {
		if runIsSuccessfulResponseTest(http.StatusInternalServerError) {
			t.Error("Expected response to be unsuccessful, but got true")
		}
	})
}
