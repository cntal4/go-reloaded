package testutils

import (
	"go-reloaded/internal/pipeline"
	"go-reloaded/pkg/tokenizer"
	"testing"
)

// FormatText applies the full formatting pipeline to input text
func FormatText(input string) string {
	tokens := tokenizer.Tokenize(input)
	pl := pipeline.New()
	processed := pl.Process(tokens)

	var result string
	for _, token := range processed {
		result += token.Value
	}
	return result
}

func TestIdempotency(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{
			name:  "simple case transformation",
			input: "hello (cap) world",
		},
		{
			name:  "hex and bin conversion",
			input: "add 42 (hex) and 10 (bin)",
		},
		{
			name:  "article correction",
			input: "a apple and a orange",
		},
		{
			name:  "punctuation normalization",
			input: "wait ... what !? really",
		},
		{
			name:  "quote handling",
			input: "he said ' hello world '",
		},
		{
			name:  "complex mixed transformations",
			input: "there (cap) once was a hero named link (cap, 3), he carried 1e (hex) rupees. he said: ' this is a honor ' before entering a old temple (up, 2) ... amazing !?",
		},
		{
			name:  "contractions and quotes",
			input: "I can't believe ' it isn't working ' properly.",
		},
		{
			name:  "already formatted text",
			input: "This is already formatted text with proper punctuation, quotes and spacing.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// First formatting pass
			firstPass := FormatText(tt.input)

			// Second formatting pass (should be identical)
			secondPass := FormatText(firstPass)

			if firstPass != secondPass {
				t.Errorf("Idempotency test failed for %s:\n"+
					"Input:       %q\n"+
					"First pass:  %q\n"+
					"Second pass: %q\n"+
					"Expected first and second pass to be identical",
					tt.name, tt.input, firstPass, secondPass)
			}
		})
	}
}

func TestMultiplePassIdempotency(t *testing.T) {
	// Test that multiple passes (3+) also yield identical results
	input := "there (cap) once was a hero named link (cap, 3), he carried 1e (hex) rupees and 10 (bin) arrows. he said: ' this is a honor ' before entering a old temple (up, 2) ... amazing !?"

	passes := make([]string, 5)
	passes[0] = input

	for i := 1; i < 5; i++ {
		passes[i] = FormatText(passes[i-1])
	}

	// All passes after the first should be identical
	for i := 2; i < 5; i++ {
		if passes[1] != passes[i] {
			t.Errorf("Multiple pass idempotency failed:\n"+
				"Pass 1: %q\n"+
				"Pass %d: %q\n"+
				"Expected all passes after first to be identical", passes[1], i, passes[i])
		}
	}
}

func TestIdempotencyWithEdgeCases(t *testing.T) {
	edgeCases := []string{
		"",                           // empty string
		" ",                          // single space
		"word",                       // single word
		"(cap)",                      // marker only
		"' '",                        // empty quotes
		"...",                        // punctuation only
		"a a a a a",                  // repeated articles
		"isn't can't won't",          // multiple contractions
		"' nested ' quotes ' test '", // multiple quotes
	}

	for _, input := range edgeCases {
		t.Run("edge_case_"+input, func(t *testing.T) {
			firstPass := FormatText(input)
			secondPass := FormatText(firstPass)

			if firstPass != secondPass {
				t.Errorf("Edge case idempotency failed:\n"+
					"Input: %q\n"+
					"First:  %q\n"+
					"Second: %q", input, firstPass, secondPass)
			}
		})
	}
}
