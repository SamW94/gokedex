package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "   CHARMANDER bulbasaur PiKaChU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("length of actual slice and expected slice differ.")
			return
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("word in slice does not match expected word at index: %d", i)
				return
			}
		}
	}
}

func TestCommandHelp(t *testing.T) {
	err := commandHelp()
	if err != nil {
		t.Errorf("error returned from commandHelp function: %v", err)
	}
}
