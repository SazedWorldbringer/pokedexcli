package main

import (
	"fmt"
	"io"
	"net/http"
)

func commandMap() error {
	// Get location areas
	res, err := http.Get("https://pokeapi.co/api/v2/location/")
	if err != nil {
		return err
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", body)
	return nil
}
