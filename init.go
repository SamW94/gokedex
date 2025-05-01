package main

import (
	"fmt"
)

var supportedCommands map[string]cliCommand

func init() {
	fmt.Println("Initialising Pokedex...")
	supportedCommands = make(map[string]cliCommand)
	supportedCommands["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	}

	supportedCommands["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	}

	supportedCommands["map"] = cliCommand{
		name:        "map",
		description: "Display the names of 20 location areas in the Pokemon world.",
		callback:    commandMap,
	}
}
