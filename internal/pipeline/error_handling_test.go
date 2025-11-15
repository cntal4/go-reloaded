package pipeline

import (
	"testing"

	"go-reloaded/pkg/tokenizer"
)

func TestPipelineErrorHandling(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect string // Should still produce reasonable output
	}{
		{
			name:   "invalid hex marker",
			input:  "invalid ZZ (hex) number",
			expect: "invalid ZZ number", // Should remove marker even if conversion fails
		},
		{
			name:   "invalid bin marker",
			input:  "invalid 22 (bin) number",
			expect: "invalid 22 number", // Should remove marker even if conversion fails
		},
		{
			name:   "malformed marker",
			input:  "test (invalid marker) text",
			expect: "test (invalid marker) text", // Should leave unknown markers alone
		},
		{
			name:   "empty quotes",
			input:  "test ' ' empty",
			expect: "test'' empty", // Should handle empty quotes gracefully
		},
		{
			name:   "unmatched quotes",
			input:  "test ' unmatched quote",
			expect: "test' unmatched quote", // Should handle gracefully
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokens := tokenizer.Tokenize(tt.input)
			pl := New()

			// Should not panic
			processed := pl.Process(tokens)

			// Rebuild string
			var result string
			for _, token := range processed {
				result += token.Value
			}

			if result != tt.expect {
				t.Errorf("Expected %q, got %q", tt.expect, result)
			}
		})
	}
}

func TestPipelineStability(t *testing.T) {
	// Test that pipeline doesn't crash on edge cases
	edgeCases := []string{
		"",          // empty string
		"(hex)",     // marker only
		"' ",        // incomplete quote
		"...",       // punctuation only
		"(up,abc)",  // invalid number in marker
		"(cap,-1)",  // negative number
		"(low,999)", // very large number
	}

	for _, input := range edgeCases {
		t.Run("edge_case", func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("Pipeline panicked on input %q: %v", input, r)
				}
			}()

			tokens := tokenizer.Tokenize(input)
			pl := New()
			processed := pl.Process(tokens)

			// Should produce some output (even if empty)
			var result string
			for _, token := range processed {
				result += token.Value
			}

			// Just verify it doesn't crash
			t.Logf("Input: %q -> Output: %q", input, result)
		})
	}
}
