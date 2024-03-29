package main

import "os"

func commandExit(cfg *config, arg string) error {
	// exit the repl
	os.Exit(0)
	return nil
}
