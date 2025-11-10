package tokenizer

import (
	"strings"
)

type Token struct {
	Type  int
	Value string
}

const (
	Word = iota
	Marker
	Punct
	Space
)

func Tokenize(s string) []Token {
	runes := []rune(s)
	var tokens []Token
	var current strings.Builder

	flush := func(t int) {
		if current.Len() > 0 {
			tokens = append(tokens, Token{Type: t, Value: current.String()})
			current.Reset()
		}
	}

	detectType := func(w string) int {
		if strings.TrimSpace(w) == "" {
			return Space
		}
		return Word
	}

	for i := 0; i < len(runes); i++ {
		ch := runes[i]

		// Handle punctuation
		// Handle punctuation groups (e.g., "..." or "!?")
if strings.ContainsRune(".,!?;:", ch) {
    flush(detectType(current.String()))

    start := i
    // Combine consecutive punctuation characters
    for i+1 < len(runes) && strings.ContainsRune(".,!?;:", runes[i+1]) {
        i++
    }

    punctGroup := string(runes[start : i+1])
    tokens = append(tokens, Token{Type: Punct, Value: punctGroup})
    continue
}


		// Handle spaces
		if ch == ' ' {
			flush(detectType(current.String()))
			tokens = append(tokens, Token{Type: Space, Value: " "})
			continue
		}

		// Handle markers like (up), (low,2), (hex)
		if ch == '(' {
			flush(detectType(current.String()))
			start := i
			for i < len(runes) && runes[i] != ')' {
				i++
			}
			if i < len(runes) {
				marker := string(runes[start : i+1])
				if isMarker(marker) {
					tokens = append(tokens, Token{Type: Marker, Value: marker})
					continue
				}
			}
			tokens = append(tokens, Token{Type: Word, Value: string(runes[start : i+1])})
			continue
		}

		// Default: part of a word
		current.WriteRune(ch)
	}

	// Flush any remaining buffer
	flush(detectType(current.String()))
	return tokens
}

func isMarker(s string) bool {
	s = strings.ToLower(strings.TrimSpace(s))
	return strings.HasPrefix(s, "(up") ||
		strings.HasPrefix(s, "(low") ||
		strings.HasPrefix(s, "(cap") ||
		strings.HasPrefix(s, "(hex") ||
		strings.HasPrefix(s, "(bin")
}


