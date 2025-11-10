package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestCLIIntegration(t *testing.T) {
	// Build the binary for testing
	tmpDir := t.TempDir()
	binaryPath := filepath.Join(tmpDir, "textfmt")

	cmd := exec.Command("go", "build", "-o", binaryPath, ".")
	if err := cmd.Run(); err != nil {
		t.Fatalf("Failed to build binary: %v", err)
	}

	tests := []struct {
		name           string
		args           []string
		inputContent   string
		expectedOutput string
		expectedExit   int
		expectStderr   bool
	}{
		{
			name:           "successful processing",
			args:           []string{"input.txt", "output.txt"},
			inputContent:   "hello (cap) world, a amazing ' test '",
			expectedOutput: "Hello world, an amazing' test'",
			expectedExit:   0,
			expectStderr:   false,
		},
		{
			name:         "missing arguments",
			args:         []string{},
			expectedExit: 1,
			expectStderr: false, // Current implementation just exits
		},
		{
			name:         "too few arguments",
			args:         []string{"input.txt"},
			expectedExit: 1,
			expectStderr: false,
		},
		{
			name:         "too many arguments",
			args:         []string{"input.txt", "output.txt", "extra.txt"},
			expectedExit: 1,
			expectStderr: false,
		},
		{
			name:         "nonexistent input file",
			args:         []string{"nonexistent.txt", "output.txt"},
			expectedExit: 1,
			expectStderr: true,
		},
		{
			name:           "empty input file",
			args:           []string{"input.txt", "output.txt"},
			inputContent:   "",
			expectedOutput: "",
			expectedExit:   0,
			expectStderr:   false,
		},
		{
			name:           "complex transformation",
			args:           []string{"input.txt", "output.txt"},
			inputContent:   "there (cap) once was a hero named link (cap, 3), he carried 1e (hex) rupees and 10 (bin) arrows.",
			expectedOutput: "There once was an Hero Named Link, he carried 30 rupees and 2 arrows.",
			expectedExit:   0,
			expectStderr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testDir := t.TempDir()

			var inputPath, outputPath string
			if len(tt.args) >= 1 {
				inputPath = filepath.Join(testDir, tt.args[0])
			}
			if len(tt.args) >= 2 {
				outputPath = filepath.Join(testDir, tt.args[1])
			}

			// Create input file if content is provided
			if tt.inputContent != "" || tt.expectedOutput != "" {
				if err := os.WriteFile(inputPath, []byte(tt.inputContent), 0644); err != nil {
					t.Fatalf("Failed to create input file: %v", err)
				}
			}

			// Prepare command arguments
			var cmdArgs []string
			for i, arg := range tt.args {
				if i == 0 && inputPath != "" {
					cmdArgs = append(cmdArgs, inputPath)
				} else if i == 1 && outputPath != "" {
					cmdArgs = append(cmdArgs, outputPath)
				} else {
					cmdArgs = append(cmdArgs, arg)
				}
			}

			// Run the command
			cmd := exec.Command(binaryPath, cmdArgs...)
			output, err := cmd.CombinedOutput()

			// Check exit code
			exitCode := 0
			if err != nil {
				if exitError, ok := err.(*exec.ExitError); ok {
					exitCode = exitError.ExitCode()
				} else {
					t.Fatalf("Unexpected error running command: %v", err)
				}
			}

			if exitCode != tt.expectedExit {
				t.Errorf("Expected exit code %d, got %d", tt.expectedExit, exitCode)
			}

			// Check stderr output
			if tt.expectStderr && len(output) == 0 {
				t.Errorf("Expected stderr output, but got none")
			}

			// Check output file content if successful
			if tt.expectedExit == 0 && tt.expectedOutput != "" {
				if outputPath == "" {
					t.Fatal("No output path specified for successful test")
				}

				content, err := os.ReadFile(outputPath)
				if err != nil {
					t.Fatalf("Failed to read output file: %v", err)
				}

				got := strings.TrimSpace(string(content))
				if got != tt.expectedOutput {
					t.Errorf("Expected output %q, got %q", tt.expectedOutput, got)
				}
			}
		})
	}
}

func TestCLIErrorMessages(t *testing.T) {
	tmpDir := t.TempDir()
	binaryPath := filepath.Join(tmpDir, "textfmt")

	// Build binary
	cmd := exec.Command("go", "build", "-o", binaryPath, ".")
	if err := cmd.Run(); err != nil {
		t.Fatalf("Failed to build binary: %v", err)
	}

	// Test error message for nonexistent file
	cmd = exec.Command(binaryPath, "nonexistent.txt", "output.txt")
	output, err := cmd.CombinedOutput()

	if err == nil {
		t.Error("Expected command to fail for nonexistent input file")
	}

	if !strings.Contains(string(output), "Error:") {
		t.Errorf("Expected error message to contain 'Error:', got: %s", string(output))
	}
}

func TestCLIUsage(t *testing.T) {
	tmpDir := t.TempDir()
	binaryPath := filepath.Join(tmpDir, "textfmt")

	// Build binary
	cmd := exec.Command("go", "build", "-o", binaryPath, ".")
	if err := cmd.Run(); err != nil {
		t.Fatalf("Failed to build binary: %v", err)
	}

	// Test with no arguments
	cmd = exec.Command(binaryPath)
	err := cmd.Run()

	if err == nil {
		t.Error("Expected command to fail with no arguments")
	}

	if exitError, ok := err.(*exec.ExitError); ok {
		if exitError.ExitCode() != 1 {
			t.Errorf("Expected exit code 1, got %d", exitError.ExitCode())
		}
	}
}
