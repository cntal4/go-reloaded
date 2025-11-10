package processors

import "go-reloaded/pkg/tokenizer"

// QuoteProcessor trims spaces immediately inside paired single quotes.
type QuoteProcessor struct{}

func (p QuoteProcessor) Process(tokens []tokenizer.Token) []tokenizer.Token {
	out := make([]tokenizer.Token, 0, len(tokens))
	inQuote := false

	for i := 0; i < len(tokens); i++ {
		tok := tokens[i]

		if tok.Type == tokenizer.Punct && tok.Value == "'" {
			if inQuote {
				// We're closing a quote → remove preceding space if any
				if len(out) > 0 && out[len(out)-1].Type == tokenizer.Space {
					out = out[:len(out)-1]
				}
				out = append(out, tok)
				inQuote = false
			} else {
				// Opening quote → append it
				out = append(out, tok)
				// Skip next token if it's a space (inside opening quote)
				if i+1 < len(tokens) && tokens[i+1].Type == tokenizer.Space {
					i++
				}
				inQuote = true
			}
			continue
		}

		out = append(out, tok)
	}
	return out
}
