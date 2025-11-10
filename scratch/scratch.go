package main

import (
	"fmt"
	"go-reloaded/pkg/tokenizer"
	"go-reloaded/pkg/processors"
)

func main() {
	text := "it (cap) was the best of times, it was the worst of times (up, 2)."
	tokens := tokenizer.Tokenize(text)
	fmt.Println("---- INPUT ----")
	for _, t := range tokens {
		fmt.Printf("%d %q\n", t.Type, t.Value)
	}

	caseProc := processors.CaseProcessor{}
	out := caseProc.Process(tokens)

	fmt.Println("---- OUTPUT ----")
	for _, t := range out {
		fmt.Printf("%d %q\n", t.Type, t.Value)
	}
}
