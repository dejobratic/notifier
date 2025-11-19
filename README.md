# notifier

## Swagger Documentation

This project uses [swaggo/swag](https://github.com/swaggo/swag) for API documentation.

### Accessing Swagger UI

When you start the application, Swagger UI will automatically open in your browser at:
- http://localhost:8080/swagger/index.html

### Regenerating Documentation

After adding or modifying API endpoints with Swagger annotations, regenerate the docs:

```bash
swag init
```

Or if `swag` is not in your PATH:

```bash
go run github.com/swaggo/swag/cmd/swag@latest init
```

### Adding Swagger Annotations

Add annotations to your handlers following the [swaggo documentation](https://github.com/swaggo/swag#declarative-comments-format).