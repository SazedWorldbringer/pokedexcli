package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex> ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		command, exists := getCommands()[commandName]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}
		err := command.callback()
		if err != nil {
			fmt.Println(err)
		}
	}
}

// turn input to lowercase and split it into a slice
func cleanInput(input string) []string {
	output := strings.ToLower(input)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

// return all commands
func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Display help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name: "map",
			description: "Explore the world of Pokemon",
			callback: commandMap,
		},
	}
}
