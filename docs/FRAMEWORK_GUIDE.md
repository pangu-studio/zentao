# Framework Guide

This document explains the framework components provided by this template and how to use them effectively.

## Table of Contents

1. [Configuration System](#configuration-system)
2. [Output Formatting](#output-formatting)
3. [CLI Command Structure](#cli-command-structure)
4. [Testing](#testing)
5. [Build System](#build-system)

---

## Configuration System

Location: `internal/config/`

The configuration system provides a flexible way to manage application settings with multiple sources.

### Configuration Priority

1. **Environment Variables** (highest priority)
2. **Configuration Files**
3. **Default Values** (lowest priority)

### Basic Usage

```go
import "github.com/pangu-studio/zentao/internal/config"

// Load configuration
cfg, err := config.Load()
if err != nil {
    // Handle error - typically means API key not configured
}

// Access configuration
apiKey := cfg.API.APIKey
apiHost := cfg.API.APIHost
```

### Configuration Structure

```go
type Config struct {
    API APIConfig
}

type APIConfig struct {
    APIKey  string  // Required: API authentication key
    APIHost string  // Optional: API endpoint host
}
```

### Environment Variables

Environment variables are built from the skill name in uppercase:

- `{SKILLNAME}_API_KEY` - API authentication key
- `{SKILLNAME}_API_HOST` - API endpoint host

For example, if `SkillName = "zentao"`:

- `MYSKILL_API_KEY`
- `MYSKILL_API_HOST`

### Configuration Files

Configuration files are stored in:

- **Linux/macOS**: `~/.config/awesome-skill/{skillname}/`
- **XDG-compliant**: `$XDG_CONFIG_HOME/awesome-skill/{skillname}/`

Files:

- `api_key` - Contains the API key (one line)
- `api_host` - Contains the API host (one line)

### Setting Configuration

#### Via Environment Variables

```bash
export MYSKILL_API_KEY="your-api-key-here"
export MYSKILL_API_HOST="api.example.com"
```

#### Via Configuration Files

```bash
# Using the config command
zentao config init --interactive
zentao config set-api-key your-api-key
zentao config set-api-host api.example.com

# Manual setup
mkdir -p ~/.config/awesome-skill/zentao
echo "your-api-key" > ~/.config/awesome-skill/zentao/api_key
echo "api.example.com" > ~/.config/awesome-skill/zentao/api_host
chmod 600 ~/.config/awesome-skill/zentao/api_key
```

### Customizing Configuration

#### Adding New Fields

Edit `internal/config/config.go`:

```go
type Config struct {
    API    APIConfig
    Custom CustomConfig  // Add your custom section
}

type CustomConfig struct {
    Timeout  int
    Region   string
    Debug    bool
}
```

#### Loading Custom Fields

Update the `LoadForSkill()` function:

```go
func LoadForSkill(skillName string) (*Config, error) {
    cfg := &Config{
        API: APIConfig{
            APIHost: getEnvWithDefault(envPrefix+"_API_HOST", "api.example.com"),
        },
        Custom: CustomConfig{
            Timeout: getEnvIntWithDefault(envPrefix+"_TIMEOUT", 30),
            Region:  getEnvWithDefault(envPrefix+"_REGION", "us-east-1"),
            Debug:   getEnvBoolWithDefault(envPrefix+"_DEBUG", false),
        },
    }
    // ... rest of loading logic
}
```

#### Helper Functions

Implement additional helper functions for different types:

```go
func getEnvIntWithDefault(key string, defaultValue int) int {
    if value := os.Getenv(key); value != "" {
        if intVal, err := strconv.Atoi(value); err == nil {
            return intVal
        }
    }
    return defaultValue
}

func getEnvBoolWithDefault(key string, defaultValue bool) bool {
    if value := os.Getenv(key); value != "" {
        if boolVal, err := strconv.ParseBool(value); err == nil {
            return boolVal
        }
    }
    return defaultValue
}
```

### Functions Reference

- `Load() (*Config, error)` - Load configuration for the default skill
- `LoadForSkill(skillName string) (*Config, error)` - Load configuration for a specific skill
- `EnsureConfigDir() error` - Create configuration directory if it doesn't exist
- `GetConfigDir() (string, error)` - Get the configuration directory path
- `SetAPIKey(apiKey string) error` - Save API key to configuration file
- `SetAPIHost(apiHost string) error` - Save API host to configuration file

---

## Output Formatting

Location: `internal/output/`

The output formatting system provides multiple output formats for displaying data.

### Supported Formats

1. **Text** - Human-readable text output
2. **JSON** - Machine-readable JSON output
3. **Table** - ASCII table format

### Basic Usage

```go
import (
    "os"
    "github.com/pangu-studio/zentao/internal/output"
)

// Create formatter based on format string
formatter, err := output.NewFormatter("json")
if err != nil {
    return err
}

// Format data
data := map[string]interface{}{
    "status": "success",
    "count": 42,
}

err = formatter.Format(data, os.Stdout)
```

### JSON Formatter

The JSON formatter works with any data type that can be marshaled to JSON:

```go
formatter := &output.JSONFormatter{}

// Works with structs
type Result struct {
    Status string `json:"status"`
    Data   int    `json:"data"`
}
result := Result{Status: "ok", Data: 123}
formatter.Format(result, os.Stdout)

// Works with maps
data := map[string]interface{}{"key": "value"}
formatter.Format(data, os.Stdout)
```

### Text Formatter

The text formatter provides simple text output. Customize it for your data types:

```go
// Edit internal/output/text.go
func (f *TextFormatter) FormatMyData(data *MyData, w io.Writer) error {
    fmt.Fprintf(w, "Title: %s\n", data.Title)
    fmt.Fprintf(w, "Value: %d\n", data.Value)
    return nil
}
```

### Table Formatter

The table formatter provides ASCII table output:

```go
// Using the generic Format method
data := map[string]interface{}{
    "Name": "John",
    "Age": 30,
    "City": "NYC",
}
formatter.Format(data, os.Stdout)

// Using helper functions for custom tables
headers := []string{"Name", "Age", "City"}
rows := [][]string{
    {"John", "30", "NYC"},
    {"Jane", "25", "LA"},
}
widths := []int{10, 5, 10}
output.PrintTableWithHeaders(os.Stdout, headers, rows, widths)
```

### Integrating with Commands

Use the global `formatFlag` from root command:

```go
func runMyCommand(cmd *cobra.Command, args []string) error {
    // Get format from parent/root flags
    format := formatFlag
    
    formatter, err := output.NewFormatter(format)
    if err != nil {
        return err
    }
    
    // Your data
    data := fetchData()
    
    // Output
    return formatter.Format(data, os.Stdout)
}
```

---

## CLI Command Structure

Location: `cmd/zentao/cmd/`

The CLI is built using [Cobra](https://github.com/spf13/cobra), a powerful library for building CLI applications.

### Command Structure

```
zentao (root command)
├── config (configuration management)
│   ├── init (initialize configuration)
│   ├── set-api-key (set API key)
│   └── set-api-host (set API host)
└── example (example command)
```

### Root Command

File: `cmd/zentao/cmd/root.go`

The root command provides:

- Global flags (format, verbose)
- Version information
- Entry point for all subcommands

```go
var rootCmd = &cobra.Command{
    Use:   "zentao",
    Short: "Short description",
    Long:  `Long description`,
    Version: "0.1.0",
}
```

### Global Flags

Available to all commands:

- `--format, -f` - Output format (text, json, table)
- `--verbose, -v` - Verbose output

### Creating a New Command

1. Create a new file: `cmd/zentao/cmd/mycommand.go`

```go
package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
)

var myCmd = &cobra.Command{
    Use:   "mycommand [args]",
    Short: "Short description",
    Long:  `Long description with examples`,
    Args:  cobra.ExactArgs(1),  // Validation
    RunE:  runMyCommand,         // Run function
}

func init() {
    // Register command
    rootCmd.AddCommand(myCmd)
    
    // Add command-specific flags
    myCmd.Flags().StringP("option", "o", "default", "Option description")
}

func runMyCommand(cmd *cobra.Command, args []string) error {
    // Get flag value
    option, _ := cmd.Flags().GetString("option")
    
    // Implement command logic
    fmt.Printf("Running mycommand with %s\n", args[0])
    
    return nil
}
```

### Command Best Practices

1. **Use RunE instead of Run** - Return errors instead of handling them
2. **Validate early** - Use `Args` validators or validate in `RunE` start
3. **Use flags for options** - Named flags are clearer than positional args
4. **Provide examples** - Include examples in the `Long` description
5. **Handle context** - Use `cmd.Context()` for cancellation support

### Argument Validation

Cobra provides built-in validators:

```go
cobra.NoArgs           // No arguments allowed
cobra.ExactArgs(n)     // Exactly n arguments
cobra.MinimumNArgs(n)  // At least n arguments
cobra.MaximumNArgs(n)  // At most n arguments
cobra.RangeArgs(min, max)  // Between min and max arguments
```

Custom validation:

```go
Args: func(cmd *cobra.Command, args []string) error {
    if len(args) < 1 {
        return fmt.Errorf("requires at least one argument")
    }
    // Custom validation logic
    return nil
},
```

---

## Testing

The template uses [testify](https://github.com/stretchr/testify) for testing.

### Running Tests

```bash
# Run all tests
go test ./...

# Run with verbose output
go test -v ./...

# Run tests with coverage
go test -cover ./...

# Generate coverage report
make test-cover
```

### Test Structure

Follow table-driven test patterns:

```go
func TestMyFunction(t *testing.T) {
    testCases := []struct {
        name     string
        input    string
        expected string
        wantErr  bool
    }{
        {"valid input", "test", "TEST", false},
        {"empty input", "", "", true},
    }
    
    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            result, err := MyFunction(tc.input)
            
            if tc.wantErr {
                assert.Error(t, err)
            } else {
                require.NoError(t, err)
                assert.Equal(t, tc.expected, result)
            }
        })
    }
}
```

### Testing Commands

Test command logic by calling the run functions directly:

```go
func TestRunMyCommand(t *testing.T) {
    // Create a test command
    cmd := &cobra.Command{}
    
    // Set up flags if needed
    cmd.Flags().String("option", "default", "")
    
    // Run the command
    err := runMyCommand(cmd, []string{"arg1"})
    
    assert.NoError(t, err)
}
```

### Testing with Temporary Config

```go
func TestWithConfig(t *testing.T) {
    // Create temporary directory
    tmpDir := t.TempDir()
    
    // Set HOME to temp directory
    originalHome := os.Getenv("HOME")
    os.Setenv("HOME", tmpDir)
    defer os.Setenv("HOME", originalHome)
    
    // Now config will use tmpDir
    config, err := config.Load()
    // ...
}
```

---

## Build System

The Makefile provides common development tasks.

### Available Commands

```bash
make build        # Build binaries
make build-linux  # Build for Linux AMD64
make install      # Install to $GOPATH/bin
make test         # Run tests
make test-cover   # Run tests with coverage
make fmt          # Format code
make vet          # Run go vet
make lint         # Run all linters
make tidy         # Tidy dependencies
make clean        # Clean build artifacts
make help         # Show help
```

### Build Configuration

Edit Makefile variables:

```makefile
BINARY_DIR := bin
BINARIES := zentao
BINARY_LINUX := zentao-linux-amd64
```

### Cross-Compilation

Build for different platforms:

```bash
# Linux AMD64
GOOS=linux GOARCH=amd64 go build -o bin/zentao-linux-amd64 ./cmd/zentao

# macOS ARM64
GOOS=darwin GOARCH=arm64 go build -o bin/zentao-darwin-arm64 ./cmd/zentao

# Windows AMD64
GOOS=windows GOARCH=amd64 go build -o bin/zentao-windows-amd64.exe ./cmd/zentao
```

### Adding Build Targets

Add custom targets to Makefile:

```makefile
.PHONY: docker
docker:
 @echo "Building Docker image..."
 docker build -t zentao:latest .
```

---

## Additional Resources

- **Go Documentation**: <https://golang.org/doc/>
- **Cobra Documentation**: <https://cobra.dev/>
- **Testify Documentation**: <https://github.com/stretchr/testify>
- **Go Code Review Comments**: <https://go.dev/wiki/CodeReviewComments>

---

For more information on creating a skill, see [CREATING_A_SKILL.md](CREATING_A_SKILL.md).
