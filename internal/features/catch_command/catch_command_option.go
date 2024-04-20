package catch_command

import (
	"time"

	bpc "bills_pc"
)

// CatchCommandOption is a type alias for a function that sets a field in the CatchCommand struct.
type CatchCommandOption func(*CatchCommand)

// WithAPIEndpoint sets the apiEndpoint field in the CatchCommand struct.
//
// Parameters:
//   - apiEndpoint: The URL of the API endpoint to fetch the location area data.
//
// Returns:
//   - CatchCommandOption: A function that sets the apiEndpoint field in the CatchCommand struct.
//
// Example usage:
//
//	c := catch_command.NewCatchCommand(
//	  catch_command.WithAPIEndpoint("https://pokeapi.co/api/v2/location-area/"),
//	)
func WithAPIEndpoint(apiEndpoint string) CatchCommandOption {
	return func(c *CatchCommand) {
		c.apiEndpoint = apiEndpoint
	}
}

// WithCacheTTL sets the cacheTTL field in the CatchCommand struct.
//
// Parameters:
//   - cacheTTL: The time-to-live (TTL) determines how long FetchEncounters' cached responses get to exist before being discarded.
//
// Returns:
//   - CatchCommandOption: A function that sets the cacheTTL field in the CatchCommand struct.
//
// Example usage:
//
//	c := catch_command.NewCatchCommand(
//	  catch_command.WithCacheTTL(time.Minute * 5),
//	)
func WithCacheTTL(cacheTTL time.Duration) CatchCommandOption {
	return func(c *CatchCommand) {
		c.cacheTTL = cacheTTL
	}
}

// WithCatchPokemon sets the catchPokemon field in the CatchCommand struct.
//
// Parameters:
//   - catchPokemon: A function that fetches locations/maps from the Pokemon API.
//
// Returns:
//   - CatchCommandOption: A function that sets the catchPokemon field in the CatchCommand struct.
//
// Example usage:
//
//	c := catch_command.NewCatchCommand(
//	  catch_command.WithCatchPokemon(func() {}),
//	)
func WithCatchPokemon(catchPokemon CatchPokemonFunc) CatchCommandOption {
	return func(c *CatchCommand) {
		c.catchPokemon = catchPokemon
	}
}

// WithCheckYourLuck sets the checkYourLuck field in the CatchCommand struct.
//
// Parameters:
//   - checkYourLuck: A function that returns the chance of catching a Pokemon.
//
// Returns:
//   - CatchCommandOption: A function that sets the checkYourLuck field in the CatchCommand struct.
//
// Example usage:
//
//	c := catch_command.NewCatchCommand(
//	  catch_command.WithCheckYourLuck(func() {}),
//	)
func WithCheckYourLuck(checkYourLuck CheckYourLuckFunc) CatchCommandOption {
	return func(c *CatchCommand) {
		c.checkYourLuck = checkYourLuck
	}
}

// WithDifficulty sets the difficulty field in the CatchCommand struct.
//
// Parameters:
//   - difficulty: The challenge level at catching any given Pokemon.
//
// Returns:
//   - CatchCommandOption: A function that sets the difficulty field in the CatchCommand struct.
//
// Example usage:
//
//	c := catch_command.NewCatchCommand(
//	  catch_command.WithDifficulty(100),
//	)
func WithDifficulty(difficulty int) CatchCommandOption {
	return func(c *CatchCommand) {
		c.difficulty = difficulty
	}
}

// WithDescription sets the description field in the CatchCommand struct.
//
// Parameters:
//   - description: The purpose of the ExploreCommand.
//
// Returns:
//   - CatchCommandOption: A function that sets the description field in the CatchCommand struct.
//
// Example usage:
//
//	c := catch_command.NewCatchCommand(
//	  catch_command.WithDescription("Catch a Pokemon."),
//	)
func WithDescription(description string) CatchCommandOption {
	return func(c *CatchCommand) {
		c.description = description
	}
}

// WithEscapedNotification sets the escapedNotification field in the CatchCommand struct.
//
// Parameters:
//   - escapedNotification: A message to display when a Pokemon escapes.
//
// Returns:
//   - CatchCommandOption: A function that sets the escapedNotification field in the CatchCommand struct.
//
// Example usage:
//
//	c := catch_command.NewCatchCommand(
//	  catch_command.WithEscapedNotification("The Pokemon escaped!"),
//	)
func WithEscapedNotification(escapedNotification string) CatchCommandOption {
	return func(c *CatchCommand) {
		c.escapedNotification = escapedNotification
	}
}

// WithName sets the name field in the CatchCommand struct.
//
// Parameters:
//   - name: The name of the ExploreCommand.
//
// Returns:
//   - CatchCommandOption: A function that sets the name field in the CatchCommand struct.
//
// Example usage:
//
//	c := catch_command.NewCatchCommand(
//	  catch_command.WithName("catch"),
//	)
func WithName(name string) CatchCommandOption {
	return func(c *CatchCommand) {
		c.name = name
	}
}

// WithNoEnteredArgsErrorMessage sets the name field in the CatchCommand struct.
//
// Parameters:
//   - name: The name of the ExploreCommand.
//
// Returns:
//   - CatchCommandOption: A function that sets the name field in the CatchCommand struct.
//
// Example usage:
//
//	c := catch_command.NewCatchCommand(
//	  catch_command.WithNoEnteredArgsErrorMessage("catch"),
//	)
func WithNoEnteredArgsErrorMessage(noEnteredArgsErrorMessage string) CatchCommandOption {
	return func(c *CatchCommand) {
		c.noEnteredArgsErrorMessage = noEnteredArgsErrorMessage
	}
}

// WithPC sets the pc field in the CatchCommand struct.
//
// Parameters:
//   - pc: All the Pokemon the user has caught.
//
// Returns:
//   - CatchCommandOption: A function that sets the pc field in the CatchCommand struct.
//
// Example usage:
//
//	c := catch_command.NewCatchCommand(
//	  catch_command.WithPC(&bpc.BillsPC{}),
//	)
func WithPC(pc *bpc.BillsPC) CatchCommandOption {
	return func(c *CatchCommand) {
		c.pc = pc
	}
}

// WithSuccessfulCatchNotification sets the successfulCatchNotification field in the CatchCommand struct.
//
// Parameters:
//   - successfulCatchNotification: A message to display when a Pokemon is caught.
//
// Returns:
//   - CatchCommandOption: A function that sets the successfulCatchNotification field in the CatchCommand struct.
//
// Example usage:
//
//	c := catch_command.NewCatchCommand(
//	  catch_command.WithSuccessfulCatchNotification("You caught a Pokemon!"),
//	)
func WithSuccessfulCatchNotification(successfulCatchNotification string) CatchCommandOption {
	return func(c *CatchCommand) {
		c.successfulCatchNotification = successfulCatchNotification
	}
}

// WithThrowBallNotification sets the throwBallNotification field in the CatchCommand struct.
//
// Parameters:
//   - throwBallNotification: A message to display when a Pokemon is thrown a ball.
//
// Returns:
//   - CatchCommandOption: A function that sets the throwBallNotification field in the CatchCommand struct.
//
// Example usage:
//
//	c := catch_command.NewCatchCommand(
//	  catch_command.WithThrowBallNotification("You threw a ball!"),
//	)
func WithThrowBallNotification(throwBallNotification string) CatchCommandOption {
	return func(c *CatchCommand) {
		c.throwBallNotification = throwBallNotification
	}
}
