/* This script on: ./internal/logger/logger.go
======================================================

then on the main file initialize it as:

package main

import (
"go-router/internal/logger"
)

func () {
	logger.Init()

	logger.Debug("This is a debug message")
	logger.Debug(fmt.Sprintf("This is a debug message with a variable formatted: %d", num))
	}

======================================================
then on any other package or script dont initialize it again, just call it

package other_package

import (
	"go-router/internal/logger"
)

func () {
	logger.Debug("This is a debug message")
	logger.Debug(fmt.Sprintf("This is a debug message with a variable formatted: %d", num))
}
======================================================
*/

package logger

import (
	"io"
	"log/slog"
	"os"
	"path/filepath"

	"gopkg.in/natefinch/lumberjack.v2"
)

// logger is the private instance of slog.Logger
var logger *slog.Logger

// Init initializes the logger with file rotation
func Init() {
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
		MaxBackups: 3,       // Maximum number of old log files to retain
		Compress:   true,    // Compress rotated files
	}

	// Multiple writers: console and rotating file
	writers := io.MultiWriter(os.Stderr, rotatingLogger)

	// Configure slog with the rotating writer
	logger = slog.New(
		slog.NewTextHandler(
			writers,
			&slog.HandlerOptions{
				Level: slog.LevelDebug,
				// AddSource: true,
			},
		),
	)

	// Set as default logger
	slog.SetDefault(logger)
}

// Debug logs a debug message
func Debug(msg string, args ...any) {
	if logger == nil {
		Init()
	}
	logger.Debug(msg, args...)
}

// Info logs an info message
func Info(msg string, args ...any) {
	if logger == nil {
		Init()
	}
	logger.Info(msg, args...)
}

// Warn logs a warning message
func Warn(msg string, args ...any) {
	if logger == nil {
		Init()
	}
	logger.Warn(msg, args...)
}

// Error logs an error message
func Error(msg string, args ...any) {
	if logger == nil {
		Init()
	}
	logger.Error(msg, args...)
}

// GetLogger returns the logger instance for advanced use cases
func GetLogger() *slog.Logger {
	if logger == nil {
		Init()
	}
	return logger
}
