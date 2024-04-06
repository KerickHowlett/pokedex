module toochain

go 1.22.1

require (
	command v1.0.0
	test_tools v1.0.0
)

replace (
	command => ../command
	test_tools => ../../test_tools
)
