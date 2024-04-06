module maps

go 1.22.1

require (
	command v1.0.0
	query v1.0.0
	testtools v1.0.0
)

replace (
	command => ../../ui/command
	query => ../../query
	testtools => ../../testtools
)
