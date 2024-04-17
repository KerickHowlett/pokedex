module pokemon

go 1.22.1

require (
	entity_link v1.0.0 // indirect
	pokemon_stat v1.0.0
	pokemon_type v1.0.0
	test_tools v1.0.0
)

replace (
	entity_link => ../entity_link
	pokemon_stat => ../pokemon_stat
	pokemon_type => ../pokemon_type
	test_tools => ../../test_tools
)
