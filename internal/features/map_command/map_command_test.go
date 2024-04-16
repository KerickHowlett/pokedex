package map_command

import (
	"fmt"
	"testing"
	"time"

	l "entities/location"
	ms "map_command/state"
	f "test_tools/fixtures"
	"test_tools/utils"
)

func TestMapCommand_Execute(t *testing.T) {
	// Define the response types for the test cases.
	const (
		// Empty represents the response type for when the fetchLocations method returns an empty list of locations.
		Empty = "empty"
		// Error represents the response type for when the fetchLocations method returns an error.
		Error = "error"
		// NilURL represents the response type for when the MapCommand.NextURL field is nil.
		NilURL = "nil-url"
		// Success represents the response type for when the fetchLocations method executes successfully.
		Success = "success"
	)

	runMapsExecuteTest := func(responseType string) (command *MapCommand, output string, expected string, err error) {
		expected = ""

		command = &MapCommand{
			listMarker:              " -",
			listTitle:               "Locations",
			noMoreMapsMessage:       "No maps were found.",
			noMapsFoundErrorMessage: "No locations were found.",
			paginationDirection:     Next,
			state:                   ms.NewMapsState(ms.WithNextURL(&f.APIEndpoint)),
		}

		switch responseType {
		case Empty:
			expected = command.noMapsFoundErrorMessage
			command.fetchLocations = func(url string, state *ms.MapsState, cacheTTL ...time.Duration) error {
				command.state.Locations = []l.Location{}
				return nil
			}
		case Error:
			expected = "error fetching locations"
			command.fetchLocations = func(url string, state *ms.MapsState, cacheTTL ...time.Duration) error {
				return fmt.Errorf(expected)
			}
		case NilURL:
			expected = fmt.Sprintf("%s\n", command.noMoreMapsMessage)
			command.state.NextURL = nil
		case Success:
			expected = fmt.Sprintf("%s\n - %s\n", command.listTitle, f.StarterTown)
			command.fetchLocations = func(url string, state *ms.MapsState, cacheTTL ...time.Duration) error {
				state.Locations = append(state.Locations, l.Location{Name: f.StarterTown})
				return nil
			}
		default:
			panic("invalid response type")
		}

		stdout := utils.NewPrintStorage()
		output, err = stdout.CaptureWithError(command.Execute)

		return command, output, expected, err
	}

	t.Run("Given a nil Maps API Endpoint", func(t *testing.T) {
		t.Parallel()
		t.Run("should return nil error value", func(t *testing.T) {
			t.Parallel()
			if _, _, _, err := runMapsExecuteTest(NilURL); err != nil {
				t.Errorf("expected err to be nil, got %v", err)
			}
		})

		t.Run("should print the set 'noMoreMapsMessage' string value to stdout.", func(t *testing.T) {
			t.Parallel()
			if _, output, expected, _ := runMapsExecuteTest(NilURL); output != expected {
				t.Errorf("expected stdout to be %q, got %q", expected, output)
			}
		})
	})

	t.Run("Given the fetchLocations function returns with a non-nil error", func(t *testing.T) {
		t.Parallel()
		t.Run("should return a non-nil error value.", func(t *testing.T) {
			t.Parallel()
			if _, _, _, err := runMapsExecuteTest(Error); err == nil {
				t.Error("expected err to be non-nil")
			}
		})

		t.Run("should return the correct error message.", func(t *testing.T) {
			t.Parallel()
			_, _, expected, err := runMapsExecuteTest(Error)
			if errorMessage := err.Error(); errorMessage != expected {
				t.Errorf("expected error to be %q, got %q", expected, errorMessage)
			}
		})
	})

	t.Run("Given the post-fetched state has no locations", func(t *testing.T) {
		t.Parallel()
		t.Run("should return the correct error message.", func(t *testing.T) {
			_, _, expected, err := runMapsExecuteTest(Empty)
			if errorMessage := err.Error(); errorMessage != expected {
				t.Errorf("expected error to be %q, got %q", expected, errorMessage)
			}
		})
	})

	t.Run("Given the function executes successfully", func(t *testing.T) {
		t.Parallel()
		t.Run("should return nil error value", func(t *testing.T) {
			t.Parallel()
			if _, _, _, err := runMapsExecuteTest(Success); err != nil {
				t.Errorf("expected err to be a non-nil value, but instead got %v", err)
			}
		})

		t.Run("should print the Pokemon maps/locations results to stdout.", func(t *testing.T) {
			t.Parallel()
			if _, output, expected, _ := runMapsExecuteTest(Success); output != expected {
				t.Errorf("expected stdout to be %q, got %q", expected, output)
			}
		})
	})
}

func TestGetMapsAPIEndpoint(t *testing.T) {
	setup := func(paginationDirection string) *MapCommand {
		nextURL := f.APIEndpoint + "/" + Next
		previousURL := f.APIEndpoint + "/" + Previous

		return &MapCommand{
			paginationDirection: paginationDirection,
			state: ms.NewMapsState(
				ms.WithNextURL(&nextURL),
				ms.WithPreviousURL(&previousURL),
			),
		}
	}

	t.Run("Given a 'pagination direction' argument of 'next'", func(t *testing.T) {
		t.Parallel()
		t.Run("should return the field value for MapCommand.NextURL", func(t *testing.T) {
			command := setup(Next)
			if url := command.getMapsAPIEndpoint(); url != command.state.NextURL {
				t.Errorf("expected url to be %q, got %s", f.APIEndpoint, *url)
			}
		})
	})

	t.Run("Given a 'pagination direction' argument of 'previous'", func(t *testing.T) {
		t.Parallel()
		t.Run("should return the field value for MapCommand.PreviousURL", func(t *testing.T) {
			command := setup(Previous)
			if url := command.getMapsAPIEndpoint(); url != command.state.PreviousURL {
				t.Errorf("expected url to be %q, got %s", f.APIEndpoint, *url)
			}
		})
	})

	t.Run("Given an invalid 'pagination direction' argument", func(t *testing.T) {
		t.Parallel()
		t.Run("should return nil", func(t *testing.T) {
			command := setup(f.Invalid)
			utils.ExpectPanic(t, command.getMapsAPIEndpoint)
		})
	})
}
