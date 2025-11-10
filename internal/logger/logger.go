package logger

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Logger struct {
	output io.Writer
}

var defaultLogger = New(os.Stderr)

func New(output io.Writer) *Logger {
	return &Logger{output: output}
}

func (l *Logger) log(level, message string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	fmt.Fprintf(l.output, "[%s] %s: %s\n", timestamp, level, message)
}

func (l *Logger) Info(message string) {
	l.log("INFO", message)
}

func (l *Logger) Error(message string) {
	l.log("ERROR", message)
}

func (l *Logger) Debug(message string) {
	l.log("DEBUG", message)
}

// Package-level functions for convenience
func Info(message string) {
	defaultLogger.Info(message)
}

func Error(message string) {
	defaultLogger.Error(message)
}

func Debug(message string) {
	defaultLogger.Debug(message)
}
