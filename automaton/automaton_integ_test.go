package automaton

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Automata(t *testing.T) {
	testCases := map[string]struct {
		inputFile string
		language  map[string]bool
	}{
		"encoded automata recognizing c++ integer": {
			inputFile: "integer.yaml",
			language: map[string]bool{
				"1234567890":   true,
				"+1234567890":  true,
				"-1234567890":  true,
				"--1234567890": false,
				"1234567890-":  false,
				"987654,3210":  false,
				"abcdefg":      false,
			},
		},
		"encoded automata recognizing odd integers": {
			inputFile: "odd_digit.yaml",
			language: map[string]bool{
				"1":           true,
				"2":           false,
				"22222286003": true,
				"ab3391821":   false,
				"12345678909": true,
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			input, err := ioutil.ReadFile(filepath.Join("testdata", tc.inputFile))
			require.NoError(t, err, "read automaton yaml")
			automaton, err := EncodeAutomaton(input)
			require.NoError(t, err, "encode automaton")

			actual := make(map[string]bool)
			for word, _ := range tc.language {
				actual[word] = automaton.Recognize(word)
			}

			require.Equal(t, tc.language, actual)
		})
	}
}
