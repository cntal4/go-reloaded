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
		{Type: tokenizer.Punctuation, Value: "'"},
		{Type: tokenizer.Space, Value: " "},
		{Type: tokenizer.Word, Value: "idea"},
		{Type: tokenizer.Space, Value: " "},
		{Type: tokenizer.Punctuation, Value: "'"},
		{Type: tokenizer.Punctuation, Value: "."},
	}

	want := []tokenizer.Token{
		{Type: tokenizer.Word, Value: "an"},
		{Type: tokenizer.Space, Value: " "},
		{Type: tokenizer.Word, Value: "amazing"},
		{Type: tokenizer.Space, Value: " "},
		{Type: tokenizer.Punctuation, Value: "'"},
		{Type: tokenizer.Word, Value: "idea"},
		{Type: tokenizer.Punctuation, Value: "'"},
		{Type: tokenizer.Punctuation, Value: "."},
	}

	pl := New()
	got := pl.Process(input)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Pipeline Process():\n got  %#v\n want %#v", got, want)
	}
}

