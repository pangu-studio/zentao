package output

import (
	"fmt"
	"io"
)

// TextFormatter formats output as human-readable text
type TextFormatter struct{}

// Format formats data as human-readable text
// For complex data structures, you may want to implement custom formatting logic
// based on the type of data passed in
func (f *TextFormatter) Format(data interface{}, w io.Writer) error {
	// Example: Simple string representation
	// You can use type assertion to handle different data types
	// and provide custom formatting for each

	switch v := data.(type) {
	case string:
		fmt.Fprintf(w, "%s\n", v)
	case map[string]interface{}:
		// Example: format map data
		for key, val := range v {
			fmt.Fprintf(w, "%s: %v\n", key, val)
		}
	default:
		// Fallback: use fmt's default formatting
		fmt.Fprintf(w, "%+v\n", data)
	}

	return nil
}

// TODO: Add custom formatting methods for your specific data types
// Example:
// func (f *TextFormatter) FormatMyData(data *MyDataType, w io.Writer) error {
//     fmt.Fprintf(w, "Field1: %s\n", data.Field1)
//     fmt.Fprintf(w, "Field2: %d\n", data.Field2)
//     return nil
// }
