package queryfetch

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	qs "query/state"
)

func QueryFetch[TResult any](url string, queryState *qs.QueryState[TResult]) error {
	response, err := http.Get(url)

	if err != nil {
		return fmt.Errorf("error while attempting to fetch query: %v", err)
	}

	body, err := io.ReadAll(response.Body)
	response.Body.Close()

	if err != nil {
		return fmt.Errorf("error from reading response body: %v", err)
	}

	if !isSuccessfulResponse(response.StatusCode) {
		return fmt.Errorf("error with response: %s", body)
	}

	if err = json.Unmarshal(body, &queryState); err != nil {
		return fmt.Errorf("error with parsing payload %v", err)
	}

	return nil
}
