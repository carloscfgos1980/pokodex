package main

import (
	"time"

	"github.com/carloscfgos1980/pokodex/commands"
	"github.com/carloscfgos1980/pokodex/internal/pokeapi"
)

func main() {

	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := commands.Config{
		CaughtPokemon: map[string]pokeapi.Pokemon{},
		PokeapiClient: pokeClient,
	}
	commands.StartRepl(&cfg)
}
