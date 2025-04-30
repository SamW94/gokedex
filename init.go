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
}
