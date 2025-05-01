package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type locationAreaEndpointResponse struct {
	Count    int                 `json:"count"`
	Next     string              `json:"next"`
	Previous string              `json:"previous"`
	Results  []map[string]string `json:"results"`
}

func CallLocationAreaEndpoint() ([]string, error) {
	locationAreaEndpointResponse := locationAreaEndpointResponse{}
	var locations []string
	resp, err := http.Get("https://pokeapi.co/api/v2/location-area/")
	if err != nil {
		return nil, fmt.Errorf("http GET to Pokeapi location-area endpoint failed: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		return nil, fmt.Errorf("http response status code was: %d", resp.StatusCode)
	}

	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&locationAreaEndpointResponse); err != nil {
		return nil, fmt.Errorf("error decoding response body: %w", err)
	}

	for _, v := range locationAreaEndpointResponse.Results {
		locations = append(locations, v["name"])
	}

	return locations, nil
}
