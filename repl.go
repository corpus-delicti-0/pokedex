package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
		if len(result) != 0 {
			fmt.Printf("Your command was: %s\n", result[0])
		}
	}
}
