package main

import (
	"fmt"

	"github.com/SazedWorldbringer/pokedexcli/internals"
)

func commandMap() error {
	// Get location areas
	locations, err := internals.GetLocations()
	if err != nil {
		return err
	}

	for i := range locations {
		fmt.Println(locations[i])
	}

	return nil
}
