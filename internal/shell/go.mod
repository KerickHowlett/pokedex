module shell

go 1.22.1

require (
	catch_command v1.0.0
	command v1.0.0 // indirect
	exit_command v1.0.0
	explore_command v1.0.0
	help_command v1.0.0
	location v1.0.0 // indirect
	location_area v1.0.0 // indirect
	map_command v1.0.0
	pokemon v1.0.0 // indirect
	pokemon_encounter v1.0.0 // indirect
	query_fetch v1.0.0 // indirect
	repl v1.0.0
	test_tools v1.0.0 // indirect
	toochain v1.0.0
)

replace (
	catch_command => ../features/catch_command
	command => ../ui/command
	exit_command => ../features/exit_command
	explore_command => ../features/explore_command
	help_command => ../features/help_command
	location => ../entities/location
	location_area => ../entities/location_area
	map_command => ../features/map_command
	pokemon => ../entities/pokemon
	pokemon_encounter => ../entities/pokemon_encounter
	query_fetch => ../query_fetch
	repl => ../ui/repl
	shell => .
	system => ../features/system
	test_tools => ../test_tools
	toochain => ../ui/toolchain
)
