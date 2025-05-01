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
	callback    func() error
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

		err := output.callback()
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

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Printf("# of CMDs: %d\n", len(commands))
	for key, _ := range commands {
		formattedOutput := fmt.Sprintf("%v: %v", commands[key].name, commands[key].description)
		fmt.Println(formattedOutput)
	}
	return nil
}

func commandMap() error {
	return nil
}
