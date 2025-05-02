package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, pokedex *Pokedex, args ...string) error {
	if len(args) < 1 {
		return errors.New("you must provide a Pokemon name")
	}

	name := args[0]

	if pokedex.HaveCaught(name) {
		fmt.Printf("You already have a %s in your Pokedex\n", name)
		return nil
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	chanceToCatch := rand.Intn(pokemon.BaseExperience)

	if chanceToCatch > 40 {
		fmt.Printf("%s escaped!\n", name)
	} else {
		fmt.Printf("%s was caught!\n", name)
		pokedex.AddToPokedex(name, pokemon)
		return nil
	}

	return nil
}
