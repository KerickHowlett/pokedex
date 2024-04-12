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
//	TPayload: The schematic of the parsed query result.
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
func decode[TPayload any](body []byte, payload *TPayload) error {
	if err := json.Unmarshal(body, payload); err != nil {
		return fmt.Errorf("error with parsing payload %v", err)
	}

	return nil
}
