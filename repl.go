package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type config struct {
	next     string
	previous string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

var registry = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
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
	cfg := &config{}
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
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
