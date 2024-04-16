module pokemon_encounter

go 1.22.1

require (
	pokemon v1.0.0
	test_tools v1.0.0
)

replace (
	pokemon => ../pokemon
	test_tools => ../../test_tools
)
