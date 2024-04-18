module shell

go 1.22.1

require (
	bills_pc v1.0.0
	catch_command v1.0.0
	exit_command v1.0.0
	explore_command v1.0.0
	help_command v1.0.0
	inspect_command v1.0.0
	map_command v1.0.0
	pokedex_command v1.0.0
	repl v1.0.0
	toochain v1.0.0
)

require (
	command v1.0.0 // indirect
	entity_link v1.0.0 // indirect
	location v1.0.0 // indirect
	location_area v1.0.0 // indirect
	pokemon v1.0.0 // indirect
	pokemon_encounter v1.0.0 // indirect
	pokemon_stat v1.0.0 // indirect
	pokemon_type v1.0.0 // indirect
	query_fetch v1.0.0 // indirect
	test_tools v1.0.0 // indirect
)

replace (
	bills_pc => ../entities/bills_pc
	catch_command => ../features/catch_command
	command => ../ui/command
	entity_link => ../entities/entity_link
	exit_command => ../features/exit_command
	explore_command => ../features/explore_command
	help_command => ../features/help_command
	inspect_command => ../features/inspect_command
	location => ../entities/location
	location_area => ../entities/location_area
	map_command => ../features/map_command
	pokedex_command => ../features/pokedex_command
	pokemon => ../entities/pokemon
	pokemon_encounter => ../entities/pokemon_encounter
	pokemon_stat => ../entities/pokemon_stat
	pokemon_type => ../entities/pokemon_type
	query_fetch => ../query_fetch
	repl => ../ui/repl
	shell => .
	system => ../features/system
	test_tools => ../test_tools
	toochain => ../ui/toolchain
)
