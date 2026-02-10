# Stargazing

Go + Echo + Datastar bootstrap for real-time web apps with SSE.

## Features

- **Echo v4** - Fast HTTP framework
- **Datastar** - Hypermedia framework for reactive UIs via SSE
- **slog** - Structured logging with colored terminal output + JSON file logs
- **Air** - Hot reload for development
- **Graceful shutdown** - Clean SIGINT/SIGTERM handling

## Quick Start

```bash
# Clone and setup
cp .env.example .env

# Development with hot reload
make dev

# Or build and run
make run
```

Open http://localhost:8080

## Project Structure

```
cmd/server/         # Application entrypoint
internal/
  config/           # Environment configuration
  handlers/         # HTTP handlers (index, sse, health)
  logger/           # Colored slog handler + JSON file logging
  middleware/       # Request ID middleware
  store/            # In-memory data store with client management
  utils/            # Shared state (store instance)
web/
  static/js/        # JavaScript (Datastar vendored)
  templates/        # Go HTML templates
examples/
  todo.go           # Example todo handler
```

## Configuration

Environment variables (see `.env.example`):

| Variable | Default | Description |
|----------|---------|-------------|
| `PORT` | `8080` | Server port |
| `LOG_LEVEL` | `debug` | Log level (debug, info, warn, error) |
| `LOG_FILE` | `logs/app.log` | JSON log file path |

## Make Commands

```bash
make dev      # Hot reload development
make build    # Build binary to bin/
make run      # Build and run
make test     # Run tests
make fmt      # Format code
make vet      # Vet code
make check    # fmt + vet + test
make clean    # Remove build artifacts
```

## Docker

```bash
# Build and run
docker compose up --build

# Or just build
docker build -t stargazing .
docker run -p 8080:8080 stargazing
```

## How It Works

1. Browser loads page, connects to `/sse` endpoint
2. SSE connection stays open, receives HTML patches
3. User actions trigger POST requests (e.g., `/todo`)
4. POST handler updates state, signals SSE to push update
5. Datastar morphs the DOM with new HTML

## License

MIT
