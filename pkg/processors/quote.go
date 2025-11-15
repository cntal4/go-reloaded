package processors

import "go-reloaded/pkg/tokenizer"

// QuoteProcessor trims spaces immediately inside paired quotes (both single and double).
type QuoteProcessor struct{}

func (p QuoteProcessor) Process(tokens []tokenizer.Token) []tokenizer.Token {
	out := make([]tokenizer.Token, 0, len(tokens))
	var openQuote string // Track which type of quote is open

	for i := 0; i < len(tokens); i++ {
		tok := tokens[i]

		if tok.Type == tokenizer.Punct && (tok.Value == "'" || tok.Value == "\"") {
			if openQuote == tok.Value {
				// Closing quote - remove preceding space if any
				if len(out) > 0 && out[len(out)-1].Type == tokenizer.Space {
					out = out[:len(out)-1]
				}
				out = append(out, tok)
				openQuote = ""
			} else if openQuote == "" {
				// Opening quote
				out = append(out, tok)
				// Skip following space if present
				if i+1 < len(tokens) && tokens[i+1].Type == tokenizer.Space {
					i++
				}
				openQuote = tok.Value
			} else {
				// Different quote type while another is open - treat as regular punctuation
				out = append(out, tok)
			}
			continue
		}

		out = append(out, tok)
	}
	return out
}
