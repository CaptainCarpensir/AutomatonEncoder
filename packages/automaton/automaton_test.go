package automaton

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_EncodeAutomaton(t *testing.T) {
	testCases := map[string]struct {
		inStream        string
		wantedAutomaton *Automaton
		wantedErrPrefix string
	}{
		"fail to unmarshal": {
			inStream: `states: 3
initial: 1
finalstates: [3]
transitions:
  - input: "a"
    outputs: [2,3,1]
  - input: "b
    outputs: [1,2,3]`,
			wantedErrPrefix: "unmarshal automaton yaml:",
		},
		"success with basic automata": {
			inStream: `states: 3
initial: 1
finalstates: [3]
transitions:
  - input: "a"
    outputs: [2,3,1]
  - input: "b"
    outputs: [1,2,3]`,
			wantedAutomaton: &Automaton{
				StartState:   1,
				AcceptStates: []int{3},
				TransitionFunction: map[rune][]int{
					'a': {2, 3, 1},
					'b': {1, 2, 3},
				},
			},
		},
		"success with complex automata": {
			inStream: `states: 3
initial: 1
finalstates: [3]
transitions:
  - input: "ac"
    outputs: [2,3,1]
  - input: "b"
    outputs: [1,2,3]`,
			wantedAutomaton: &Automaton{
				StartState:   1,
				AcceptStates: []int{3},
				TransitionFunction: map[rune][]int{
					'a': {2, 3, 1},
					'b': {1, 2, 3},
					'c': {2, 3, 1},
				},
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got, gotErr := EncodeAutomaton([]byte(tc.inStream))
			if tc.wantedErrPrefix != "" {
				require.ErrorContains(t, gotErr, tc.wantedErrPrefix)
			} else {
				require.NoError(t, gotErr)
				require.Equal(t, tc.wantedAutomaton, got)
			}
		})
	}
}
