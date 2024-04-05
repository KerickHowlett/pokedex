module shell

go 1.22.1

require (
	command v1.0.0
	maps v1.0.0
	query v1.0.0
	repl v1.0.0
	toochain v1.0.0
	testtools v1.0.0 // indirect
)

replace (
	command => ../ui/command
	maps => ../features/maps
	query => ../utils/query
	repl => ../ui/repl
	shell => .
	toochain => ../ui/toolchain
	testtools => ../testtools
)
