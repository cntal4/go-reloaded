package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestCLIBasicFunctionality(t *testing.T) {
	// Test the run function directly
	tmpDir := t.TempDir()
	inputPath := filepath.Join(tmpDir, "input.txt")
	outputPath := filepath.Join(tmpDir, "output.txt")

	// Create input file
	inputContent := "hello (cap) world, a amazing ' test '"
	if err := os.WriteFile(inputPath, []byte(inputContent), 0644); err != nil {
		t.Fatalf("Failed to create input file: %v", err)
	}

	// Run the formatter
	if err := run(inputPath, outputPath); err != nil {
		t.Fatalf("run() failed: %v", err)
	}

	// Check output
	output, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatalf("Failed to read output file: %v", err)
	}

	got := strings.TrimSpace(string(output))
	expected := "Hello world, an amazing' test'"

	if got != expected {
		t.Errorf("Expected %q, got %q", expected, got)
	}
}

func TestCLIErrorHandling(t *testing.T) {
	// Test error handling for nonexistent input file
	tmpDir := t.TempDir()
	inputPath := filepath.Join(tmpDir, "nonexistent.txt")
	outputPath := filepath.Join(tmpDir, "output.txt")

	err := run(inputPath, outputPath)
	if err == nil {
		t.Error("Expected error for nonexistent input file")
	}

	if !strings.Contains(err.Error(), "reading input") {
		t.Errorf("Expected error message about reading input, got: %v", err)
	}
}

func TestCLIEmptyFile(t *testing.T) {
	// Test handling of empty input file
	tmpDir := t.TempDir()
	inputPath := filepath.Join(tmpDir, "empty.txt")
	outputPath := filepath.Join(tmpDir, "output.txt")

	// Create empty input file
	if err := os.WriteFile(inputPath, []byte(""), 0644); err != nil {
		t.Fatalf("Failed to create empty input file: %v", err)
	}

	// Run the formatter
	if err := run(inputPath, outputPath); err != nil {
		t.Fatalf("run() failed on empty file: %v", err)
	}

	// Check output
	output, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatalf("Failed to read output file: %v", err)
	}

	if string(output) != "" {
		t.Errorf("Expected empty output, got %q", string(output))
	}
}

func TestCLIComplexTransformation(t *testing.T) {
	// Test complex transformation
	tmpDir := t.TempDir()
	inputPath := filepath.Join(tmpDir, "complex.txt")
	outputPath := filepath.Join(tmpDir, "output.txt")

	inputContent := "there (cap) once was a hero named link (cap, 3), he carried 1e (hex) rupees and 10 (bin) arrows."
	if err := os.WriteFile(inputPath, []byte(inputContent), 0644); err != nil {
		t.Fatalf("Failed to create input file: %v", err)
	}

	// Run the formatter
	if err := run(inputPath, outputPath); err != nil {
		t.Fatalf("run() failed: %v", err)
	}

	// Check output
	output, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatalf("Failed to read output file: %v", err)
	}

	got := strings.TrimSpace(string(output))
	expected := "There once was an Hero Named Link, he carried 30 rupees and 2 arrows."

	if got != expected {
		t.Errorf("Expected %q, got %q", expected, got)
	}
}
