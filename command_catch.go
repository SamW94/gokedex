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

	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	pokemonBaseXP := pokemon.BaseExperience
	chanceToCatch := rand.Intn(1000)

	if pokemonBaseXP > chanceToCatch {
		fmt.Printf("%s escaped!\n", name)
	} else {
		if pokedex.HaveCaught(name) {
			fmt.Printf("You already have a %s in your Pokedex\n", name)
			return nil
		} else {
			fmt.Printf("Adding %s to your Pokedex!\n", name)
			pokedex.AddToPokedex(name, pokemon)
			return nil
		}
	}

	return nil
}
