package commands

import (
	"fmt"
	"os"
)

// CommandExit is the callback function for the "exit" command in the Pokedex CLI. It prints a goodbye message and exits the program with a status code of 0, indicating a successful termination.
func CommandExit(cfg *Config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
