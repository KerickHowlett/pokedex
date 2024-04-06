module query

require test_tools v1.0.0

replace (
	query => ./
	test_tools => ../test_tools
)

go 1.22.1
