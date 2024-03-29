package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New(`you must provide a location name
usage: explore <location_name>
`)
	}

	areaName := args[0]
	locationRes, err := cfg.pokeapiClient.GetLocation(areaName)
	if err != nil {
		return err
	}

	for _, encounter := range locationRes.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}
	return nil
}
