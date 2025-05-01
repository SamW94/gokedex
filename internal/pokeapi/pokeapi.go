package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetLocationsNoID() (locationAreaEndpointResponse, error) {
	locationAreaEndpointResponse := locationAreaEndpointResponse{}
	resp, err := http.Get("https://pokeapi.co/api/v2/location-area/")
	if err != nil {
		return locationAreaEndpointResponse, fmt.Errorf("http GET to Pokeapi location-area endpoint failed: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		return locationAreaEndpointResponse, fmt.Errorf("http response status code was: %d", resp.StatusCode)
	}

	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&locationAreaEndpointResponse); err != nil {
		return locationAreaEndpointResponse, fmt.Errorf("error decoding response body: %w", err)
	}

	return locationAreaEndpointResponse, nil
}
