package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("request failed: %s", response.Status)
	}

	locationData := locationAreaResponse{}

	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&locationData)
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

func commandMap(cfg *config) error {
	url := cfg.next
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area"
	}

	return fetchMap(cfg, url)
}

func commandMapb(cfg *config) error {
	if cfg.previous == "" {
		fmt.Println("You're on the first page")
		return nil
	}

	return fetchMap(cfg, cfg.previous)
}
