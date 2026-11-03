package main
import "fmt"

func commandHelp(cfg *Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")
	for _, command := range registry {
		fmt.Println(command.name + ": " + command.description)
	}
	return nil
}