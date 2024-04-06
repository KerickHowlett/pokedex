package main

import (
	s "shell"

	"github.com/KerickHowlett/pokedexcli/cmd/cli/config"
)

// main is the entry point for the Pokedex CLI application.
//
// Example usage:
//
//	main()
func main() {
	s.PokedexCLIShell(config.PokedexCLIConfig)
}
