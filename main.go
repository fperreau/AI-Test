package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/yourusername/dockerfile-to-distrobuilder/internal/optimizer"
	"github.com/yourusername/dockerfile-to-distrobuilder/internal/parser"
	"github.com/yourusername/dockerfile-to-distrobuilder/internal/translator"
)

func main() {
	// Check if a Dockerfile path is provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: dockerfile-to-distrobuilder <dockerfile>")
		os.Exit(1)
	}

	dockerfilePath := os.Args[1]

	// Read the Dockerfile content
	content, err := ioutil.ReadFile(dockerfilePath)
	if err != nil {
		fmt.Printf("Error reading Dockerfile: %v\n", err)
		os.Exit(1)
	}

	// Parse the Dockerfile
	dockerfile, err := parser.ParseDockerfile(string(content))
	if err != nil {
		fmt.Printf("Error parsing Dockerfile: %v\n", err)
		os.Exit(1)
	}

	// Translate the Dockerfile to intermediate YAML
	yamlIntermediate, err := translator.TranslateDockerfile(dockerfile)
	if err != nil {
		fmt.Printf("Error translating Dockerfile: %v\n", err)
		os.Exit(1)
	}

	// Optimize the YAML
	finalYAML, err := optimizer.OptimizeYAML(yamlIntermediate)
	if err != nil {
		fmt.Printf("Error optimizing YAML: %v\n", err)
		os.Exit(1)
	}

	// Write the final YAML to stdout
	fmt.Println(finalYAML)
}
