package mapslist

import (
	"fmt"
	"testing"

	utils "internal/tests/utils"

	l "github.com/KerickHowlett/pokedexcli/cmd/location"
)

var stdout = utils.NewPrintStorage()
var starterTown = "Pallet Town"

func TestPrintLocations_NoLocationsFound(t *testing.T) {
	mapsList := NewMapsList()

	if _, err := stdout.CaptureWithError(mapsList.PrintLocations); err == nil {
		t.Errorf("Expected error, but got nil")
	}
}

func TestPrintLocations_NoMapsFound_ErrorMessage(t *testing.T) {
	const expectedErr = "no maps found"
	mapsList := NewMapsList()

	if _, err := stdout.CaptureWithError(mapsList.PrintLocations); err.Error() != expectedErr {
		t.Errorf("Expected error message '%s', but got '%s'", expectedErr, err.Error())
	}
}

func TestPrintLocations_WithoutError(t *testing.T) {
	mapsList, _ := setupSuccessfulMapsListTest(starterTown)

	_, err := stdout.CaptureWithError(mapsList.PrintLocations)

	if err != nil {
		t.Errorf("Expected no error, but got '%s'", err.Error())
	}
}

func TestPrintLocations_Success(t *testing.T) {
	mapsList, expectedOutput := setupSuccessfulMapsListTest(starterTown)

	output, _ := stdout.CaptureWithError(mapsList.PrintLocations)

	if output != expectedOutput {
		t.Errorf("Expected output to be %q, but got %q", expectedOutput, output)
	}
}

func setupSuccessfulMapsListTest(locationName string) (mapsList *MapsList, expectedLocation string) {
	location := l.NewLocation(l.WithName(locationName))

	mapsList = NewMapsList(AddLocation(*location))
	expectedLocation = fmt.Sprintf("%s\n", starterTown)

	return mapsList, expectedLocation
}
