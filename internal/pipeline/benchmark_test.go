package pipeline

import (
	"go-reloaded/pkg/tokenizer"
	"testing"
)

func BenchmarkPipeline(b *testing.B) {
	input := "there (cap) once was a hero named link (cap, 3), he carried 1e (hex) rupees and 10 (bin) arrows. he said: ' this is a honor ' before entering a old temple (up, 2), where legends said time (low) STOPS (low, 2) ... a mysterious voice whispered: ' welcome ,hero ' and the adventure began !"

	tokens := tokenizer.Tokenize(input)
	pl := New()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = pl.Process(tokens)
	}
}

func BenchmarkTokenizeAndProcess(b *testing.B) {
	input := "there (cap) once was a hero named link (cap, 3), he carried 1e (hex) rupees and 10 (bin) arrows. he said: ' this is a honor ' before entering a old temple (up, 2), where legends said time (low) STOPS (low, 2) ... a mysterious voice whispered: ' welcome ,hero ' and the adventure began !"

	pl := New()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tokens := tokenizer.Tokenize(input)
		processed := pl.Process(tokens)

		// Rebuild string to simulate full pipeline
		var result string
		for _, token := range processed {
			result += token.Value
		}
	}
}
