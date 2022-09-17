package automaton

import (
	"fmt"

	"github.com/dustin/go-humanize/english"
)

func (i IntermediateAutomaton) validate() error {
	var usedChars []rune
	start := int(*i.StartState)
	end := *i.States

	if start <= 0 || start > end {
		return fmt.Errorf("start state must be a valid state from [0-%v]: %v", end, start)
	}

	for _, acceptState := range i.AcceptStates {
		if acceptState <= 0 || int(acceptState) > end {
			return fmt.Errorf("accept states must be valid states from [0-%v]: %v", end, acceptState)
		}
	}

	for _, transitions := range i.Transitions {
		if len(transitions.OutStates) != end {
			return fmt.Errorf("transition functions must map %s %s exactly once: %v",
				english.PluralWord(end, "", "each"), english.PluralWord(end, "states", "state"), len(transitions.OutStates))
		}
		for _, outputState := range transitions.OutStates {
			if outputState <= 0 || int(outputState) > end {
				return fmt.Errorf("output state must be a valid state from [0-%v]: %v", end, outputState)
			}
		}
		for _, letter := range transitions.Input {
			for _, used := range usedChars {
				if letter == used {
					return fmt.Errorf("letters must have exactly one transition function: %U", letter)
				}
			}
			usedChars = append(usedChars, letter)
		}
	}

	return nil
}
