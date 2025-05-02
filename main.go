package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(config *cliConfig) error
}

type cliConfig struct {
	Next     string
	Previous string
}

var commands = map[string]cliCommand{
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
}

func init() {
	commands["help"] = cliCommand{
		name:        "help",
		description: "Displays all commands",
		callback:    commandHelp,
	}
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	userconfig := cliConfig{
		Previous: "",
		Next:     "",
	}

	for {
		fmt.Print("Pokedex > ")
		ok := scanner.Scan()
		if !ok {
			continue
		}

		cleanedInput := (cleanInput(scanner.Text()))[0]

		output, ok := commands[cleanedInput]
		if !ok {
			fmt.Println("Unknown Command")
			continue
		}

		fmt.Println()

		err := output.callback(&userconfig)
		if err != nil {
			fmt.Println("Unknown Command")
			continue
		}
	}

}

func cleanInput(text string) []string {
	lowercasedText := strings.ToLower(text)
	words := strings.Fields(lowercasedText)

	return words
}

func commandExit(config *cliConfig) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *cliConfig) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Printf("# of CMDs: %d\n", len(commands))
	for key, _ := range commands {
		formattedOutput := fmt.Sprintf("%v: %v", commands[key].name, commands[key].description)
		fmt.Println(formattedOutput)
	}
	return nil
}

func commandMap(config *cliConfig) error {
	locations, prev, next, err := getLocationNames(config.Next)
	if err != nil {
		return err
	}

	config.Previous = prev
	config.Next = next

	for _, location := range locations {
		fmt.Println(location)
	}

	return nil
}

func commandMapb(config *cliConfig) error {
	locations, prev, next, err := getLocationNames(config.Previous)
	if err != nil {
		return err
	}

	config.Previous = prev
	config.Next = next

	for _, location := range locations {
		fmt.Println(location)
	}
	return nil
}
