package optimizer

import (
	"fmt"
	"strings"

	"github.com/yourusername/dockerfile-to-distrobuilder/internal/translator"
)

func OptimizeYAML(yamlIntermediate *translator.YAMLIntermediate) (string, error) {
	// Merge consecutive RUN commands
	mergedScripts := mergeScripts(yamlIntermediate.Scripts)

	// Generate the YAML content
	var yamlBuilder strings.Builder

	// Write the image section
	yamlBuilder.WriteString("image:\n")
	yamlBuilder.WriteString("  distribution: ubuntu\n")
	yamlBuilder.WriteString("  release: jammy\n")
	yamlBuilder.WriteString("  arch: amd64\n")

	// Write the source section
	yamlBuilder.WriteString("source:\n")
	yamlBuilder.WriteString("  type: download\n")
	yamlBuilder.WriteString("  url: http://cloud-images.ubuntu.com/jammy/current/jammy-server-cloudimg-amd64-root.tar.xz\n")

	// Write the targets section
	yamlBuilder.WriteString("targets:\n")
	yamlBuilder.WriteString("  lxd:\n")

	// Write the config section with merged scripts
	yamlBuilder.WriteString("    config:\n")
	yamlBuilder.WriteString("      - type: all\n")
	yamlBuilder.WriteString("        before: 5\n")
	yamlBuilder.WriteString("        content: |\n")
	yamlBuilder.WriteString("          #!/bin/sh\n")
	for _, script := range mergedScripts {
		yamlBuilder.WriteString(fmt.Sprintf("          %s\n", script))
	}

	// Write the post_files section
	if len(yamlIntermediate.Files) > 0 {
		yamlBuilder.WriteString("    post_files:\n")
		for path, content := range yamlIntermediate.Files {
			yamlBuilder.WriteString(fmt.Sprintf("      - path: %s\n", path))
			yamlBuilder.WriteString("        content: |\n")
			lines := strings.Split(content, "\n")
			for _, line := range lines {
				yamlBuilder.WriteString(fmt.Sprintf("          %s\n", line))
			}
		}
	}

	// Write the post_scripts section
	if len(yamlIntermediate.PostScripts) > 0 {
		yamlBuilder.WriteString("    post_scripts:\n")
		for _, script := range yamlIntermediate.PostScripts {
			yamlBuilder.WriteString("      - |\n")
			lines := strings.Split(script, "\n")
			for _, line := range lines {
				yamlBuilder.WriteString(fmt.Sprintf("        %s\n", line))
			}
		}
	}

	return yamlBuilder.String(), nil
}

func mergeScripts(scripts []string) []string {
	var merged []string
	var currentScript strings.Builder

	for _, script := range scripts {
		if currentScript.Len() > 0 {
			currentScript.WriteString(" && ")
		}
		currentScript.WriteString(script)
	}

	if currentScript.Len() > 0 {
		merged = append(merged, currentScript.String())
	}

	return merged
}