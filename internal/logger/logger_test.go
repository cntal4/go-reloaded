package logger

import (
	"bytes"
	"strings"
	"testing"
)

func TestLogger_Info(t *testing.T) {
	var buf bytes.Buffer
	logger := New(&buf)

	logger.Info("test message")

	output := buf.String()
	if !strings.Contains(output, "INFO") {
		t.Errorf("Expected INFO in output, got: %s", output)
	}
	if !strings.Contains(output, "test message") {
		t.Errorf("Expected 'test message' in output, got: %s", output)
	}
}

func TestLogger_Error(t *testing.T) {
	var buf bytes.Buffer
	logger := New(&buf)

	logger.Error("error message")

	output := buf.String()
	if !strings.Contains(output, "ERROR") {
		t.Errorf("Expected ERROR in output, got: %s", output)
	}
	if !strings.Contains(output, "error message") {
		t.Errorf("Expected 'error message' in output, got: %s", output)
	}
}

func TestLogger_Debug(t *testing.T) {
	var buf bytes.Buffer
	logger := New(&buf)

	logger.Debug("debug message")

	output := buf.String()
	if !strings.Contains(output, "DEBUG") {
		t.Errorf("Expected DEBUG in output, got: %s", output)
	}
	if !strings.Contains(output, "debug message") {
		t.Errorf("Expected 'debug message' in output, got: %s", output)
	}
}

func TestLogger_DefaultLogger(t *testing.T) {
	// Test that default logger doesn't panic
	Info("test info")
	Error("test error")
	Debug("test debug")
}
