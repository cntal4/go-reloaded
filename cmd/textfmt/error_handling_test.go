package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestErrorHandlingGraceful(t *testing.T) {
	// Test that the program handles errors gracefully without panicking
	tmpDir := t.TempDir()

	tests := []struct {
		name        string
		setupInput  func() (string, string) // returns input path, output path
		expectError bool
	}{
		{
			name: "nonexistent input file",
			setupInput: func() (string, string) {
				return filepath.Join(tmpDir, "nonexistent.txt"), filepath.Join(tmpDir, "out.txt")
			},
			expectError: true,
		},
		{
			name: "invalid output directory",
			setupInput: func() (string, string) {
				inputPath := filepath.Join(tmpDir, "input.txt")
				os.WriteFile(inputPath, []byte("test"), 0644)
				return inputPath, "/invalid/path/output.txt"
			},
			expectError: true,
		},
		{
			name: "empty input file",
			setupInput: func() (string, string) {
				inputPath := filepath.Join(tmpDir, "empty.txt")
				outputPath := filepath.Join(tmpDir, "output.txt")
				os.WriteFile(inputPath, []byte(""), 0644)
				return inputPath, outputPath
			},
			expectError: false,
		},
		{
			name: "malformed input",
			setupInput: func() (string, string) {
				inputPath := filepath.Join(tmpDir, "malformed.txt")
				outputPath := filepath.Join(tmpDir, "output.txt")
				// Input with various edge cases that shouldn't crash
				os.WriteFile(inputPath, []byte("(hex) ' (up,abc) (cap,-1) ZZ (hex)"), 0644)
				return inputPath, outputPath
			},
			expectError: false, // Should handle gracefully
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputPath, outputPath := tt.setupInput()

			err := run(inputPath, outputPath)

			if tt.expectError && err == nil {
				t.Errorf("Expected error but got none")
			}
			if !tt.expectError && err != nil {
				t.Errorf("Expected no error but got: %v", err)
			}
		})
	}
}

func TestRecoveryFromPanic(t *testing.T) {
	// Test that the program doesn't panic on extreme edge cases
	tmpDir := t.TempDir()
	inputPath := filepath.Join(tmpDir, "extreme.txt")
	outputPath := filepath.Join(tmpDir, "output.txt")

	// Create input with extreme edge cases
	extremeInput := string(make([]byte, 1000)) + "(cap)" + string(make([]byte, 1000))
	os.WriteFile(inputPath, []byte(extremeInput), 0644)

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Program panicked: %v", r)
		}
	}()

	// Should not panic
	_ = run(inputPath, outputPath)
}
