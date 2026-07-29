package main

import "strings"

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
