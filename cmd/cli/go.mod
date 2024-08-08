module github.com/KerickHowlett/pokedexcli

go 1.22.1

require shell v1.0.0

replace shell => ../../internal/shell

require (
	bills_pc v1.0.0 // indirect
	catch_command v1.0.0 // indirect
	command v1.0.0 // indirect
	entity_link v1.0.0 // indirect
	exit_command v1.0.0 // indirect
	explore_command v1.0.0 // indirect
	help_command v1.0.0 // indirect
	inspect_command v1.0.0 // indirect
	location v1.0.0 // indirect
	location_area v1.0.0 // indirect
	map_command v1.0.0 // indirect
	pokedex_command v1.0.0 // indirect
	pokemon v1.0.0 // indirect
	pokemon_encounter v1.0.0 // indirect
	pokemon_stat v1.0.0 // indirect
	pokemon_type v1.0.0 // indirect
	query_fetch v1.0.0 // indirect
	repl v1.0.0 // indirect
	test_tools v1.0.0 // indirect
	toochain v1.0.0 // indirect
)

replace (
	bills_pc => ../../internal/entities/bills_pc
	catch_command => ../../internal/features/catch_command
	command => ../../internal/ui/command
	entity_link => ../../internal/entities/entity_link
	exit_command => ../../internal/features/exit_command
	explore_command => ../../internal/features/explore_command
	help_command => ../../internal/features/help_command
	inspect_command => ../../internal/features/inspect_command
	location => ../../internal/entities/location
	location_area => ../../internal/entities/location_area
	map_command => ../../internal/features/map_command
	pokedex_command => ../../internal/features/pokedex_command
	pokemon => ../../internal/entities/pokemon
	pokemon_encounter => ../../internal/entities/pokemon_encounter
	pokemon_stat => ../../internal/entities/pokemon_stat
	pokemon_type => ../../internal/entities/pokemon_type
	query_fetch => ../../internal/query_fetch
	repl => ../../internal/ui/repl
	test_tools => ../../internal/test_tools
	toochain => ../../internal/ui/toolchain
)
