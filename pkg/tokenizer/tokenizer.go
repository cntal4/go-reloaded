package tokenizer

import (
	"strings"
	"unicode"
)

type TokenType int

const (
	Word TokenType = iota
	Marker
	Punctuation
	Quote
	Space
)

type Token struct {
	Type  TokenType
	Value string
}

func Tokenize(text string) []Token {
	if text == "" {
		return nil
	}

	var tokens []Token
	var current strings.Builder
	currentType := Space

	flush := func() {
		if current.Len() > 0 {
			tokens = append(tokens, Token{Type: currentType, Value: current.String()})
			current.Reset()
		}
	}

	runes := []rune(text)
	for i := 0; i < len(runes); {
		r := runes[i]

		// Handle space
		if unicode.IsSpace(r) {
			flush()
			tokens = append(tokens, Token{Type: Space, Value: string(r)})
			i++
			currentType = Space
			continue
		}

		// Handle marker (starts with '(' and ends with ')')
		if r == '(' {
			flush()
			j := i + 1
			for j < len(runes) && runes[j] != ')' {
				j++
			}
			if j < len(runes) {
				segment := string(runes[i : j+1])
				tokens = append(tokens, Token{Type: Marker, Value: segment})
				i = j + 1
				currentType = Marker
				continue
			}
		}

		// Handle punctuation groups
		if strings.ContainsRune("!?.,:;", r) {
			flush()
			j := i + 1
			for j < len(runes) && strings.ContainsRune("!?.,:;", runes[j]) {
				j++
			}
			group := string(runes[i:j])
			tokens = append(tokens, Token{Type: Punctuation, Value: group})
			i = j
			currentType = Punctuation
			continue
		}

		// Handle quotes
		if r == '\'' || r == '"' {
			flush()
			tokens = append(tokens, Token{Type: Quote, Value: string(r)})
			i++
			currentType = Quote
			continue
		}

		// Otherwise, it's part of a word
		if current.Len() == 0 {
			currentType = Word
		}
		current.WriteRune(r)
		i++
	}

	flush()
	return tokens
}
