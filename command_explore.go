package main

import (
	"encoding/json"
	"fmt"
)

type pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type pokemonEncounter struct {
	Pokemon pokemon `json:"pokemon"`
}

type locationAreaDetails struct {
	PokemonEncounters []pokemonEncounter `json:"pokemon_encounters"`
}

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("you must provide a location area")
	}

	areaName := args[0]
	fmt.Printf("Exploring %s...\n", areaName)

	url := "https://pokeapi.co/api/v2/location-area/" + areaName + "/"

	data, err := fetchData(cfg, url)
	if err != nil {
		return err
	}

	locationData := locationAreaDetails{}

	err = json.Unmarshal(data, &locationData)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokémon:")
	for _, encounter := range locationData.PokemonEncounters {
		fmt.Printf("- %s\n", encounter.Pokemon.Name)
	}

	return nil
}
