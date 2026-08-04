package main

import "fmt"

func commandPokedex(cfg *config, _ ...string) error {
	fmt.Println("Your Pokédex:")

	for name := range cfg.pokedex {
		fmt.Printf("- %s\n", name)
	}

	return nil
}
