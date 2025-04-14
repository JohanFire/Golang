package main

import (
	"log/slog"
	"os"
)

func main() {
	println()

	// handlerOptions := &slog.HandlerOptions{
	// 	Level: slog.LevelDebug,
	// }

	logger := slog.New(
		slog.NewTextHandler(
			os.Stderr,
			// handlerOptions,
			&slog.HandlerOptions{
				Level: slog.LevelDebug,
				// AddSource: true,
			},
		),
	)

	logger.Info("Hello World")
}
