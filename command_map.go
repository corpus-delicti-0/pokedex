package main

import (
	"encoding/json"
	"fmt"
)

type locationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type locationAreaResponse struct {
	Next     string         `json:"next"`
	Previous string         `json:"previous"`
	Results  []locationArea `json:"results"`
}

func fetchMap(cfg *config, url string) error {
	data, err := fetchData(cfg, url)
	if err != nil {
		return err
	}

	locationData := locationAreaResponse{}

	/*decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&locationData)
	if err != nil {
		return err
	}*/

	err = json.Unmarshal(data, &locationData)
	if err != nil {
		return err
	}

	cfg.next = locationData.Next
	cfg.previous = locationData.Previous

	for _, location := range locationData.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMap(cfg *config, _ ...string) error {
	url := cfg.next
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"
	}

	return fetchMap(cfg, url)
}

func commandMapb(cfg *config, _ ...string) error {
	if cfg.previous == "" {
		fmt.Println("You're on the first page")
		return nil
	}

	return fetchMap(cfg, cfg.previous)
}
