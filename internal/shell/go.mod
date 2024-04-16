module shell

go 1.22.1

require (
	command v1.0.0 // indirect
	entities v1.0.0 // indirect
	exit_command v1.0.0
	explore_command v1.0.0
	help_command v1.0.0
	map_command v1.0.0
	query_fetch v1.0.0 // indirect
	repl v1.0.0
	test_tools v1.0.0 // indirect
	toochain v1.0.0
)

replace (
	command => ../ui/command
	entities => ../entities
	exit_command => ../features/exit_command
	explore_command => ../features/explore_command
	help_command => ../features/help_command
	map_command => ../features/map_command
	query_fetch => ../query_fetch
	repl => ../ui/repl
	shell => .
	system => ../features/system
	test_tools => ../test_tools
	toochain => ../ui/toolchain
)
