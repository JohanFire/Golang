package main

import (
	"io"
	"log/slog"
	"os"
)

func main() {
	println()

	// LOG_FILE := os.Getenv("LOG_FILE")

	file, err := os.OpenFile("logs.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writers := io.MultiWriter(os.Stderr, file)

	logger := slog.New(
		slog.NewTextHandler(
			// slog.NewJSONHandler(
			writers,
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
