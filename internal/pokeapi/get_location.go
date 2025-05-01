package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type locationAreaEndpointResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (wc *WebClient) GetLocations(pageURL *string) (locationAreaEndpointResponse, error) {
	url := baseURL + "/location-area"

	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return locationAreaEndpointResponse{}, fmt.Errorf("error creating HTTP GET request: %w", err)
	}

	resp, err := wc.httpClient.Do(req)
	if err != nil {
		return locationAreaEndpointResponse{}, fmt.Errorf("error making HTTP GET request to %s: %w", url, err)
	}

	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		return locationAreaEndpointResponse{}, fmt.Errorf("response status code was: %d", resp.StatusCode)
	}

	locationAreaEndpointResponse := locationAreaEndpointResponse{}
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&locationAreaEndpointResponse); err != nil {
		return locationAreaEndpointResponse, fmt.Errorf("error decoding response body: %w", err)
	}

	return locationAreaEndpointResponse, nil
}
