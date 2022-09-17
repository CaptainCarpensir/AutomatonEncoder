package automaton

import "fmt"

// Automaton is a 3-Tuple representing a deterministic finite automaton.
type Automaton struct {
	StartState         int
	AcceptStates       []int
	TransitionFunction map[rune][]int
}

// Recognize returns true if the given word is recognized by the language represented by the Automaton.
func (a *Automaton) Recognize(word string) bool {
	state := a.StartState

	// Move through finite state machine on word input
	for _, letter := range word {
		stateFunction, ok := a.TransitionFunction[letter]
		if !ok {
			return false
		}
		state = stateFunction[state-1] // Offset state as yaml is 1-indexed but Golang is 0-indexed
	}

	// Check if state is in the set of accept states
	for _, acceptState := range a.AcceptStates {
		if state == acceptState {
			return true
		}
	}
	return false
}

// EncodeAutomaton reads a yaml ByteStream definition of an automata and logically encodes it into an Automata interface.
func EncodeAutomaton(inStream []byte) (*Automaton, error) {
	intermediate, err := UnmarshalAutomaton(inStream)
	if err != nil {
		return nil, fmt.Errorf("unmarshal automaton: %e", err)
	}

	automaton, err := convertToAutomaton(intermediate)
	if err != nil {
		return nil, fmt.Errorf("encode automaton: %e", err)
	}

	return automaton, nil
}
