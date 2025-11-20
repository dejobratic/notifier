# notifier

## Quick Start

### Development

For development, use the convenience command that generates docs and starts the server:

```bash
make dev
```

### Available Commands

- `make dev` - Generate Swagger docs and run the application (development mode)
- `make build` - Build the application binary
- `make docs` - Generate Swagger documentation
- `make run` - Build and run the application
- `make test` - Run tests
- `make clean` - Clean build artifacts and generated docs

## Swagger Documentation

This project uses [swaggo/swag](https://github.com/swaggo/swag) for API documentation.

### Accessing Swagger UI

When you start the application, Swagger UI will automatically open in your browser at:
- http://localhost:8080/swagger/index.html

### Regenerating Documentation

After adding or modifying API endpoints with Swagger annotations, regenerate the docs:

```bash
make docs
```

Or manually:

```bash
go run github.com/swaggo/swag/cmd/swag@latest init
```

### Adding Swagger Annotations

Add annotations to your handlers following the [swaggo documentation](https://github.com/swaggo/swag#declarative-comments-format).