package main

import (
	"fmt"
)

func commandHelp(_ *config, _ ...string) error {
	fmt.Println(
		`Welcome to the Pokédex!
Usage:

help: 					Displays a help message
map: 					Displays the names of 20 location areas (forward)
mapb: 					Displays the names of 20 location areas (backward)
explore <area_name>:    		Displays Pokémon found in a location area
catch <pokemon_name>:   		Attempts to catch a Pokémon 
inspect <pokemon_name>: 		Displays details about a caught Pokémon
pokedex: 				Displays all caught Pokémon
exit: 					Exit the Pokédex`)
	return nil
}
