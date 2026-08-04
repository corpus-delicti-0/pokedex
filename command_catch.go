package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
)

type Pokemon struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
}

func commandCatch(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("you must provide a Pokémon name")
	}

	pokemonName := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	url := "http://pokeapi.co/api/v2/pokemon/" + pokemonName + "/"

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

	cfg.pokedex[pokemon.Name] = pokemon
	fmt.Printf("%s was caught!\n", pokemon.Name)

	return nil
}
