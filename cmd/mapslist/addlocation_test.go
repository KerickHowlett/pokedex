package mapslist

import (
	"testing"

	l "github.com/KerickHowlett/pokedexcli/cmd/location"
)

func TestAddLocation(t *testing.T) {
	setup := func() (location *l.Location, mapsList *MapsList) {
		const starterTown = "Pallet Town"
		location = l.NewLocation(l.WithName(starterTown))
		mapsList = &MapsList{}

		return location, mapsList
	}

	t.Run("should add location to the mapsList.Locations field.", func(t *testing.T) {
		location, mapsList := setup()
		expectedLength := len(mapsList.Locations) + 1

		AddLocation(*location)(mapsList)

		if len(mapsList.Locations) != expectedLength {
			t.Errorf("Expected mapsList.results to have 1 item, but got %d", len(mapsList.Locations))
		}
	})

	t.Run("should add Location with unchanged values.", func(t *testing.T) {
		location, mapsList := setup()

		AddLocation(*location)(mapsList)

		if mapsList.Locations[0] != *location {
			t.Errorf("Expected mapsList.results[0] to be %v, but got %v", location, mapsList.Locations[0])
		}
	})
}
