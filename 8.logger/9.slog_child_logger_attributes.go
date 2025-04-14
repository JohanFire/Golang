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
			// slog.NewJSONHandler(
			os.Stderr,
			// handlerOptions,
			&slog.HandlerOptions{
				Level: slog.LevelDebug,
				// AddSource: true,
			},
		),
	)
	slog.SetDefault(logger)
	println()

	logger.Debug("Hello World 1")
	slog.Debug("Hello World 2")

	println()

	requestGroup := slog.Group(
		"request",
		slog.String("method", "GET"), // imaginary GET request
	)

	requestLogger := logger.With(requestGroup, slog.String("extraKey", "this key is not in a group"))

	requestLogger.Info("Hi from requestLogger !")

}
