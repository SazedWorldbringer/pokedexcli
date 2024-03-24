package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationAreas struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous *string    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

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

	locations := LocationAreas{}
	// parse(unmarshal?) the json
	json.Unmarshal(body, &locations)

	for i := range locations.Results {
		fmt.Println(locations.Results[i].Name)
	}

	return nil
}
