package inspect_command

import (
	bpc "bills_pc"
	"fmt"
	"reflect"
	"testing"

	el "entity_link"
	p "pokemon"
	ps "pokemon_stat"
	pt "pokemon_type"
	f "test_tools/fixtures"
	"test_tools/utils"
)

func TestInspectCommand_Execute(t *testing.T) {
	const (
		InvalidArgs     = "invalid-args"
		NoCaughtPokemon = "no-caught-pokemon"
		PokemonNotFound = "pokemon-not-found"
		Success         = "success"
	)

	runInspectCommand_ExecuteTest := func(responseType string) (output string, expectedOutput string, err error) {
		pokemon := &p.Pokemon{
			Name:   f.PokemonName,
			Height: 4,
			Weight: 60,
			Types: []*pt.PokemonType{{
				Type: &el.EntityLink{
					Name: "Electric",
				},
			}},
			Stats: []*ps.PokemonStat{{
				BaseStat: 35,
				Stat:     &el.EntityLink{Name: "hp"},
			}},
		}

		command := &InspectCommand{
			args:                        []string{},
			noCaughtPokemonMessage:      "No Caught Pokemon",
			invalidArgsErrorMessage:     "Invalid Args",
			pokemonNotFoundErrorMessage: "Pokemon Not Found",
			pc:                          bpc.NewBillsPC(),
		}

		switch responseType {
		case NoCaughtPokemon:
			expectedOutput = fmt.Sprintf("%s\n", command.noCaughtPokemonMessage)
		case InvalidArgs:
			command.pc.Deposit(pokemon)
			expectedOutput = command.invalidArgsErrorMessage
		case PokemonNotFound:
			command.args = []string{f.Invalid}
			command.pc.Deposit(pokemon)
			expectedOutput = fmt.Sprintf("%s %s", command.pokemonNotFoundErrorMessage, f.Invalid)
		case Success:
			command.args = []string{f.PokemonName}
			command.pc.Deposit(pokemon)
			expectedOutput = fmt.Sprintf("Name: %s\nHeight: %d\nWeight: %d\n", pokemon.Name, pokemon.Height, pokemon.Weight)
			expectedOutput += fmt.Sprintf("Stats:\n - %s: %d\n", pokemon.Stats[0].Stat.Name, pokemon.Stats[0].BaseStat)
			expectedOutput += fmt.Sprintf("Types:\n - %s\n", pokemon.Types[0].Type.Name)
		default:
			panic("invalid response type")
		}

		stdout := utils.NewPrintStorage()
		output, err = stdout.CaptureWithError(command.Execute)

		return output, expectedOutput, err
	}

	t.Run("Given the user has not caught any Pokemon", func(t *testing.T) {
		t.Parallel()
		t.Run("Then it should return a nil error value", func(t *testing.T) {
			if _, _, err := runInspectCommand_ExecuteTest(NoCaughtPokemon); err != nil {
				t.Errorf("Expected error to be nil, but instead got: %v", err)
			}
		})

		t.Run("And it should let the user know they have not caught any Pokemon", func(t *testing.T) {
			if actual, expected, _ := runInspectCommand_ExecuteTest(NoCaughtPokemon); actual != expected {
				t.Errorf("Expected: %s, but instead got: %s", expected, actual)
			}
		})
	})

	t.Run("Given the user provides invalid arguments", func(t *testing.T) {
		t.Run("Then it should return a non-nil error value", func(t *testing.T) {
			if _, _, err := runInspectCommand_ExecuteTest(InvalidArgs); err == nil {
				t.Errorf("Expected error to be non-nil, but instead got: %v", err)
			}
		})

		t.Run("And it should let the user know they provided invalid arguments", func(t *testing.T) {
			if _, expected, err := runInspectCommand_ExecuteTest(InvalidArgs); err.Error() != expected {
				t.Errorf("Expected: %s, but instead got: %s", expected, err.Error())
			}
		})
	})

	t.Run("Given the user tries to inspect a Pokemon they haven't caught", func(t *testing.T) {
		t.Run("Then it should return a non-nil error value", func(t *testing.T) {
			if _, _, err := runInspectCommand_ExecuteTest(PokemonNotFound); err == nil {
				t.Errorf("Expected error to be non-nil, but instead got: %v", err)
			}
		})

		t.Run("And it should let the user know the Pokemon was not found", func(t *testing.T) {
			if _, expected, err := runInspectCommand_ExecuteTest(PokemonNotFound); err.Error() != expected {
				t.Errorf("Expected: %s, but instead got: %s", expected, err.Error())
			}
		})
	})

	t.Run("Given the user tries to inspect a Pokemon they have caught", func(t *testing.T) {
		t.Run("Then it should return a nil error value", func(t *testing.T) {
			if _, _, err := runInspectCommand_ExecuteTest(Success); err != nil {
				t.Errorf("Expected error to be nil, but instead got: %v", err)
			}
		})

		t.Run("Then it should display the Pokemon's information", func(t *testing.T) {
			if output, expectedOutput, _ := runInspectCommand_ExecuteTest(Success); output != expectedOutput {
				t.Errorf("\nExpected:\n\n%s\nActual:\n\n%s\n", expectedOutput, output)
			}
		})
	})
}

func TestInspectCommand_GetArgs(t *testing.T) {
	setup := func() (actual string, expected string) {
		expected = f.PokemonName
		command := &InspectCommand{args: []string{expected}}
		actual = command.GetArgs()[0]
		return actual, expected
	}

	t.Run("should return the arguments of the InspectCommand", func(t *testing.T) {
		if actual, expected := setup(); actual != expected {
			t.Errorf("Expected: %s, but instead got: %s", expected, actual)
		}
	})
}

func TestInspectCommand_GetDescription(t *testing.T) {
	setup := func() (actual string, expected string) {
		expected = "inspect description"
		command := &InspectCommand{description: expected}
		actual = command.GetDescription()
		return actual, expected
	}
	t.Run("should return the command's description", func(t *testing.T) {
		if actual, expected := setup(); actual != expected {
			t.Errorf("Expected: %s, but instead got: %s", expected, actual)
		}
	})
}

func TestInspectCommand_GetName(t *testing.T) {
	setup := func() (actual string, expected string) {
		expected = "inspect"
		command := &InspectCommand{name: expected}
		actual = command.GetName()
		return actual, expected
	}
	t.Run("should return the command's name", func(t *testing.T) {
		if actual, expected := setup(); actual != expected {
			t.Errorf("Expected: %s, but instead got: %s", expected, actual)
		}
	})
}

func TestInspectCommand_PrintHelp(t *testing.T) {
	stdout := utils.NewPrintStorage()
	setup := func() string {
		command := &InspectCommand{description: "foobar", name: "foobar"}
		command.PrintHelp()
		return stdout.Capture(command.PrintHelp)

	}
	t.Run("should call the PrintHelp function.", func(t *testing.T) {
		if output := setup(); output == "" {
			t.Error("Expected PrintHelp to be called, but it wasn't.")
		}
	})
}

func TestInspectCommand_SetArgs(t *testing.T) {
	setup := func() (actual []string, expected []string) {
		expected = []string{"foo", "bar"}
		command := &InspectCommand{}
		command.SetArgs(expected)
		actual = command.args
		return actual, expected
	}
	t.Run("should do nothing.", func(t *testing.T) {
		if actual, expected := setup(); !reflect.DeepEqual(actual, expected) {
			t.Errorf("Expected '%s', but instead got: '%s'", expected, actual)
		}
	})
}

func TestHasValidArgs(t *testing.T) {
	const (
		// emptyArgs setups a scenario where the InspectCommand.args[0] field contains an empty string.
		emptyArgs = ""
		// nilArgs setups a scenario where the InspectCommand.args field is nil.
		nilArgs = "nil"
		// validArgs setups a scenario where the InspectCommand.args[0] field contains a valid (non-empty) string.
		validArgs = "mock"
	)

	setup := func(arg string) *InspectCommand {
		if arg == nilArgs {
			return &InspectCommand{}
		}
		return &InspectCommand{args: []string{arg}}
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
