package mapslist

import (
	"testing"

	l "github.com/KerickHowlett/pokedexcli/cmd/location"
)

func TestAddLocation(t *testing.T) {
	location := l.NewLocation(l.WithName("Pallet Town"))
	mapsList := &MapsList{}

	AddLocation(*location)(mapsList)

	if len(mapsList.Locations) != 1 {
		t.Errorf("Expected mapsList.results to have 1 item, but got %d", len(mapsList.Locations))
	}

	if mapsList.Locations[0] != *location {
		t.Errorf("Expected mapsList.results[0] to be %v, but got %v", location, mapsList.Locations[0])
	}
}
