package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	// Global flags
	formatFlag  string
	verboseFlag bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "myskill",
	Short: "A CLI tool template for skill development",
	Long: `A template for building command-line skills with Cobra.

This template provides:
- Configuration management (API keys, hosts)
- Multiple output formats (text, json, table)
- Command structure and examples

Replace "myskill" with your actual skill name and implement your commands.`,
	Version: "0.1.0",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Global flags
	rootCmd.PersistentFlags().StringVarP(&formatFlag, "format", "f", "text", "Output format: text, json, table")
	rootCmd.PersistentFlags().BoolVarP(&verboseFlag, "verbose", "v", false, "Verbose output")
}

// printError prints error message to stderr
func printError(err error) {
	fmt.Fprintf(os.Stderr, "Error: %v\n", err)
}
