module github.com/KerickHowlett/pokedexcli

go 1.22.1

require maps v1.0.0

replace maps => ./internal/feature/maps

require query v1.0.0

replace query => ./internal/utils/query

require testtools v1.0.0

replace testtools => ./internal/testtools
