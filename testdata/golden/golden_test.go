package golden

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"go-reloaded/internal/pipeline"
	"go-reloaded/pkg/tokenizer"
)

// formatText applies the full formatting pipeline to input text
func formatText(input string) string {
	tokens := tokenizer.Tokenize(input)
	pl := pipeline.New()
	processed := pl.Process(tokens)

	var result string
	for _, token := range processed {
		result += token.Value
	}
	return result
}

func TestGoldenTests(t *testing.T) {
	goldenDir := "."

	// Find all input files
	files, err := filepath.Glob(filepath.Join(goldenDir, "*_input.txt"))
	if err != nil {
		t.Fatalf("Failed to find golden input files: %v", err)
	}

	if len(files) == 0 {
		t.Skip("No golden test files found")
	}

	for _, inputFile := range files {
		// Derive output file name
		baseName := strings.TrimSuffix(filepath.Base(inputFile), "_input.txt")
		outputFile := filepath.Join(goldenDir, baseName+"_output.txt")

		t.Run(baseName, func(t *testing.T) {
			// Read input
			input, err := os.ReadFile(inputFile)
			if err != nil {
				t.Fatalf("Failed to read input file %s: %v", inputFile, err)
			}

			// Read expected output
			expectedBytes, err := os.ReadFile(outputFile)
			if err != nil {
				t.Fatalf("Failed to read output file %s: %v", outputFile, err)
			}
			expected := strings.TrimSpace(string(expectedBytes))

			// Process input
			got := formatText(string(input))
			got = strings.TrimSpace(got)

			// Compare results
			if got != expected {
				t.Errorf("Golden test %s failed:\nInput:    %q\nExpected: %q\nGot:      %q\n\nDiff:\n%s",
					baseName, string(input), expected, got, generateDiff(expected, got))
			}
		})
	}
}

// generateDiff creates a simple unified diff format
func generateDiff(expected, got string) string {
	if expected == got {
		return "No differences"
	}

	var diff strings.Builder
	diff.WriteString("--- expected\n")
	diff.WriteString("+++ got\n")

	expectedLines := strings.Split(expected, "\n")
	gotLines := strings.Split(got, "\n")

	maxLines := len(expectedLines)
	if len(gotLines) > maxLines {
		maxLines = len(gotLines)
	}

	for i := 0; i < maxLines; i++ {
		var expectedLine, gotLine string
		if i < len(expectedLines) {
			expectedLine = expectedLines[i]
		}
		if i < len(gotLines) {
			gotLine = gotLines[i]
		}

		if expectedLine != gotLine {
			if expectedLine != "" {
				diff.WriteString("- " + expectedLine + "\n")
			}
			if gotLine != "" {
				diff.WriteString("+ " + gotLine + "\n")
			}
		}
	}

	return diff.String()
}
