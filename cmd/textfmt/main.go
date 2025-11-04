package main

import (
	"errors"
	"fmt"
	"os"
)

// main is the entry point of the CLI.
func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

// run validates arguments and will later trigger the formatter pipeline.
func run() error {
	if len(os.Args) != 3 {
		return errors.New("usage: textfmt <input> <output>")
	}
	return nil
}
