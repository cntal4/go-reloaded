package pipeline

import (
	"testing"

	"go-reloaded/pkg/tokenizer"
)

func TestIntegrationPipeline(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "hex and bin conversion",
			input: "Simply add 42 (hex) and 10 (bin) and you will see the result is 68.",
			want:  "Simply add 66 and 2 and you will see the result is 68.",
		},
		{
			name:  "case transformations",
			input: "it (cap) was the best of times, it was the worst of times (up), it was the age of wisdom, it was the age of foolishness (cap, 6).",
			want:  "It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness.",
		},
		{
			name:  "article correction",
			input: "There is no greater agony than bearing a untold story inside you.",
			want:  "There is no greater agony than bearing an untold story inside you.",
		},
		{
			name:  "punctuation spacing",
			input: "Punctuation tests are ... kinda boring ,what do you think ?",
			want:  "Punctuation tests are... kinda boring, what do you think?",
		},
		{
			name:  "quote handling",
			input: "I am exactly how they describe me: ' awesome '",
			want:  "I am exactly how they describe me: 'awesome'",
		},
		{
			name:  "comprehensive example",
			input: "there (cap) once was a hero named link (cap, 3), he carried 1e (hex) rupees and 10 (bin) arrows. he said: ' this is a honor ' before entering a old temple (up, 2).",
			want:  "There once was an Hero Named Link, he carried 30 rupees and 2 arrows. he said:' this is an honor' before entering an OLD TEMPLE.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Tokenize input
			tokens := tokenizer.Tokenize(tt.input)

			// Process through pipeline
			pl := New()
			processed := pl.Process(tokens)

			// Rebuild string
			var result string
			for _, token := range processed {
				result += token.Value
			}

			if result != tt.want {
				t.Errorf("%s:\n got  %q\n want %q", tt.name, result, tt.want)
			}
		})
	}
}
