package processors

import (
	"reflect"
	"testing"

	"go-reloaded/pkg/tokenizer"
)

func TestArticleProcessor_Process(t *testing.T) {
	tests := []struct {
		name  string
		input []tokenizer.Token
		want  []tokenizer.Token
	}{
		{
			name: "simple vowel",
			input: []tokenizer.Token{
				{Type: tokenizer.Word, Value: "a"},
				{Type: tokenizer.Space, Value: " "},
				{Type: tokenizer.Word, Value: "apple"},
			},
			want: []tokenizer.Token{
				{Type: tokenizer.Word, Value: "an"},
				{Type: tokenizer.Space, Value: " "},
				{Type: tokenizer.Word, Value: "apple"},
			},
		},
		{
			name: "capitalized hour",
			input: []tokenizer.Token{
				{Type: tokenizer.Word, Value: "A"},
				{Type: tokenizer.Space, Value: " "},
				{Type: tokenizer.Word, Value: "hour"},
			},
			want: []tokenizer.Token{
				{Type: tokenizer.Word, Value: "An"},
				{Type: tokenizer.Space, Value: " "},
				{Type: tokenizer.Word, Value: "hour"},
			},
		},
		{
			name: "consonant remains",
			input: []tokenizer.Token{
				{Type: tokenizer.Word, Value: "a"},
				{Type: tokenizer.Space, Value: " "},
				{Type: tokenizer.Word, Value: "banana"},
			},
			want: []tokenizer.Token{
				{Type: tokenizer.Word, Value: "a"},
				{Type: tokenizer.Space, Value: " "},
				{Type: tokenizer.Word, Value: "banana"},
			},
		},
		{
			name: "edge punctuation",
			input: []tokenizer.Token{
				{Type: tokenizer.Word, Value: "a"},
				{Type: tokenizer.Space, Value: " "},
				{Type: tokenizer.Word, Value: "honest"},
				{Type: tokenizer.Punct, Value: "."},
			},
			want: []tokenizer.Token{
				{Type: tokenizer.Word, Value: "an"},
				{Type: tokenizer.Space, Value: " "},
				{Type: tokenizer.Word, Value: "honest"},
				{Type: tokenizer.Punct, Value: "."},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ArticleProcessor{}
			got := p.Process(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%s:\n got  %#v\n want %#v", tt.name, got, tt.want)
			}
		})
	}
}
