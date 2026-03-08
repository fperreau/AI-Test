# Dockerfile to Distrobuilder Converter

This tool converts Dockerfiles to YAML configurations compatible with [distrobuilder](https://github.com/lxc/distrobuilder), used for building Incus/LXD images.

## Features

- Convert Dockerfile instructions to distrobuilder YAML
- Handle common Dockerfile instructions (FROM, RUN, COPY, EXPOSE, CMD)
- Generate post-build scripts for Incus configuration
- Optimize the generated YAML

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/fperreau/AI-Test.git
   cd AI-Test
   ```

2. Build the application:
   ```bash
   go build -o AI-Test
   ```

## Usage

```bash
./AI-Test nginx.dkb > nginx.yaml
```

## Example

### Input Dockerfile (nginx.dkb)

```dockerfile
FROM ubuntu:22.04

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
CMD ["nginx", "-g", "daemon off;"]
```

### Output YAML (nginx.yaml)

```yaml
image:
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
        incus exec ${CONTAINER_NAME} -- nginx -g daemon off;
```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

This project is licensed under the MIT License.