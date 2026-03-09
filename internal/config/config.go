package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
)

// SkillName is the name of the skill
// TODO: Replace "myskill" with your actual skill name
const SkillName = "myskill"

// Config holds the application configuration
type Config struct {
	API APIConfig
}

// APIConfig contains API configuration
type APIConfig struct {
	APIKey  string
	APIHost string
}

// Load loads configuration from environment variables and config files
// Environment variables use uppercase skill name: {SKILLNAME}_API_KEY, {SKILLNAME}_API_HOST
func Load() (*Config, error) {
	return LoadForSkill(SkillName)
}

// LoadForSkill loads configuration for a specific skill name
func LoadForSkill(skillName string) (*Config, error) {
	// Load .env file if it exists
	// This will not error if .env doesn't exist, which is fine
	_ = godotenv.Load()

	// Build environment variable names
	envPrefix := strings.ToUpper(skillName)
	apiKeyEnv := envPrefix + "_API_KEY"
	apiHostEnv := envPrefix + "_API_HOST"

	cfg := &Config{
		API: APIConfig{
			// Default host can be customized per skill
			APIHost: getEnvWithDefault(apiHostEnv, "api.example.com"),
		},
	}

	// Try to load API key from environment variable first
	apiKey := os.Getenv(apiKeyEnv)
	if apiKey != "" {
		cfg.API.APIKey = apiKey
		return cfg, nil
	}

	// Try to load from config file
	configDir, err := getConfigDir()
	if err != nil {
		return nil, fmt.Errorf("get config directory: %w", err)
	}

	apiKeyPath := filepath.Join(configDir, skillName, "api_key")
	apiKeyBytes, err := os.ReadFile(apiKeyPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("API key not found. Please set %s environment variable or create %s", apiKeyEnv, apiKeyPath)
		}
		return nil, fmt.Errorf("read API key file: %w", err)
	}

	cfg.API.APIKey = strings.TrimSpace(string(apiKeyBytes))
	if cfg.API.APIKey == "" {
		return nil, fmt.Errorf("API key file is empty: %s", apiKeyPath)
	}

	// Load API host from config file if not set via environment variable
	if os.Getenv(apiHostEnv) == "" {
		apiHostPath := filepath.Join(configDir, skillName, "api_host")
		if apiHostBytes, err := os.ReadFile(apiHostPath); err == nil {
			if host := strings.TrimSpace(string(apiHostBytes)); host != "" {
				cfg.API.APIHost = host
			}
		}
	}

	return cfg, nil
}

// getConfigDir returns the configuration directory path
func getConfigDir() (string, error) {
	// Check for XDG_CONFIG_HOME first
	if xdgConfig := os.Getenv("XDG_CONFIG_HOME"); xdgConfig != "" {
		return filepath.Join(xdgConfig, "awesome-skill"), nil
	}

	// Fall back to ~/.config
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("get user home directory: %w", err)
	}

	return filepath.Join(homeDir, ".config", "awesome-skill"), nil
}

// getEnvWithDefault gets environment variable with a default value
func getEnvWithDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// EnsureConfigDir creates the config directory if it doesn't exist
func EnsureConfigDir() error {
	return EnsureConfigDirForSkill(SkillName)
}

// EnsureConfigDirForSkill creates the config directory for a specific skill if it doesn't exist
func EnsureConfigDirForSkill(skillName string) error {
	configDir, err := getConfigDir()
	if err != nil {
		return err
	}

	skillDir := filepath.Join(configDir, skillName)
	if err := os.MkdirAll(skillDir, 0755); err != nil {
		return fmt.Errorf("create config directory: %w", err)
	}

	return nil
}

// GetConfigDir returns the config directory path
func GetConfigDir() (string, error) {
	return getConfigDir()
}

// SetAPIKey saves the API key to the config file
func SetAPIKey(apiKey string) error {
	return SetAPIKeyForSkill(SkillName, apiKey)
}

// SetAPIKeyForSkill saves the API key to the config file for a specific skill
func SetAPIKeyForSkill(skillName, apiKey string) error {
	if err := EnsureConfigDirForSkill(skillName); err != nil {
		return err
	}

	configDir, err := getConfigDir()
	if err != nil {
		return err
	}

	apiKeyPath := filepath.Join(configDir, skillName, "api_key")
	if err := os.WriteFile(apiKeyPath, []byte(apiKey), 0600); err != nil {
		return fmt.Errorf("write API key file: %w", err)
	}

	return nil
}

// SetAPIHost saves the API host to the config file
func SetAPIHost(apiHost string) error {
	return SetAPIHostForSkill(SkillName, apiHost)
}

// SetAPIHostForSkill saves the API host to the config file for a specific skill
func SetAPIHostForSkill(skillName, apiHost string) error {
	if err := EnsureConfigDirForSkill(skillName); err != nil {
		return err
	}

	configDir, err := getConfigDir()
	if err != nil {
		return err
	}

	apiHostPath := filepath.Join(configDir, skillName, "api_host")
	if err := os.WriteFile(apiHostPath, []byte(apiHost), 0600); err != nil {
		return fmt.Errorf("write API host file: %w", err)
	}

	return nil
}
