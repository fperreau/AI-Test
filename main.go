package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <dockerfile-path>")
		os.Exit(1)
	}

	dockerfilePath := os.Args[1]
	content, err := os.ReadFile(dockerfilePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Parse the Dockerfile
	dockerfile, err := ParseDockerfile(string(content))
	if err != nil {
		fmt.Printf("Error parsing Dockerfile: %v\n", err)
		os.Exit(1)
	}

	// Convert the Dockerfile
	yamlIntermediate, err := TranslateDockerfile(dockerfile)
	if err != nil {
		fmt.Printf("Error translating Dockerfile: %v\n", err)
		os.Exit(1)
	}

	converted := strings.Join(yamlIntermediate.Scripts, "\n")

	// Print the parsed instructions
	fmt.Println("=== Parsed Dockerfile ===")
	for _, instruction := range dockerfile.Instructions {
		fmt.Printf("Line %d: %s - %s\n", instruction.Line, instruction.Type, instruction.Content)
	}

	// Print the converted content
	fmt.Println("\n=== Converted Dockerfile ===")
	fmt.Println(converted)
}
