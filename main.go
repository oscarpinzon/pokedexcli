package main

import (
	"time"

	"github.com/oscarpinzon/pokedexcli/internal/pokeapi"
)

func main() {
	pokeapiClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeapiClient,
	}
	startRepl(cfg)
}
