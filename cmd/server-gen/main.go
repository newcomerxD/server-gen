// File: cmd/server-gen/main.go
package main

import (
    "context"
    "fmt"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"

    "github.com/bocaletto-luca/server-gen/internal/config"
    "github.com/bocaletto-luca/server-gen/internal/mailer"
    "github.com/bocaletto-luca/server-gen/internal/sysinfo"
    "github.com/robfig/cron/v3"
    "github.com/rs/zerolog"
    "github.com/rs/zerolog/log"
    "github.com/spf13/cobra"
)

var cfgFile string

func main() {
    // Setup logger
    log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
    root := &cobra.Command{
        Use:   "server-gen",
        Short: "Collect & email system info on schedule",
        RunE:  run,
    }
    root.PersistentFlags().StringVar(&cfgFile, "config", "configs/config.yaml", "config file")
    if err := root.Execute(); err != nil {
        log.Fatal().Err(err).Msg("command failed")
    }
}

func run(cmd *cobra.Command, args []string) error {
    log.Info().Msg("Loading config")
    cfg, err := config.Load(cfgFile)
    if err != nil {
        return err
    }
    log.Info().Msgf("Config loaded, schedule=%s", cfg.Schedule)

    // HTTP health endpoint
    mux := http.NewServeMux()
    mux.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("ok"))
    })
    srv := &http.Server{Addr: cfg.HTTP.Addr, Handler: mux}
    go func() {
        log.Info().Msgf("Starting HTTP server on %s", cfg.HTTP.Addr)
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatal().Err(err).Msg("HTTP server failed")
        }
    }()

    // Cron scheduler
    c := cron.New()
    _, err = c.AddFunc(cfg.Schedule, func() {
        log.Info().Msg("Job started")
        d := sysinfo.Collect(cfg.Modules)
        if err := mailer.Send(cfg.SMTP, d); err != nil {
            log.Error().Err(err).Msg("mailer.Send failure")
        } else {
            log.Info().Msg("Email sent successfully")
        }
    })
    if err != nil {
        return fmt.Errorf("invalid schedule: %w", err)
    }
    c.Start()
    log.Info().Msg("Scheduler started")

    // Graceful shutdown
    stop := make(chan os.Signal, 1)
    signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
    <-stop
    log.Info().Msg("Shutting down...")

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    c.Stop()
    srv.Shutdown(ctx)
    return nil
}
