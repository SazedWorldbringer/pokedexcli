package main

import "os"

func commandExit(cfg *config, args ...string) error {
	// exit the repl
	os.Exit(0)
	return nil
}
