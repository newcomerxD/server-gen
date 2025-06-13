// File: cmd/server-gen/main.go
package main

import (
    "fmt"
    "os"
    "os/signal"
    "syscall"

    "github.com/bocaletto-luca/server-gen/internal/config"
    "github.com/bocaletto-luca/server-gen/internal/mailer"
    "github.com/bocaletto-luca/server-gen/internal/sysinfo"
    "github.com/robfig/cron/v3"
    "github.com/spf13/cobra"
)

var cfgFile string

func main() {
    root := &cobra.Command{
        Use:   "server-gen",
        Short: "Collects system info and emails it on a schedule",
        RunE:  run,
    }
    root.PersistentFlags().StringVar(&cfgFile, "config", "configs/config.yaml",
        "path to configuration file")
    if err := root.Execute(); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}

func run(cmd *cobra.Command, args []string) error {
    // Load configuration
    cfg, err := config.Load(cfgFile)
    if err != nil {
        return fmt.Errorf("load config: %w", err)
    }

    // Initialize cron scheduler
    scheduler := cron.New(
        cron.WithParser(cron.Minute|cron.Hour|cron.Dom|cron.Month|cron.Dow),
    )
    _, err = scheduler.AddFunc(cfg.Schedule, func() {
        data := sysinfo.Collect(cfg.Modules)
        if err := mailer.Send(cfg.SMTP, data); err != nil {
            fmt.Fprintf(os.Stderr, "mail send error: %v\n", err)
        }
    })
    if err != nil {
        return fmt.Errorf("schedule job: %w", err)
    }
    scheduler.Start()
    fmt.Println("🚀 server-gen started. Press Ctrl+C to stop.")

    // Graceful shutdown on SIGINT/SIGTERM
    sig := make(chan os.Signal, 1)
    signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
    <-sig

    fmt.Println("Shutting down…")
    scheduler.Stop()
    return nil
}
