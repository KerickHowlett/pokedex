package mapslist

import (
	"testing"

	l "github.com/KerickHowlett/pokedexcli/cmd/location"
)

func TestNewMapsList(t *testing.T) {
	maps := NewMapsList()

	expectURLToBeNil(t, maps.NextURL, "maps.next")
	expectURLToBeNil(t, maps.PreviousURL, "maps.previous")
	expectResultsToBeEmpty(t, maps.Locations)
}

func expectURLToBeNil(t *testing.T, url *string, label string) {
	if url == nil {
		return
	}
	t.Errorf("Expected %s to be nil, but got non-nil", label)
}

func expectResultsToBeEmpty(t *testing.T, results []l.Location) {
	if len(results) == 0 {
		return
	}
	t.Errorf("Expected results to be empty, but got %d elements", len(results))
}
