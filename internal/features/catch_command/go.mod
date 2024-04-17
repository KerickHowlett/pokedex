module catch_command

go 1.22.1

require (
	command v1.0.0
	pokemon v1.0.0
	query_fetch v1.0.0
	test_tools v1.0.0
)

replace (
	command => ../../ui/command
	pokemon => ../../entities/pokemon
	query_fetch => ../../query_fetch
	test_tools => ../../test_tools
)
