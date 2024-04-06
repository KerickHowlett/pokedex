package main

import (
	s "shell"

	"github.com/KerickHowlett/pokedexcli/cmd/cli/config"
)

func main() {
	s.PokedexCLIShell(config.PokedexCLIConfig)
}
