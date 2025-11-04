package processors

import (
	"reflect"
	"testing"

	"go-reloaded/pkg/tokenizer"
)

func TestHexBinProcessor_Process(t *testing.T) {
	tests := []struct {
		name  string
		input []tokenizer.Token
		want  []tokenizer.Token
	}{
		{
			name: "hex to decimal",
			input: []tokenizer.Token{
				{Type: tokenizer.Word, Value: "1E"},
				{Type: tokenizer.Space, Value: " "},
				{Type: tokenizer.Marker, Value: "(hex)"},
			},
			want: []tokenizer.Token{
				{Type: tokenizer.Word, Value: "30"},
				{Type: tokenizer.Space, Value: " "},
			},
		},
		{
			name: "bin to decimal",
			input: []tokenizer.Token{
				{Type: tokenizer.Word, Value: "10"},
				{Type: tokenizer.Space, Value: " "},
				{Type: tokenizer.Marker, Value: "(bin)"},
			},
			want: []tokenizer.Token{
				{Type: tokenizer.Word, Value: "2"},
				{Type: tokenizer.Space, Value: " "},
			},
		},
	}

	for _, tt := range tests {
		p := HexBinProcessor{}
		got := p.Process(tt.input)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%s: got %#v, want %#v", tt.name, got, tt.want)
		}
	}
}
