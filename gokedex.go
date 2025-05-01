package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		print("Pokedex>")
		scanner.Scan()
		userInput := scanner.Text()

		cleanedInput := cleanInput(userInput)

		if len(cleanedInput) > 0 {
			inputCommand := cleanedInput[0]
			command, ok := supportedCommands[inputCommand]
			if !ok {
				fmt.Println("Unknown command")
			} else {
				err := command.callback()
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}
}
