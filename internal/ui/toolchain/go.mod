module tc

go 1.22.1

require (
	command v1.0.0
	testtools v1.0.0
)

replace (
	command => ../command
	testtools => ../../testtools
)
