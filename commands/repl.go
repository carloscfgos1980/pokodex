package commands

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/carloscfgos1980/pokodex/internal/pokeapi"
)

type Config struct {
	PokeapiClient    pokeapi.Client
	NextLocationsURL *string
	PrevLocationsURL *string
	CaughtPokemon    map[string]pokeapi.Pokemon
}

// StartRepl starts the REPL loop for the Pokedex CLI. It reads user input, parses commands, and executes the corresponding callbacks.
func StartRepl(cfg *Config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}
		command, exists := GetCommands()[commandName]
		if exists {
			err := command.Callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

// cleanInput takes a string input, converts it to lowercase, and splits it into words. It returns a slice of cleaned words.
func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

// cliCommand represents a command in the Pokedex CLI. It contains the command's name, description, and a callback function that executes the command's logic.
type cliCommand struct {
	Name        string
	Description string
	Callback    func(cfg *Config, args ...string) error
}

// GetCommands returns a map of available commands in the Pokedex CLI. Each command is associated with its name and a cliCommand struct that contains the command's description and callback function.
func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    CommandHelp,
		},
		"catch": {
			Name:        "catch <pokemon_name>",
			Description: "Attempt to catch a pokemon",
			Callback:    CommandCatch,
		},
		"inspect": {
			Name:        "inspect <pokemon_name>",
			Description: "View details about a caught Pokemon",
			Callback:    CommandInspect,
		},
		"explore": {
			Name:        "explore <location_name>",
			Description: "Explore a location",
			Callback:    CommandExplore,
		},
		"map": {
			Name:        "map",
			Description: "Get the next page of locations",
			Callback:    CommandMapf,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Get the previous page of locations",
			Callback:    CommandMapb,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "See all the pokemon you've caught",
			Callback:    CommandPokedex,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    CommandExit,
		},
	}
}
