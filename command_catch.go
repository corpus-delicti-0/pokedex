package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
)

type PokemonStat struct {
	BaseStat int `json:"base_stat"`
	Stat     struct {
		Name string `json:"name"`
	} `json:"stat"`
}

type PokemonType struct {
	Type struct {
		Name string `json:"name"`
	} `json:"type"`
}

type Pokemon struct {
	Name           string        `json:"name"`
	BaseExperience int           `json:"base_experience"`
	Height         int           `json:"height"`
	Weight         int           `json:"weight"`
	Stats          []PokemonStat `json:"stats"`
	Types          []PokemonType `json:"types"`
}

func commandCatch(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("you must provide a Pokémon name")
	}

	pokemonName := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	url := "https://pokeapi.co/api/v2/pokemon/" + pokemonName + "/"

	data, err := fetchData(cfg, url)
	if err != nil {
		return err
	}

	pokemon := Pokemon{}
	if err := json.Unmarshal(data, &pokemon); err != nil {
		return err
	}

	if pokemon.BaseExperience <= 0 {
		return fmt.Errorf("invalid base experience for %s", pokemon.Name)
	}

	const catchThreshold = 50
	roll := rand.Intn(pokemon.BaseExperience) // for lower exp pokemon the possible roll range is <=50. for high exp rolling above 50 is more likely

	if roll > catchThreshold {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	cfg.pokedex[pokemon.Name] = pokemon //stores a pokemon in pokedex
	fmt.Printf("%s was caught!\n", pokemon.Name)

	return nil
}
