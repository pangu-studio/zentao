package output

import (
	"encoding/json"
	"fmt"
	"io"
)

// Format defines the output format type
type Format string

const (
	// FormatText outputs in human-readable text format
	FormatText Format = "text"
	// FormatJSON outputs in JSON format
	FormatJSON Format = "json"
	// FormatTable outputs in table format
	FormatTable Format = "table"
)

// Formatter formats data for output
// Implement this interface with your custom data types
type Formatter interface {
	// Format formats any data structure to the writer
	Format(data interface{}, w io.Writer) error
}

// NewFormatter creates a new formatter based on the format type
func NewFormatter(format string) (Formatter, error) {
	switch Format(format) {
	case FormatText:
		return &TextFormatter{}, nil
	case FormatJSON:
		return &JSONFormatter{}, nil
	case FormatTable:
		return &TableFormatter{}, nil
	default:
		return nil, fmt.Errorf("unsupported format: %s (supported: text, json, table)", format)
	}
}

// JSONFormatter formats output as JSON
type JSONFormatter struct{}

// Format formats any data as JSON
func (f *JSONFormatter) Format(data interface{}, w io.Writer) error {
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	return encoder.Encode(data)
}
