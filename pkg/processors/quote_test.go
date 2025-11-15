package processors

import (
	"reflect"
	"testing"

	"go-reloaded/pkg/tokenizer"
)

func TestQuoteProcessor_Process(t *testing.T) {
	tests := []struct {
		name  string
		input []tokenizer.Token
		want  []tokenizer.Token
	}{
		{
			name: "single word in quotes",
			input: []tokenizer.Token{
				{Type: tokenizer.Punct, Value: "'"},
				{Type: tokenizer.Space, Value: " "},
				{Type: tokenizer.Word, Value: "awesome"},
				{Type: tokenizer.Space, Value: " "},
				{Type: tokenizer.Punct, Value: "'"},
			},
			want: []tokenizer.Token{
				{Type: tokenizer.Punct, Value: "'"},
				{Type: tokenizer.Word, Value: "awesome"},
				{Type: tokenizer.Punct, Value: "'"},
			},
		},
		{
			name: "multi-word quote",
			input: []tokenizer.Token{
				{Type: tokenizer.Punct, Value: "'"},
				{Type: tokenizer.Space, Value: " "},
				{Type: tokenizer.Word, Value: "I"},
				{Type: tokenizer.Space, Value: " "},
				{Type: tokenizer.Word, Value: "am"},
				{Type: tokenizer.Space, Value: " "},
				{Type: tokenizer.Word, Value: "the"},
				{Type: tokenizer.Space, Value: " "},
				{Type: tokenizer.Word, Value: "boss"},
				{Type: tokenizer.Space, Value: " "},
				{Type: tokenizer.Punct, Value: "'"},
			},
			want: []tokenizer.Token{
				{Type: tokenizer.Punct, Value: "'"},
				{Type: tokenizer.Word, Value: "I"},
				{Type: tokenizer.Space, Value: " "},
				{Type: tokenizer.Word, Value: "am"},
				{Type: tokenizer.Space, Value: " "},
				{Type: tokenizer.Word, Value: "the"},
				{Type: tokenizer.Space, Value: " "},
				{Type: tokenizer.Word, Value: "boss"},
				{Type: tokenizer.Punct, Value: "'"},
			},
		},
		{
			name: "double quotes with spaces",
			input: []tokenizer.Token{
				{Type: tokenizer.Punct, Value: "\""},
				{Type: tokenizer.Space, Value: " "},
				{Type: tokenizer.Word, Value: "It"},
				{Type: tokenizer.Space, Value: " "},
				{Type: tokenizer.Word, Value: "was"},
				{Type: tokenizer.Space, Value: " "},
				{Type: tokenizer.Word, Value: "fun"},
				{Type: tokenizer.Punct, Value: "!"},
				{Type: tokenizer.Space, Value: " "},
				{Type: tokenizer.Punct, Value: "\""},
			},
			want: []tokenizer.Token{
				{Type: tokenizer.Punct, Value: "\""},
				{Type: tokenizer.Word, Value: "It"},
				{Type: tokenizer.Space, Value: " "},
				{Type: tokenizer.Word, Value: "was"},
				{Type: tokenizer.Space, Value: " "},
				{Type: tokenizer.Word, Value: "fun"},
				{Type: tokenizer.Punct, Value: "!"},
				{Type: tokenizer.Punct, Value: "\""},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := QuoteProcessor{}
			got := p.Process(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%s:\n got  %#v\n want %#v", tt.name, got, tt.want)
			}
		})
	}
}
