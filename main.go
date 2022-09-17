package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/CaptainCarpensir/AutomatonEncoder/packages/automaton"
)

func main() {
	var (
		input   string
		machine *automaton.Automaton
	)

	if len(os.Args) != 2 {
		fmt.Println("command requires one filename argument")
		return
	}

	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("read file: %e\n", err)
		return
	}

	machine, err = automaton.EncodeAutomaton(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print("Enter a word to be recognized by the automaton: ")
	for {
		fmt.Print("\033[s")
		fmt.Scanln(&input)

		inLang := machine.Recognize(input)

		color := 32
		if !inLang {
			color = 31
		}

		fmt.Printf("Automaton recognizes '%s': \033[%vm%v", input, color, inLang)
		fmt.Print("\033[K")
		fmt.Print("\033[u")
		fmt.Print("\033[K")
	}
}
