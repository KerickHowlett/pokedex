package pokedex_command

import (
	"fmt"
	"reflect"
	"testing"

	bpc "bills_pc"
	p "pokemon"
	f "test_tools/fixtures"
	"test_tools/utils"
)

func TestPokedexCommand_Execute(t *testing.T) {
	const (
		NoPokemon   = "no-pokemon"
		PokemonList = "pokemon-list"
	)

	setup := func(responseType string) (actual string, expected string) {
		const charmander = "Charmander"
		command := &PokedexCommand{
			pc:                     bpc.NewBillsPC(),
			pokedexTitle:           "Your Pokedex:",
			noCaughtPokemonMessage: "You haven't caught any Pokemon yet!",
		}

		switch responseType {
		case NoPokemon:
			command.pc = bpc.NewBillsPC()
			expected = fmt.Sprintf("%s\n", command.noCaughtPokemonMessage)
		case PokemonList:
			command.pc = bpc.NewBillsPC(
				bpc.WithCaughtPokemon(&p.Pokemon{Name: f.PokemonName}),
				bpc.WithCaughtPokemon(&p.Pokemon{Name: charmander}),
			)
			expected = fmt.Sprintf("%s\n- %s\n- %s\n", command.pokedexTitle, f.PokemonName, charmander)
		default:
			panic("invalid response type")
		}

		stdout := utils.NewPrintStorage()
		actual, _ = stdout.CaptureWithError(command.Execute)

		return actual, expected
	}

	t.Run("Given the user has not caught any Pokemon", func(t *testing.T) {
		t.Run("Then it should print a message saying the user has not caught any Pokemon", func(t *testing.T) {
			if actual, expected := setup(NoPokemon); actual != expected {
				t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
			}
		})
	})

	t.Run("Given the user has caught some Pokemon", func(t *testing.T) {
		t.Run("Then it should print the list of caught Pokemon", func(t *testing.T) {
			if actual, expected := setup(PokemonList); actual != expected {
				t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
			}
		})
	})
}

func TestPokedexCommand_GetArgs(t *testing.T) {
	t.Run("should return an empty slice.", func(t *testing.T) {
		command := &PokedexCommand{}
		if reflect.DeepEqual(command.GetArgs(), []string{}) == false {
			t.Error("Expected GetArgs to return an empty slice, but it didn't.")
		}
	})
}

func TestPokedexCommand_GetDescription(t *testing.T) {
	setup := func() (actual string, expected string) {
		expected = "Show help for the Pokemon CLI commands."
		command := &PokedexCommand{description: expected}
		actual = command.GetDescription()
		return actual, expected
	}
	t.Run("should return the description of the PokedexCommand.", func(t *testing.T) {
		if actual, expected := setup(); actual != expected {
			t.Errorf("Expected GetDescription to return '%s', but got '%s'", expected, actual)
		}
	})
}

func TestPokedexCommand_GetName(t *testing.T) {
	setup := func() (actual string, expected string) {
		expected = "help"
		command := &PokedexCommand{name: expected}
		actual = command.GetName()
		return actual, expected
	}
	t.Run("should return the name of the PokedexCommand.", func(t *testing.T) {
		if actual, expected := setup(); actual != expected {
			t.Errorf("Expected GetName to return '%s', but got '%s'", expected, actual)
		}
	})
}

func TestPokedexCommand_PrintHelp(t *testing.T) {
	stdout := utils.NewPrintStorage()
	setup := func() string {
		command := &PokedexCommand{description: "foobar", name: "foobar"}
		command.PrintHelp()
		return stdout.Capture(command.PrintHelp)

	}
	t.Run("should call the PrintHelp function.", func(t *testing.T) {
		if output := setup(); output == "" {
			t.Error("Expected PrintHelp to be called, but it wasn't.")
		}
	})
}

func TestPokedexCommand_SetArgs(t *testing.T) {
	setup := func() []string {
		command := &PokedexCommand{}
		command.SetArgs([]string{"foo", "bar"})
		actual := command.GetArgs()
		return actual
	}
	t.Run("should do nothing.", func(t *testing.T) {
		if actual := setup(); !reflect.DeepEqual(actual, []string{}) {
			t.Error("Expected SetArgs to do nothing, but it didn't.")
		}
	})
}
