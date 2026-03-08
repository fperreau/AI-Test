package main

import (
	"fmt"
	"strings"
)

type YAMLIntermediate struct {
	Scripts   []string
	BaseImage string
}

func TranslateDockerfile(dockerfile *Dockerfile) (*YAMLIntermediate, error) {
	var scripts []string
	var baseImage string
	for _, instruction := range dockerfile.Instructions {
		switch instruction.Type {
		case "FROM":
			baseImage = instruction.Content
		case "RUN":
			script, err := translateRun(instruction)
			if err != nil {
				return nil, err
			}
			scripts = append(scripts, script)
		// Ajouter d'autres cas pour les autres instructions Dockerfile
		default:
			// Gérer les instructions non supportées
			scripts = append(scripts, fmt.Sprintf("# Unsupported instruction: %s", instruction.Type))
		}
	}

	return &YAMLIntermediate{
		Scripts:   scripts,
		BaseImage: baseImage,
	}, nil
}

func translateRun(instruction Instruction) (string, error) {
	content := instruction.Content
	if !strings.Contains(content, "apt-get update") {
		content = "apt-get update && " + content
	}
	return content, nil
}
