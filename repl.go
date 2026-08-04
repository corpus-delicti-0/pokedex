package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/corpus-delicti-0/pokedex/internal/pokecache"
)

type config struct {
	next     string
	previous string
	cache    *pokecache.Cache
	pokedex  map[string]Pokemon
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error //...string means the command can receive zero or more arguments
}

var registry = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokédex",
		callback:    commandExit,
	},
	"help": {
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	},
	"map": {
		name:        "map",
		description: "Displays the next 20 location areas",
		callback:    commandMap,
	},
	"mapb": {
		name:        "mapb",
		description: "Displays the previous 20 location areas",
		callback:    commandMapb,
	},
	"explore": {
		name:        "explore",
		description: "Explores a location area",
		callback:    commandExplore,
	},
	"catch": {
		name:        "catch",
		description: "Attempts to catch a Pokémon",
		callback:    commandCatch,
	},
}

func cleanInput(text string) []string {
	result := []string{}
	for part := range strings.SplitSeq(text, " ") {
		if len(part) == 0 {
			continue
		}
		result = append(result, strings.ToLower(part))
	}
	return result
}

func runREPL() error {
	cfg := &config{
		cache:   pokecache.NewCache(5 * time.Minute),
		pokedex: make(map[string]Pokemon),
	}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			return scanner.Err()
		}
		input := scanner.Text()
		result := cleanInput(input)
		if len(result) == 0 {
			continue
		}
		command, ok := registry[result[0]]
		if !ok {
			fmt.Println("Unknown command")
		} else {
			err := command.callback(cfg, result[1:]...) //result[1:] produces the arg slice and ... passes its elements into variadic func
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
