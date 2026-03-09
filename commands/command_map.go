package commands

import (
	"errors"
	"fmt"
)

// CommandMapf is the callback function for the "map" command in the Pokedex CLI. It retrieves the next page of locations from the PokeAPI and prints their names. It also updates the next and previous location URLs in the configuration.
func CommandMapf(cfg *Config, args ...string) error {
	locationsResp, err := cfg.PokeapiClient.ListLocations(cfg.NextLocationsURL)
	if err != nil {
		return err
	}

	cfg.NextLocationsURL = locationsResp.Next
	cfg.PrevLocationsURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

// commandMapb is the callback function for the "mapb" command in the Pokedex CLI. It retrieves the previous page of locations from the PokeAPI and prints their names. It also updates the next and previous location URLs in the configuration. If there is no previous page, it returns an error message.
func CommandMapb(cfg *Config, args ...string) error {
	if cfg.PrevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	locationResp, err := cfg.PokeapiClient.ListLocations(cfg.PrevLocationsURL)
	if err != nil {
		return err
	}

	cfg.NextLocationsURL = locationResp.Next
	cfg.PrevLocationsURL = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
