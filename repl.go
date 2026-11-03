package main
import "strings"
import "bufio"
import "fmt"
import "os"

//Formats input
func cleanInput(text string) []string {
	txt := strings.ToLower(text)
	newTxt := strings.Fields(txt)
	return newTxt
}

func startRepl(cfg *Config) {

	registry = map[string]cliCommand{
		"exit": {
		name:	"exit",
		description: "Exit the Pokedex",
		callback: 	commandExit,
		},

		"help": {
		name:	"help",
		description: "Displays a help message",
		callback:	commandHelp,
		},
	}

	//Take user input
	scanner := bufio.NewScanner(os.Stdin)

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
			errs := command.callback(cfg)
			if errs != nil{
				fmt.Println(errs)
			}
		}
	}
}

type cliCommand struct {
	name		string
	description	string
	callback	func(*Config) error
}

var registry map[string]cliCommand

type Config struct {
	state 	string
	registry 	map[string]cliCommand
}