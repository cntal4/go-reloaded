package pipeline

import (
	"reflect"
	"testing"

	"go-reloaded/pkg/tokenizer"
)

type dummyProcessor struct{}

func (d dummyProcessor) Process(tokens []tokenizer.Token) []tokenizer.Token {
	for i := range tokens {
		if tokens[i].Type == tokenizer.Word {
			tokens[i].Value = "X" + tokens[i].Value
		}
	}
	return tokens
}

func TestPipeline_Run(t *testing.T) {
	p := New()
	p.Add(dummyProcessor{})

	input := []tokenizer.Token{
		{tokenizer.Word, "Hello"},
		{tokenizer.Space, " "},
		{tokenizer.Word, "world"},
	}
	want := []tokenizer.Token{
		{tokenizer.Word, "XHello"},
		{tokenizer.Space, " "},
		{tokenizer.Word, "Xworld"},
	}

	got := p.Run(input)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Pipeline.Run() = %#v, want %#v", got, want)
	}
}
