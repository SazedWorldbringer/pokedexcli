package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func()
}

var commands = map[string]cliCommand{
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
}

func commandHelp() {
	fmt.Print(`
Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex

`)
}

func commandExit() {
	// exit the repl
	os.Exit(0)
}

func main() {
	for true {
		scanner := bufio.NewScanner(os.Stdin)

		// read for input
		fmt.Print("Pokedex> ")
		scanner.Scan()
		input := scanner.Text()

		// execute command
		for _, command := range commands {
			if input == command.name {
				command.callback()
			}
		}
	}
}
