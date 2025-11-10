package processors

import (
	"strings"

	"go-reloaded/pkg/tokenizer"
)

type ArticleProcessor struct{}

func (p ArticleProcessor) Process(tokens []tokenizer.Token) []tokenizer.Token {
	if len(tokens) < 2 {
		return tokens
	}

	out := make([]tokenizer.Token, 0, len(tokens))

	for i := 0; i < len(tokens); i++ {
		tok := tokens[i]

		// Look ahead to the next non-space word
		if tok.Type == tokenizer.Word &&
			(strings.EqualFold(tok.Value, "a")) &&
			i+1 < len(tokens) {

			// Find next word (skip spaces)
			j := i + 1
			for j < len(tokens) && tokens[j].Type == tokenizer.Space {
				j++
			}

			if j < len(tokens) && tokens[j].Type == tokenizer.Word {
				nextWord := tokens[j].Value
				if startsWithVowelOrH(nextWord) {
					if tok.Value == "A" {
						tok.Value = "An"
					} else {
						tok.Value = "an"
					}
				}
			}
		}

		out = append(out, tok)
	}

	return out
}

func startsWithVowelOrH(s string) bool {
	if s == "" {
		return false
	}
	r := strings.ToLower(string([]rune(s)[0]))
	return strings.ContainsAny(r, "aeiouh")
}
