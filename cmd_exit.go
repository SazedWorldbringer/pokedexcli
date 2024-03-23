package main

import "os"

func commandExit() error {
	// exit the repl
	os.Exit(0)
	return nil
}
