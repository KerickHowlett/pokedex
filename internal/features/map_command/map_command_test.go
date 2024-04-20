package map_command

import (
	"fmt"
	"testing"

	l "location"
	pd "map_command/pagination_direction"
	ms "map_command/state"
	qec "query_fetch/query_cache/cache_eviction_config"
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
			paginationDirection:     pd.Next,
			state:                   ms.NewMapsState(ms.WithNextURL(&f.APIEndpoint)),
		}

		switch responseType {
		case Empty:
			expected = command.noMapsFoundErrorMessage
			command.fetchLocations = func(url string, config ...*qec.QueryEvictionConfig) (state *ms.MapsState, err error) {
				state = ms.NewMapsState(ms.WithNextURL(&f.APIEndpoint))
				return state, nil
			}
		case Error:
			expected = "error fetching locations"
			command.fetchLocations = func(url string, config ...*qec.QueryEvictionConfig) (state *ms.MapsState, err error) {
				return state, fmt.Errorf(expected)
			}
		case NilURL:
			expected = fmt.Sprintf("%s\n", command.noMoreMapsMessage)
			command.state.NextURL = nil
		case Success:
			expected = fmt.Sprintf("%s\n - %s\n", command.listTitle, f.StarterTown)
			command.fetchLocations = func(url string, config ...*qec.QueryEvictionConfig) (state *ms.MapsState, err error) {
				location := l.Location{Name: f.StarterTown}
				state = ms.NewMapsState(
					ms.WithNextURL(&f.APIEndpoint),
					ms.WithLocation(location),
				)
				return state, nil
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

func TestMapCommand_GetArgs(t *testing.T) {
	t.Run("should return the arguments of the MapCommand.", func(t *testing.T) {
		command := &MapCommand{}
		if args := command.GetArgs(); len(args) != 0 {
			t.Error("expected args to be an empty slice")
		}
	})
}

func TestMapCommand_GetDescription(t *testing.T) {
	t.Run("should return the description of the MapCommand.", func(t *testing.T) {
		command := &MapCommand{description: "Fetches locations from the Pokemon API."}
		if description := command.GetDescription(); description != "Fetches locations from the Pokemon API." {
			t.Errorf("expected description to be %q, got %q", "Fetches locations from the Pokemon API.", description)
		}
	})
}

func TestMapCommand_GetName(t *testing.T) {
	t.Run("should return the name of the MapCommand.", func(t *testing.T) {
		command := &MapCommand{name: "map"}
		if name := command.GetName(); name != "map" {
			t.Errorf("expected name to be %q, got %q", "map", name)
		}
	})
}

func TestMapCommand_PrintHelp(t *testing.T) {
	t.Run("should print the help message for the MapCommand.", func(t *testing.T) {
		stdout := utils.NewPrintStorage()
		t.Run("should call the PrintHelp function.", func(t *testing.T) {
			command := &MapCommand{
				description: "Fetches locations from the Pokemon API.",
				name:        "map",
			}
			if output := stdout.Capture(command.PrintHelp); output == "" {
				t.Error("Expected PrintHelp to be called, but it wasn't.")
			}
		})
	})
}

func TestMapCommand_SetArgs(t *testing.T) {
	t.Run("should do nothing.", func(t *testing.T) {
		command := &MapCommand{}
		command.SetArgs([]string{"arg1", "arg2"})
		if args := command.GetArgs(); len(args) != 0 {
			t.Error("expected args to be an empty slice")
		}
	})
}

func TestMapCommand_getMapsAPIEndpoint(t *testing.T) {
	setup := func(paginationDirection string) *MapCommand {
		nextURL := f.APIEndpoint + "/" + pd.Next
		previousURL := f.APIEndpoint + "/" + pd.Previous

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
			command := setup(pd.Next)
			if url := command.getMapsAPIEndpoint(); url != command.state.NextURL {
				t.Errorf("expected url to be %q, got %s", f.APIEndpoint, *url)
			}
		})
	})

	t.Run("Given a 'pagination direction' argument of 'previous'", func(t *testing.T) {
		t.Parallel()
		t.Run("should return the field value for MapCommand.PreviousURL", func(t *testing.T) {
			command := setup(pd.Previous)
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
