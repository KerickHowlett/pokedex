module pokedex_command

go 1.22.1

require (
	bills_pc v1.0.0
	command v1.0.0
	pokemon v1.0.0
	test_tools v1.0.0
)

require (
	entity_link v1.0.0 // indirect
	pokemon_stat v1.0.0 // indirect
	pokemon_type v1.0.0 // indirect
)

replace (
	bills_pc => ../../entities/bills_pc
	command => ../../ui/command
	entity_link => ../../entities/entity_link
	pokemon => ../../entities/pokemon
	pokemon_stat => ../../entities/pokemon_stat
	pokemon_type => ../../entities/pokemon_type
	test_tools => ../../test_tools
)
