package mapslist

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (m *MapsList) FetchMapsList(url string) error {
	response, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error fetching maps list: %v", err)
	}

	body, err := io.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return fmt.Errorf("error reading response body: %v", err)
	}

	if !m.isSuccessfulResponse(response) {
		return fmt.Errorf("error with response: %s", body)
	}

	if err = json.Unmarshal(body, &m); err != nil {
		return fmt.Errorf("error with parsing payload %v", err)
	}

	return nil
}

func (m *MapsList) isSuccessfulResponse(response *http.Response) bool {
	code := response.StatusCode
	return code >= http.StatusOK && code < http.StatusMultipleChoices
}
