package testutils

import (
	"testing"
)

func BenchmarkIdempotency(b *testing.B) {
	input := "there (cap) once was a hero named link (cap, 3), he carried 1e (hex) rupees and 10 (bin) arrows. he said: ' this is a honor ' before entering a old temple (up, 2) ... amazing !?"

	// First pass to get formatted text
	formatted := FormatText(input)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Test that formatting already-formatted text is fast and consistent
		result := FormatText(formatted)
		if result != formatted {
			b.Fatalf("Idempotency failed in benchmark")
		}
	}
}

func BenchmarkFirstPass(b *testing.B) {
	input := "there (cap) once was a hero named link (cap, 3), he carried 1e (hex) rupees and 10 (bin) arrows. he said: ' this is a honor ' before entering a old temple (up, 2) ... amazing !?"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = FormatText(input)
	}
}
