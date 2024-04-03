package mapslist

import (
	"fmt"
	"testing"

	f "testtools/fixtures"

	l "github.com/KerickHowlett/pokedexcli/cmd/location"
)

func TestMapsList(t *testing.T) {
	setup := func() (location *l.Location, mapsList *MapsList) {
		location = l.NewLocation(l.WithName("Pallet Town"))
		mapsList = NewMapsList(
			AddLocation(*location),
			WithNextURL(&f.APIEndpoint),
			WithPreviousURL(&f.APIEndpoint),
		)

		return location, mapsList
	}

	t.Run("should return a non-nil MapsList struct.", func(t *testing.T) {
		fmt.Println("should return a non-nil MapsList struct.")
		_, mapsList := setup()

		if mapsList == nil {
			t.Errorf("Expected NewMapsList to return a MapsList, but got nil")
		}
	})

	t.Run("should return the NextURL value set in the MapsList field.", func(t *testing.T) {
		_, mapsList := setup()
		nextURL := *mapsList.NextURL

		if nextURL != f.APIEndpoint {
			t.Errorf("Expected GetNextURL to return %s, but got %s", f.APIEndpoint, nextURL)
		}
	})

	t.Run("should return the PreviousURL field set in the MapsList field.", func(t *testing.T) {
		_, mapsList := setup()
		previousURL := *mapsList.PreviousURL

		if previousURL != f.APIEndpoint {
			t.Errorf("Expected GetPreviousURL to return %s, but got %s", f.APIEndpoint, previousURL)
		}
	})

	t.Run("should return a list of locations stored in the MapsList field.", func(t *testing.T) {
		location, mapsList := setup()
		result := mapsList.Locations[0]

		if result != *location {
			expectedName, actualName := location.Name, result.Name
			t.Errorf("Expected Locations[0].Name to be %s, but got %s", actualName, expectedName)
		}
	})
}
