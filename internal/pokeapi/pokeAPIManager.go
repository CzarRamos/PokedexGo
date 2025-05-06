package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const _pokeAPIURL = "https://pokeapi.co/api/v2"
const _locations = "/location-area"

func (c CliConfig) GetLocationNames(url string) (ResLocations, error) {

	if url == "" {
		url = _pokeAPIURL + _locations
	}

	// find if we have already have existing data
	existingData, found := c.CachedInfo.Get(url)
	if found {
		resLocations := ResLocations{}
		err := json.Unmarshal(existingData, &resLocations)
		if err != nil {
			return ResLocations{}, err
		}
		return resLocations, nil
	}

	fmt.Println("CONNECTING...")
	res, err := http.Get(url)
	if err != nil {
		return ResLocations{}, err
	}
	defer res.Body.Close()

	jsonRaw, err := io.ReadAll(res.Body)
	if err != nil {
		return ResLocations{}, err
	}

	resLocations := ResLocations{}
	err = json.Unmarshal(jsonRaw, &resLocations)
	if err != nil {
		return ResLocations{}, err
	}

	//add info to cache
	c.CachedInfo.Add(url, jsonRaw)

	return resLocations, nil
}
