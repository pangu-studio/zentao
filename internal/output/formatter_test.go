package output

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewFormatter(t *testing.T) {
	testCases := []struct {
		name        string
		format      string
		expectError bool
		expectType  interface{}
	}{
		{"text format", "text", false, &TextFormatter{}},
		{"json format", "json", false, &JSONFormatter{}},
		{"table format", "table", false, &TableFormatter{}},
		{"invalid format", "invalid", true, nil},
		{"empty format", "", true, nil},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			formatter, err := NewFormatter(tc.format)

			if tc.expectError {
				assert.Error(t, err)
				assert.Nil(t, formatter)
			} else {
				require.NoError(t, err)
				assert.IsType(t, tc.expectType, formatter)
			}
		})
	}
}

func TestJSONFormatter_Format(t *testing.T) {
	formatter := &JSONFormatter{}

	testCases := []struct {
		name string
		data interface{}
	}{
		{
			name: "simple map",
			data: map[string]interface{}{
				"key1": "value1",
				"key2": 123,
			},
		},
		{
			name: "string",
			data: "test string",
		},
		{
			name: "struct",
			data: struct {
				Name  string
				Value int
			}{
				Name:  "test",
				Value: 42,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var buf bytes.Buffer
			err := formatter.Format(tc.data, &buf)
			require.NoError(t, err)

			// Verify it's valid JSON
			var result interface{}
			err = json.Unmarshal(buf.Bytes(), &result)
			assert.NoError(t, err, "Output should be valid JSON")
		})
	}
}

func TestTextFormatter_Format(t *testing.T) {
	formatter := &TextFormatter{}

	testCases := []struct {
		name           string
		data           interface{}
		expectedInclue string
	}{
		{
			name:           "string",
			data:           "hello world",
			expectedInclue: "hello world",
		},
		{
			name: "map",
			data: map[string]interface{}{
				"key1": "value1",
				"key2": "value2",
			},
			expectedInclue: "key",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var buf bytes.Buffer
			err := formatter.Format(tc.data, &buf)
			require.NoError(t, err)

			output := buf.String()
			assert.NotEmpty(t, output)
			if tc.expectedInclue != "" {
				assert.Contains(t, output, tc.expectedInclue)
			}
		})
	}
}

func TestTableFormatter_Format(t *testing.T) {
	formatter := &TableFormatter{}

	testCases := []struct {
		name string
		data interface{}
	}{
		{
			name: "simple map",
			data: map[string]string{
				"Key1": "Value1",
				"Key2": "Value2",
			},
		},
		{
			name: "interface map",
			data: map[string]interface{}{
				"Name":  "Test",
				"Count": 42,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var buf bytes.Buffer
			err := formatter.Format(tc.data, &buf)
			require.NoError(t, err)

			output := buf.String()
			assert.NotEmpty(t, output)
			// Table should contain separator lines
			assert.True(t, strings.Contains(output, "="), "Table should contain separator lines")
		})
	}
}

func TestPrintTable(t *testing.T) {
	var buf bytes.Buffer
	rows := [][]string{
		{"Name", "John Doe"},
		{"Age", "30"},
		{"City", "New York"},
	}

	PrintTable(&buf, rows)

	output := buf.String()
	assert.Contains(t, output, "Name")
	assert.Contains(t, output, "John Doe")
	assert.Contains(t, output, "Age")
	assert.Contains(t, output, "30")
}

func TestPrintTableWithHeaders(t *testing.T) {
	var buf bytes.Buffer
	headers := []string{"Name", "Age", "City"}
	rows := [][]string{
		{"John", "30", "NYC"},
		{"Jane", "25", "LA"},
	}
	widths := []int{10, 5, 10}

	PrintTableWithHeaders(&buf, headers, rows, widths)

	output := buf.String()
	assert.Contains(t, output, "Name")
	assert.Contains(t, output, "Age")
	assert.Contains(t, output, "City")
	assert.Contains(t, output, "John")
	assert.Contains(t, output, "Jane")
	// Should contain separator line
	assert.True(t, strings.Contains(output, "-"), "Table should contain separator line")
}

func TestFormatConstants(t *testing.T) {
	assert.Equal(t, Format("text"), FormatText)
	assert.Equal(t, Format("json"), FormatJSON)
	assert.Equal(t, Format("table"), FormatTable)
}
