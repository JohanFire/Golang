package main

import (
	"os"
	"runtime"

	"log/slog"
	// "golang.org/x/exp/slog"
)

func main() {
	println()

	logger := slog.New(
		slog.NewJSONHandler(
			os.Stderr,
			&slog.HandlerOptions{Level: slog.LevelDebug},
		),
	)
	logger = logger.With(slog.String("app_version", "1.0.0")) // adds a defined attribute to every log

	logger.Debug("Hello there", slog.String("version", runtime.Version()))
	logger.Info(
		"slog Groups, like a nested JSON",
		slog.Group("OS Info",
			slog.String("OS", runtime.GOOS),
			slog.Int("CPUs", runtime.NumCPU()),
			slog.String("Arch", runtime.GOARCH),
		),
	)

	logger.Info("Hi!")
}
