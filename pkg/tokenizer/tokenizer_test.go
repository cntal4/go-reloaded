package tokenizer

import (
	"reflect"
	"testing"
)

func TestTokenize_BasicSentence(t *testing.T) {
	input := "It (cap) was the best of times, it was the worst of times (up)"
	want := []Token{
		{Word, "It"},
		{Space, " "},
		{Marker, "(cap)"},
		{Space, " "},
		{Word, "was"},
		{Space, " "},
		{Word, "the"},
		{Space, " "},
		{Word, "best"},
		{Space, " "},
		{Word, "of"},
		{Space, " "},
		{Word, "times"},
		{Punctuation, ","},
		{Space, " "},
		{Word, "it"},
		{Space, " "},
		{Word, "was"},
		{Space, " "},
		{Word, "the"},
		{Space, " "},
		{Word, "worst"},
		{Space, " "},
		{Word, "of"},
		{Space, " "},
		{Word, "times"},
		{Space, " "},
		{Marker, "(up)"},
	}
	got := Tokenize(input)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Tokenize() = %#v, want %#v", got, want)
	}
}

func TestTokenize_PunctuationGroup(t *testing.T) {
	input := "Wait... what!?"
	want := []Token{
		{Word, "Wait"},
		{Punctuation, "..."},
		{Space, " "},
		{Word, "what"},
		{Punctuation, "!?"},
	}
	got := Tokenize(input)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Tokenize() = %#v, want %#v", got, want)
	}
}

func TestTokenize_MarkersWithNumbers(t *testing.T) {
	input := "Wow (up,2)"
	want := []Token{
		{Word, "Wow"}, {Space, " "}, {Marker, "(up,2)"},
	}
	got := Tokenize(input)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Tokenize() = %#v, want %#v", got, want)
	}
}
