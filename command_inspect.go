package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, pokedex *Pokedex, args ...string) error {
	if len(args) < 1 {
		return errors.New("you must provide a Pokemon name")
	}

	name := args[0]

	if !pokedex.HaveCaught(name) {
		return errors.New("you have not caught that pokemon")
	}

	attributes := pokedex.GetPokemonAttributes(name)
	for _, attribute := range attributes {
		fmt.Println(attribute)
	}

	return nil
}
