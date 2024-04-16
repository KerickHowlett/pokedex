module github.com/KerickHowlett/pokedexcli

go 1.22.1

require (
	command v1.0.0 // indirect
	entities v1.0.0 // indirect
	exit_command v1.0.0 // indirect
	explore_command v1.0.0 // indirect
	help_command v1.0.0 // indirect
	map_command v1.0.0 // indirect
	query_fetch v1.0.0 // indirect
	repl v1.0.0 // indirect
	shell v1.0.0
	test_tools v1.0.0 // indirect
	toochain v1.0.0 // indirect
)

replace (
	command => ./internal/ui/command
	entities => ./internal/entities
	exit_command => ./internal/features/exit_command
	explore_command => ./internal/features/explore_command
	help_command => ./internal/features/help_command
	map_command => ./internal/features/map_command
	query_fetch => ./internal/query_fetch
	repl => ./internal/ui/repl
	shell => ./internal/shell
	system => ./internal/features/system
	test_tools => ./internal/test_tools
	toochain => ./internal/ui/toolchain
)
