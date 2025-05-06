package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	pokeapi "github.com/CzarRamos/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(config *pokeapi.CliConfig, args []string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Lists all locations in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "map back",
			description: "Displays the previous 20 locations shown",
			callback:    commandMapb,
		},
		"help": {
			name:        "help",
			description: "Displays all commands",
			callback:    commandHelp,
		},
		"explore": {
			name:        "explore",
			description: "Lists all pokemon located here",
			callback:    commandExplore,
		},
	}
}

func main() {

	client := pokeapi.NewClient()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		ok := scanner.Scan()
		if !ok {
			continue
		}

		cleanedInput := cleanInput(scanner.Text())
		command := cleanedInput[0]
		args := make([]string, 0)
		if len(cleanedInput) > 1 {
			args = cleanedInput[1:]
		}

		output, ok := getCommands()[command]
		if !ok {
			fmt.Println("Unknown Command")
			continue
		}

		err := output.callback(&client.Config, args)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Unknown Command")
			continue
		}
	}

}

func cleanInput(text string) []string {
	lowercasedText := strings.ToLower(text)
	words := strings.Fields(lowercasedText)
	fmt.Printf("WORDS:%d\n", len(words))
	fmt.Println(words)
	return words
}

func commandExit(config *pokeapi.CliConfig, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *pokeapi.CliConfig, args []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	commandList := getCommands()
	fmt.Printf("# of CMDs: %d\n", len(commandList))
	for key, val := range commandList {
		formattedOutput := fmt.Sprintf("\n%v (%v): %v", val.name, key, val.description)
		fmt.Println(formattedOutput)
	}
	return nil
}

func commandMap(config *pokeapi.CliConfig, args []string) error {
	results, err := config.GetLocationNames(config.Next)
	if err != nil {
		return err
	}

	config.Previous = results.Previous
	config.Next = results.Next

	for _, location := range results.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapb(config *pokeapi.CliConfig, args []string) error {
	results, err := config.GetLocationNames(config.Previous)
	if err != nil {
		return err
	}

	config.Previous = results.Previous
	config.Next = results.Next

	for _, location := range results.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandExplore(config *pokeapi.CliConfig, args []string) error {
	results, err := config.GetPokemonInArea(args[0])
	if err != nil {
		return err
	}

	for _, pokemonType := range results.PokemonEncounters {
		fmt.Println(pokemonType.Pokemon.Name)
	}

	return nil
}
