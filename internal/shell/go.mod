module shell

go 1.22.1

require (
	command v1.0.0 // indirect
	maps v1.0.0
	query v1.0.0
	repl v1.0.0
	system v1.0.0
	testtools v1.0.0 // indirect
	toochain v1.0.0
)

replace (
	command => ../ui/command
	maps => ../features/maps
	query => ../utils/query
	repl => ../ui/repl
	shell => .
	system => ../features/system
	testtools => ../testtools
	toochain => ../ui/toolchain
)
