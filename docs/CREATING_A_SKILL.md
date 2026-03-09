# Creating a Skill from this Template

This guide will walk you through using this template to create your own CLI skill.

## Quick Start

### Option 1: Using the Initialization Script (Recommended)

```bash
# Clone or download this template
git clone <this-repo> my-new-skill
cd my-new-skill

# Run the initialization script
./scripts/init-skill.sh your-skill-name github.com/yourname/your-skill-name

# Start implementing your logic
code .
```

### Option 2: Manual Setup

If you prefer to manually configure the template, follow these steps:

## Step 1: Find and Replace Placeholders

Replace the following placeholder values throughout the project:

| Placeholder | Replace With | Example |
|------------|--------------|---------|
| `zentao` | Your skill name (lowercase) | `weatherapi` |
| `github.com/awesome-skill/zentao` | Your Go module path | `github.com/yourname/weatherapi` |
| `api.example.com` | Your default API host | `api.weatherapi.com` |

### Files to Update

1. **go.mod**

   ```go
   module github.com/yourname/your-skill-name
   ```

2. **internal/config/config.go**

   ```go
   const SkillName = "yourskillname"
   ```

   - Update default API host in `LoadForSkill()` if needed

3. **cmd/zentao/** (rename directory)

   ```bash
   mv cmd/zentao cmd/yourskillname
   ```

4. **All import statements**
   - Update all imports from `github.com/awesome-skill/zentao` to your module path
   - Find and replace across all `.go` files

5. **Makefile**

   ```makefile
   BINARIES := yourskillname
   BINARY_LINUX := yourskillname-linux-amd64
   ```

6. **cmd/yourskillname/cmd/root.go**
   - Update `Use`, `Short`, and `Long` descriptions

7. **cmd/yourskillname/cmd/config.go**
   - Update environment variable names (e.g., `YOURSKILL_API_KEY`)
   - Update configuration messages and examples

## Step 2: Run Go Module Commands

```bash
# Update module name and fetch dependencies
go mod edit -module github.com/yourname/your-skill-name
go mod tidy

# Verify everything compiles
go build ./...
```

## Step 3: Customize Configuration

The template provides a flexible configuration system. By default, it supports:

- API keys
- API hosts

### Adding Custom Configuration Fields

Edit `internal/config/config.go`:

```go
type Config struct {
    API APIConfig
    // Add your custom fields
    Custom CustomConfig
}

type CustomConfig struct {
    Region  string
    Timeout int
}
```

Update the `Load()` function to read your custom configuration from environment variables or files.

## Step 4: Implement Your Commands

### Create a New Command

1. Create a new file in `cmd/yourskillname/cmd/`, e.g., `fetch.go`:

```go
package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "github.com/yourname/your-skill-name/internal/config"
)

var fetchCmd = &cobra.Command{
    Use:   "fetch [resource]",
    Short: "Fetch a resource from the API",
    Args:  cobra.ExactArgs(1),
    RunE:  runFetch,
}

func init() {
    rootCmd.AddCommand(fetchCmd)
    // Add command-specific flags here
    fetchCmd.Flags().StringP("format", "f", "json", "Output format")
}

func runFetch(cmd *cobra.Command, args []string) error {
    // Load configuration
    cfg, err := config.Load()
    if err != nil {
        return fmt.Errorf("load config: %w", err)
    }

    // Implement your command logic
    resource := args[0]
    fmt.Printf("Fetching %s with API key: %s\n", resource, cfg.API.APIKey)
    
    // TODO: Implement your API call logic
    
    return nil
}
```

1. The command is automatically registered via `init()` function

### Using the Output Formatters

The template includes text, JSON, and table formatters:

```go
import (
    "os"
    "github.com/yourname/your-skill-name/internal/output"
)

func runFetch(cmd *cobra.Command, args []string) error {
    // Get format flag from root command
    format, _ := cmd.Flags().GetString("format")
    
    // Create formatter
    formatter, err := output.NewFormatter(format)
    if err != nil {
        return err
    }
    
    // Your data
    data := map[string]interface{}{
        "status": "success",
        "data":   "some value",
    }
    
    // Format and output
    return formatter.Format(data, os.Stdout)
}
```

### Custom Formatters

For complex data types, implement custom formatting:

**internal/output/custom.go**:

```go
package output

import (
    "fmt"
    "io"
)

// Add methods to existing formatters or create new types
func (f *TextFormatter) FormatMyData(data *MyDataType, w io.Writer) error {
    fmt.Fprintf(w, "Custom Format:\n")
    fmt.Fprintf(w, "Field1: %s\n", data.Field1)
    return nil
}
```

## Step 5: Implement API Client (Optional)

If your skill needs to call external APIs:

1. Create `internal/client/` directory:

```bash
mkdir -p internal/client
```

1. Implement your client:

```go
// internal/client/client.go
package client

import (
    "context"
    "encoding/json"
    "fmt"
    "net/http"
    "time"
)

type Client struct {
    apiKey  string
    apiHost string
    client  *http.Client
}

func NewClient(apiKey, apiHost string) *Client {
    return &Client{
        apiKey:  apiKey,
        apiHost: apiHost,
        client: &http.Client{
            Timeout: 30 * time.Second,
        },
    }
}

func (c *Client) Fetch(ctx context.Context, resource string) (*Response, error) {
    url := fmt.Sprintf("https://%s/api/%s", c.apiHost, resource)
    
    req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
    if err != nil {
        return nil, err
    }
    
    req.Header.Set("Authorization", "Bearer "+c.apiKey)
    
    resp, err := c.client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    
    var result Response
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return nil, err
    }
    
    return &result, nil
}
```

## Step 6: Add Tests

Follow Go testing best practices:

```go
// cmd/yourskillname/cmd/fetch_test.go
package cmd

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestFetchCommand(t *testing.T) {
    // Test your command logic
    assert.NotNil(t, fetchCmd)
}
```

Run tests:

```bash
go test ./...
```

## Step 7: Build and Test Your Skill

```bash
# Format code
make fmt

# Run linters
make lint

# Run tests
make test

# Build binary
make build

# Test your CLI
./bin/yourskillname --help
./bin/yourskillname config init
./bin/yourskillname config set-api-key your-key
./bin/yourskillname fetch something
```

## Step 8: Create Skill Definition (Optional)

If this skill is part of a larger skill ecosystem:

1. Create skill metadata:

```bash
mkdir -p skills/yourskillname
```

1. Create `skills/yourskillname/skill.toml`:

```toml
[skill]
name = "yourskillname"
version = "0.1.0"
description = "Your skill description"
author = "Your Name"

[skill.runtime]
type = "cli"
command = "yourskillname"

[skill.config]
api_key_required = true
config_path = "~/.config/awesome-skill/yourskillname"
```

1. Create `skills/yourskillname/SKILL.md` with detailed documentation

## Next Steps

- **Documentation**: Update README.md with your skill's purpose and usage
- **Examples**: Add example commands and use cases
- **CI/CD**: Set up GitHub Actions or other CI/CD pipelines
- **Distribution**: Publish binaries or add to package managers (Homebrew, etc.)
- **Error Handling**: Improve error messages and add retries for API calls
- **Logging**: Add structured logging for debugging

## Tips and Best Practices

1. **Follow Go conventions**: Use `gofmt`, `go vet`, and `golangci-lint`
2. **Keep commands simple**: Each command should do one thing well
3. **Validate input early**: Check arguments and flags before making API calls
4. **Handle errors gracefully**: Provide clear, actionable error messages
5. **Test thoroughly**: Write unit tests and integration tests
6. **Document everything**: Add godoc comments for all exported functions
7. **Use context**: Pass `context.Context` for cancellation and timeouts

## Troubleshooting

### Import errors after renaming

```bash
# Clean module cache and rebuild
go clean -modcache
go mod tidy
go build ./...
```

### Binary not found after build

```bash
# Check binary directory
ls -la bin/

# Make sure Makefile BINARIES variable is correct
```

### Tests failing

```bash
# Run with verbose output
go test -v ./...

# Run specific test
go test -v -run TestName ./path/to/package
```

## Getting Help

- Check `AGENTS.md` for development guidelines
- Review `docs/FRAMEWORK_GUIDE.md` for framework details
- See the example command in `cmd/zentao/cmd/example.go`

---

**Happy skill building!** 🚀
