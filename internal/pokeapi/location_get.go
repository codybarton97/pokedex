package pokeapi

import (
	"encoding/json"
	"io"
	"fmt"
)

func (c *Client) GetLocation(locationName string) (Location, error) {
	url := "https://pokeapi.co/api/v2/location-area/" + locationName

	if data, found := c.cache.Get(url); found {
		
		locationsResp := Location{}
		err := json.Unmarshal(data, &locationsResp)
		if err != nil {
			return Location{}, err
		}
		return locationsResp, nil
	}

	res, err := c.httpClient.Get(url)
	
	if err != nil {
		return Location{}, err
	}
	defer res.Body.Close()
	if res.StatusCode > 299 {
    	return Location{}, fmt.Errorf("bad status code: %v", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Location{}, err
	}

	locationsResp := Location{}
	err = json.Unmarshal(data, &locationsResp)
	if err != nil {
		return Location{}, err
	}

	c.cache.Add(url, data)

	return locationsResp, nil
}