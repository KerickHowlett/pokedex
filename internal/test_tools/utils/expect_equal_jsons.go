package utils

import "encoding/json"

func ExpectEqualJSONs[TStruct any](a TStruct, b TStruct) bool {
	aJSON, _ := json.Marshal(a)
	bJSON, _ := json.Marshal(b)
	return string(aJSON) == string(bJSON)
}
