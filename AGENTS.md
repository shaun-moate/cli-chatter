# AGENTS.md - Coding Guidelines for cli-chatter

## Build/Test Commands
- **Build**: `go build`
- **Run**: `go run main.go`
- **Test all**: `go test ./...`
- **Test single package**: `go test -v`
- **Format code**: `go fmt ./...`
- **Check imports**: `go mod tidy`

## Code Style Guidelines

### Formatting & Imports
- Use `go fmt` for consistent formatting
- Group imports: standard library, then third-party, then local
- Use blank lines between import groups

### Naming Conventions
- **Functions/Methods**: PascalCase for exported, camelCase for unexported
- **Variables**: camelCase, descriptive names
- **Constants**: PascalCase for exported, camelCase for unexported
- **Types**: PascalCase for exported structs/interfaces

### Error Handling
- Use `if err != nil` pattern immediately after operations
- Return errors early, don't wrap unnecessarily
- Use descriptive error messages

### Code Structure
- Keep functions focused and under 50 lines
- Use meaningful variable names over comments
- Follow Go's idiomatic patterns (e.g., bufio.Scanner for input)

### Best Practices
- Handle all errors appropriately
- Use defer for cleanup when needed
- Prefer explicit error returns over panics
- Write testable code with clear interfaces