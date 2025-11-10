package processors

import (
	
	"go-reloaded/pkg/tokenizer"
)

type PunctuationProcessor struct{}

func (p PunctuationProcessor) Process(tokens []tokenizer.Token) []tokenizer.Token {
	out := make([]tokenizer.Token, 0, len(tokens))

	for i := 0; i < len(tokens); i++ {
		tok := tokens[i]

		// remove space before punctuation
		if tok.Type == tokenizer.Punct {
			// drop any trailing space before punctuation
			if len(out) > 0 && out[len(out)-1].Type == tokenizer.Space {
				out = out[:len(out)-1]
			}
			out = append(out, tok)

			// add a single space after punctuation if next is a word
			if i+1 < len(tokens) && tokens[i+1].Type == tokenizer.Word {
				out = append(out, tokenizer.Token{Type: tokenizer.Space, Value: " "})
			}
			continue
		}

		// collapse multiple spaces to one
		if tok.Type == tokenizer.Space {
			if len(out) > 0 && out[len(out)-1].Type == tokenizer.Space {
				continue
			}
		}

		out = append(out, tok)
	}

	// trim trailing spaces
	for len(out) > 0 && out[len(out)-1].Type == tokenizer.Space {
		out = out[:len(out)-1]
	}

	return out
}



