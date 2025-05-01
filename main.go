package main

import (
	"github.com/SamW94/gokedex/internal/pokeapi"
)

func main() {
	webClient := pokeapi.NewClient()
	config := &config{
		pokeapiClient: webClient,
	}

	startRepl(config)
}
