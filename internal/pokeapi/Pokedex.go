package pokeapi

import "sync"

type Pokedex struct {
	Pokemons map[string]ResPokemon
	mu       *sync.RWMutex
}

func NewPokedex() Pokedex {
	newPokedex := Pokedex{
		make(map[string]ResPokemon),
		&sync.RWMutex{},
	}

	return newPokedex
}

func (p Pokedex) Add(newPokemon ResPokemon) {
	p.Pokemons[newPokemon.Name] = newPokemon
}
