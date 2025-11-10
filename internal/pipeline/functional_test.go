package pipeline

import (
	"go-reloaded/pkg/tokenizer"
	"testing"
)

func TestPipelineEndToEnd(t *testing.T) {
	// Test that the pipeline correctly processes a complex input
	input := "it (cap) was 42 (hex) and 10 (bin), a honor to ' test ' this."

	// Tokenize
	tokens := tokenizer.Tokenize(input)

	// Process through pipeline
	pl := New()
	processed := pl.Process(tokens)

	// Rebuild string
	var result string
	for _, token := range processed {
		result += token.Value
	}

	// Verify transformations occurred:
	// - it (cap) -> It
	// - 42 (hex) -> 66
	// - 10 (bin) -> 2
	// - a honor -> an honor
	// - ' test ' -> 'test'
	expected := "It was 66 and 2, an honor to' test' this."

	if result != expected {
		t.Errorf("Pipeline end-to-end test failed:\n got  %q\n want %q", result, expected)
	}
}

func TestPipelineProcessorOrder(t *testing.T) {
	// Test that processors run in correct order
	// HexBin -> Case -> Article -> Quote -> Punctuation

	input := "a 1e (hex) (cap) test"
	tokens := tokenizer.Tokenize(input)

	pl := New()
	result := pl.Process(tokens)

	// Rebuild string
	var output string
	for _, token := range result {
		output += token.Value
	}

	// Should be: "an 30 Test" (article correction, hex conversion, capitalization)
	expected := "an 30 Test"

	if output != expected {
		t.Errorf("Processor order test failed:\n got  %q\n want %q", output, expected)
	}
}
