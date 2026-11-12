package pokeapi

import (
	"encoding/json"
	"io"
	"fmt"
)

func (c *Client) ListLocations(pageURL *string) (LocationArea, error) {
	url := "https://pokeapi.co/api/v2/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if data, found := c.cache.Get(url); found {
		fmt.Println("cache hit!")
		locationsResp := LocationArea{}
		err := json.Unmarshal(data, &locationsResp)
		if err != nil {
			return LocationArea{}, err
		}
		return locationsResp, nil
	}

	res, err := c.httpClient.Get(url)
	fmt.Println("making network request...")
	if err != nil {
		return LocationArea{}, err
	}
	defer res.Body.Close()
	if res.StatusCode > 299 {
    	return LocationArea{}, fmt.Errorf("bad status code: %v", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationArea{}, err
	}

	locationsResp := LocationArea{}
	err = json.Unmarshal(data, &locationsResp)
	if err != nil {
		return LocationArea{}, err
	}

	c.cache.Add(url, data)

	return locationsResp, nil
}