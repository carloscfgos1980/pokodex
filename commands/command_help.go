package commands

import (
	"fmt"
)

// CommandHelp is the callback function for the "help" command in the Pokedex CLI. It prints a welcome message and a list of available commands along with their descriptions.
func CommandHelp(cfg *Config, args ...string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range GetCommands() {
		fmt.Printf("%s: %s\n", cmd.Name, cmd.Description)
	}
	fmt.Println()
	return nil
}
