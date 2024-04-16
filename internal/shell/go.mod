module shell

go 1.22.1

require (
	command v1.0.0 // indirect
	entities v1.0.0 // indirect
	exit v1.0.0
	explore v1.0.0
	help v1.0.0
	map v1.0.0
	query v1.0.0 // indirect
	repl v1.0.0
	test_tools v1.0.0 // indirect
	toochain v1.0.0
)

replace (
	command => ../ui/command
	entities => ../entities
	exit => ../features/exit_command
	explore => ../features/explore_command
	help => ../features/help_command
	map => ../features/map_command
	query => ../query
	repl => ../ui/repl
	shell => .
	system => ../features/system
	test_tools => ../test_tools
	toochain => ../ui/toolchain
)
