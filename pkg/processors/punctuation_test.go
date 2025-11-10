package processors

import (
	"reflect"
	"testing"

	"go-reloaded/pkg/tokenizer"
)

func TestPunctuationProcessor_Process(t *testing.T) {
	tests := []struct {
		name  string
		input []tokenizer.Token
		want  []tokenizer.Token
	}{
		{
			name: "simple comma",
			input: []tokenizer.Token{
				{Type: tokenizer.Word, Value: "word"},
				{Type: tokenizer.Space, Value: " "},
				{Type: tokenizer.Punct, Value: ","},
				{Type: tokenizer.Word, Value: "next"},
			},
			want: []tokenizer.Token{
				{Type: tokenizer.Word, Value: "word"},
				{Type: tokenizer.Punct, Value: ","},
				{Type: tokenizer.Space, Value: " "},
				{Type: tokenizer.Word, Value: "next"},
			},
		},
		{
			name: "grouped punctuation ...",
			input: []tokenizer.Token{
				{Type: tokenizer.Word, Value: "Wait"},
				{Type: tokenizer.Space, Value: " "},
				{Type: tokenizer.Punct, Value: "..."},
				{Type: tokenizer.Space, Value: " "},
				{Type: tokenizer.Word, Value: "what"},
			},
			want: []tokenizer.Token{
				{Type: tokenizer.Word, Value: "Wait"},
				{Type: tokenizer.Punct, Value: "..."},
				{Type: tokenizer.Space, Value: " "},
				{Type: tokenizer.Word, Value: "what"},
			},
		},
		{
			name: "mixed punctuation !?",
			input: []tokenizer.Token{
				{Type: tokenizer.Word, Value: "Really"},
				{Type: tokenizer.Space, Value: " "},
				{Type: tokenizer.Punct, Value: "!?"},
				{Type: tokenizer.Space, Value: " "},
				{Type: tokenizer.Word, Value: "no"},
			},
			want: []tokenizer.Token{
				{Type: tokenizer.Word, Value: "Really"},
				{Type: tokenizer.Punct, Value: "!?"},
				{Type: tokenizer.Space, Value: " "},
				{Type: tokenizer.Word, Value: "no"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PunctuationProcessor{}
			got := p.Process(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%s:\n got  %#v\n want %#v", tt.name, got, tt.want)
			}
		})
	}
}
