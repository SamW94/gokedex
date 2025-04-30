package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		print("Pokedex >")
		scanner.Scan()
		userInput := scanner.Text()

		cleanedInput := cleanInput(userInput)
		inputCommand := cleanedInput[0]

		command, ok := supportedCommands[inputCommand]
		if !ok {
			fmt.Println("Unknown command")
		} else {
			command.callback()
		}
	}
}
