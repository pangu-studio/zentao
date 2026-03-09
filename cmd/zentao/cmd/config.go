package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/pangu-studio/zentao/internal/config"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage configuration",
	Long: `Manage API configuration.

Configuration files are stored in:
  $HOME/.config/awesome-skill/zentao/api_key   - API Key
  $HOME/.config/awesome-skill/zentao/api_host  - API Host (optional)

Priority order for API Key:
  1. Environment variable ZENTAO_API_KEY
  2. Configuration file

Priority order for API Host:
  1. Environment variable ZENTAO_API_HOST
  2. Configuration file
  3. Default: api.example.com`,
	RunE: runConfig,
}

var configInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize configuration directory",
	Long: `Initialize configuration directory and optionally set API key and host.

Examples:
  zentao config init
  zentao config init --interactive`,
	RunE: runConfigInit,
}

var configSetAPIKeyCmd = &cobra.Command{
	Use:   "set-api-key <key>",
	Short: "Set API key",
	Long: `Set the API key.

The API key will be saved to:
  $HOME/.config/awesome-skill/zentao/api_key

Examples:
  zentao config set-api-key your-api-key-here`,
	Args: cobra.ExactArgs(1),
	RunE: runConfigSetAPIKey,
}

var configSetAPIHostCmd = &cobra.Command{
	Use:   "set-api-host <host>",
	Short: "Set API host",
	Long: `Set the API host.

The API host will be saved to:
  $HOME/.config/awesome-skill/zentao/api_host

Examples:
  zentao config set-api-host api.production.com`,
	Args: cobra.ExactArgs(1),
	RunE: runConfigSetAPIHost,
}

var (
	configInteractive bool
)

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(configInitCmd)
	configCmd.AddCommand(configSetAPIKeyCmd)
	configCmd.AddCommand(configSetAPIHostCmd)

	configInitCmd.Flags().BoolVarP(&configInteractive, "interactive", "i", false, "Interactive mode to set API key and host")
}

func runConfig(cmd *cobra.Command, args []string) error {
	configDir, err := config.GetConfigDir()
	if err != nil {
		printError(err)
		return err
	}

	fmt.Printf("Configuration directory: %s/%s\n\n", configDir, config.SkillName)

	cfg, err := config.Load()
	if err != nil {
		fmt.Printf("Status: Not configured\n\n")
		fmt.Printf("To configure, run:\n")
		fmt.Printf("  zentao config init\n")
		fmt.Printf("  zentao config set-api-key <your-api-key>\n")
		return nil
	}

	fmt.Printf("Status: Configured\n\n")
	fmt.Printf("API Key: %s...\n", maskAPIKey(cfg.API.APIKey))
	fmt.Printf("API Host: %s\n", cfg.API.APIHost)

	return nil
}

func runConfigInit(cmd *cobra.Command, args []string) error {
	if err := config.EnsureConfigDir(); err != nil {
		printError(err)
		return err
	}

	configDir, err := config.GetConfigDir()
	if err != nil {
		printError(err)
		return err
	}

	fmt.Printf("Config directory created: %s/%s\n\n", configDir, config.SkillName)

	if !configInteractive {
		fmt.Printf("To set your API key, run:\n")
		fmt.Printf("  zentao config set-api-key <your-api-key>\n")
		return nil
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Enter your API key (or press Enter to skip): ")
	apiKey, err := reader.ReadString('\n')
	if err != nil {
		printError(err)
		return err
	}
	apiKey = strings.TrimSpace(apiKey)

	if apiKey != "" {
		if err := config.SetAPIKey(apiKey); err != nil {
			printError(err)
			return err
		}
		fmt.Printf("API key saved.\n\n")
	}

	fmt.Printf("Enter your API host (or press Enter to use default: api.example.com): ")
	apiHost, err := reader.ReadString('\n')
	if err != nil {
		printError(err)
		return err
	}
	apiHost = strings.TrimSpace(apiHost)

	if apiHost != "" {
		if err := config.SetAPIHost(apiHost); err != nil {
			printError(err)
			return err
		}
		fmt.Printf("API host saved: %s\n", apiHost)
	}

	fmt.Printf("\nConfiguration complete!\n")
	return nil
}

func runConfigSetAPIKey(cmd *cobra.Command, args []string) error {
	apiKey := args[0]

	if err := config.SetAPIKey(apiKey); err != nil {
		printError(err)
		return err
	}

	configDir, err := config.GetConfigDir()
	if err != nil {
		printError(err)
		return err
	}

	fmt.Printf("API key saved to: %s/%s/api_key\n", configDir, config.SkillName)
	return nil
}

func runConfigSetAPIHost(cmd *cobra.Command, args []string) error {
	apiHost := args[0]

	if err := config.SetAPIHost(apiHost); err != nil {
		printError(err)
		return err
	}

	configDir, err := config.GetConfigDir()
	if err != nil {
		printError(err)
		return err
	}

	fmt.Printf("API host saved to: %s/%s/api_host\n", configDir, config.SkillName)
	return nil
}

func maskAPIKey(key string) string {
	if len(key) <= 8 {
		return strings.Repeat("*", len(key))
	}
	return key[:4] + strings.Repeat("*", len(key)-8) + key[len(key)-4:]
}
