module explore_command

go 1.22.1

require (
	command v1.0.0
	entities v1.0.0
	query v1.0.0
	test_tools v1.0.0
)

replace (
	command => ../../ui/command
	entities => ../../entities
	query => ../../query
	test_tools => ../../test_tools
)
