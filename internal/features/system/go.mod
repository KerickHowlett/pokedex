module system

go 1.22.1

require (
	command v1.0.0
	testtools v1.0.0 // indirect
)

replace (
	command => ../../ui/command
	testtools => ../../testtools
)
