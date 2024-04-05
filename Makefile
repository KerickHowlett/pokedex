PWD:=$(shell pwd)
bin_file:=$(PWD)/bin/pokedexcli
main_go:=$(PWD)/cmd/cli/app

clean:
	rm -rf $(bin_file)

build: clean
	go build -C $(main_go) -o $(bin_file)

run: build
	clear && $(bin_file)
