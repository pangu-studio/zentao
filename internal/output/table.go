package output

import (
	"fmt"
	"io"
	"strings"
)

// TableFormatter formats output as ASCII tables
type TableFormatter struct{}

// Format formats data as an ASCII table
// For simple data, it will format it as a key-value table
// For more complex tables, implement custom formatting methods
func (f *TableFormatter) Format(data interface{}, w io.Writer) error {
	switch v := data.(type) {
	case map[string]interface{}:
		// Format map as key-value table
		return f.formatKeyValueTable(v, w)
	case map[string]string:
		// Format string map as key-value table
		converted := make(map[string]interface{})
		for k, val := range v {
			converted[k] = val
		}
		return f.formatKeyValueTable(converted, w)
	default:
		// Fallback to simple string representation
		fmt.Fprintf(w, "%+v\n", data)
	}

	return nil
}

// formatKeyValueTable formats a map as a key-value table
func (f *TableFormatter) formatKeyValueTable(data map[string]interface{}, w io.Writer) error {
	if len(data) == 0 {
		fmt.Fprintf(w, "No data to display.\n")
		return nil
	}

	// Find maximum key length for alignment
	maxKeyLen := 0
	for key := range data {
		if len(key) > maxKeyLen {
			maxKeyLen = len(key)
		}
	}

	// Print table
	fmt.Fprintf(w, "%s\n", strings.Repeat("=", maxKeyLen+30))
	for key, val := range data {
		fmt.Fprintf(w, "%-*s: %v\n", maxKeyLen, key, val)
	}
	fmt.Fprintf(w, "%s\n", strings.Repeat("=", maxKeyLen+30))

	return nil
}

// TODO: Add custom table formatting methods for your specific data types
// Example helper methods for building custom tables:

// PrintTable prints a simple key-value table
func PrintTable(w io.Writer, rows [][]string) {
	if len(rows) == 0 {
		return
	}

	maxKeyLen := 0
	for _, row := range rows {
		if len(row) > 0 && len(row[0]) > maxKeyLen {
			maxKeyLen = len(row[0])
		}
	}

	for _, row := range rows {
		if len(row) >= 2 {
			fmt.Fprintf(w, "%-*s: %s\n", maxKeyLen, row[0], row[1])
		}
	}
}

// PrintTableWithHeaders prints a table with headers and multiple columns
func PrintTableWithHeaders(w io.Writer, headers []string, rows [][]string, columnWidths []int) {
	if len(headers) == 0 || len(rows) == 0 {
		return
	}

	// Print header
	format := buildFormatString(columnWidths)
	args := make([]interface{}, len(headers))
	for i, h := range headers {
		args[i] = h
	}
	fmt.Fprintf(w, format+"\n", args...)

	// Print separator
	totalWidth := 0
	for _, w := range columnWidths {
		totalWidth += w + 1
	}
	fmt.Fprintf(w, "%s\n", strings.Repeat("-", totalWidth))

	// Print rows
	for _, row := range rows {
		args := make([]interface{}, len(row))
		for i, cell := range row {
			args[i] = cell
		}
		fmt.Fprintf(w, format+"\n", args...)
	}
}

// buildFormatString builds a format string for table printing
func buildFormatString(widths []int) string {
	var parts []string
	for _, w := range widths {
		parts = append(parts, fmt.Sprintf("%%-%ds", w))
	}
	return strings.Join(parts, " ")
}
