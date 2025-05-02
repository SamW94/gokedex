package main

import (
	"fmt"

	"github.com/SamW94/gokedex/internal/pokeapi"
)

type Pokedex struct {
	CaughtPokemon map[string]pokeapi.Pokemon
}

func NewPokedex() Pokedex {
	return Pokedex{
		CaughtPokemon: make(map[string]pokeapi.Pokemon),
	}
}

func (p *Pokedex) AddToPokedex(pokemonName string, pokemon pokeapi.Pokemon) {
	p.CaughtPokemon[pokemonName] = pokemon
}

func (p *Pokedex) HaveCaught(pokemonName string) bool {
	_, ok := p.CaughtPokemon[pokemonName]
	return ok
}

func (p *Pokedex) GetPokemonAttributes(pokemonName string) []string {
	pokemonAttributes := make([]string, 0)

	pokemon := p.CaughtPokemon[pokemonName]
	name := fmt.Sprintf("Name: %s", pokemonName)
	height := fmt.Sprintf("Height: %d", pokemon.Height)
	weight := fmt.Sprintf("Weight: %d", pokemon.Weight)
	statsSlice := make([]string, 0)
	statsSlice = append(statsSlice, "Stats:")

	for _, stat := range pokemon.Stats {
		statsSlice = append(statsSlice, fmt.Sprintf("	-%v: %v", stat.Stat.Name, stat.BaseStat))
	}

	typesSlice := make([]string, 0)
	typesSlice = append(typesSlice, "Types:")

	for _, pokeType := range pokemon.Types {
		typesSlice = append(typesSlice, fmt.Sprintf("	- %v", pokeType.Type.Name))
	}

	pokemonAttributes = append(pokemonAttributes, []string{name, height, weight}...)
	pokemonAttributes = append(pokemonAttributes, statsSlice...)
	pokemonAttributes = append(pokemonAttributes, typesSlice...)

	return pokemonAttributes
}
