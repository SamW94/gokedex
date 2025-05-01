package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/SamW94/gokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.WebClient
	nextLocationsURL *string
	prevLocationsURL *string
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		print("Pokedex>")
		scanner.Scan()

		cleanedInput := cleanInput(scanner.Text())
		if len(cleanedInput) == 0 {
			continue
		}

		inputCommand := cleanedInput[0]

		command, ok := getCommands()[inputCommand]
		if !ok {
			fmt.Println("Unknown command")
			continue
		} else {
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Display the next 20 location areas in the Pokemon world.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the previous 20 location areas in the Pokemon world.",
			callback:    commandMapb,
		},
	}
}
