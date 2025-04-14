package main

import (
	"log/slog"
	"math/rand"
	"os"
)

func main() {
	println()

	// handlerOptions := &slog.HandlerOptions{
	// 	Level: slog.LevelDebug,
	// }

	logger := slog.New(
		// slog.NewTextHandler(
		slog.NewJSONHandler(
			os.Stderr,
			// handlerOptions,
			&slog.HandlerOptions{
				Level: slog.LevelDebug,
				// AddSource: true,
			},
		),
	)

	logger.Debug("Hello World")
	logger.Info("This is an Info message", slog.String("key", "value"))
	logger.Info("This is an Info message", slog.String("user", "JohanFire"))
	logger.Warn(
		"This is a Warn message",
		slog.Int("id", rand.Intn(100)),
		slog.String("user", "JohanFire"),
	)
}
