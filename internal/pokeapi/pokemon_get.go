package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName

	if body, ok := c.cache.Get(url); ok {
		pokemonRes := Pokemon{}
		err := json.Unmarshal(body, &pokemonRes)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemonRes, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	// make get request
	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	// parse json data
	pokemonRes := Pokemon{}
	err = json.Unmarshal(body, &pokemonRes)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, body)
	return pokemonRes, nil
}
