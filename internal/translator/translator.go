package translator

import (
	"fmt"
	"strings"

	"github.com/yourusername/dockerfile-to-distrobuilder/internal/parser"
)

type YAMLIntermediate struct {
	Scripts     []string
	Files       map[string]string
	PostScripts []string
}

func TranslateDockerfile(dockerfile *parser.Dockerfile) (*YAMLIntermediate, error) {
	yaml := &YAMLIntermediate{
		Files: make(map[string]string),
	}

	for _, instruction := range dockerfile.Instructions {
		switch instruction.Type {
		case "RUN":
			script, err := translateRun(instruction)
			if err != nil {
				return nil, err
			}
			yaml.Scripts = append(yaml.Scripts, script)
		case "COPY":
			err := translateCopy(instruction, yaml)
			if err != nil {
				return nil, err
			}
		case "EXPOSE":
			script := translateExpose(instruction)
			yaml.PostScripts = append(yaml.PostScripts, script)
		case "CMD":
			script := translateCMD(instruction)
			yaml.PostScripts = append(yaml.PostScripts, script)
		default:
			// Handle unsupported instructions
			yaml.Scripts = append(yaml.Scripts, fmt.Sprintf("# Unsupported instruction: %s", instruction.Type))
		}
	}

	return yaml, nil
}

func translateRun(instruction parser.Instruction) (string, error) {
	// Extract the command from the RUN instruction
	parts := strings.SplitN(instruction.Content, " ", 2)
	if len(parts) < 2 {
		return "", fmt.Errorf("invalid RUN instruction: %s", instruction.Content)
	}
	cmd := parts[1]

	// Return the command as a script
	return cmd, nil
}

func translateCopy(instruction parser.Instruction, yaml *YAMLIntermediate) error {
	// Parse the COPY instruction
	parts := strings.Fields(instruction.Content)
	if len(parts) < 3 {
		return fmt.Errorf("invalid COPY instruction: %s", instruction.Content)
	}

	src := parts[1]
	dest := parts[2]

	// For simplicity, we'll assume the source is a file in the current directory
	// In a real implementation, you would need to read the file content
	content := fmt.Sprintf("# Content of %s\n", src)

	// Add the file to the YAML intermediate structure
	yaml.Files[dest] = content

	return nil
}

func translateExpose(instruction parser.Instruction) string {
	// Parse the EXPOSE instruction
	parts := strings.Fields(instruction.Content)
	if len(parts) < 2 {
		return fmt.Sprintf("# Invalid EXPOSE instruction: %s", instruction.Content)
	}

	port := parts[1]

	// Generate the post-script to configure the port
	return fmt.Sprintf(`#!/bin/sh
incus config device add ${CONTAINER_NAME} nginx-port proxy listen=tcp:0.0.0.0:%s connect=tcp:127.0.0.1:%s`, port, port)
}

func translateCMD(instruction parser.Instruction) string {
	// Parse the CMD instruction
	parts := strings.Fields(instruction.Content)
	if len(parts) < 2 {
		return fmt.Sprintf("# Invalid CMD instruction: %s", instruction.Content)
	}

	// Extract the command and arguments
	cmd := parts[1]
	args := strings.Join(parts[2:], " ")

	// Generate the post-script to start the service
	return fmt.Sprintf(`#!/bin/sh
incus exec ${CONTAINER_NAME} -- %s %s`, cmd, args)
}