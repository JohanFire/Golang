package main

import (
	"fmt"
	"log"
	"os"

	"log/slog"
	// "golang.org/x/exp/slog"
)

func main() {
	println()

	var name string = "Johan"

	log.Println("Hello, World!")

	slog.Debug("This is a debug message")
	slog.Info("This is an Info message")
	slog.Info("This is a log with a variable", "name", name)
	/* In slog, to print variables, they are key=value attributes */
	slog.Error("This is an error message", slog.String("key", "value")) // easier to track key value variables
	slog.Warn("This is a warning message")

	println()
	fmt.Println("Creating slog text handler:")

	// Creating slog text handler
	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))

	logger.Info("This is an ugly as fuck log message")

	// set our new logger as default slog, in case u want to use slog
	slog.SetDefault(logger)
	slog.Error("Hi this is a slog but setting logger handler as default slog")

	println()
	fmt.Println("Setting a logger handler with DEBUG level")

	logger2 := slog.New(slog.NewTextHandler(
		os.Stderr,
		&slog.HandlerOptions{Level: slog.LevelDebug},
	))

	logger2.Debug("Now you can see DEBUG logs")
}
