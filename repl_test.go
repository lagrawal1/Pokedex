package main

import "testing"

func TestCleanInput(t *testing.T) {

	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello  world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " HeLlo World BulbAsaUR ",
			expected: []string{"hello", "world", "bulbasaur"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(c.expected) != len(actual) {
			t.Errorf("slice lengths don't match")
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("words don't match.\n Expected Word: %s \nActual Word: %s", expectedWord, word)
			}
		}
	}

}
