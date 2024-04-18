package inspect_command

import (
	"testing"

	bpc "bills_pc"
	f "test_tools/fixtures"
)

func TestWithDescription(t *testing.T) {
	runWithDescription := func() (command *InspectCommand, description string) {
		command = &InspectCommand{}
		description = "Inspect Description"
		WithDescription(description)(command)
		return command, description
	}

	t.Run("should set the description field", func(t *testing.T) {
		if command, description := runWithDescription(); command.description != description {
			t.Errorf("expected description to be %s, got %s", description, command.description)
		}
	})
}

func TestWithInvalidArgsErrorMessage(t *testing.T) {
	runWithInvalidArgsErrorMessage := func() (command *InspectCommand, invalidArgsErrorMessage string) {
		command = &InspectCommand{}
		invalidArgsErrorMessage = "Invalid Arguments"
		WithInvalidArgsErrorMessage(invalidArgsErrorMessage)(command)
		return command, invalidArgsErrorMessage
	}

	t.Run("should set the invalidArgsErrorMessage field", func(t *testing.T) {
		if command, invalidArgsErrorMessage := runWithInvalidArgsErrorMessage(); command.invalidArgsErrorMessage != invalidArgsErrorMessage {
			t.Errorf("expected invalidArgsErrorMessage to be %s, got %s", invalidArgsErrorMessage, command.invalidArgsErrorMessage)
		}
	})
}

func TestWithName(t *testing.T) {
	runWithName := func() (command *InspectCommand) {
		command = &InspectCommand{}
		WithName(f.PokemonName)(command)
		return command
	}

	t.Run("should set the name field", func(t *testing.T) {
		if command := runWithName(); command.name != f.PokemonName {
			t.Errorf("expected name to be %s, got %s", f.PokemonName, command.name)
		}
	})
}

func TestWithNoCaughtPokemonMessage(t *testing.T) {
	runWithNoCaughtPokemonMessage := func() (command *InspectCommand, noCaughtPokemonMessage string) {
		command = &InspectCommand{}
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

func TestWithPC(t *testing.T) {
	runWithPC := func() (command *InspectCommand, pc *bpc.BillsPC) {
		command = &InspectCommand{}
		pc = &bpc.BillsPC{}
		WithPC(pc)(command)
		return command, pc
	}

	t.Run("should set the pc field", func(t *testing.T) {
		if command, pc := runWithPC(); command.pc != pc {
			t.Errorf("expected pc to be %v, got %v", pc, command.pc)
		}
	})
}

func TestWithPokemonNotFoundErrorMessage(t *testing.T) {
	runWithPokemonNotFoundErrorMessage := func() (command *InspectCommand, pokemonNotFoundErrorMessage string) {
		command = &InspectCommand{}
		pokemonNotFoundErrorMessage = "Pokemon not found"
		WithPokemonNotFoundErrorMessage(pokemonNotFoundErrorMessage)(command)
		return command, pokemonNotFoundErrorMessage
	}

	t.Run("should set the pokemonNotFoundErrorMessage field", func(t *testing.T) {
		if command, pokemonNotFoundErrorMessage := runWithPokemonNotFoundErrorMessage(); command.pokemonNotFoundErrorMessage != pokemonNotFoundErrorMessage {
			t.Errorf("expected pokemonNotFoundErrorMessage to be %s, got %s", pokemonNotFoundErrorMessage, command.pokemonNotFoundErrorMessage)
		}
	})
}
