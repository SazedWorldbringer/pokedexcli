package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocation(areaName string) (Location, error) {
	url := baseURL + "/location-area/" + areaName

	if body, ok := c.cache.Get(url); ok {
		locationRes := Location{}
		err := json.Unmarshal(body, &locationRes)
		if err != nil {
			return Location{}, err
		}
		return locationRes, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, err
	}

	// make get request
	res, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Location{}, err
	}

	// parse json data
	locationRes := Location{}
	err = json.Unmarshal(body, &locationRes)
	if err != nil {
		return Location{}, err
	}

	c.cache.Add(url, body)
	return locationRes, nil
}
