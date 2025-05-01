package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
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

	if cachedData, found := wc.cache.Get(url); found {
		var response locationAreaEndpointResponse
		if err := json.Unmarshal(cachedData, &response); err != nil {
			return locationAreaEndpointResponse{}, fmt.Errorf("error unmarshalling cached data: %w", err)
		}
		return response, nil
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

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return locationAreaEndpointResponse{}, fmt.Errorf("error reading response body: %w", err)
	}

	wc.cache.Add(url, bodyBytes)

	var response locationAreaEndpointResponse
	if err := json.Unmarshal(bodyBytes, &response); err != nil {
		return locationAreaEndpointResponse{}, fmt.Errorf("error unmarshalling response body: %w", err)
	}

	return response, nil
}
