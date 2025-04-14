package main

import (
	"io"
	"log/slog"
	"os"
	"path/filepath"

	"gopkg.in/natefinch/lumberjack.v2" // for log rotation
)

func main() {
	println()

	// Create logs directory if it doesn't exist
	logsDir := "logs"
	if err := os.MkdirAll(logsDir, 0755); err != nil {
		panic(err)
	}

	// Configure lumberjack for log rotation
	logFile := filepath.Join(logsDir, "logs.log")
	rotatingLogger := &lumberjack.Logger{
		Filename:   logFile, // Path to the log file
		MaxSize:    10,      // Maximum size in megabytes before rotating
		MaxBackups: 5,       // Maximum number of old log files to retain
		// MaxAge:     30,      // Maximum number of days to retain old log files
		Compress: true, // Compress rotated files
	}

	// Multiple writers: console and rotating file
	writers := io.MultiWriter(os.Stderr, rotatingLogger)

	// Configure slog with the rotating writer
	logger := slog.New(
		slog.NewTextHandler(
			writers,
			&slog.HandlerOptions{
				Level:     slog.LevelDebug,
				AddSource: true,
			},
		),
	)

	// Set as default logger
	slog.SetDefault(logger)
	println()

	// Usage examples
	logger.Debug("This is a debug message")
	logger.Info("This is an info message")
	logger.Warn("This is a warning message")
	logger.Error("This is an error message")
}
