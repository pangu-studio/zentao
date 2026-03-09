# AGENTS.md - Guide for AI Coding Agents

This document provides guidelines for AI coding agents (including Claude, GPT, Copilot, etc.) working in this repository.

## Language Requirement

**IMPORTANT: All interactions with users MUST be conducted in Chinese (Simplified Chinese).**

AI agents working in this repository must:
- Respond to users in Chinese
- Provide explanations, comments, and documentation in Chinese
- Use Chinese for error messages and user-facing output
- Use Chinese when asking questions or requesting clarifications

Code, comments within code files, commit messages, and technical documentation should follow standard English conventions as per Go community standards.

## Repository Overview

**Project**: CLI Skill Development Template  
**Language**: Go 1.25.5  
**Module**: `github.com/pangu-studio/zentao`  
**Framework**: [Cobra](https://github.com/spf13/cobra) for CLI, [Testify](https://github.com/stretchr/testify) for testing

This is a template for building command-line skills with configuration management, output formatters, and comprehensive testing.

## Build, Test, and Lint Commands

### Using Makefile (Recommended)

```bash
# Build all CLIs
make build                 # Outputs to bin/ directory

# Build for Linux AMD64
make build-linux

# Install to $GOPATH/bin
make install

# Testing
make test                  # Run all tests
make test-cover            # Run tests with coverage report (generates coverage.html)

# Code quality
make fmt                   # Format code with gofmt
make vet                   # Run go vet
make lint                  # Run fmt, vet, and golangci-lint

# Maintenance
make tidy                  # Tidy dependencies
make clean                 # Remove bin/, coverage files
make help                  # Show all available targets
```

### Direct Go Commands

```bash
# Build
go build -v -o bin/zentao ./cmd/zentao

# Testing
go test ./...                              # Run all tests
go test -v ./...                           # Verbose output
go test -run TestFunctionName ./internal/config  # Run specific test
go test ./internal/config/...              # Test specific package
go test -cover ./...                       # With coverage
go test -race ./...                        # With race detection

# Code quality
go fmt ./...
go vet ./...
golangci-lint run
go mod tidy
```

## Project Structure

```
cmd/zentao/          # CLI entry point and commands
internal/config/      # Configuration management (env vars, files)
internal/output/      # Output formatters (text, json, table)
skills/zentao/       # Skill definitions and docs
```

## Code Style Guidelines

### Import Formatting

**Always use three groups separated by blank lines:**

```go
import (
    // 1. Standard library
    "context"
    "fmt"
    "os"

    // 2. External dependencies
    "github.com/spf13/cobra"
    "github.com/stretchr/testify/assert"

    // 3. Internal packages
    "github.com/pangu-studio/zentao/internal/config"
)
```

### Naming Conventions

- **Functions/Variables**: `camelCase` for private, `PascalCase` for exported
- **Files**: `snake_case.go` (e.g., `config_test.go`, `formatter.go`)
- **Test files**: Must end with `_test.go`
- **Interfaces**: Use `-er` suffix for single-method (e.g., `Formatter`, `Writer`)
- **Constants**: `PascalCase` (e.g., `FormatJSON`, `SkillName`)

### Error Handling

```go
// Good: Wrap errors with context using %w
if err := saveData(data); err != nil {
    return fmt.Errorf("save data for user %s: %w", userID, err)
}

// Good: Check errors immediately
result, err := operation()
if err != nil {
    return fmt.Errorf("operation failed: %w", err)
}

// Bad: Never ignore errors
_ = operation() // ❌ Don't do this
```

### Function Design

- **Keep functions small**: Under 50 lines ideally
- **Context first**: `func Do(ctx context.Context, arg string) error`
- **Error last**: Return error as the last return value
- **Document exported functions**: Use godoc comments

```go
// Load loads configuration from environment variables and config files.
// Environment variables use uppercase skill name: {SKILLNAME}_API_KEY, {SKILLNAME}_API_HOST
func Load() (*Config, error) {
    // Implementation
}
```

### Testing Patterns

**Use Arrange-Act-Assert pattern with testify:**

```go
func TestLoad_FromEnv(t *testing.T) {
    // Arrange
    expectedKey := "test-api-key"
    os.Setenv("MYSKILL_API_KEY", expectedKey)
    defer os.Unsetenv("MYSKILL_API_KEY")

    // Act
    config, err := Load()

    // Assert
    require.NoError(t, err)
    assert.Equal(t, expectedKey, config.API.APIKey)
}
```

**Use table-driven tests for multiple cases:**

```go
func TestFormatValidation(t *testing.T) {
    tests := []struct {
        name    string
        format  string
        wantErr bool
    }{
        {"valid text", "text", false},
        {"valid json", "json", false},
        {"invalid", "invalid", true},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            _, err := NewFormatter(tt.format)
            if tt.wantErr {
                require.Error(t, err)
            } else {
                require.NoError(t, err)
            }
        })
    }
}
```

## Project-Specific Guidelines

### Configuration Loading Priority

1. Environment variables (highest priority)
2. Config files in `~/.config/awesome-skill/{skillname}/`
3. Default values

Environment variables: `{SKILLNAME}_API_KEY`, `{SKILLNAME}_API_HOST` (uppercase)

### Adding New Commands

Create new file in `cmd/zentao/cmd/` and register in `init()`:

```go
var myCmd = &cobra.Command{
    Use:   "mycommand [args]",
    Short: "Brief description",
    Args:  cobra.ExactArgs(1),
    RunE:  runMyCommand,
}

func init() {
    rootCmd.AddCommand(myCmd)
}
```

## Pre-Commit Checklist

```bash
make fmt && make vet && make test
```

Or use the comprehensive check:

```bash
make lint && make test-cover
```

## Common Patterns

**Using formatters:**
```go
formatter, err := output.NewFormatter(formatFlag)
if err != nil {
    return err
}
return formatter.Format(data, os.Stdout)
```

**Loading config:**
```go
cfg, err := config.Load()
if err != nil {
    return err
}
```

---

**Last Updated**: March 9, 2026
