package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/caarlos0/env/v11"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"{{MODULE_PATH}}/frontend"
)

type config struct {
	Port string `env:"PORT" envDefault:"{{PORT}}"`
}

func main() {
	// Root context for the whole application
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	var cfg config
	if err := env.Parse(&cfg); err != nil {
		logger.ErrorContext(ctx, "failed to parse env", "error", err)
		os.Exit(1)
	}

	mux := http.NewServeMux()

	// Health check endpoint
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	// Mount static assets at root
	mux.Handle("/", http.FileServer(http.FS(frontend.Assets)))

	server := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: h2c.NewHandler(mux, &http2.Server{}),
	}

	// Run server in a goroutine
	go func() {
		logger.InfoContext(ctx, "server starting", "port", cfg.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.ErrorContext(ctx, "server failed", "error", err)
			os.Exit(1)
		}
	}()

	// Wait for interruption signal
	<-ctx.Done()
	logger.InfoContext(ctx, "shutting down gracefully...")

	// Shutdown HTTP server
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(shutdownCtx); err != nil {
		logger.ErrorContext(ctx, "server forced to shutdown", "error", err)
	}

	logger.InfoContext(ctx, "shutdown complete")
}
