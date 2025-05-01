package pokeapi

import (
	"net/http"
)

type WebClient struct {
	httpClient http.Client
}

func NewClient() WebClient {
	return WebClient{
		httpClient: http.Client{},
	}
}
