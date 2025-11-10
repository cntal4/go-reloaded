// Copyright (c) 2024 go-reloaded contributors
// Licensed under the MIT License. See LICENSE file for details.

package main

import (
	"fmt"
	"os"

	"go-reloaded/internal/logger"
	"go-reloaded/internal/pipeline"
	"go-reloaded/pkg/tokenizer"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s <input-file> <output-file>\n", os.Args[0])
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
	logger.Info(fmt.Sprintf("Processing file: %s -> %s", inPath, outPath))

	input, err := os.ReadFile(inPath)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to read input file: %v", err))
		return fmt.Errorf("reading input: %w", err)
	}

	// Tokenize
	tokens := tokenizer.Tokenize(string(input))
	logger.Debug(fmt.Sprintf("Tokenized into %d tokens", len(tokens)))

	// Process through pipeline
	pl := pipeline.New()
	processed := pl.Process(tokens)
	logger.Debug(fmt.Sprintf("Pipeline processed %d tokens", len(processed)))

	// Rebuild string from tokens
	var result string
	for _, t := range processed {
		result += t.Value
	}

	// Write output
	if err := os.WriteFile(outPath, []byte(result), 0644); err != nil {
		logger.Error(fmt.Sprintf("Failed to write output file: %v", err))
		return fmt.Errorf("writing output: %w", err)
	}

	logger.Info("Processing completed successfully")
	return nil
}
