package query_fetch

import (
	"encoding/json"
	"fmt"
	qs "query/state"
)

// decode parses the HTTP Response's JSON body and populates the
// queryState it's values.
//
// Parameters:
//
//	body: The HTTP Response's body.
//	queryState: The query state to populate with the response's values.
//
// Returns:
//
//	error: An error if the response body cannot be parsed.
//
// Example:
//
//	body := []byte(`{
//		"NextURL": null,
//		"PreviousURL": null,
//		"Results": [{"name": "value"}]
//	}`)
//	state := qs.NewQueryState[result]()
//	err := decode(body, state)
func decode[TResult any](body []byte, queryState *qs.QueryState[TResult]) error {
	if err := json.Unmarshal(body, queryState); err != nil {
		return fmt.Errorf("error with parsing payload %v", err)
	}

	return nil
}
