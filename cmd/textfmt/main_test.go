package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestCLI_Run(t *testing.T) {
	tmpDir := t.TempDir()
	inPath := filepath.Join(tmpDir, "in.txt")
	outPath := filepath.Join(tmpDir, "out.txt")

	// input simulates: "a amazing idea"
	if err := os.WriteFile(inPath, []byte("a amazing idea"), 0644); err != nil {
		t.Fatalf("failed to write input: %v", err)
	}

	// run the program
	if err := run(inPath, outPath); err != nil {
		t.Fatalf("run() error = %v", err)
	}

	out, err := os.ReadFile(outPath)
	if err != nil {
		t.Fatalf("reading out: %v", err)
	}

	got := strings.TrimSpace(string(out))
	want := "an amazing idea"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

