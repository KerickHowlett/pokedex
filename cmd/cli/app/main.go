package main

import (
	s "shell"

	conf "github.com/KerickHowlett/pokedexcli/cmd/cli/config"
)

func main() {
	s.PokedexCLIShell(conf.PokedexCLIConfig)
}
