package main

import (
	"errors"
	"fmt"
)

func commandMapf(cfg *config, arg string) error {
	locationRes, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsUrl)
	if err != nil {
		return err
	}

	cfg.nextLocationsUrl = locationRes.Next
	cfg.prevLocationsUrl = locationRes.Previous

	for _, location := range locationRes.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapb(cfg *config, arg string) error {
	if cfg.prevLocationsUrl == nil {
		return errors.New("you're on the first page")
	}

	locationRes, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsUrl)
	if err != nil {
		return err
	}

	cfg.nextLocationsUrl = locationRes.Next
	cfg.prevLocationsUrl = locationRes.Previous

	for _, location := range locationRes.Results {
		fmt.Println(location.Name)
	}

	return nil
}
