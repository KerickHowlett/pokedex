module github.com/KerickHowlett/pokedexcli

go 1.22.1

require (
	catch_command v1.0.0 // indirect
	command v1.0.0 // indirect
	entity_link v1.0.0 // indirect
	exit_command v1.0.0 // indirect
	explore_command v1.0.0 // indirect
	help_command v1.0.0 // indirect
	location v1.0.0 // indirect
	location_area v1.0.0 // indirect
	map_command v1.0.0 // indirect
	pokemon v1.0.0 // indirect
	pokemon_encounter v1.0.0 // indirect
	pokemon_type v1.0.0 // indirect
	query_fetch v1.0.0 // indirect
	repl v1.0.0 // indirect
	shell v1.0.0
	test_tools v1.0.0 // indirect
	toochain v1.0.0 // indirect
)

replace (
	catch_command => ./internal/features/catch_command
	command => ./internal/ui/command
	entity_link => ./internal/entities/entity_link
	exit_command => ./internal/features/exit_command
	explore_command => ./internal/features/explore_command
	help_command => ./internal/features/help_command
	location => ./internal/entities/location
	location_area => ./internal/entities/location_area
	map_command => ./internal/features/map_command
	pokemon => ./internal/entities/pokemon
	pokemon_encounter => ./internal/entities/pokemon_encounter
	pokemon_type => ./internal/entities/pokemon_type
	query_fetch => ./internal/query_fetch
	repl => ./internal/ui/repl
	shell => ./internal/shell
	system => ./internal/features/system
	test_tools => ./internal/test_tools
	toochain => ./internal/ui/toolchain
)
