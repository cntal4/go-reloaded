package docs

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestREADMEExamples(t *testing.T) {
	// Test that the basic usage example from README works
	t.Run("basic_usage", func(t *testing.T) {
		tmpDir := t.TempDir()
		inputPath := filepath.Join(tmpDir, "input.txt")
		outputPath := filepath.Join(tmpDir, "output.txt")

		// Create sample input
		sampleInput := "hello (cap) world, a amazing ' test '"
		if err := os.WriteFile(inputPath, []byte(sampleInput), 0644); err != nil {
			t.Fatalf("Failed to create input file: %v", err)
		}

		// Run the command as documented in README
		cmd := exec.Command("go", "run", "./cmd/textfmt", inputPath, outputPath)
		cmd.Dir = ".." // Go up one directory to project root
		if err := cmd.Run(); err != nil {
			t.Fatalf("README example command failed: %v", err)
		}

		// Verify output exists and is correct
		output, err := os.ReadFile(outputPath)
		if err != nil {
			t.Fatalf("Failed to read output file: %v", err)
		}

		expected := "Hello world, an amazing' test'"
		got := strings.TrimSpace(string(output))
		if got != expected {
			t.Errorf("README example output mismatch:\nExpected: %q\nGot:      %q", expected, got)
		}
	})

	t.Run("examples_directory", func(t *testing.T) {
		// Test that examples in docs/examples/ work correctly
		examplesDir := "examples"

		// Find example input files
		files, err := filepath.Glob(filepath.Join(examplesDir, "*_input.txt"))
		if err != nil {
			t.Fatalf("Failed to find example files: %v", err)
		}

		if len(files) == 0 {
			t.Skip("No example files found")
		}

		for _, inputFile := range files {
			baseName := strings.TrimSuffix(filepath.Base(inputFile), "_input.txt")
			expectedOutputFile := filepath.Join(examplesDir, baseName+"_output.txt")

			// Check if expected output exists
			if _, err := os.Stat(expectedOutputFile); os.IsNotExist(err) {
				t.Errorf("Missing expected output file: %s", expectedOutputFile)
				continue
			}

			// Run formatter on example
			tmpOutput := filepath.Join(t.TempDir(), "output.txt")
			cmd := exec.Command("go", "run", "./cmd/textfmt", inputFile, tmpOutput)
			cmd.Dir = ".." // Go up to project root
			if err := cmd.Run(); err != nil {
				t.Errorf("Example %s failed: %v", baseName, err)
				continue
			}

			// Compare with expected output
			expected, err := os.ReadFile(expectedOutputFile)
			if err != nil {
				t.Errorf("Failed to read expected output for %s: %v", baseName, err)
				continue
			}

			actual, err := os.ReadFile(tmpOutput)
			if err != nil {
				t.Errorf("Failed to read actual output for %s: %v", baseName, err)
				continue
			}

			if strings.TrimSpace(string(expected)) != strings.TrimSpace(string(actual)) {
				t.Errorf("Example %s output mismatch:\nExpected: %q\nActual:   %q",
					baseName, strings.TrimSpace(string(expected)), strings.TrimSpace(string(actual)))
			}
		}
	})
}

func TestREADMECommandsWork(t *testing.T) {
	// Test that all commands mentioned in README actually work
	commands := []struct {
		name string
		cmd  []string
		dir  string
	}{
		{
			name: "go_test_all",
			cmd:  []string{"go", "test", "./..."},
			dir:  "..",
		},
		{
			name: "go_test_golden",
			cmd:  []string{"go", "test", "./testdata/golden"},
			dir:  "..",
		},
		{
			name: "go_build",
			cmd:  []string{"go", "build", "./cmd/textfmt"},
			dir:  "..",
		},
	}

	for _, tc := range commands {
		t.Run(tc.name, func(t *testing.T) {
			cmd := exec.Command(tc.cmd[0], tc.cmd[1:]...)
			cmd.Dir = tc.dir
			if err := cmd.Run(); err != nil {
				t.Errorf("README command %v failed: %v", tc.cmd, err)
			}
		})
	}
}
