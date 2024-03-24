package internals

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationAreas struct {
	Count    int     `json:"count"`
	Next     string  `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocations() ([]string, error) {
	res, err := http.Get("https://pokeapi.co/api/v2/location")
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return nil, fmt.Errorf("Response failed with status code: %d and\nbody: %s", res.StatusCode, body)
	}
	if err != nil {
		return nil, err
	}

	locationsInfo := LocationAreas{}
	// parse json response
	json.Unmarshal(body, &locationsInfo)

	locations := []string{}

	for i := range locationsInfo.Results {
		locations = append(locations, locationsInfo.Results[i].Name)
	}

	return locations, nil
}
