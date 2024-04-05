module github.com/KerickHowlett/pokedexcli

go 1.22.1

require (
	command v1.0.0 // indirect
	maps v1.0.0 // indirect
	query v1.0.0 // indirect
	repl v1.0.0 // indirect
	shell v1.0.0
	testtools v1.0.0 // indirect
	toochain v1.0.0 // indirect
)

replace (
	command => ./internal/ui/command
	maps => ./internal/features/maps
	query => ./internal/utils/query
	repl => ./internal/ui/repl
	shell => ./internal/shell
	testtools => ./internal/testtools
	toochain => ./internal/ui/toolchain
)
