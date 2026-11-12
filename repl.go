package main

import ( 
	"strings"
	"bufio"
	"fmt"
	"os"
	"github.com/codybarton97/pokedex/internal/pokeapi"
)

//Formats input
func cleanInput(text string) []string {
	txt := strings.ToLower(text)
	newTxt := strings.Fields(txt)
	return newTxt
}

func startRepl(cfg *Config) {
	//Register commands to execute
	registry = map[string]cliCommand{
		"exit": {
		name:		 "exit",
		description: "Exit the Pokedex",
		callback: 	 commandExit,
		},

		"help": {
		name:		 "help",
		description: "Displays a help message",
		callback:	 commandHelp,
		},

		"map": {
			name:		 "map",
			description: "Get the next page of locations", 
			callback:	commandMapf,
		},

		"mapb": {
			name: 		 "mapb",
			description: "Get the previous page of locations",
			callback: 	 commandMapb,
		},

		"explore": {
			name: 		 "explore",
			description: "Explore a specific area location",
			callback: 	 commandExplore,
		},

		"catch": {
			name:        "catch",
			description: "Attempt to catch a pokemon",
			callback:    commandCatch,
		},

		"inspect": {
			name: 		 "inspect",
			description: "Inspect a pokemon's stats",
			callback: 	 commandInspect,
		},

		"pokedex": {
			name: "pokedex",
			description: "List all pokemon you've caught",
			callback: commandPokedex,
		},

	}

	//Take user input
	scanner := bufio.NewScanner(os.Stdin)

	//Begin forever loop
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}

		input := scanner.Text()
		cleaned := cleanInput(input)
		if len(cleaned) == 0 {
			continue
		}

		command, exists := registry[cleaned[0]]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}else{
			errs := command.callback(cfg, cleaned[1:]...)
			if errs != nil{
				fmt.Println(errs)
			}
		}
	}
}

type cliCommand struct {
	name		string
	description	string
	callback	func(*Config, ...string) error
}

var registry map[string]cliCommand

type Config struct {
	state 			 string
	registry 		 map[string]cliCommand
	pokeapiClient	 pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	caughtPokemon    map[string]pokeapi.Pokemon
}