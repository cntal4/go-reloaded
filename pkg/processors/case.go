package processors

import (
	"fmt"
	"strings"

	"go-reloaded/pkg/tokenizer"
)

type CaseProcessor struct{}

func (p CaseProcessor) Process(tokens []tokenizer.Token) []tokenizer.Token {
	for i := 0; i < len(tokens); i++ {
		tok := tokens[i]
		if tok.Type != tokenizer.Marker {
			continue
		}

		val := strings.ToLower(strings.TrimSpace(tok.Value))
		count := 1

		// handle patterns like (up, 2)
		if strings.Contains(val, ",") {
			var parsed int
			fmt.Sscanf(val, "(%*[^,], %d)", &parsed)
			if parsed > 0 {
				count = parsed
			}
		}

		// Walk backward for target words
		for j := i - 1; j >= 0 && count > 0; j-- {
			if tokens[j].Type != tokenizer.Word {
				continue
			}

			switch {
			case strings.HasPrefix(val, "(up"):
				tokens[j].Value = strings.ToUpper(tokens[j].Value)
			case strings.HasPrefix(val, "(low"):
				tokens[j].Value = strings.ToLower(tokens[j].Value)
			case strings.HasPrefix(val, "(cap"):
				if len(tokens[j].Value) > 0 {
					first := strings.ToUpper(tokens[j].Value[:1])
					rest := ""
					if len(tokens[j].Value) > 1 {
						rest = strings.ToLower(tokens[j].Value[1:])
					}
					tokens[j].Value = first + rest
				}
			}
			count--
		}

		// Remove the marker token
		tokens = append(tokens[:i], tokens[i+1:]...)
		i-- // adjust index after removal
	}

	return tokens
}









