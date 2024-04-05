module github.com/KerickHowlett/pokedexcli

go 1.22.1

require (
	command v1.0.0
	maps v1.0.0
	query v1.0.0
	repl v1.0.0
	toochain v1.0.0
	testtools v1.0.0 // indirect
)

replace (
	pokedex => ./cmd/commands
	command => ./internal/ui/command
	maps => ./internal/feature/maps
	query => ./internal/utils/query
	repl => ./internal/ui/repl
	toochain => ./internal/ui/toolchain
	testtools => ./internal/testtools
)
