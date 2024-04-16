module map_command

go 1.22.1

require (
	command v1.0.0
	location v1.0.0
	query_fetch v1.0.0
	test_tools v1.0.0
)

replace (
	command => ../../ui/command
	location => ../../entities/location
	query_fetch => ../../query_fetch
	test_tools => ../../test_tools
)
