package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/SazedWorldbringer/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsUrl *string
	prevLocationsUrl *string
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := words[1:]

		command, exists := getCommands()[commandName]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}
		err := command.callback(cfg, args[0])
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
	callback    func(*config, string) error
}

// return all commands
func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Display help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Explore the world of Pokemon",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the previous locations",
			callback:    commandMapb,
		},
		"explore": {
			name: "explore",
			description: "Display the list of Pokemon available in given area",
			callback: commandExplore,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
