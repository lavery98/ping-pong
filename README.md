# ping-pong

[![Go](https://img.shields.io/github/go-mod/go-version/lavery98/ping-pong)](https://go.dev/)
[![License](https://img.shields.io/github/license/lavery98/ping-pong)](LICENSE)
[![Release](https://img.shields.io/github/v/release/lavery98/ping-pong)](https://github.com/lavery98/ping-pong/releases)

A lightweight HTTP server that responds to `/ping` with `pong`. Perfect for health checks, endpoint monitoring, and service availability testing.

## Installation

### Using Go

```bash
go install github.com/lavery98/ping-pong@latest
```

### Using Docker

```bash
docker pull ghcr.io/lavery98/ping-pong:latest
```

### Download Pre-built Binary

Download the latest release for your platform from the [releases page](https://github.com/lavery98/ping-pong/releases).

## Usage

### Run Locally

Default port (4000):
```bash
ping-pong
```

Custom port:
```bash
ping-pong -port 8080
```

### Run with Docker

```bash
docker run -p 4000:4000 ghcr.io/lavery98/ping-pong:latest
```

### Test the Endpoint

```bash
curl http://localhost:4000/ping
# Response: pong
```

## Building from Source

### Prerequisites

- Go 1.23 or later

### Build Binary

```bash
make build
```

Or manually:
```bash
CGO_ENABLED=0 go build -a --trimpath --installsuffix cgo --ldflags '-s' -o ping-pong
```

## Configuration

| Flag | Default | Description |
|------|---------|-------------|
| `-port` | `4000` | Port number to listen on |

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
