package mapslist

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	f "internal/tests/fixtures"
)

func TestFetchMapsList(t *testing.T) {
	t.Run("should fetch the list of maps from the API successfully.", func(t *testing.T) {
		maps, server := setupFetchMapsListTests("success")
		defer server.Close()

		err := maps.FetchMapsList(server.URL)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
	})

	t.Run("MapsList struct payload", func(t *testing.T) {
		t.Run("should set the NextURL field with the same named field's value.", func(t *testing.T) {
			maps, server := setupFetchMapsListTests("success")
			defer server.Close()

			maps.FetchMapsList(server.URL)

			if *maps.NextURL != f.APIEndpoint {
				t.Errorf("Expected maps.NextURL to be %s, but fetched: %v", f.APIEndpoint, maps.NextURL)
			}
		})

		t.Run("should set the PreviousURL with the same named field's value.", func(t *testing.T) {
			maps, server := setupFetchMapsListTests("success")
			defer server.Close()

			maps.FetchMapsList(server.URL)

			if maps.PreviousURL != nil {
				t.Errorf("Expected maps.PreviousURL to be nil, but fetched: %v", maps.PreviousURL)
			}
		})

		t.Run("should set the Locations field with the same named field's slice.", func(t *testing.T) {
			maps, server := setupFetchMapsListTests("success")
			defer server.Close()

			maps.FetchMapsList(server.URL)

			if len(maps.Locations) != 0 {
				t.Errorf("Expected maps.Locations to be empty, but fetched %d elements", len(maps.Locations))
			}
		})
	})

	t.Run("should return an error if http.Get fails.", func(t *testing.T) {
		maps, server := setupFetchMapsListTests("get-error")
		defer server.Close()

		err := maps.FetchMapsList("invalid")

		expectFetchError(t, err, "error fetching maps list")
	})

	t.Run("should return an error if HTTP Response Body cannot be read.", func(t *testing.T) {
		maps, server := setupFetchMapsListTests("read-error")
		defer server.Close()

		err := maps.FetchMapsList(server.URL)

		expectFetchError(t, err, "error reading response body")
	})

	t.Run("should return an error if the response status code is not within the 200 range.", func(t *testing.T) {
		maps, server := setupFetchMapsListTests("error-response")
		defer server.Close()

		err := maps.FetchMapsList(server.URL)

		expectFetchError(t, err, "error with response")
	})
}

func TestFetchMapsList_ErrorResponse(t *testing.T) {
	maps, server := setupFetchMapsListTests("error-response")
	defer server.Close()

	err := maps.FetchMapsList(server.URL)

	const expectedErrorMessage = "error with response: Internal Server Error"
	expectFetchError(t, err, expectedErrorMessage)
}

func TestFetchMapsList_InvalidJSON(t *testing.T) {
	maps, server := setupFetchMapsListTests("parse-json-error")
	defer server.Close()

	err := maps.FetchMapsList(server.URL)

	const expectedErrorMessage = "error with parsing payload"
	expectFetchError(t, err, expectedErrorMessage)
}

func expectFetchError(t *testing.T, err error, expectedErrorMessage string) {
	if err == nil {
		t.Error("Expected an error, but got nil")
	}

	if err.Error()[:len(expectedErrorMessage)] != expectedErrorMessage {
		t.Errorf("Expected error message to be '%s', but got: %s", expectedErrorMessage, err.Error())
	}
}

func setupFetchMapsListTests(responseType string) (maps *MapsList, server *httptest.Server) {
	server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if responseType == "get-error" {
			return
		}

		if responseType == "read-error" {
			w.Header().Set("Content-Length", "1")
			w.(http.Flusher).Flush()
			return
		}

		if responseType == "parse-json-error" {
			fmt.Fprint(w, "invalid json")
			return
		}

		if responseType == "success" {
			responseJSON := fmt.Sprintf(`{"next": "%s", "previous": null, "results": []}`, f.APIEndpoint)
			fmt.Fprint(w, responseJSON)
			return
		}

		if responseType == "error-response" {
			w.WriteHeader(http.StatusInternalServerError)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}))

	maps = NewMapsList()

	return maps, server
}
