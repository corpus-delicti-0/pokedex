package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
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
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
