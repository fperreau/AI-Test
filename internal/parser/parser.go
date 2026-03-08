package parser

import (
	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

type Dockerfile struct {
	Instructions []Instruction
}

type Instruction struct {
	Type    string
	Content string
	Line    int
}

func ParseDockerfile(content string) (*Dockerfile, error) {
	result, err := parser.Parse([]byte(content))
	if err != nil {
		return nil, err
	}

	var instructions []Instruction
	for _, child := range result.AST.Children {
		instructions = append(instructions, Instruction{
			Type:    child.Value,
			Content: child.Original,
			Line:    child.StartLine,
		})
	}

	return &Dockerfile{Instructions: instructions}, nil
}