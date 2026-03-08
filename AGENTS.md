# AGENTS.md

## Build/Lint/Test Commands

### Build

```bash
make build
```

### Lint

```bash
go vet ./...
go fmt ./...
```

### Test

```bash
go test ./...
```

### Run a Single Test

```bash
go test -run TestFunctionName ./...
```

## Code Style Guidelines

### Imports

- Group imports into three sections: standard library, third-party packages, and local packages.
- Sort imports alphabetically within each group.
- Separate groups with a blank line.

### Formatting

- Use `go fmt` to format code.
- Use `go vet` to check for common mistakes.

### Types

- Use named types for clarity and consistency.
- Use interfaces to define behavior and dependencies.

### Naming Conventions

- Use PascalCase for type names.
- Use camelCase for variable and function names.
- Use UPPER_CASE for constants.

### Error Handling

- Use error wrapping to provide context.
- Handle errors at the appropriate level.
- Avoid ignoring errors.

### Comments

- Write clear and concise comments.
- Use comments to explain why, not what.
- Avoid redundant comments.

### Testing

- Write tests for all non-trivial functions.
- Use table-driven tests for complex logic.
- Mock dependencies in tests.

### Documentation

- Write clear and concise documentation.
- Use examples to illustrate usage.
- Keep documentation up-to-date.

### Dependencies

- Use `go mod` to manage dependencies.
- Avoid using deprecated packages.
- Keep dependencies up-to-date.

### Code Reviews

- Review code changes thoroughly.
- Provide constructive feedback.
- Ensure code quality and consistency.

### Continuous Integration

- Run tests and lint checks on every commit.
- Ensure all tests pass before merging.
- Monitor build status and fix issues promptly.

### Version Control

- Use meaningful commit messages.
- Keep commits focused and atomic.
- Rebase branches before merging.

### Security

- Avoid hardcoding secrets.
- Use secure coding practices.
- Keep dependencies up-to-date.

### Performance

- Optimize code for performance.
- Avoid unnecessary allocations.
- Use efficient algorithms and data structures.

### Logging

- Use structured logging.
- Include relevant context in logs.
- Avoid logging sensitive information.

### Error Messages

- Provide clear and actionable error messages.
- Include relevant context in error messages.
- Avoid generic error messages.

### Code Organization

- Organize code into logical packages.
- Keep functions and methods focused and cohesive.
- Avoid deep nesting and complex control flow.

### Code Reviews

- Review code changes thoroughly.
- Provide constructive feedback.
- Ensure code quality and consistency.

### Documentation

- Write clear and concise documentation.
- Use examples to illustrate usage.
- Keep documentation up-to-date.

### Dependencies

- Use `go mod` to manage dependencies.
- Avoid using deprecated packages.
- Keep dependencies up-to-date.

### Continuous Integration

- Run tests and lint checks on every commit.
- Ensure all tests pass before merging.
- Monitor build status and fix issues promptly.

### Version Control

- Use meaningful commit messages.
- Keep commits focused and atomic.
- Rebase branches before merging.

### Security

- Avoid hardcoding secrets.
- Use secure coding practices.
- Keep dependencies up-to-date.

### Performance

- Optimize code for performance.
- Avoid unnecessary allocations.
- Use efficient algorithms and data structures.

### Logging

- Use structured logging.
- Include relevant context in logs.
- Avoid logging sensitive information.

### Error Messages

- Provide clear and actionable error messages.
- Include relevant context in error messages.
- Avoid generic error messages.

### Code Organization

- Organize code into logical packages.
- Keep functions and methods focused and cohesive.
- Avoid deep nesting and complex control flow.

### Code Reviews

- Review code changes thoroughly.
- Provide constructive feedback.
- Ensure code quality and consistency.

### Documentation

- Write clear and concise documentation.
- Use examples to illustrate usage.
- Keep documentation up-to-date.

### Dependencies

- Use `go mod` to manage dependencies.
- Avoid using deprecated packages.
- Keep dependencies up-to-date.

### Continuous Integration

- Run tests and lint checks on every commit.
- Ensure all tests pass before merging.
- Monitor build status and fix issues promptly.

### Version Control

- Use meaningful commit messages.
- Keep commits focused and atomic.
- Rebase branches before merging.

### Security

- Avoid hardcoding secrets.
- Use secure coding practices.
- Keep dependencies up-to-date.

### Performance

- Optimize code for performance.
- Avoid unnecessary allocations.
- Use efficient algorithms and data structures.

### Logging

- Use structured logging.
- Include relevant context in logs.
- Avoid logging sensitive information.

### Error Messages

- Provide clear and actionable error messages.
- Include relevant context in error messages.
- Avoid generic error messages.

### Code Organization

- Organize code into logical packages.
- Keep functions and methods focused and cohesive.
- Avoid deep nesting and complex control flow.

