package mapslist

import (
	"testing"

	f "internal/tests/fixtures"

	l "github.com/KerickHowlett/pokedexcli/cmd/location"
)

func TestMapsList_GetNextURL(t *testing.T) {
	_, mapsList := setupTestMapsListStructTest()
	nextURL := mapsList.NextURL

	if *mapsList.NextURL != f.APIEndpoint {
		t.Errorf("Expected GetNextURL to return %s, but got %s", f.APIEndpoint, *nextURL)
	}
}

func TestMapsList_GetPreviousURL(t *testing.T) {
	_, mapsList := setupTestMapsListStructTest()
	previousURL := mapsList.PreviousURL

	if *mapsList.PreviousURL != f.APIEndpoint {
		t.Errorf("Expected GetPreviousURL to return %s, but got %s", f.APIEndpoint, *previousURL)
	}
}

func TestMapsList_GetResults(t *testing.T) {
	location, mapsList := setupTestMapsListStructTest()
	result := mapsList.Locations[0]

	if result != *location {
		expectedName, actualName := location.Name, result.Name
		t.Errorf("Expected GetResults[0].Name to be %s, but got %s", actualName, expectedName)
	}
}

func setupTestMapsListStructTest() (location *l.Location, mapsList *MapsList) {
	location = l.NewLocation(l.WithName("Pallet Town"))
	mapsList = NewMapsList(
		AddLocation(*location),
		WithNextURL(&f.APIEndpoint),
		WithPreviousURL(&f.APIEndpoint),
	)

	return location, mapsList
}
