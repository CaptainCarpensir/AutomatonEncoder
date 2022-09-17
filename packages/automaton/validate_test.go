package automaton

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ValidateIntermediateAutomaton(t *testing.T) {
	var (
		validState       = State(2)
		invalidStateLow  = State(-1)
		invalidStateHigh = State(165)
		numStates        = 2
		input1           = "a"
		input2           = "bc"
		invalidInput2    = "ac"
	)

	testCases := map[string]struct {
		automaton       *IntermediateAutomaton
		wantedErrPrefix string
	}{
		"invalid state too low": {
			automaton: &IntermediateAutomaton{
				States:       &numStates,
				StartState:   &invalidStateLow,
				AcceptStates: []State{validState},
				Transitions: []Transition{
					{
						Input:     []rune(input1),
						OutStates: []State{validState, validState},
					},
					{
						Input:     []rune(input2),
						OutStates: []State{validState, validState},
					},
				},
			},
			wantedErrPrefix: "start state must be a valid state",
		},
		"invalid state too high": {
			automaton: &IntermediateAutomaton{
				States:       &numStates,
				StartState:   &invalidStateHigh,
				AcceptStates: []State{validState},
				Transitions: []Transition{
					{
						Input:     []rune(input1),
						OutStates: []State{validState, validState},
					},
					{
						Input:     []rune(input2),
						OutStates: []State{validState, validState},
					},
				},
			},
			wantedErrPrefix: "start state must be a valid state",
		},
		"invalid output state": {
			automaton: &IntermediateAutomaton{
				States:       &numStates,
				StartState:   &validState,
				AcceptStates: []State{validState},
				Transitions: []Transition{
					{
						Input:     []rune(input1),
						OutStates: []State{invalidStateHigh, validState},
					},
					{
						Input:     []rune(input2),
						OutStates: []State{validState, validState},
					},
				},
			},
			wantedErrPrefix: "output state must be a valid state",
		},
		"invalid accept state": {
			automaton: &IntermediateAutomaton{
				States:       &numStates,
				StartState:   &validState,
				AcceptStates: []State{invalidStateHigh},
				Transitions: []Transition{
					{
						Input:     []rune(input1),
						OutStates: []State{validState, validState},
					},
					{
						Input:     []rune(input2),
						OutStates: []State{validState, validState},
					},
				},
			},
			wantedErrPrefix: "accept states must be valid states",
		},
		"non-deterministic transition function": {
			automaton: &IntermediateAutomaton{
				States:       &numStates,
				StartState:   &validState,
				AcceptStates: []State{validState},
				Transitions: []Transition{
					{
						Input:     []rune(input1),
						OutStates: []State{validState, validState, validState},
					},
					{
						Input:     []rune(input2),
						OutStates: []State{validState, validState},
					},
				},
			},
			wantedErrPrefix: "transition functions must map each state exactly once",
		},
		"input maps to multiple transition functions": {
			automaton: &IntermediateAutomaton{
				States:       &numStates,
				StartState:   &validState,
				AcceptStates: []State{validState},
				Transitions: []Transition{
					{
						Input:     []rune(input1),
						OutStates: []State{validState, validState},
					},
					{
						Input:     []rune(invalidInput2),
						OutStates: []State{validState, validState},
					},
				},
			},
			wantedErrPrefix: "letters must have exactly one transition function",
		},
		"valid automata": {
			automaton: &IntermediateAutomaton{
				States:       &numStates,
				StartState:   &validState,
				AcceptStates: []State{validState},
				Transitions: []Transition{
					{
						Input:     []rune(input1),
						OutStates: []State{validState, validState},
					},
					{
						Input:     []rune(input2),
						OutStates: []State{validState, validState},
					},
				},
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			gotErr := tc.automaton.validate()
			if tc.wantedErrPrefix != "" {
				require.ErrorContains(t, gotErr, tc.wantedErrPrefix)
			} else {
				require.NoError(t, gotErr)
			}
		})
	}
}
