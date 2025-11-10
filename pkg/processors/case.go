package processors

import (
	"strconv"
	"strings"

	"go-reloaded/pkg/tokenizer"
)

type CaseProcessor struct{}

func (p CaseProcessor) Process(tokens []tokenizer.Token) []tokenizer.Token {
	for i := 0; i < len(tokens); i++ {
		if tokens[i].Type != tokenizer.Marker {
			continue
		}

		val := strings.TrimSpace(tokens[i].Value)
		val = strings.TrimPrefix(val, "(")
		val = strings.TrimSuffix(val, ")")
		parts := strings.Split(val, ",")
		cmd := strings.TrimSpace(strings.ToLower(parts[0]))
		count := 1
		if len(parts) > 1 {
			if n, err := strconv.Atoi(strings.TrimSpace(parts[1])); err == nil {
				count = n
			}
		}

		// Apply transformations backwards
		for j := i - 1; j >= 0 && count > 0; j-- {
			if tokens[j].Type != tokenizer.Word {
				continue
			}

			switch cmd {
			case "up":
				tokens[j].Value = strings.ToUpper(tokens[j].Value)
			case "low":
				tokens[j].Value = strings.ToLower(tokens[j].Value)
			case "cap":
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

		// Remove marker
		tokens = append(tokens[:i], tokens[i+1:]...)
		i--
	}

	return tokens
}










