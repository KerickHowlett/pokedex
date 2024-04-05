module github.com/KerickHowlett/pokedexcli

go 1.22.1

require (
	command v1.0.0
	maps v1.0.0
	query v1.0.0
	testtools v1.0.0
)

replace (
	command => ./internal/ui/command
	maps => ./internal/feature/maps
	query => ./internal/utils/query
	testtools => ./internal/testtools
)
