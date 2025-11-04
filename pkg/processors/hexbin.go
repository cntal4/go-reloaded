package processors

import (
	"fmt"
	"strconv"
	"strings"

	"go-reloaded/pkg/tokenizer"
)

// HexBinProcessor converts hexadecimal or binary numbers before markers to decimal.
type HexBinProcessor struct{}

func (p HexBinProcessor) Process(tokens []tokenizer.Token) []tokenizer.Token {
	if len(tokens) == 0 {
		return tokens
	}

	out := make([]tokenizer.Token, 0, len(tokens))
	for i := 0; i < len(tokens); i++ {
		tok := tokens[i]

		// Look ahead one or two tokens for the marker, skipping spaces
		if i+1 < len(tokens) {
			next := tokens[i+1]

			// Case 1: direct marker after word (no space)
			if next.Type == tokenizer.Marker {
				tok = convertIfNeeded(tok, next)
				i++ // skip marker
			} else if next.Type == tokenizer.Space && i+2 < len(tokens) && tokens[i+2].Type == tokenizer.Marker {
				tok = convertIfNeeded(tok, tokens[i+2])
				// keep the space
				out = append(out, tok, next)
				i += 2 // skip word + space + marker
				continue
			}
		}

		out = append(out, tok)
	}
	return out
}

// Helper that performs the numeric conversion
func convertIfNeeded(tok tokenizer.Token, marker tokenizer.Token) tokenizer.Token {
	val := strings.ToLower(strings.TrimSpace(marker.Value))
	word := strings.TrimSpace(tok.Value)

	var n int64
	var err error

	switch val {
	case "(hex)":
		n, err = strconv.ParseInt(word, 16, 64)
	case "(bin)":
		n, err = strconv.ParseInt(word, 2, 64)
	default:
		return tok
	}

	if err == nil {
		tok.Value = fmt.Sprintf("%d", n)
	}
	return tok
}
