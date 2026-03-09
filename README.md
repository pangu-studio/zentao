# Skill Development Template

A comprehensive Go template for building CLI-based skills with best practices, testing, and documentation.

## 🎯 Overview

This template provides a complete foundation for creating command-line skills with:

- **CLI Framework**: Built with [Cobra](https://github.com/spf13/cobra) for powerful command-line interfaces
- **Configuration Management**: Flexible config system supporting environment variables and files
- **Multiple Output Formats**: Text, JSON, and table formatters built-in
- **Testing**: Structured testing with [testify](https://github.com/stretchr/testify)
- **Documentation**: Complete guides and examples
- **Build System**: Makefile with common development tasks

## ✨ Features

- 🚀 **Quick Setup**: Initialize a new skill in minutes with the provided script
- 🔧 **Flexible Configuration**: Environment variables, config files, and defaults
- 📊 **Multiple Output Formats**: Text, JSON, and table output built-in
- 🧪 **Testing Framework**: Comprehensive test structure and examples
- 📖 **Extensive Documentation**: Detailed guides for skill creation and framework usage
- 🛠️ **Development Tools**: Makefile targets for building, testing, and linting
- 🎨 **Code Quality**: Follows Go best practices and conventions

## 🚀 Quick Start

### Option 1: Using the Initialization Script (Recommended)

```bash
# Clone this template
git clone https://github.com/awesome-skill/template.git my-new-skill
cd my-new-skill

# Initialize your skill
./scripts/init-skill.sh weatherapi github.com/yourname/weatherapi "Your Name"

# Start developing
code .
```

### Option 2: Manual Setup

See [docs/CREATING_A_SKILL.md](docs/CREATING_A_SKILL.md) for detailed manual setup instructions.

## 📁 Project Structure

```
.
├── cmd/
│   └── myskill/              # CLI entry point
│       ├── main.go           # Main entry
│       └── cmd/              # Command implementations
│           ├── root.go       # Root command
│           ├── config.go     # Config management command
│           └── example.go    # Example command
├── internal/
│   ├── config/               # Configuration system
│   │   ├── config.go         # Config loading and management
│   │   └── config_test.go    # Config tests
│   └── output/               # Output formatters
│       ├── formatter.go      # Formatter interface and factory
│       ├── text.go           # Text formatter
│       ├── table.go          # Table formatter
│       └── formatter_test.go # Formatter tests
├── skills/
│   └── myskill/              # Skill definition
│       ├── SKILL.md          # Skill documentation template
│       └── skill.toml        # Skill configuration template
├── docs/
│   ├── CREATING_A_SKILL.md   # Skill creation guide
│   └── FRAMEWORK_GUIDE.md    # Framework documentation
├── scripts/
│   └── init-skill.sh         # Initialization script
├── Makefile                  # Build system
├── go.mod                    # Go module definition
├── AGENTS.md                 # AI agent guidelines
└── README.md                 # This file
```

## 🎓 Documentation

- **[Creating a Skill](docs/CREATING_A_SKILL.md)** - Step-by-step guide to create a new skill
- **[Framework Guide](docs/FRAMEWORK_GUIDE.md)** - Detailed framework documentation
- **[AI Agent Guidelines](AGENTS.md)** - Guidelines for AI coding agents

## 🛠️ Development

### Prerequisites

- Go 1.20 or later
- Make (optional, but recommended)

### Building

```bash
# Build all binaries
make build

# Build for Linux
make build-linux

# Install to $GOPATH/bin
make install
```

### Testing

```bash
# Run all tests
make test

# Run tests with coverage
make test-cover

# View coverage report
open coverage.html
```

### Code Quality

```bash
# Format code
make fmt

# Run go vet
make vet

# Run all linters
make lint

# Tidy dependencies
make tidy
```

### Available Make Targets

```bash
make build          # Build all CLIs
make build-linux    # Build for Linux AMD64
make install        # Install to system
make test           # Run tests
make test-cover     # Run tests with coverage
make fmt            # Format code
make vet            # Run go vet
make lint           # Run all linters
make tidy           # Tidy dependencies
make clean          # Clean build artifacts
make help           # Show help
```

## 📦 What's Included

### Configuration System

Flexible configuration management with:

- Environment variable support
- Config file support
- Default values
- Multiple configuration sources with priority

Example:

```go
cfg, err := config.Load()
if err != nil {
    return err
}
apiKey := cfg.API.APIKey
```

### Output Formatters

Built-in formatters for different output formats:

```go
formatter, err := output.NewFormatter("json")
if err != nil {
    return err
}

data := map[string]interface{}{"status": "success"}
formatter.Format(data, os.Stdout)
```

Supported formats:

- **Text**: Human-readable output
- **JSON**: Machine-readable JSON
- **Table**: ASCII table format

### Command Structure

Uses Cobra for building CLI commands:

```go
var myCmd = &cobra.Command{
    Use:   "mycommand",
    Short: "Description",
    RunE:  runMyCommand,
}

func init() {
    rootCmd.AddCommand(myCmd)
}
```

## 🔧 Customization

### Adding Configuration Fields

Edit `internal/config/config.go`:

```go
type Config struct {
    API    APIConfig
    Custom CustomConfig  // Add your fields
}

type CustomConfig struct {
    Timeout int
    Region  string
}
```

### Adding Commands

Create a new file in `cmd/myskill/cmd/`:

```go
package cmd

import "github.com/spf13/cobra"

var newCmd = &cobra.Command{
    Use:   "new",
    Short: "New command",
    RunE:  runNew,
}

func init() {
    rootCmd.AddCommand(newCmd)
}

func runNew(cmd *cobra.Command, args []string) error {
    // Implementation
    return nil
}
```

### Custom Output Formatters

Extend formatters in `internal/output/`:

```go
func (f *TextFormatter) FormatMyData(data *MyData, w io.Writer) error {
    fmt.Fprintf(w, "Field: %s\n", data.Field)
    return nil
}
```

## 📝 Example: Creating a Weather Skill

```bash
# Initialize
./scripts/init-skill.sh weatherapi github.com/john/weatherapi "John Doe"

# Implement your API client
cat > internal/client/client.go << 'EOF'
package client

type Client struct {
    apiKey string
}

func NewClient(apiKey string) *Client {
    return &Client{apiKey: apiKey}
}

func (c *Client) GetWeather(city string) (*Weather, error) {
    // Implementation
    return &Weather{}, nil
}
EOF

# Add a command
cat > cmd/weatherapi/cmd/get.go << 'EOF'
package cmd

import (
    "github.com/spf13/cobra"
    "github.com/john/weatherapi/internal/client"
    "github.com/john/weatherapi/internal/config"
)

var getCmd = &cobra.Command{
    Use:   "get [city]",
    Short: "Get weather for a city",
    Args:  cobra.ExactArgs(1),
    RunE:  runGet,
}

func init() {
    rootCmd.AddCommand(getCmd)
}

func runGet(cmd *cobra.Command, args []string) error {
    cfg, _ := config.Load()
    client := client.NewClient(cfg.API.APIKey)
    weather, err := client.GetWeather(args[0])
    // Handle and format output
    return nil
}
EOF

# Build and test
make build
./bin/weatherapi get Beijing
```

## 🤝 Contributing

Contributions are welcome! Feel free to:

- Report bugs
- Suggest features
- Submit pull requests
- Improve documentation

## 📄 License

This project is licensed under the GNU General Public License v3.0 - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- Built with [Cobra](https://github.com/spf13/cobra) - CLI framework
- Testing with [Testify](https://github.com/stretchr/testify) - Test toolkit
- Inspired by Go best practices and community conventions

## 📞 Support

- 📖 **Documentation**: Check [docs/](docs/) directory
- 🐛 **Issues**: Report on GitHub Issues
- 💡 **Discussions**: Open a GitHub Discussion

---

**Ready to build your skill?** Start with `./scripts/init-skill.sh` or read [docs/CREATING_A_SKILL.md](docs/CREATING_A_SKILL.md)! 🚀
