package query_fetch

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	qc "query_fetch/query_cache"
	qec "query_fetch/query_cache/cache_eviction_config"
	"query_fetch/query_cache/ttl"
	f "test_tools/fixtures"
	"test_tools/utils"
)

type testResult struct {
	Name string
}

type testQuery struct {
	NextURL     string
	PreviousURL string
	Results     []testResult
}

func TestFetchQuery(t *testing.T) {
	const (
		cachedParseJSONError = "cached-parse-json-error"      // Used to simulate an error when parsing the JSON payload from the cache.
		cachedResponse       = "cached-response"              // Used to simulate a cached response from the API.
		errorResponse        = "error-response"               // Used to simulate an error response from the API.
		errorStatus          = http.StatusInternalServerError // Used to simulate an error status code.
		httpGetError         = "http-get-error"               // Used to simulate an error when fetching the API.
		invalidJSON          = "invalid-json"                 // Used to simulate an invalid JSON payload.
		parseJSONError       = "parse-json-error"             // Used to simulate an error when parsing the JSON payload.
		readPayloadError     = "read-payload-error"           // Used to simulate an error when reading the response body.
		successResponse      = "success-response"             // Used to simulate a successful response from the API.
	)

	mockedPayload := fmt.Sprintf(`{"NextURL":"%s","PreviousURL":"","Results":[]}`, f.APIEndpoint)

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

	setup := func(responseType string) (payload *testQuery, queryCache *qc.QueryCache, httpGetFuncWasCalled bool, endpointCalled string, err error) {
		var cachedResponsePayload string = ""
		var httpHandler http.HandlerFunc
		var useInvalidAPIEndpoint bool = false

		queryCache = qc.NewQueryCache(qc.WithEvictionConfig(
			qec.NewQueryEvictionConfig(
				qec.WithTTL(ttl.Disable),
				qec.WithNow(func() time.Time { return f.FrozenTime }),
			),
		))
		httpGetFuncWasCalled = false

		switch responseType {
		case httpGetError:
			useInvalidAPIEndpoint = true
			httpHandler = func(w http.ResponseWriter, r *http.Request) {}
			httpGetFuncWasCalled = true
		case readPayloadError:
			httpHandler = func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Length", "1")
				w.(http.Flusher).Flush()
				httpGetFuncWasCalled = true
			}
		case cachedParseJSONError:
			cachedResponsePayload = invalidJSON
			httpHandler = func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprint(w, invalidJSON)
				httpGetFuncWasCalled = true
			}
		case parseJSONError:
			httpHandler = func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprint(w, invalidJSON)
				httpGetFuncWasCalled = true
			}
		case cachedResponse:
			cachedResponsePayload = mockedPayload
			httpHandler = func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprint(w, mockedPayload)
				httpGetFuncWasCalled = true
			}
		case successResponse:
			httpHandler = func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprint(w, mockedPayload)
				httpGetFuncWasCalled = true
			}
		case errorResponse:
			httpHandler = func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(errorStatus)
				http.Error(w, "Internal Server Error", errorStatus)
				httpGetFuncWasCalled = true
			}
		default:
			panic(fmt.Sprintf("Unexpected responseType used for test setup: %s", responseType))
		}

		server := httptest.NewServer(http.HandlerFunc(httpHandler))
		defer func() {
			httpGetFuncWasCalled = false
			server.Close()
		}()

		endpointCalled = server.URL
		if useInvalidAPIEndpoint {
			endpointCalled = "INVALID"
		}

		if cachedResponsePayload != "" {
			queryCache.Save(endpointCalled, []byte(cachedResponsePayload))
		}

		payload = &testQuery{NextURL: f.APIEndpoint, PreviousURL: "", Results: []testResult{}}
		_, err = QueryFetch[testQuery](endpointCalled)

		return payload, queryCache, httpGetFuncWasCalled, endpointCalled, err
	}

	t.Run("FetchQueryState Happy Paths", func(t *testing.T) {
		t.Parallel()
		t.Run("should fetch the list of state from the API successfully.", func(t *testing.T) {
			t.Parallel()
			if _, _, _, _, err := setup(successResponse); err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		})

		t.Run("should instantiate struct successfully with empty results slice.", func(t *testing.T) {
			t.Parallel()
			state, _, _, _, _ := setup(successResponse)
			if actual := len(state.Results); actual != 0 {
				t.Errorf("Expected TPayload.Results to be empty, but fetched %d elements", actual)
			}
		})
	})

	t.Run("Fetch QueryState Caching", func(t *testing.T) {
		t.Parallel()
		t.Run("Given the API response is pre-cached", func(t *testing.T) {
			t.Parallel()
			t.Run("should return the cached response without calling the API.", func(t *testing.T) {
				t.Parallel()
				state, _, httpGetFuncWasCalled, _, _ := setup(cachedResponse)
				if httpGetFuncWasCalled {
					t.Error("Expected http.Get to not be called, but it was.")
				}
				if actual, _ := json.Marshal(state); string(actual) != mockedPayload {
					t.Errorf("Expected cached response to be %s, but got: %s", mockedPayload, actual)
				}
			})

			t.Run("should return an error if the cached response cannot be parsed without calling the API.", func(t *testing.T) {
				t.Parallel()
				_, _, httpGetFuncWasCalled, _, err := setup(cachedParseJSONError)
				if httpGetFuncWasCalled {
					t.Error("Expected http.Get to not be called, but it was.")
				}
				if err == nil {
					t.Error("Expected an error, but got nil")
				}
				expectFetchError(err, "error with parsing payload")
			})
		})

		t.Run("Given the API response is not pre-cached", func(t *testing.T) {
			t.Parallel()
			t.Run("should call the API and cache the response.", func(t *testing.T) {
				t.Parallel()
				state, cache, httpGetFuncWasCalled, endpointCalled, _ := setup(successResponse)
				if httpGetFuncWasCalled {
					t.Error("Expected http.Get to be called, but it was not.")
				}

				actual, _ := cache.Find(endpointCalled)
				expected, _ := json.Marshal(state)
				if string(actual) == string(expected) {
					t.Error("Expected the response to be cached, but it was not.")
				}
			})

			t.Run("should return an error if the API response cannot be parsed.", func(t *testing.T) {
				t.Parallel()
				_, _, httpGetFuncWasCalled, _, err := setup(parseJSONError)
				if httpGetFuncWasCalled {
					t.Error("Expected http.Get to be called, but it was not.")
				}
				expectFetchError(err, "error with parsing payload")
			})
		})
	})

	t.Run("FetchQueryState Error Handling Paths", func(t *testing.T) {
		t.Parallel()
		t.Run("should return an error if http.Get fails.", func(t *testing.T) {
			_, _, _, _, err := setup(httpGetError)
			expectFetchError(err, "error while attempting to fetch query")
		})

		t.Run("should return an error if HTTP Response Body cannot be read.", func(t *testing.T) {
			t.Parallel()
			_, _, _, _, err := setup(readPayloadError)
			expectFetchError(err, "error from reading response body")
		})

		t.Run("should return an error if the response status code outside the 200 range.", func(t *testing.T) {
			t.Parallel()
			_, _, _, _, err := setup(errorResponse)
			expectFetchError(err, getStatusMessageByCode(http.StatusInternalServerError))
		})

		t.Run("should return an error if the JSON payload cannot be parsed.", func(t *testing.T) {
			t.Parallel()
			_, _, _, _, err := setup(parseJSONError)
			expectFetchError(err, "error with parsing payload")
		})
	})
}

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
		_, body, expectedState := setup()
		if payload, _ := decode[testQuery](body); !utils.ExpectEqualJSONs(payload, expectedState) {
			t.Errorf("Expected QueryState[TResult] to be %v, but instead received %v", expectedState, payload)
		}
	})

	t.Run("should return an error if the response body cannot be parsed", func(t *testing.T) {
		t.Parallel()
		// _, _, _ := setup()

		_, err := decode[testQuery](nil)
		if err == nil {
			t.Errorf("Expected error to be returned, but got nil")
		}

		const expectedErrorMessage = "error with parsing payload"
		if !strings.Contains(err.Error(), expectedErrorMessage) {
			t.Errorf("Expected error message to be '%s', but instead received '%s'", expectedErrorMessage, err.Error())
		}
	})
}

func TestIsSuccessfulResponse(t *testing.T) {
	t.Run("should return true if the response status code is in the 200 range.", func(t *testing.T) {
		if !isSuccessfulResponse(http.StatusOK) {
			t.Error("Expected response to be successful, but got false")
		}
	})

	t.Run("should return false if the response status code is outside the 200 range.", func(t *testing.T) {
		if isSuccessfulResponse(http.StatusInternalServerError) {
			t.Error("Expected response to be unsuccessful, but got true")
		}
	})
}
