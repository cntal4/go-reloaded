package main

import (
	"os"
	"testing"
)

func TestRun_NoArgs(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"textfmt"}
	if err := run(); err == nil {
		t.Error("expected an error when no arguments are provided")
	}
}

func TestRun_WithArgs(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"textfmt", "input.txt", "output.txt"}
	if err := run(); err != nil {
		t.Fatalf("unexpected error with valid args: %v", err)
	}
}
