package main

import (
	"fmt"
	"os"

	"github.com/SamW94/gokedex/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for key, value := range supportedCommands {
		fmt.Printf("%s: %s\n", key, value.description)
	}
	return nil
}

func commandMap() error {
	locations, err := pokeapi.CallLocationAreaEndpoint()
	if err != nil {
		return fmt.Errorf("error returned from pokeapi.CallLocationAreaEndpoint() function: %v", err)
	}

	for _, location := range locations {
		fmt.Println(location)
	}
	return nil
}
