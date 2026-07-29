package main

import "testing"

func TestCleanInput(t *testing.T) {

	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " Charmander Bulbasaur PIKACHU ",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input:    " CharizaRD  attackS mETApod  ",
			expected: []string{"charizard", "attacks", "metapod"},
		},
		{
			input:    "ratTatA   lOSt 86  xP ",
			expected: []string{"rattata", "lost", "86", "xp"},
		},
		{
			input:    " clefAIry IS     vEry  cUTE    ",
			expected: []string{"clefairy", "is", "very", "cute"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		//check the length of the actual slice
		//if they don't match, use t.Errorf and continue to the next case
		if len(actual) != len(c.expected) {
			t.Errorf("expected: %v, got: %v", c.expected, actual)
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			//check each word in the slice
			//if they don't match, use f.Errorf to print an error message
			//and fail the test
			if word != expectedWord {
				t.Errorf("expected: %v, got: %v", c.expected, actual)
				break
			}
		}
	}
}
