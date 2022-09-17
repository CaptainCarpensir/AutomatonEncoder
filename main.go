package main

import (
	"fmt"
	"io/ioutil"

	"github.com/CaptainCarpensir/AutomatonEncoder/packages/automaton"
)

func main() {
	var (
		input   string
		machine *automaton.Automaton
	)

	fmt.Print("Enter yaml filename: ")
	fmt.Scanln(&input)

	inStream, err := ioutil.ReadFile(input)
	if err != nil {
		fmt.Printf("read file: %e\n", err)
		return
	}

	machine, err = automaton.EncodeAutomaton(inStream)
	if err != nil {
		fmt.Printf("encode automaton: %e\n", err)
		return
	}

	for true {
		fmt.Print("Enter a word to be recognized by the automaton: ")
		fmt.Scanln(&input)

		fmt.Printf("Automaton recognizes %s: %v\n\n", input, machine.Recognize(input))
	}
}
