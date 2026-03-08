package main

import (
	"fmt"
)

func GeneratePostBuildScript(containerName string) string {
	return fmt.Sprintf(`#!/bin/bash
incus config device add %s myport80 proxy listen=tcp:0.0.0.0:80 connect=tcp:127.0.0.1:80
incus exec %s -- systemctl enable nginx`, containerName, containerName)
}
