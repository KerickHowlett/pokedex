PWD:=$(shell pwd)
bin_file:=$(PWD)/bin/pokedexcli
src_code:=$(PWD)/cmd/app

clean:
	rm -rf $(bin_file)

build: clean
	go build -C $(src_code) -o $(bin_file)

run: build
	clear && $(bin_file)
