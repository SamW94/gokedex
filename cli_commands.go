package main

import (
	"fmt"
	"os"
)

func commandExit(c *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for key, value := range getCommands() {
		fmt.Printf("%s: %s\n", key, value.description)
	}
	return nil
}

func commandMap(c *config) error {
	getLocationsResponse, err := c.pokeapiClient.GetLocations(c.nextLocationsURL)
	if err != nil {
		return fmt.Errorf("error calling pokeapi GetLocations() function: %w", err)
	}

	c.nextLocationsURL = getLocationsResponse.Next
	c.prevLocationsURL = getLocationsResponse.Previous

	for _, locationResponseResult := range getLocationsResponse.Results {
		fmt.Println(locationResponseResult.Name)
	}

	return nil
}

func commandMapb(c *config) error {
	if c.prevLocationsURL == nil {
		fmt.Println("You're on the first page.")
		return nil
	}

	getLocationsResponse, err := c.pokeapiClient.GetLocations(c.prevLocationsURL)
	if err != nil {
		return fmt.Errorf("error calling pokeapi GetLocations() function: %w", err)
	}

	c.nextLocationsURL = getLocationsResponse.Next
	c.prevLocationsURL = getLocationsResponse.Previous

	for _, locationResponseResult := range getLocationsResponse.Results {
		fmt.Println(locationResponseResult.Name)
	}
	return nil
}
