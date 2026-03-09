package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// exampleCmd represents an example command
var exampleCmd = &cobra.Command{
	Use:   "example [name]",
	Short: "Example command demonstrating basic usage",
	Long: `An example command to demonstrate how to create new commands.

This command shows:
- How to accept arguments
- How to use flags
- Basic command structure
- Error handling

Replace this with your actual command implementation.`,
	Args: cobra.MaximumNArgs(1),
	RunE: runExample,
}

var (
	exampleGreeting string
)

func init() {
	rootCmd.AddCommand(exampleCmd)
	exampleCmd.Flags().StringVarP(&exampleGreeting, "greeting", "g", "Hello", "Greeting message to use")
}

func runExample(cmd *cobra.Command, args []string) error {
	name := "World"
	if len(args) > 0 {
		name = args[0]
	}

	message := fmt.Sprintf("%s, %s!", exampleGreeting, name)

	if verboseFlag {
		fmt.Printf("Verbose mode is enabled\n")
		fmt.Printf("Output format: %s\n", formatFlag)
		fmt.Printf("Message: %s\n", message)
	} else {
		fmt.Println(message)
	}

	return nil
}
