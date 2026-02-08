package main

import (
	"ddr13/riddlercore/cmd"
	"log/slog"
	"os"
	"time"

	_ "github.com/davecgh/go-spew/spew"
	"github.com/lmittmann/tint"
)

func main() {
	cmd.Execute()

	var handler slog.Handler
	if os.Getenv("RIDDLERCORE_ENV") == "production" {
		handler = slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelInfo})
	} else {
		handler = tint.NewHandler(os.Stderr, &tint.Options{
			Level:      slog.LevelDebug,
			TimeFormat: time.Kitchen,
		})
	}

	// This makes slog.Info() use your configured handler globally
	slog.SetDefault(slog.New(handler))
}
