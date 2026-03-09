#!/bin/bash

# init-skill.sh - Initialize a new skill from this template
# Usage: ./scripts/init-skill.sh <skill-name> <module-path> [author-name]

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Print colored output
print_error() {
    echo -e "${RED}Error: $1${NC}" >&2
}

print_success() {
    echo -e "${GREEN}✓ $1${NC}"
}

print_info() {
    echo -e "${YELLOW}→ $1${NC}"
}

# Display usage
usage() {
    cat << EOF
Usage: $0 <skill-name> <module-path> [author-name]

Arguments:
  skill-name    Name of your skill (lowercase, e.g., "weatherapi")
  module-path   Go module path (e.g., "github.com/yourname/weatherapi")
  author-name   Your name (optional, defaults to git user.name)

Example:
  $0 weatherapi github.com/johndoe/weatherapi "John Doe"

This script will:
  1. Rename directories and files
  2. Replace all placeholder text
  3. Update go.mod
  4. Run go mod tidy
  5. Format the code

EOF
    exit 1
}

# Validate arguments
if [ $# -lt 2 ]; then
    print_error "Insufficient arguments"
    usage
fi

SKILL_NAME="$1"
MODULE_PATH="$2"
AUTHOR_NAME="${3:-$(git config user.name 2>/dev/null || echo "Your Name")}"

# Validate skill name (lowercase, alphanumeric, hyphens)
if ! [[ "$SKILL_NAME" =~ ^[a-z0-9-]+$ ]]; then
    print_error "Skill name must be lowercase alphanumeric with hyphens only"
    exit 1
fi

# Validate module path
if ! [[ "$MODULE_PATH" =~ ^[a-z0-9./\-]+$ ]]; then
    print_error "Invalid module path format"
    exit 1
fi

# Derived values
SKILL_NAME_UPPER=$(echo "$SKILL_NAME" | tr '[:lower:]' '[:upper:]' | tr '-' '_')
USERNAME=$(echo "$MODULE_PATH" | cut -d'/' -f2)

print_info "Initializing skill: $SKILL_NAME"
print_info "Module path: $MODULE_PATH"
print_info "Author: $AUTHOR_NAME"
echo ""

# Check if we're in the template directory
if [ ! -f "go.mod" ] || [ ! -d "cmd/myskill" ]; then
    print_error "This script must be run from the template root directory"
    exit 1
fi

# Confirm before proceeding
read -p "This will modify files in the current directory. Continue? (y/N) " -n 1 -r
echo
if [[ ! $REPLY =~ ^[Yy]$ ]]; then
    print_error "Aborted by user"
    exit 1
fi

echo ""
print_info "Step 1/8: Renaming directories..."

# Rename cmd/myskill to cmd/<skill-name>
if [ -d "cmd/myskill" ]; then
    mv cmd/myskill "cmd/$SKILL_NAME"
    print_success "Renamed cmd/myskill to cmd/$SKILL_NAME"
fi

# Rename skills/myskill to skills/<skill-name>
if [ -d "skills/myskill" ]; then
    mv skills/myskill "skills/$SKILL_NAME"
    print_success "Renamed skills/myskill to skills/$SKILL_NAME"
fi

print_info "Step 2/8: Updating go.mod..."

# Update go.mod
if [ -f "go.mod" ]; then
    sed -i.bak "s|github.com/awesome-skill/template|$MODULE_PATH|g" go.mod
    rm go.mod.bak
    print_success "Updated module path in go.mod"
fi

print_info "Step 3/8: Updating import paths in Go files..."

# Update all Go files with new import paths
find . -name "*.go" -type f -exec sed -i.bak "s|github.com/awesome-skill/template|$MODULE_PATH|g" {} \;
find . -name "*.go.bak" -type f -delete
print_success "Updated import paths in all Go files"

print_info "Step 4/8: Replacing skill name placeholders..."

# Replace myskill with actual skill name in Go files
find . -name "*.go" -type f -exec sed -i.bak "s|cmd/myskill|cmd/$SKILL_NAME|g" {} \;
find . -name "*.go.bak" -type f -delete

# Update SkillName constant in config.go
if [ -f "internal/config/config.go" ]; then
    sed -i.bak "s|const SkillName = \"myskill\"|const SkillName = \"$SKILL_NAME\"|g" internal/config/config.go
    rm internal/config/config.go.bak
    print_success "Updated SkillName constant"
fi

# Update root command
if [ -f "cmd/$SKILL_NAME/cmd/root.go" ]; then
    sed -i.bak "s|Use:   \"myskill\"|Use:   \"$SKILL_NAME\"|g" "cmd/$SKILL_NAME/cmd/root.go"
    rm "cmd/$SKILL_NAME/cmd/root.go.bak"
    print_success "Updated root command name"
fi

print_info "Step 5/8: Updating Makefile..."

if [ -f "Makefile" ]; then
    sed -i.bak "s|BINARIES := myskill|BINARIES := $SKILL_NAME|g" Makefile
    sed -i.bak "s|myskill-linux-amd64|$SKILL_NAME-linux-amd64|g" Makefile
    sed -i.bak "s|cmd/myskill|cmd/$SKILL_NAME|g" Makefile
    rm Makefile.bak
    print_success "Updated Makefile"
fi

print_info "Step 6/8: Updating documentation templates..."

# Update SKILL.md
if [ -f "skills/$SKILL_NAME/SKILL.md" ]; then
    sed -i.bak "s|{{SKILL_NAME}}|$SKILL_NAME|g" "skills/$SKILL_NAME/SKILL.md"
    sed -i.bak "s|{{SKILL_NAME_UPPER}}|$SKILL_NAME_UPPER|g" "skills/$SKILL_NAME/SKILL.md"
    sed -i.bak "s|{{USERNAME}}|$USERNAME|g" "skills/$SKILL_NAME/SKILL.md"
    sed -i.bak "s|{{AUTHOR_NAME}}|$AUTHOR_NAME|g" "skills/$SKILL_NAME/SKILL.md"
    rm "skills/$SKILL_NAME/SKILL.md.bak"
    print_success "Updated SKILL.md"
fi

# Update skill.toml
if [ -f "skills/$SKILL_NAME/skill.toml" ]; then
    sed -i.bak "s|{{SKILL_NAME}}|$SKILL_NAME|g" "skills/$SKILL_NAME/skill.toml"
    sed -i.bak "s|{{SKILL_NAME_UPPER}}|$SKILL_NAME_UPPER|g" "skills/$SKILL_NAME/skill.toml"
    sed -i.bak "s|{{USERNAME}}|$USERNAME|g" "skills/$SKILL_NAME/skill.toml"
    sed -i.bak "s|{{AUTHOR_NAME}}|$AUTHOR_NAME|g" "skills/$SKILL_NAME/skill.toml"
    rm "skills/$SKILL_NAME/skill.toml.bak"
    print_success "Updated skill.toml"
fi

# Update documentation files
for doc_file in docs/*.md; do
    if [ -f "$doc_file" ]; then
        sed -i.bak "s|myskill|$SKILL_NAME|g" "$doc_file"
        sed -i.bak "s|MYSKILL|$SKILL_NAME_UPPER|g" "$doc_file"
        sed -i.bak "s|github.com/awesome-skill/template|$MODULE_PATH|g" "$doc_file"
        rm "$doc_file.bak"
    fi
done
print_success "Updated documentation files"

print_info "Step 7/8: Running go mod tidy..."

if command -v go &> /dev/null; then
    go mod tidy
    print_success "Dependencies tidied"
else
    print_error "Go not found, skipping 'go mod tidy'"
fi

print_info "Step 8/8: Formatting code..."

if command -v go &> /dev/null; then
    go fmt ./...
    print_success "Code formatted"
fi

echo ""
print_success "Skill initialization complete!"
echo ""
echo "Next steps:"
echo "  1. Review and update skills/$SKILL_NAME/SKILL.md"
echo "  2. Review and update skills/$SKILL_NAME/skill.toml"
echo "  3. Update README.md with your skill's information"
echo "  4. Implement your commands in cmd/$SKILL_NAME/cmd/"
echo "  5. Build and test:"
echo "       make build"
echo "       ./bin/$SKILL_NAME --help"
echo ""
echo "For detailed instructions, see:"
echo "  - docs/CREATING_A_SKILL.md"
echo "  - docs/FRAMEWORK_GUIDE.md"
echo ""
print_success "Happy coding! 🚀"
