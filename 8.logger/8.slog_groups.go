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

	logger.Debug("Hello World")
	logger.Info("This is an Info message", slog.String("key", "value"))
	logger.Info("This is an Info message", slog.String("user", "JohanFire"))
	logger.Warn(
		"This is a Warn message",
		slog.Int("id", rand.Intn(100)),
		slog.String("user", "JohanFire"),
	)

	println()

	/* Groups:
	- users
	- requests
	*/
	userGroup := slog.Group(
		"users",
		slog.Int("id", rand.Intn(100)),
		slog.String("username", "JohanFire"),
	)

	slog.Info("using userGroup", userGroup)

	requestGroup := slog.Group(
		"request",
		slog.String("method", "GET"), // imaginary GET request
	)

	slog.Info("This is an endpoint request", requestGroup)

}
