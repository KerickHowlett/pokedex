package explore_command

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	la "location_area"
	p "pokemon"
	pe "pokemon_encounter"
	f "test_tools/fixtures"
	"test_tools/utils"
)

func TestExploreCommand_GetArgs(t *testing.T) {
	t.Run("should return the arguments of the ExploreCommand.", func(t *testing.T) {
		expected := []string{"mock", "mock2"}
		command := &ExploreCommand{args: expected}
		if args := command.GetArgs(); reflect.DeepEqual(args, expected) {
			t.Error("expected args to be non-empty, but got an empty slice")
		}
	})
}

func TestExploreCommand_GetDescription(t *testing.T) {
	t.Run("should return the description of the ExploreCommand.", func(t *testing.T) {
		const expected = "mock description"
		command := &ExploreCommand{description: expected}
		if description := command.GetDescription(); description != expected {
			t.Errorf("expected description to be %q, but got %q", expected, description)
		}
	})
}

func TestExploreCommand_Execute(t *testing.T) {
	// Define the response types for the test cases.
	const (
		// Empty represents the response type for when the fetchEncounters method returns an empty list of locations.
		Empty = "empty"
		// Error represents the response type for when the fetchEncounters method returns an error.
		Error = "error"
		// NoArgs represents the response type for when the ExploreCommand.args field is empty.
		NoArgs = "noArgs"
		// Success represents the response type for when the fetchEncounters method executes successfully.
		Success = "success"
	)

	runMapsExecuteTest := func(responseType string) (command *ExploreCommand, output string, expected string, err error) {
		expected = ""

		command = &ExploreCommand{
			apiEndpoint: f.APIEndpoint,
			args:        []string{"mock"},
			listMarker:  " -",
			listTitle:   "Pokemon Encounters",
			state:       la.NewLocationArea(),
		}

		switch responseType {
		case Empty:
			expected = command.noEncountersFoundErrorMessage
			command.fetchEncounters = func(url string, state *la.LocationArea, cacheTTL ...time.Duration) error {
				command.state.Encounters = []pe.PokemonEncounter{}
				return nil
			}
		case Error:
			expected = "error fetching location area encounters"
			command.fetchEncounters = func(url string, state *la.LocationArea, cacheTTL ...time.Duration) error {
				return fmt.Errorf(expected)
			}
		case NoArgs:
			expected = command.noEnteredArgsErrorMessage
			command.args = []string{}
		case Success:
			expected = fmt.Sprintf("%s\n - %s\n", command.listTitle, f.PokemonName)
			command.fetchEncounters = func(url string, state *la.LocationArea, cacheTTL ...time.Duration) error {
				command.state.Encounters = []pe.PokemonEncounter{
					{Pokemon: &p.Pokemon{Name: f.PokemonName}},
				}
				return nil
			}
		default:
			panic("invalid response type")
		}

		stdout := utils.NewPrintStorage()
		output, err = stdout.CaptureWithError(command.Execute)

		return command, output, expected, err
	}

	t.Run("Given no arguments are provided", func(t *testing.T) {
		t.Parallel()
		t.Run("should return the correct error message.", func(t *testing.T) {
			_, _, expected, err := runMapsExecuteTest(NoArgs)
			if errorMessage := err.Error(); errorMessage != expected {
				t.Errorf("expected error to be %q, got %q", expected, errorMessage)
			}
		})
	})

	t.Run("Given the fetchEncounters function returns with a non-nil error", func(t *testing.T) {
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
				t.Errorf("expected stdout to be %q, but instead got %q", expected, output)
			}
		})
	})
}

func TestExploreCommand_GetName(t *testing.T) {
	t.Run("should return the name of the ExploreCommand.", func(t *testing.T) {
		const expected = "explore"
		command := &ExploreCommand{name: expected}
		if name := command.GetName(); name != expected {
			t.Errorf("expected name to be %q, but got %q", expected, name)
		}
	})
}

func TestExploreCommand_PrintHelp(t *testing.T) {
	stdout := utils.NewPrintStorage()
	t.Run("should call the PrintHelp function.", func(t *testing.T) {
		command := &ExploreCommand{
			description: "mock description",
			name:        "mock",
		}
		if output := stdout.Capture(command.PrintHelp); output == "" {
			t.Error("Expected PrintHelp to be called, but it wasn't.")
		}
	})
}

func TestHasValidArgs(t *testing.T) {
	// Options for the setup function.
	const (
		// emptyArgs setups a scenario where the ExploreCommand.args[0] field contains an empty string.
		emptyArgs = ""
		// nilArgs setups a scenario where the ExploreCommand.args field is nil.
		nilArgs = "nil"
		// validArgs setups a scenario where the ExploreCommand.args[0] field contains a valid (non-empty) string.
		validArgs = "mock"
	)

	setup := func(arg string) *ExploreCommand {
		if arg == nilArgs {
			return &ExploreCommand{}
		}
		return &ExploreCommand{args: []string{arg}}
	}

	t.Run("should return true if the command has valid arguments.", func(t *testing.T) {
		if command := setup(validArgs); !command.hasValidArgs() {
			t.Error("Expected HasValidArgs to return true, but got false.")
		}
	})

	t.Run("should return false if the command has no arguments.", func(t *testing.T) {
		if command := setup(emptyArgs); command.hasValidArgs() {
			t.Error("Expected HasValidArgs to return false, but got true.")
		}
	})

	t.Run("should return false if the command has nil arguments.", func(t *testing.T) {
		if command := setup(nilArgs); command.hasValidArgs() {
			t.Error("Expected HasValidArgs to return false, but got true.")
		}
	})
}
