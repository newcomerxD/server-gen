// File: internal/config/config.go
package config

import (
    "fmt"

    "github.com/spf13/viper"
)

// SMTPConfig holds SMTP connection and authentication settings.
type SMTPConfig struct {
    Host     string   `mapstructure:"host"`
    Port     int      `mapstructure:"port"`
    Username string   `mapstructure:"username"`
    Password string   `mapstructure:"password"`
    From     string   `mapstructure:"from"`
    To       []string `mapstructure:"to"`
}

// Config represents the full application configuration.
type Config struct {
    Schedule string     `mapstructure:"schedule"`
    SMTP     SMTPConfig `mapstructure:"smtp"`
    Modules  []string   `mapstructure:"modules"`
}

// Load reads and parses the config file at the given path.
func Load(path string) (*Config, error) {
    viper.SetConfigFile(path)
    if err := viper.ReadInConfig(); err != nil {
        return nil, fmt.Errorf("read config: %w", err)
    }
    var cfg Config
    if err := viper.Unmarshal(&cfg); err != nil {
        return nil, fmt.Errorf("unmarshal config: %w", err)
    }
    return &cfg, nil
}
