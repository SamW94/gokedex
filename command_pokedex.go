package main

import (
	"fmt"
)

func commandPokedex(cfg *config, pokedex *Pokedex, args ...string) error {
	caughtPokemon := pokedex.AllCaughtPokemon()
	if len(caughtPokemon) == 0 {
		fmt.Println("Your Pokedex is empty, catch some Pokemon!")
	} else {
		fmt.Println("Your Pokedex:")
		for _, pokemon := range caughtPokemon {
			fmt.Printf("  - %s\n", pokemon)
		}
	}

	return nil
}
