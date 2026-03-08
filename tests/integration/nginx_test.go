package integration

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNginxConversion(t *testing.T) {
	// Create a temporary directory for the test
	tempDir, err := ioutil.TempDir("", "nginx-test")
	assert.NoError(t, err)
	defer os.RemoveAll(tempDir)

	// Create a sample Dockerfile
	dockerfileContent := `FROM ubuntu:22.04

# Install dependencies
RUN apt-get update && apt-get install -y \
    nginx \
    && rm -rf /var/lib/apt/lists/*

# Copy nginx configuration
COPY nginx.conf /etc/nginx/nginx.conf

# Copy website content
COPY website/ /var/www/html/

# Expose port 80
EXPOSE 80

# Start nginx
CMD ["nginx", "-g", "daemon off;"]`

	dockerfilePath := filepath.Join(tempDir, "nginx.dkb")
	err = ioutil.WriteFile(dockerfilePath, []byte(dockerfileContent), 0644)
	assert.NoError(t, err)

	// Create a sample nginx.conf
	nginxConfContent := `user www-data;
worker_processes auto;
pid /run/nginx.pid;
events {
    worker_connections 1024;
}
http {
    include /etc/nginx/mime.types;
    default_type application/octet-stream;
    sendfile on;
    keepalive_timeout 65;
    server {
        listen 80;
        server_name localhost;
        location / {
            root /var/www/html;
            index index.html index.htm;
        }
    }
}`

	nginxConfPath := filepath.Join(tempDir, "nginx.conf")
	err = ioutil.WriteFile(nginxConfPath, []byte(nginxConfContent), 0644)
	assert.NoError(t, err)

	// Create a sample website directory
	websiteDir := filepath.Join(tempDir, "website")
	err = os.Mkdir(websiteDir, 0755)
	assert.NoError(t, err)

	// Create a sample index.html
	indexHTMLContent := `<!DOCTYPE html>
<html>
<head>
    <title>Welcome to Nginx!</title>
</head>
<body>
    <h1>Success! The Nginx server is running.</h1>
</body>
</html>`

	indexHTMLPath := filepath.Join(websiteDir, "index.html")
	err = ioutil.WriteFile(indexHTMLPath, []byte(indexHTMLContent), 0644)
	assert.NoError(t, err)

	// Run the conversion tool
	cmd := exec.Command("go", "run", "main.go", dockerfilePath)
	cmd.Dir = tempDir
	output, err := cmd.CombinedOutput()
	assert.NoError(t, err, string(output))

	// Verify the output
	expectedYAML := `image:
  distribution: ubuntu
  release: jammy
  arch: amd64
source:
  type: download
  url: http://cloud-images.ubuntu.com/jammy/current/jammy-server-cloudimg-amd64-root.tar.xz
targets:
  lxd:
    config:
      - type: all
        before: 5
        content: |
          #!/bin/sh
          apt-get update && apt-get install -y nginx && rm -rf /var/lib/apt/lists/*
    post_files:
      - path: /etc/nginx/nginx.conf
        content: |
          user www-data;
          worker_processes auto;
          pid /run/nginx.pid;
          events {
              worker_connections 1024;
          }
          http {
              include /etc/nginx/mime.types;
              default_type application/octet-stream;
              sendfile on;
              keepalive_timeout 65;
              server {
                  listen 80;
                  server_name localhost;
                  location / {
                      root /var/www/html;
                      index index.html index.htm;
                  }
              }
          }
      - path: /var/www/html/index.html
        content: |
          <!DOCTYPE html>
          <html>
          <head>
              <title>Welcome to Nginx!</title>
          </head>
          <body>
              <h1>Success! The Nginx server is running.</h1>
          </body>
          </html>
    post_scripts:
      - |
        #!/bin/sh
        incus config device add ${CONTAINER_NAME} nginx-port proxy listen=tcp:0.0.0.0:80 connect=tcp:127.0.0.1:80
      - |
        #!/bin/sh
        incus exec ${CONTAINER_NAME} -- nginx -g daemon off;`

	assert.Equal(t, expectedYAML, string(output))
}