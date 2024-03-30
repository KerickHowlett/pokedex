package mapslist

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	f "internal/tests/fixtures"
)

func TestFetchMapsList_Success_NoError(t *testing.T) {
	maps, server := setupFetchMapsListTests("success")
	defer server.Close()

	err := maps.FetchMapsList(server.URL)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestFetchMapsList_Success_NextURL(t *testing.T) {
	maps, server := setupFetchMapsListTests("success")
	defer server.Close()

	maps.FetchMapsList(server.URL)

	if *maps.NextURL != f.APIEndpoint {
		t.Errorf("Expected maps.next to be nil, but got: %v", maps.NextURL)
	}
}

func TestFetchMapsList_Success_PreviousURL(t *testing.T) {
	maps, server := setupFetchMapsListTests("success")
	defer server.Close()

	maps.FetchMapsList(server.URL)

	if maps.PreviousURL != nil {
		t.Errorf("Expected maps.previous to be nil, but got: %v", maps.PreviousURL)
	}
}

func TestFetchMapsList_Success_Locations(t *testing.T) {
	maps, server := setupFetchMapsListTests("success")
	defer server.Close()

	maps.FetchMapsList(server.URL)

	if len(maps.Locations) != 0 {
		t.Errorf("Expected maps.results to be empty, but got %d elements", len(maps.Locations))
	}
}

func TestFetchMapsList_HTTPError(t *testing.T) {
	maps, server := setupFetchMapsListTests("get-error")
	defer server.Close()

	err := maps.FetchMapsList("invalid")

	const expectedErrorMessage = "error fetching maps list"
	expectFetchError(t, err, expectedErrorMessage)
}

func TestFetchMapsList_ReadError(t *testing.T) {
	maps, server := setupFetchMapsListTests("read-error")
	defer server.Close()

	err := maps.FetchMapsList(server.URL)

	const expectedErrorMessage = "error reading response body"
	expectFetchError(t, err, expectedErrorMessage)
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
