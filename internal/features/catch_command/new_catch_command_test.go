package catch_command

import (
	"testing"

	"query_fetch/query_cache/ttl"
	f "test_tools/fixtures"
	"test_tools/utils"
)

func TestNewCatchCommand(t *testing.T) {
	setup := func() *CatchCommand {
		return NewCatchCommand(WithAPIEndpoint(f.APIEndpoint))
	}
	t.Run("should return a new CatchCommand", func(t *testing.T) {
		t.Parallel()
		if command := setup(); command == nil {
			t.Error("expected a new CatchCommand, got nil")
		}
	})

	t.Run("should set the cacheTTL field", func(t *testing.T) {
		t.Parallel()
		if command := setup(); command.cacheTTL != ttl.OneDay {
			t.Errorf("expected cacheTTL to be %d, got %d", ttl.OneDay, command.cacheTTL)
		}
	})

	t.Run("should set the catchPokemon field", func(t *testing.T) {
		t.Parallel()
		// @NOTE: There's no way to test if it's the same function given how its implemented.
		if command := setup(); command.catchPokemon == nil {
			t.Error("expected catchPokemon to be set, got nil")
		}
	})

	t.Run("should set the checkYourLuck field", func(t *testing.T) {
		t.Parallel()
		// @NOTE: There's no way to test if it's the same function given how its implemented.
		if command := setup(); command.checkYourLuck == nil {
			t.Error("expected checkYourLuck to be set, got nil")
		}
	})

	t.Run("should set the description field", func(t *testing.T) {
		t.Parallel()
		if command := setup(); command.description != "Catch That Pokemon! Usage: catch <pokemon-name>" {
			t.Errorf("expected description to be 'Catch That Pokemon! Usage: catch <pokemon-name>', got %s", command.description)
		}
	})

	t.Run("should set the escapedNotification field", func(t *testing.T) {
		t.Parallel()
		if command := setup(); command.escapedNotification != "escaped!" {
			t.Errorf("expected escapedNotification to be 'escaped!', got %s", command.escapedNotification)
		}
	})

	t.Run("should set the name field", func(t *testing.T) {
		t.Parallel()
		if command := setup(); command.name != "catch" {
			t.Errorf("expected name to be 'catch', got %s", command.name)
		}
	})

	t.Run("should set the noEnteredArgsErrorMessage field", func(t *testing.T) {
		t.Parallel()
		if command := setup(); command.noEnteredArgsErrorMessage != "a pokemon name is required" {
			t.Errorf("expected noEnteredArgsErrorMessage to be 'a pokemon name is required', got %s", command.noEnteredArgsErrorMessage)
		}
	})

	t.Run("should set the pc field", func(t *testing.T) {
		t.Parallel()
		// @NOTE: There's no way to test if it's the same function given how its implemented.
		if command := setup(); command.pc == nil {
			t.Error("expected pc to be set, got nil")
		}
	})

	t.Run("should set the successfulCatchNotification field", func(t *testing.T) {
		t.Parallel()
		if command := setup(); command.successfulCatchNotification != "You caught" {
			t.Errorf("expected successfulCatchNotification to be 'You caught', got %s", command.successfulCatchNotification)
		}
	})

	t.Run("should set the throwBallNotification field", func(t *testing.T) {
		t.Parallel()
		if command := setup(); command.throwBallNotification != "You threw a pokeball at" {
			t.Errorf("expected throwBallNotification to be 'You threw a pokeball at', got %s", command.throwBallNotification)
		}
	})

	t.Run("should panic if the apiEndpoint field is not set", func(t *testing.T) {
		utils.ExpectPanic(t, NewCatchCommand)
	})
}
