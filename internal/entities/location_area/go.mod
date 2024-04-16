module location_area

go 1.22.1

require (
	pokemon v1.0.0
	pokemon_encounter v1.0.0
	test_tools v1.0.0
)

replace (
	pokemon => ../pokemon
	pokemon_encounter => ../pokemon_encounter
	test_tools => ../../test_tools
)
