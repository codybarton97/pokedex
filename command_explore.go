package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *Config, args ...string) error {
	// 1. Make sure exactly one area name was typed
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}
	name := args[0]

	// 2. Tell the user what's happening
	fmt.Printf("Exploring %s...\n", name)

	// 3. Ask the API (or the cache) about this area
	location, err := cfg.pokeapiClient.GetLocation(name)
	if err != nil {
		return err
	}

	// 4. Print every Pokemon in the area
	fmt.Println("Found Pokemon:")
	for _, enc := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", enc.Pokemon.Name)
	}
	return nil
}