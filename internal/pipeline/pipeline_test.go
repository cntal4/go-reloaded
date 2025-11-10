package pipeline

import (
	"reflect"
	"testing"

	"go-reloaded/pkg/tokenizer"
)

func TestPipeline_Process(t *testing.T) {
	input := []tokenizer.Token{
		{Type: tokenizer.Word, Value: "a"},
		{Type: tokenizer.Space, Value: " "},
		{Type: tokenizer.Word, Value: "amazing"},
		{Type: tokenizer.Space, Value: " "},
		{Type: tokenizer.Punct, Value: "'"},
		{Type: tokenizer.Space, Value: " "},
		{Type: tokenizer.Word, Value: "idea"},
		{Type: tokenizer.Space, Value: " "},
		{Type: tokenizer.Punct, Value: "'"},
		{Type: tokenizer.Punct, Value: "."},
	}

	want := []tokenizer.Token{
		{Type: tokenizer.Word, Value: "an"},
		{Type: tokenizer.Space, Value: " "},
		{Type: tokenizer.Word, Value: "amazing"},
		{Type: tokenizer.Space, Value: " "},
		{Type: tokenizer.Punct, Value: "'"},
		{Type: tokenizer.Word, Value: "idea"},
		{Type: tokenizer.Punct, Value: "'"},
		{Type: tokenizer.Punct, Value: "."},
	}

	pl := New()
	got := pl.Process(input)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Pipeline Process():\n got  %#v\n want %#v", got, want)
	}
}
