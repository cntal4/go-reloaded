package pipeline

import "go-reloaded/pkg/tokenizer"

type Processor interface {
	Process(tokens []tokenizer.Token) []tokenizer.Token
}

type Pipeline struct {
	steps []Processor
}

func New() *Pipeline {
	return &Pipeline{}
}

func (p *Pipeline) Add(step Processor) {
	p.steps = append(p.steps, step)
}

func (p *Pipeline) Run(tokens []tokenizer.Token) []tokenizer.Token {
	out := tokens
	for _, step := range p.steps {
		out = step.Process(out)
	}
	return out
}
