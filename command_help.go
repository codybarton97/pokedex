package main
import "fmt"

func commandHelp(cfg *Config, args ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for _, command := range registry {
		fmt.Println(command.name + ": " + command.description)
	}
	return nil
}