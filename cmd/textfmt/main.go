package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"go-reloaded/internal/pipeline"
	"go-reloaded/pkg/tokenizer"
)

func main() {
	if len(os.Args) != 3 {
		os.Exit(1)
	}

	inPath := os.Args[1]
	outPath := os.Args[2]

	if err := run(inPath, outPath); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run(inPath, outPath string) error {
	input, err := ioutil.ReadFile(inPath)
	if err != nil {
		return fmt.Errorf("reading input: %w", err)
	}

	// Tokenize
	tokens := tokenizer.Tokenize(string(input))
for _, t := range tokens {
    fmt.Printf("%v %q\n", t.Type, t.Value)

}
	// Process through pipeline
	pl := pipeline.New()
	processed := pl.Process(tokens)

	// Rebuild string from tokens
	var result string
	for _, t := range processed {
		result += t.Value
	}

	// Write output
	if err := os.WriteFile(outPath, []byte(result), 0644); err != nil {
		return fmt.Errorf("writing output: %w", err)
	}
	return nil
}
