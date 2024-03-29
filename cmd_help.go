package main

import "fmt"

func commandHelp(cfg *config, arg string) error {
	fmt.Print(`
Welcome to the Pokedex!
Usage:
	`)
	fmt.Println()
	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Println()

	return nil
}
