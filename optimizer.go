package main

import (
	"fmt"
	"strings"
)

func OptimizeYAML(yamlIntermediate *YAMLIntermediate) (string, error) {
	// Fusionner les scripts RUN consécutifs
	var optimizedScripts []string
	var currentScript strings.Builder
	for _, script := range yamlIntermediate.Scripts {
		if strings.HasPrefix(script, "apt-get update") {
			if currentScript.Len() > 0 {
				optimizedScripts = append(optimizedScripts, currentScript.String())
				currentScript.Reset()
			}
			optimizedScripts = append(optimizedScripts, script)
		} else {
			if currentScript.Len() > 0 {
				currentScript.WriteString(" && ")
			}
			currentScript.WriteString(script)
		}
	}
	if currentScript.Len() > 0 {
		optimizedScripts = append(optimizedScripts, currentScript.String())
	}

	// Générer le YAML final
	var yamlBuilder strings.Builder
	yamlBuilder.WriteString("image:\n")
	yamlBuilder.WriteString(fmt.Sprintf("  distribution: %s\n", yamlIntermediate.BaseImage))
	yamlBuilder.WriteString("  release: jammy\n")
	yamlBuilder.WriteString("  arch: amd64\n")
	yamlBuilder.WriteString("source:\n")
	yamlBuilder.WriteString("  type: download\n")
	yamlBuilder.WriteString("  url: http://cloud-images.ubuntu.com/jammy/current/jammy-server-cloudimg-amd64-root.tar.xz\n")
	yamlBuilder.WriteString("targets:\n")
	yamlBuilder.WriteString("  lxd:\n")
	yamlBuilder.WriteString("    config:\n")
	yamlBuilder.WriteString("      - type: all\n")
	yamlBuilder.WriteString("        before: 5\n")
	yamlBuilder.WriteString("        content: |\n")
	yamlBuilder.WriteString("          #!/bin/sh\n")
	for _, script := range optimizedScripts {
		yamlBuilder.WriteString(fmt.Sprintf("          %s\n", script))
	}

	return yamlBuilder.String(), nil
}
