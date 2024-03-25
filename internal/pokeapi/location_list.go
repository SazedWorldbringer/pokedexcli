package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (ResShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if body, ok := c.cache.Get(url); ok {
		locationRes := ResShallowLocations{}
		err := json.Unmarshal(body, &locationRes)
		if err != nil {
			return ResShallowLocations{}, err
		}
		return locationRes, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ResShallowLocations{}, err
	}

	// make get request
	res, err := c.httpClient.Do(req)
	if err != nil {
		return ResShallowLocations{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return ResShallowLocations{}, err
	}

	// parse json data
	locationRes := ResShallowLocations{}
	err = json.Unmarshal(body, &locationRes)
	if err != nil {
		return ResShallowLocations{}, err
	}

	c.cache.Add(url, body)
	return locationRes, nil
}
