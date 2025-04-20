package config

import (
	"fmt"
	"os"
	"strings"
	"time"

)

// LogLevel represents the level of logging
type LogLevel int 

const(
	// DebugLevel logs detailed information for debugging
	DebugLevel LogLevel = iota

	// InfoLevel logs general operational information
	InfoLevel

	// WarnLevel logs potentially harmful situations
	WarnLevel

	// ErrorLevel logs error events that might still allow the application to continue running
	ErrorLevel
)

// Logger provides structured logging capabilities
type Logger struct{
	level LogLevel
}

// NewLogger creates a new logger with specified level
func NewLogger(level string) *Logger{
	var logLevel LogLevel

	switch strings.ToLower(level){
	case "debug":
		logLevel = DebugLevel

	case "info":
		logLevel = InfoLevel

	case "warn":
		logLevel = WarnLevel

	case "error":
		logLevel = ErrorLevel

	default:
		logLevel = InfoLevel
	}

	return &Logger{level: logLevel}
}

// logMessage logs a message with specified level
func (l *Logger) logMessage(level LogLevel, format string, args ...interface{}){
	if level < l.level{
		return
	}

	var levelStr string

	switch level {
	case DebugLevel:
		levelStr = "DEBUG"

	case InfoLevel:
		levelStr = "INFO"

	case WarnLevel:
		levelStr = "WARN"

	case ErrorLevel:
		levelStr = "ERROR"
	}

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	message := fmt.Sprintf(format, args...)
	logLine := fmt.Sprintf("[%s] %s: %s\n", timestamp, levelStr, message)

	if level >= ErrorLevel {
		fmt.Fprint(os.Stderr, logLine)
	} else {
		fmt.Fprint(os.Stdout, logLine)
	}
}

// Debug logs a debug message
func (l *Logger) Debug(format string, args ...interface{}){
	l.logMessage(DebugLevel, format, args...)
}

// Info logs an info message
func (l *Logger) Info(format string, args ...interface{}){
	l.logMessage(InfoLevel, format, args...)
}

// Warn logs a warning message
func (l *Logger) Warn(format string, args ...interface{}){
	l.logMessage(WarnLevel, format, args...)
}

// Error logs an error message
func (l *Logger) Error(format string, args ...interface{}){
	l.logMessage(ErrorLevel, format, args...)
}