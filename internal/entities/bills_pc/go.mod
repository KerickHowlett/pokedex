module bills_pc

go 1.22.1

require (
	pokemon v0.0.0
	test_tools v1.0.0
)

require (
	entity_link v1.0.0 // indirect
	pokemon_stat v1.0.0 // indirect
	pokemon_type v1.0.0 // indirect
)

replace (
	entity_link => ../entity_link
	pokemon => ../pokemon
	pokemon_stat => ../pokemon_stat
	pokemon_type => ../pokemon_type
	test_tools => ../../test_tools
)
