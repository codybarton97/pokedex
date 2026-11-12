package main
import (
    "time"

    "github.com/codybarton97/pokedex/internal/pokeapi"
)

func main() {
	cfg := &Config{
    	pokeapiClient: pokeapi.NewClient(5 * time.Minute),
		caughtPokemon: map[string]pokeapi.Pokemon{},
	}
	startRepl(cfg)
}









