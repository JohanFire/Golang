package main

import (
	"log"

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
	slog.Error("This is an error message", slog.String("key", "value"))
	slog.Warn("This is a warning message")

	println()
}
