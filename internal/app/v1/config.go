package v1

import (
	configparser "aoc/internal/utils/configparser"
)

type loggerConfig struct {
	Level  string `mapstructure:"level"`
	Format string `mapstructure:"format"`
	Source string `mapstructure:"source"`
}

// Keeps all the needed configs for running the app.
// It is sent as a parameter when a new app is created
type config struct {
	Logger loggerConfig `mapstructure:"logger"`
}

const (
	defaultLoggerLevel  = "debug"
	defaultLoggerFormat = "pretty"
)

// Parse and load the configuration from the config file.
func newConfig() (config, error) {
	cfg := config{}

	defaults := map[string]interface{}{
		"logger": map[string]interface{}{
			"level":  defaultLoggerLevel,
			"format": defaultLoggerFormat,
			"source": appID,
		},
	}

	if err := configparser.Parse(configFile, &cfg, defaults); err != nil {
		return cfg, err
	}

	return cfg, nil
}
