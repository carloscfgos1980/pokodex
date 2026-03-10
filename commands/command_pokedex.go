package commands

import (
	"fmt"
)

func CommandPokedex(cfg *Config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for _, p := range cfg.CaughtPokemon {
		fmt.Printf(" - %s\n", p.Name)
	}
	return nil
}
