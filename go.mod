module github.com/KerickHowlett/pokedexcli

go 1.22.1

require (
	command v1.0.0 // indirect
	maps v1.0.0 // indirect
	query v1.0.0 // indirect
	repl v1.0.0 // indirect
	shell v1.0.0
	system v1.0.0 // indirect
	test_tools v1.0.0 // indirect
	toochain v1.0.0 // indirect
)

replace (
	command => ./internal/ui/command
	explore => ./internal/features/explore
	maps => ./internal/features/maps
	query => ./internal/query
	repl => ./internal/ui/repl
	shell => ./internal/shell
	system => ./internal/features/system
	test_tools => ./internal/test_tools
	toochain => ./internal/ui/toolchain
)
