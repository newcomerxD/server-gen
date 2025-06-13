// File: internal/config/config_test.go
package config

import (
    "os"
    "testing"
)

func TestLoad_Valid(t *testing.T) {
    os.Setenv("SMTP_HOST", "smtp.test.com")
    os.Setenv("SMTP_PORT", "1025")
    os.Setenv("SMTP_USER", "u")
    os.Setenv("SMTP_PASS", "p")
    os.Setenv("SMTP_FROM", "from@test.com")
    os.Setenv("SMTP_TO1", "to@test.com")

    cfg, err := Load("../configs/config.yaml")
    if err != nil {
        t.Fatalf("Load failed: %v", err)
    }
    if cfg.SMTP.Host != "smtp.test.com" {
        t.Errorf("expected host smtp.test.com, got %s", cfg.SMTP.Host)
    }
}
