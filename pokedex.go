package main

import (
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
