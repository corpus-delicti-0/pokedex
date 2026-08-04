package main

import (
	"fmt"
	"os"
)

func commandExit(_ *config, _ ...string) error {
	fmt.Println("Closing the Pokédex... Goodbye!")
	os.Exit(0)
	return nil
}
