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

func (c CliConfig) GetPokemonInArea(area string) (ResExplore, error) {
	if area == "" {
		return ResExplore{}, fmt.Errorf("no location entered")
	}

	url := _pokeAPIURL + _locations + "/" + area

	// find if we have already have existing data
	existingData, found := c.CachedInfo.Get(url)
	if found {
		resExplore := ResExplore{}
		err := json.Unmarshal(existingData, &resExplore)
		if err != nil {
			return ResExplore{}, err
		}
		return resExplore, nil
	}

	fmt.Println("CONNECTING...")
	res, err := http.Get(url)
	if err != nil {
		return ResExplore{}, fmt.Errorf("cannot connect to %s", url)
	}
	defer res.Body.Close()

	jsonRaw, err := io.ReadAll(res.Body)
	if err != nil {
		return ResExplore{}, fmt.Errorf("cannot read content")
	}

	resExplore := ResExplore{}
	err = json.Unmarshal(jsonRaw, &resExplore)
	if err != nil {
		return ResExplore{}, err
	}

	//add info to cache
	c.CachedInfo.Add(url, jsonRaw)

	return resExplore, nil
}
