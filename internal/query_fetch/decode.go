package query_fetch

import (
	"encoding/json"
	"fmt"
)

// decode parses the HTTP Response's JSON body and populates the
// payload it's values.
//
// Generic Constraints:
//
//	TQuery: The schematic of the intended query payload.
//
// Parameters:
//
//	body: The HTTP Response's body.
//	payload: The payload struct to populate with the response body's values.
//
// Returns:
//
//	error: An error if the response body cannot be parsed.
//
// Example:
//
//	err := decode(body, state)
func decode[TQuery any](body []byte, query *TQuery) error {
	if err := json.Unmarshal(body, query); err != nil {
		return fmt.Errorf("error with parsing payload %v", err)
	}

	return nil
}
