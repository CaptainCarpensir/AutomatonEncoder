package automaton

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_UnmarshalAutomaton(t *testing.T) {
	var (
		states = 3
		init   = State(1)
		input1 = "a"
		input2 = "b"
		input3 = "ac"
	)

	testCases := map[string]struct {
		inStream        string
		wantedAutomaton *IntermediateAutomaton
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
			wantedAutomaton: &IntermediateAutomaton{
				States:       &states,
				StartState:   &init,
				AcceptStates: []State{3},
				Transitions: []Transition{
					{
						Input:     []rune(input1),
						OutStates: []State{2, 3, 1},
					},
					{
						Input:     []rune(input2),
						OutStates: []State{1, 2, 3},
					},
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
			wantedAutomaton: &IntermediateAutomaton{
				States:       &states,
				StartState:   &init,
				AcceptStates: []State{3},
				Transitions: []Transition{
					{
						Input:     []rune(input3),
						OutStates: []State{2, 3, 1},
					},
					{
						Input:     []rune(input2),
						OutStates: []State{1, 2, 3},
					},
				},
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got, gotErr := UnmarshalAutomaton([]byte(tc.inStream))
			if tc.wantedErrPrefix != "" {
				require.ErrorContains(t, gotErr, tc.wantedErrPrefix)
			} else {
				require.NoError(t, gotErr)
				require.Equal(t, tc.wantedAutomaton, got)
			}
		})
	}
}

func Test_ConvertToAutomaton(t *testing.T) {
	var (
		validState   = State(2)
		invalidState = State(165)
		numStates    = 2
		input1       = "a"
		input2       = "d"
		input3       = "bc"
		rune1        = 'a'
		rune2        = 'd'
		rune3        = 'b'
		rune4        = 'c'
	)

	testCases := map[string]struct {
		inputAutomaton  IntermediateAutomaton
		outputAutomaton *Automaton
		wantedErrPrefix string
	}{
		"invalid automaton": {
			inputAutomaton: IntermediateAutomaton{
				States:       &numStates,
				StartState:   &invalidState,
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
			wantedErrPrefix: "validate automaton:",
		},
		"success with basic automaton": {
			inputAutomaton: IntermediateAutomaton{
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
			outputAutomaton: &Automaton{
				StartState:   int(validState),
				AcceptStates: []int{int(validState)},
				TransitionFunction: map[rune][]int{
					rune1: {2, 2},
					rune2: {2, 2},
				},
			},
		},
		"success with complex automaton": {
			inputAutomaton: IntermediateAutomaton{
				States:       &numStates,
				StartState:   &validState,
				AcceptStates: []State{validState},
				Transitions: []Transition{
					{
						Input:     []rune(input1),
						OutStates: []State{validState, validState},
					},
					{
						Input:     []rune(input3),
						OutStates: []State{validState, validState},
					},
				},
			},
			outputAutomaton: &Automaton{
				StartState:   int(validState),
				AcceptStates: []int{int(validState)},
				TransitionFunction: map[rune][]int{
					rune1: {2, 2},
					rune3: {2, 2},
					rune4: {2, 2},
				},
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got, gotErr := convertToAutomaton(&tc.inputAutomaton)
			if tc.wantedErrPrefix != "" {
				require.ErrorContains(t, gotErr, tc.wantedErrPrefix)
			} else {
				require.NoError(t, gotErr)
				require.Equal(t, tc.outputAutomaton, got)
			}
		})
	}
}
