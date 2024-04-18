package pokedex_command

import "testing"

func TestWithDescription(t *testing.T) {
	runWithDescriptionTest := func() (actual string, expected string) {
		command := &PokedexCommand{}
		expected = "Help Description."
		WithDescription(expected)(command)
		return command.description, expected

	}

	t.Run("should set the description of the PokedexCommand", func(t *testing.T) {
		if actual, expected := runWithDescriptionTest(); actual != expected {
			t.Errorf("Expected description to be '%s', but got '%s'", expected, actual)
		}
	})
}

func TestWithName(t *testing.T) {
	runWithNameTest := func() (actual string, expected string) {
		command := &PokedexCommand{}
		expected = "help-name"
		WithName(expected)(command)
		return command.name, expected

	}

	t.Run("should set the name of the PokedexCommand", func(t *testing.T) {
		if actual, expected := runWithNameTest(); actual != expected {
			t.Errorf("Expected name to be '%s', but got '%s'", expected, actual)
		}
	})
}

func TestWithNoCaughtPokemonMessage(t *testing.T) {
	runWithNoCaughtPokemonMessage := func() (command *PokedexCommand, noCaughtPokemonMessage string) {
		command = &PokedexCommand{}
		noCaughtPokemonMessage = "No Pokemon"
		WithNoCaughtPokemonMessage(noCaughtPokemonMessage)(command)
		return command, noCaughtPokemonMessage
	}

	t.Run("should set the noCaughtPokemonMessage field", func(t *testing.T) {
		if command, noCaughtPokemonMessage := runWithNoCaughtPokemonMessage(); command.noCaughtPokemonMessage != noCaughtPokemonMessage {
			t.Errorf("expected noCaughtPokemonMessage to be %s, got %s", noCaughtPokemonMessage, command.noCaughtPokemonMessage)
		}
	})
}

func TestWithTitle(t *testing.T) {
	runWithTitleTest := func() (actual string, expected string) {
		expected = "Pokedex Title"
		command := &PokedexCommand{}
		WithPokedexTitle(expected)(command)
		actual = command.pokedexTitle
		return actual, expected
	}
	t.Run("should set the title for the help message.", func(t *testing.T) {
		if actual, expected := runWithTitleTest(); actual != expected {
			t.Errorf("Expected title to be %q but got %q", expected, actual)
		}
	})
}
