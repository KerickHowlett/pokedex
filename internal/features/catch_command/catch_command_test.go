package catch_command

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
	"time"

	"catch_command/bills_pc"
	p "pokemon"
	f "test_tools/fixtures"
	"test_tools/utils"
)

func TestCatchCommand_Execute(t *testing.T) {
	// Define the response types for the test cases.
	const (
		// Error represents the response type for when the fetchEncounters method returns an error.
		Error = "error"
		// Response for when the pokemon escapes.
		Escape = "escape"
		// NoArgs represents the response type for when the CatchCommand.args field is empty.
		NoArgs = "noArgs"
		// Success represents the response type for when the fetchEncounters method executes successfully.
		Success = "success"
	)

	runMapsExecuteTest := func(responseType string) (command *CatchCommand, output string, expected string, err error) {
		expected = ""

		command = &CatchCommand{
			apiEndpoint:                 f.APIEndpoint,
			args:                        []string{f.PokemonName},
			checkYourLuck:               func(_ int) int { return 100 },
			escapedNotification:         "escaped!",
			noEnteredArgsErrorMessage:   "a pokemon name is required",
			pc:                          bills_pc.NewBillsPC(),
			successfulCatchNotification: "You caught",
			throwBallNotification:       "You threw a pokeball at",
		}

		throwMessage := fmt.Sprintf("%s %s...", command.throwBallNotification, f.PokemonName)
		escapeMessage := fmt.Sprintf("%s %s", f.PokemonName, command.escapedNotification)
		successMessage := fmt.Sprintf("%s %s", command.successfulCatchNotification, f.PokemonName)

		switch responseType {
		case Error:
			expected = "error fetching pokemon"
			command.catchPokemon = func(url string, query *p.Pokemon, cacheTTL ...time.Duration) error {
				return fmt.Errorf(expected)
			}
		case Escape:
			expected = fmt.Sprintf("%s\n%s\n", throwMessage, escapeMessage)
			command.battleOpponent = &p.Pokemon{Name: f.PokemonName, BaseExperience: 1_000_000}
			command.catchPokemon = func(url string, query *p.Pokemon, cacheTTL ...time.Duration) error {
				return nil
			}
		case NoArgs:
			expected = command.noEnteredArgsErrorMessage
			command.args = []string{}
		case Success:
			expected = fmt.Sprintf("%s\n%s!\n", throwMessage, successMessage)
			pokemon := &p.Pokemon{Name: f.PokemonName, BaseExperience: 0}
			command.battleOpponent = pokemon
			command.catchPokemon = func(url string, query *p.Pokemon, cacheTTL ...time.Duration) error {
				return nil
			}
			command.pc.Deposit(pokemon)
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

		t.Run("should display throw ball message to stdout.", func(t *testing.T) {
			t.Parallel()
			command, output, _, _ := runMapsExecuteTest(Error)
			if expected := fmt.Sprintf("%s %s...\n", command.throwBallNotification, f.PokemonName); !strings.Contains(output, expected) {
				t.Errorf("expected stdout to be %q, but instead got %q", expected, output)
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

	t.Run("Given the Pokemon escapes", func(t *testing.T) {
		t.Parallel()
		t.Run("should return a nil error value", func(t *testing.T) {
			t.Parallel()
			if _, _, _, err := runMapsExecuteTest(Escape); err != nil {
				t.Errorf("expected err to be a non-nil value, but instead got %v", err)
			}
		})

		t.Run("should print the Pokemon escape message to stdout.", func(t *testing.T) {
			t.Parallel()
			if _, output, expected, _ := runMapsExecuteTest(Escape); output != expected {
				t.Errorf("expected stdout to be %q, but instead got %q", expected, output)
			}
		})
	})

	t.Run("Given a successful catch", func(t *testing.T) {
		t.Parallel()
		t.Run("should return a nil error value", func(t *testing.T) {
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

func TestCatchCommand_GetArgs(t *testing.T) {
	t.Parallel()
	setup := func() (command *CatchCommand, expected []string) {
		expected = []string{f.PokemonName}
		command = &CatchCommand{args: expected}
		return command, expected
	}

	t.Run("should return the arguments provided to the CatchCommand.", func(t *testing.T) {
		command, expected := setup()
		if args := command.GetArgs(); !reflect.DeepEqual(args, expected) {
			t.Errorf("expected args to be %v, got %v", expected, args)
		}
	})
}

func TestCatchCommand_GetDescription(t *testing.T) {
	t.Parallel()
	setup := func() (command *CatchCommand, expected string) {
		expected = "catch description"
		command = &CatchCommand{description: expected}
		return command, expected
	}

	t.Run("should return the description of the CatchCommand.", func(t *testing.T) {
		command, expected := setup()
		if description := command.GetDescription(); description != expected {
			t.Errorf("expected description to be %q, got %q", expected, description)
		}
	})
}

func TestCatchCommand_GetName(t *testing.T) {
	t.Parallel()
	setup := func() (command *CatchCommand, expected string) {
		expected = "catch"
		command = &CatchCommand{name: expected}
		return command, expected
	}

	t.Run("should return the name of the CatchCommand.", func(t *testing.T) {
		command, expected := setup()
		if name := command.GetName(); name != expected {
			t.Errorf("expected name to be %q, got %q", expected, name)
		}
	})
}

func TestCatchCommand_PrintHelp(t *testing.T) {
	t.Parallel()
	runCatchCommand_PrintHelpTest := func() (output string) {
		stdout := utils.NewPrintStorage()
		command := &CatchCommand{}
		output = stdout.Capture(command.PrintHelp)
		return output
	}

	t.Run("should print the help message for the CatchCommand.", func(t *testing.T) {
		if output := runCatchCommand_PrintHelpTest(); len(output) == 0 {
			t.Error("Expected PrintHelp to be called, but it wasn't.")
		}
	})
}

func TestCatchCommand_SetArgs(t *testing.T) {
	t.Parallel()
	runCatchCommand_SetArgsTest := func() (command *CatchCommand, expected []string) {
		expected = []string{f.PokemonName}
		command = &CatchCommand{}
		command.SetArgs(expected)
		return command, expected
	}

	t.Run("should set the arguments for the CatchCommand.", func(t *testing.T) {
		if command, expected := runCatchCommand_SetArgsTest(); !reflect.DeepEqual(command.args, expected) {
			t.Errorf("expected args to be %v, got %v", expected, command.args)
		}
	})
}

func TestCatchCommand_isCatchSuccessful(t *testing.T) {
	const (
		Botched  = 0
		Critical = 1_000
	)
	runCatchCommand_isCatchSuccessfulTest := func(luck int) (command *CatchCommand, expected bool) {
		switch luck {
		case Botched:
			expected = false
		case Critical:
			expected = true
		default:
			panic("invalid luck value")
		}
		command = &CatchCommand{
			battleOpponent: &p.Pokemon{BaseExperience: 100},
			checkYourLuck:  func(_ int) int { return luck },
		}
		return command, expected
	}

	t.Run("should return TRUE if the CatchCommand is successful in catching a Pokemon.", func(t *testing.T) {
		t.Parallel()
		if command, expected := runCatchCommand_isCatchSuccessfulTest(Critical); command.isCatchSuccessful() != expected {
			t.Errorf("expected isCatchSuccessful to return %t, got %t", expected, !expected)
		}
	})

	t.Run("should return FALSE if the CatchCommand is unsuccessful in catching a Pokemon.", func(t *testing.T) {
		t.Parallel()
		if command, expected := runCatchCommand_isCatchSuccessfulTest(Botched); command.isCatchSuccessful() != expected {
			t.Errorf("expected isCatchSuccessful to return %t, got %t", expected, !expected)
		}
	})
}

func TestHasValidArgs(t *testing.T) {
	// Options for the setup function.
	const (
		// emptyArgs setups a scenario where the CatchCommand.args[0] field contains an empty string.
		emptyArgs = ""
		// nilArgs setups a scenario where the CatchCommand.args field is nil.
		nilArgs = "nil"
		// validArgs setups a scenario where the CatchCommand.args[0] field contains a valid (non-empty) string.
		validArgs = "mock"
	)

	setup := func(arg string) *CatchCommand {
		if arg == nilArgs {
			return &CatchCommand{}
		}
		return &CatchCommand{args: []string{arg}}
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
