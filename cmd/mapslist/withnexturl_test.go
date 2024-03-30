package mapslist

import (
	"testing"

	f "internal/tests/fixtures"
)

func TestWithNextURL(t *testing.T) {
	mapsList := performWithNextURLTest(&f.APIEndpoint)

	if *mapsList.NextURL != f.APIEndpoint {
		t.Errorf("Expected mapsList.next to be %q, but got %q", f.APIEndpoint, *mapsList.NextURL)
	}
}

func TestWithNextURL_Nil(t *testing.T) {
	var nilString *string = nil
	mapsList := performWithNextURLTest(nilString, &f.APIEndpoint)

	if mapsList.NextURL != nilString {
		t.Errorf("Expected mapsList.next to be nil, but got non-nil")
	}
}

func performWithNextURLTest(url *string, defaultNextURL ...*string) *MapsList {
	mapsList := createNextURLMapsList(defaultNextURL...)

	WithNextURL(url)(mapsList)

	return mapsList
}

func createNextURLMapsList(defaultNextURL ...*string) *MapsList {
	if len(defaultNextURL) != 0 && defaultNextURL[0] != nil {
		return &MapsList{NextURL: defaultNextURL[0]}
	}
	return &MapsList{}
}
