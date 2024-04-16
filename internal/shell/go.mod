module shell

go 1.22.1

require (
	command v1.0.0 // indirect
	entities v1.0.0 // indirect
	explore v1.0.0
	maps v1.0.0
	query v1.0.0 // indirect
	repl v1.0.0
	system v1.0.0
	test_tools v1.0.0 // indirect
	toochain v1.0.0
)

replace (
	command => ../ui/command
	entities => ../entities
	explore => ../features/explore
	maps => ../features/maps
	query => ../query
	repl => ../ui/repl
	shell => .
	system => ../features/system
	test_tools => ../test_tools
	toochain => ../ui/toolchain
)
