package processors

import (
	"go-reloaded/pkg/tokenizer"
)

type PunctuationProcessor struct{}

func (p PunctuationProcessor) Process(tokens []tokenizer.Token) []tokenizer.Token {
	out := make([]tokenizer.Token, 0, len(tokens))

	for i := 0; i < len(tokens); i++ {
		tok := tokens[i]

		if tok.Type == tokenizer.Punct {
			// Remove space before punctuation
			if len(out) > 0 && out[len(out)-1].Type == tokenizer.Space {
				out = out[:len(out)-1]
			}
			out = append(out, tok)
			continue
		}

		// Add space after punctuation if current token is not space and previous was punctuation
		if len(out) > 0 && out[len(out)-1].Type == tokenizer.Punct && tok.Type != tokenizer.Space {
			out = append(out, tokenizer.Token{Type: tokenizer.Space, Value: " "})
		}

		// Skip multiple consecutive spaces
		if tok.Type == tokenizer.Space && len(out) > 0 && out[len(out)-1].Type == tokenizer.Space {
			continue
		}

		out = append(out, tok)
	}

	return out
}
