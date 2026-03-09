# {{SKILL_NAME}} Skill

<!-- TODO: Replace {{SKILL_NAME}} with your actual skill name -->

## Overview

{{SHORT_DESCRIPTION}}

<!-- TODO: Provide a brief description of what your skill does -->

## Features

- Feature 1
- Feature 2
- Feature 3

<!-- TODO: List key features -->

## Installation

### From Source

```bash
go install github.com/{{USERNAME}}/{{SKILL_NAME}}/cmd/{{SKILL_NAME}}@latest
```

### From Binary

Download the latest release from the [releases page](https://github.com/{{USERNAME}}/{{SKILL_NAME}}/releases).

## Configuration

### API Key Setup

```bash
# Initialize configuration
{{SKILL_NAME}} config init --interactive

# Or set via environment variable
export {{SKILL_NAME_UPPER}}_API_KEY="your-api-key-here"

# Or set via command
{{SKILL_NAME}} config set-api-key your-api-key-here
```

### Configuration Files

Configuration is stored in:
- Linux/macOS: `~/.config/awesome-skill/{{SKILL_NAME}}/`
- Windows: `%APPDATA%\awesome-skill\{{SKILL_NAME}}\`

Files:
- `api_key` - Your API key
- `api_host` - API endpoint (optional)

## Usage

### Basic Commands

```bash
# Show help
{{SKILL_NAME}} --help

# Configure API key
{{SKILL_NAME}} config init

# Example command
{{SKILL_NAME}} example World

# Example with options
{{SKILL_NAME}} example World --greeting "Hi"
```

### Output Formats

The skill supports multiple output formats:

```bash
# Text format (default)
{{SKILL_NAME}} example World

# JSON format
{{SKILL_NAME}} example World --format json

# Table format
{{SKILL_NAME}} example World --format table
```

### Command Reference

<!-- TODO: Document your actual commands -->

#### `{{SKILL_NAME}} config`

Manage configuration settings.

**Subcommands:**
- `init` - Initialize configuration directory
- `set-api-key <key>` - Set API key
- `set-api-host <host>` - Set API host

**Examples:**
```bash
{{SKILL_NAME}} config init
{{SKILL_NAME}} config set-api-key abc123
```

#### `{{SKILL_NAME}} example`

Example command showing basic usage.

**Usage:**
```bash
{{SKILL_NAME}} example [name]
```

**Flags:**
- `-g, --greeting` - Greeting message (default: "Hello")

**Examples:**
```bash
{{SKILL_NAME}} example World
{{SKILL_NAME}} example Alice --greeting "Hi"
```

<!-- TODO: Add your actual commands here -->

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `{{SKILL_NAME_UPPER}}_API_KEY` | API authentication key | (required) |
| `{{SKILL_NAME_UPPER}}_API_HOST` | API endpoint host | `api.example.com` |

<!-- TODO: Update with your actual environment variables -->

## Examples

### Example 1: Basic Usage

```bash
{{SKILL_NAME}} example World
# Output: Hello, World!
```

### Example 2: Custom Greeting

```bash
{{SKILL_NAME}} example Alice --greeting "Hi"
# Output: Hi, Alice!
```

### Example 3: JSON Output

```bash
{{SKILL_NAME}} example World --format json
```

<!-- TODO: Add real-world examples -->

## Development

See [CREATING_A_SKILL.md](../docs/CREATING_A_SKILL.md) for development guide.

### Building from Source

```bash
# Clone repository
git clone https://github.com/{{USERNAME}}/{{SKILL_NAME}}.git
cd {{SKILL_NAME}}

# Install dependencies
go mod download

# Build
make build

# Run tests
make test

# Run linters
make lint
```

## Troubleshooting

### API Key Not Found

```
Error: API key not found. Please set {{SKILL_NAME_UPPER}}_API_KEY environment variable or create ~/.config/awesome-skill/{{SKILL_NAME}}/api_key
```

**Solution:** Configure your API key using one of the methods in the Configuration section.

### Connection Timeout

If you're experiencing timeouts, check:
1. Your internet connection
2. The API host is correct
3. Your API key is valid

<!-- TODO: Add common issues and solutions -->

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

<!-- TODO: Add your license information -->

This project is licensed under the GNU General Public License v3.0 - see the [LICENSE](../LICENSE) file for details.

## Acknowledgments

<!-- TODO: Add acknowledgments -->

- Built with [Cobra](https://github.com/spf13/cobra)
- Template from [awesome-skill/template](https://github.com/awesome-skill/template)

## Support

<!-- TODO: Add support information -->

- Issues: https://github.com/{{USERNAME}}/{{SKILL_NAME}}/issues
- Documentation: https://github.com/{{USERNAME}}/{{SKILL_NAME}}/wiki
- Email: your-email@example.com

---

**Note:** This is a template. Replace all `{{PLACEHOLDER}}` values with your actual information.
