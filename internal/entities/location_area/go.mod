module location_area

go 1.22.1

require (
	entity_link v1.0.0
	pokemon_encounter v1.0.0
	test_tools v1.0.0
)

replace (
	entity_link => ../entity_link
	pokemon_encounter => ../pokemon_encounter
	test_tools => ../../test_tools
)
