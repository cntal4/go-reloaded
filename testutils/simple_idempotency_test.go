package testutils

import (
	"testing"
)

func TestSimpleIdempotency(t *testing.T) {
	// Test basic idempotency with known inputs
	testCases := []string{
		"Hello world",                       // already formatted
		"hello (cap) world",                 // needs formatting
		"Simply add 66 and 2.",              // already formatted numbers
		"Simply add 42 (hex) and 10 (bin).", // needs hex/bin conversion
	}

	for _, input := range testCases {
		first := FormatText(input)
		second := FormatText(first)

		if first != second {
			t.Errorf("Idempotency failed:\nInput: %q\nFirst: %q\nSecond: %q",
				input, first, second)
		}
	}
}

func TestIdempotencyVerification(t *testing.T) {
	// Verify that our formatter produces the expected behavior
	input := "hello (cap) world"
	expected := "Hello world"

	result := FormatText(input)
	if result != expected {
		t.Errorf("Formatting failed: got %q, want %q", result, expected)
	}

	// Test idempotency
	secondResult := FormatText(result)
	if result != secondResult {
		t.Errorf("Idempotency failed: first %q, second %q", result, secondResult)
	}
}
