package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
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
		"catch": {
			name:        "catch",
			description: "Attempts to catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspects a previously caught pokemon type",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Shows all pokemon in pokedex",
			callback:    commandViewPokedex,
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
			continue
		}
	}

}

func cleanInput(text string) []string {
	lowercasedText := strings.ToLower(text)
	words := strings.Fields(lowercasedText)
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
	if len(args) == 0 {
		return fmt.Errorf("you need to specify which area to explore")
	}

	results, err := config.GetPokemonInArea(args[0])
	if err != nil {
		return err
	}

	for _, pokemonType := range results.PokemonEncounters {
		fmt.Println(pokemonType.Pokemon.Name)
	}

	return nil
}

func commandCatch(config *pokeapi.CliConfig, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("you need to specify which Pokemon to catch")
	}

	results, err := config.GetPokemon(args[0])
	if err != nil {
		return err
	}

	x := float32(results.BaseExperience) * 0.01
	chance := float32(rand.IntN(5))

	fmt.Printf("Throwing a Pokeball at %s...\n", results.Name)
	if chance >= x {
		fmt.Printf("You caught %s!\n", results.Name)
		config.Pokedex.Add(results)
	} else {
		fmt.Println("Try again next time!")
	}

	return nil
}

func commandInspect(config *pokeapi.CliConfig, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("you need to specify which Pokemon to inspect")
	}

	pokemonName := args[0]
	results, found := config.Pokedex.Pokemons[pokemonName]
	if !found {
		fmt.Printf("You have not caught a %s yet.\n", pokemonName)
		return nil
	}

	fmt.Printf("Name: %s\nHeight: %d\nWeight: %d\n",
		results.Name,
		results.Height,
		results.Weight)

	fmt.Printf("Stats:\n")
	for _, val := range results.Stats {
		fmt.Printf("  -%s: %d\n", val.Stat.Name, val.BaseStat)
	}

	fmt.Printf("Types:\n")
	for _, val := range results.Types {
		fmt.Printf("  -%s\n", val.Type.Name)
	}

	return nil
}

func commandViewPokedex(config *pokeapi.CliConfig, args []string) error {
	fmt.Println("Your Pokedex:")
	for _, val := range config.Pokedex.Pokemons {
		fmt.Printf("  - %s\n", val.Name)
	}

	return nil
}
