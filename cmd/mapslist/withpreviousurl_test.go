package mapslist

import (
	"testing"

	f "internal/tests/fixtures"
)

func TestWithPreviousURL(t *testing.T) {
	mapsList := performWithPreviousURLTest(&f.APIEndpoint)

	if *mapsList.PreviousURL != f.APIEndpoint {
		t.Errorf("Expected mapsList.previous to be %q, but got %q", f.APIEndpoint, *mapsList.PreviousURL)
	}
}

func TestWithPreviousURL_Nil(t *testing.T) {
	var nilString *string = nil
	mapsList := performWithPreviousURLTest(nilString, &f.APIEndpoint)

	if mapsList.NextURL != nilString {
		t.Errorf("Expected mapsList.previous to be nil, but got non-nil")
	}
}

func performWithPreviousURLTest(url *string, defaultPreviousURL ...*string) *MapsList {
	mapsList := createPreviousURLMapsList(defaultPreviousURL...)

	WithPreviousURL(url)(mapsList)

	return mapsList
}

func createPreviousURLMapsList(defaultPreviousURL ...*string) *MapsList {
	if len(defaultPreviousURL) != 0 && defaultPreviousURL[0] != nil {
		return &MapsList{NextURL: defaultPreviousURL[0]}
	}
	return &MapsList{}
}
