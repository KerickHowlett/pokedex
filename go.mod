module github.com/KerickHowlett/pokedexcli

go 1.22.1

require (
	command v1.0.0 // indirect
	entities v1.0.0 // indirect
	exit v1.0.0 // indirect
	explore v1.0.0 // indirect
	help v1.0.0 // indirect
	map v1.0.0 // indirect
	query v1.0.0 // indirect
	repl v1.0.0 // indirect
	shell v1.0.0
	test_tools v1.0.0 // indirect
	toochain v1.0.0 // indirect
)

replace (
	command => ./internal/ui/command
	entities => ./internal/entities
	exit => ./internal/features/exit_command
	explore => ./internal/features/explore_command
	help => ./internal/features/help_command
	map => ./internal/features/map_command
	query => ./internal/query
	repl => ./internal/ui/repl
	shell => ./internal/shell
	system => ./internal/features/system
	test_tools => ./internal/test_tools
	toochain => ./internal/ui/toolchain
)
