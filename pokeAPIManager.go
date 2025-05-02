package main

import (
	"encoding/json"
	"io"
	"net/http"
)

const _pokeAPIURL = "https://pokeapi.co/api/v2"
const _locations = "/location-area"

func getLocationNames(url string) ([]string, string, string, error) {
	//fullURL := _pokeAPIURL + _locations
	if url == "" {
		url = _pokeAPIURL + _locations
	}

	res, err := http.Get(url)
	if err != nil {
		return []string{}, "", "", err
	}
	defer res.Body.Close()

	jsonRaw, err := io.ReadAll(res.Body)
	if err != nil {
		return []string{}, "", "", err
	}

	var URLData map[string]string
	json.Unmarshal(jsonRaw, &URLData)

	NextURL, ok := URLData["next"]
	if !ok {
		return []string{}, "", "", err
	}

	PrevURL, ok := URLData["previous"]
	if !ok {
		return []string{}, "", "", err
	}

	var resultsData map[string][]any
	json.Unmarshal(jsonRaw, &resultsData)

	rawResults, ok := resultsData["results"]
	if !ok {
		return []string{}, "", "", err
	}
	results := make([]string, 0)
	for _, innerResources := range rawResults {
		value := innerResources.(map[string]interface{})["name"]
		//fmt.Printf("name of place is: %s\n", value)
		results = append(results, value.(string))
	}

	return results, PrevURL, NextURL, nil
}
