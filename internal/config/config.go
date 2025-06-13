// File: internal/config/config.go
package config

import (
    "fmt"
    "os"
    "time"

    "github.com/go-playground/validator/v10"
    "github.com/spf13/viper"
)

// SMTPConfig holds SMTP settings.
type SMTPConfig struct {
    Host     string   `mapstructure:"host" validate:"required,hostname|ip"`
    Port     int      `mapstructure:"port" validate:"required,min=1,max=65535"`
    Username string   `mapstructure:"username" validate:"required"`
    Password string   `mapstructure:"password" validate:"required"`
    From     string   `mapstructure:"from" validate:"required,email"`
    To       []string `mapstructure:"to" validate:"required,min=1,dive,email"`
}

// HTTPConfig holds health endpoint settings.
type HTTPConfig struct {
    Addr string `mapstructure:"addr" validate:"required"`
}

// Config is the application configuration.
type Config struct {
    Schedule string     `mapstructure:"schedule" validate:"required"`
    SMTP     SMTPConfig `mapstructure:"smtp"`
    Modules  []string   `mapstructure:"modules" validate:"required,min=1,dive,oneof=ip os cpu mem users"`
    HTTP     HTTPConfig `mapstructure:"http"`
}

// Load reads, validates config, supports env overrides.
func Load(path string) (*Config, error) {
    viper.AutomaticEnv()
    viper.SetConfigFile(path)
    viper.SetEnvKeyReplacer(nil) // default
    if err := viper.ReadInConfig(); err != nil {
        return nil, fmt.Errorf("read config: %w", err)
    }
    viper.SetDefault("schedule", "@every 30m")

    var cfg Config
    if err := viper.Unmarshal(&cfg); err != nil {
        return nil, fmt.Errorf("unmarshal: %w", err)
    }
    // validate
    validate := validator.New()
    if err := validate.Struct(&cfg); err != nil {
        return nil, fmt.Errorf("config validation: %w", err)
    }
    // check cron parse
    if _, err := cron.ParseStandard(cfg.Schedule); err != nil {
        return nil, fmt.Errorf("parse schedule: %w", err)
    }
    return &cfg, nil
}
