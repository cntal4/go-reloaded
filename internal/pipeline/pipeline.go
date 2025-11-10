package pipeline

import (
	"go-reloaded/pkg/processors"
	"go-reloaded/pkg/tokenizer"
)

type Pipeline struct {
	stages []Processor
}

type Processor interface {
	Process(tokens []tokenizer.Token) []tokenizer.Token
}

func New() *Pipeline {
	return &Pipeline{
		stages: []Processor{
			processors.HexBinProcessor{},
			processors.CaseProcessor{},
			processors.ArticleProcessor{},
			processors.QuoteProcessor{},
			processors.PunctuationProcessor{},
		},
	}
}

func (p *Pipeline) Process(tokens []tokenizer.Token) []tokenizer.Token {
	result := tokens
	for _, stage := range p.stages {
		result = stage.Process(result)
	}
	return result
}
