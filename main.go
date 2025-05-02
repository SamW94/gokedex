package main

import (
	"time"

	"github.com/SamW94/gokedex/internal/pokeapi"
)

func main() {
	webClient := pokeapi.NewClient(time.Second*5, time.Minute*5)
	config := &config{
		pokeapiClient: webClient,
	}

	startRepl(config)
}
