module exit_command

go 1.22.1

require (
	command v1.0.0
	test_tools v1.0.0 // indirect
)

replace (
	command => ../../ui/command
	test_tools => ../../test_tools
)
