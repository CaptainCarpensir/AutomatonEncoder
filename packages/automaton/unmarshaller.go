package automaton

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

// State represents an encoded digit for an automaton.
type State int

// Transition represents the set of state transitions a finite automaton can make given a specified input character.
type Transition struct {
	Input     []rune  `yaml:"input"`
	OutStates []State `yaml:"outputs"`
}

// IntermediateAutomaton is a human accessible representation of a finite automaton.
type IntermediateAutomaton struct {
	States       *int    `yaml:"states"`
	StartState   *State  `yaml:"initial"`
	AcceptStates []State `yaml:"finalstates"`

	Transitions []Transition `yaml:"transitions"`
}

// UnmarshalAutomaton deserializes the YAML input stream into an automaton object.
// If an error occurs during deserialization, then returns the error.
func UnmarshalAutomaton(in []byte) (*IntermediateAutomaton, error) {
	var a IntermediateAutomaton
	if err := yaml.Unmarshal(in, &a); err != nil {
		return nil, fmt.Errorf("unmarshal automaton yaml: %w", err)
	}
	return &a, nil
}

// UnmarshalYAML overrides the default YAML unmarshaling logic for the Transition struct allowing it to unmarshal input into a rune slice
// This method implements the yaml.Unmarshaler (v3) interface.
func (t *Transition) UnmarshalYAML(value *yaml.Node) error {
	type yamlTransition struct {
		InputString *string `yaml:"input"`
		OutStates   []State `yaml:"outputs"`
	}

	var tempStruct yamlTransition

	if err := value.Decode(&tempStruct); err != nil {
		return err
	}

	if tempStruct.InputString != nil {
		t.Input = []rune(*tempStruct.InputString)
	}
	t.OutStates = tempStruct.OutStates

	return nil
}

func convertToAutomaton(i *IntermediateAutomaton) (*Automaton, error) {
	if err := i.validate(); err != nil {
		return nil, fmt.Errorf("validate automaton: %w", err)
	}

	var a Automaton
	a.StartState = int(*i.StartState)
	for _, acceptState := range i.AcceptStates {
		a.AcceptStates = append(a.AcceptStates, int(acceptState))
	}
	a.TransitionFunction = make(map[rune][]int)
	for _, transitionGroup := range i.Transitions {
		for _, letter := range transitionGroup.Input {
			for _, outState := range transitionGroup.OutStates {
				a.TransitionFunction[letter] = append(a.TransitionFunction[letter], int(outState))
			}
		}
	}

	return &a, nil
}
