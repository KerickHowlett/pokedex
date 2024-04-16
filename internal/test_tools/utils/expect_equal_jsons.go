package utils

import "encoding/json"

// ExpectEqualJSONs compares two JSON objects and returns true if they are equal.
//
// Generic Constraints:
//   - TStruct: The type of the JSON objects to compare.
//
// Parameters:
//   - a: The first JSON object to compare.
//   - b: The second JSON object to compare.
//
// Returns:
//   - true if the JSON objects are equal, and false otherwise.
//
// Example usage:
//
// type SomeStruct struct {Name string `json:"name"`}
// a := SomeStruct{Name string `json:"name"`}
// b := SomeStruct{ Name string `json:"name"`}
// isMatching := ExpectEqualJSONs(a, b) // Returns true
func ExpectEqualJSONs[TStruct any](a TStruct, b TStruct) bool {
	aJSON, _ := json.Marshal(a)
	bJSON, _ := json.Marshal(b)
	return string(aJSON) == string(bJSON)
}
