package processors

import (
	"reflect"
	"testing"
	"go-reloaded/pkg/tokenizer"
)

func TestCaseProcessor_Process(t *testing.T) {
	tests := []struct {
		name  string
		input []tokenizer.Token
		want  []tokenizer.Token
	}{
		{
			name: "simple uppercase",
			input: []tokenizer.Token{
				{Type: tokenizer.Word, Value: "go"},
				{Type: tokenizer.Space, Value: " "},
				{Type: tokenizer.Marker, Value: "(up)"},
			},
			want: []tokenizer.Token{
				{Type: tokenizer.Word, Value: "GO"},
				{Type: tokenizer.Space, Value: " "},
			},
		},
		{
			name: "lowercase numeric",
			input: []tokenizer.Token{
				{Type: tokenizer.Word, Value: "Go"},
				{Type: tokenizer.Space, Value: " "},
				{Type: tokenizer.Word, Value: "Lang"},
				{Type: tokenizer.Space, Value: " "},
				{Type: tokenizer.Marker, Value: "(low,2)"},
			},
			want: []tokenizer.Token{
				{Type: tokenizer.Word, Value: "go"},
				{Type: tokenizer.Space, Value: " "},
				{Type: tokenizer.Word, Value: "lang"},
				{Type: tokenizer.Space, Value: " "},
			},
		},
		{
			name: "capitalize",
			input: []tokenizer.Token{
				{Type: tokenizer.Word, Value: "hello"},
				{Type: tokenizer.Space, Value: " "},
				{Type: tokenizer.Word, Value: "world"},
				{Type: tokenizer.Space, Value: " "},
				{Type: tokenizer.Marker, Value: "(cap,2)"},
			},
			want: []tokenizer.Token{
				{Type: tokenizer.Word, Value: "Hello"},
				{Type: tokenizer.Space, Value: " "},
				{Type: tokenizer.Word, Value: "World"},
				{Type: tokenizer.Space, Value: " "},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := CaseProcessor{}
			got := p.Process(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%s: got %#v, want %#v", tt.name, got, tt.want)
			}
		})
	}
}
