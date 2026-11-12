package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *Config, args ...string) error {
	
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}
	name := args[0]

	
	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	
	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	randNum := rand.Intn(pokemon.BaseExperience)
	if randNum > 40 {
		fmt.Printf("%s escaped!\n", name)
		return nil
	}

	fmt.Printf("%s was caught!\n", name)
	cfg.caughtPokemon[name] = pokemon
	return nil
}