module github.com/KerickHowlett/pokedexcli

go 1.22.1

require testtools v1.0.0

replace testtools => ./internal/testtools

require query v1.0.0

replace query => ./internal/query

require location v1.0.0

replace location => ./internal/location
